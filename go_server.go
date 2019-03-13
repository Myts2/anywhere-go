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

func dial_host(hostname string,conn net.Conn)(err int){
	if remote_conn,ok := m[hostname]; ok{
		conn.Write([]byte("OK! START!\n"))
		stop := make(chan bool)

		go relay(remote_conn, conn, stop)
		go relay(conn, remote_conn, stop)
		select {
		case <-stop:
			return 0
		}
	}else{
		return -1
	}
}

func relay(src net.Conn, dst net.Conn, stop chan bool) {
	io.Copy(dst, src)
	dst.Close()
	src.Close()
	stop <- true
	return
}
func store_conn(dict map[string]net.Conn,conn net.Conn) (string){
	signByte := []byte(conn.RemoteAddr().String())
	hash := md5.New()
	hash.Write(signByte)
	hashname := hex.EncodeToString(hash.Sum(nil))[:8]
	dict[hashname] = conn
	return hashname

}



func connHandler(c net.Conn) {
	if c == nil {
		return
	}

	hostname := store_conn(m,c)
	defer delete(m,hostname)
	c.Write([]byte("Your hostname is :"+hostname+"\n"))
	for {
		c.Write([]byte("You want to call who?\n"))
		buf := make([]byte, 4096)
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		raw_hostname_input := strings.TrimSpace(string(buf[0:cnt]))
		hostname_input := strings.Split(raw_hostname_input," ")
		result := dial_host(hostname_input[0],c)
		if result != -1 {
			break
		} else {
			c.Write([]byte("Wrong,again\n"))
			continue
		}
		fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
		defer c.Close()
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
