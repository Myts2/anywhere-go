package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"strings"
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
func storeConn(dict map[string]net.Conn, conn net.Conn) string {
	signByte := []byte(conn.RemoteAddr().String())
	hash := md5.New()
	hash.Write(signByte)
	hostName := hex.EncodeToString(hash.Sum(nil))[:8]
	dict[hostName] = conn
	return hostName
}

func connHandler(c net.Conn) {
	if c == nil {
		return
	}

	hostname := storeConn(m, c)
	defer delete(m, hostname)
	_, e := c.Write([]byte("HM" + hostname))
	if e != nil {
		return
	}
	for {
		buf := make([]byte, 4096)
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			e = c.Close()
			if e != nil {
				return
			}
			break
		}
		if string(buf[0:2]) == "OK" {
			_, _ = io.Copy(m[string(buf[2:10])], c)
			break
		} else {
			rawHostnameInput := strings.TrimSpace(string(buf[0:cnt]))
			hostnameInputs := strings.Split(rawHostnameInput, " ")
			result := dialHost(hostnameInputs[0], hostname)
			if result != -1 {
				break
			} else {
				_, e = c.Write([]byte("WA"))
				if e != nil {
					return
				}
				continue
			}
		}
	}
	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
	e = c.Close()
	if e != nil {
		return
	}
}

func main() {
	server, err := net.Listen("tcp", ":1208")
	if err != nil {
		fmt.Printf("Fail to start server, %s\n", err)
	}
	fmt.Println("Server Started ...")
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Printf("Fail to connect, %s\n", err)
			break
		}
		go connHandler(conn)
	}
}
