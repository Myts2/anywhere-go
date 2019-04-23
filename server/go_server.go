package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	p "github.com/Myts2/anywhere-go/common/proto"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
)

var m = make(map[string]net.Conn)

func dialHost(hostname string, comConn string) (err int) {
	conn, _ := m[comConn]
	if remoteConn, ok := m[hostname]; ok {
		_, e := remoteConn.Write([]byte("FM" + comConn))
		if e != nil {
			return -1
		}
		_, e = conn.Write([]byte("OK" + hostname))
		if e != nil {
			return -1
		}
		//stop := make(chan bool)
		_, _ = io.Copy(remoteConn, conn)
		return 0
	} else {
		return -1
	}
}

//func relay(src net.Conn, dst net.Conn, stop chan bool) {
//	_, ok :=io.Copy(dst, src)
//	if ok != nil {
//		return
//	}
//	stop <- true
//	return
//}
func storeConn(conn net.Conn) string {
	signByte := []byte(conn.RemoteAddr().String())
	hash := md5.New()
	hash.Write(signByte)
	hostName := hex.EncodeToString(hash.Sum(nil))[:8]
	m[hostName] = conn
	return hostName
}

func connHandler(c net.Conn, h chan string) {
	if c == nil {
		return
	}
	defer c.Close()
	var hostname string
	for {
		buf := make([]byte, 40960)
		cnt, err := c.Read(buf)
		if err != nil {
			break
		}
		pkg := &p.Pkg{}
		err = proto.Unmarshal(buf[:cnt], pkg)
		if err != nil {
			continue
		}
		switch pkg.Type {
		case p.Type_Hello:
			hostname = <- h
			pkgToSend := &p.Pkg{
				From: "",
				To:   hostname,
				Type: p.Type_Hello,
				Data: nil,
			}
			b, _ := pkgToSend.Marshal()
			_, e := c.Write(b)
			if e != nil {
				break
			}
		case p.Type_OK ,p.Type_PassWay, p.Type_ConnectClosed, p.Type_CantConnectSocks:
			if co, ok := m[pkg.To]; ok {
				_, err := co.Write(buf[:cnt])
				if err != nil {
					pkgToSend := &p.Pkg{
						From: pkg.To,
						To:   hostname,
						Type: p.Type_ConnectClosed,
						Data: nil,
					}
					b, _ := pkgToSend.Marshal()
					_, err = c.Write(b)
					if err != nil {
						break
					}
				}
			}
			continue
		case p.Type_WantToConnection:
			if co, ok := m[pkg.To]; ok {
				pkgToSend := &p.Pkg{
					From: hostname,
					To:   pkg.To,
					Type: p.Type_ConnectionFrom,
					Data: nil,
				}
				b, _ := pkgToSend.Marshal()
				_, err = co.Write(b)
				if err != nil {
					pkgToSend := &p.Pkg{
						From: pkg.To,
						To:   hostname,
						Type: p.Type_ConnectClosed,
						Data: nil,
					}
					b, _ := pkgToSend.Marshal()
					_, err = c.Write(b)
					if err != nil {
						break
					}
				}
			}else {
				pkgToSend := &p.Pkg{
					From: pkg.To,
					To:   hostname,
					Type: p.Type_CantFind,
					Data: nil,
				}
				b, _ := pkgToSend.Marshal()
				_, err = c.Write(b)
				if err != nil {
					break
				}
			}
			continue
		default:
			continue
		}
	}
	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
	delete(m, hostname)
	return
}

func main() {
	server, err := net.Listen("tcp", ":1208")
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
	}
	fmt.Println("Server Started ...")
	for {
		h := make(chan string)
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		go connHandler(conn, h)
		go func() {h <- storeConn(conn);close(h)}()
	}
}
