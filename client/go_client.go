package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var hostName = ""

func connReader(c net.Conn, stop chan bool) {
	buf := make([]byte, 1024)
	for {
		cnt, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			break
		}
		fmt.Print(string(buf[0:cnt]))
	}
	stop <- false
	return
}
func connWriter(c net.Conn, stop chan bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		_, e := c.Write([]byte(input))
		if e != nil {
			fmt.Printf("Fail to write data, %s\n", e)
			break
		}
	}
	stop <- false
	return
}
func connTryer(c net.Conn, connected chan bool) {
	buf := make([]byte, 1024)
	for {
		_, err := c.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			return
		}
		switch string(buf[0:2]) {
		case "HM":
			hostName = string(buf[2:10])
			fmt.Printf("Your hostname is %s.\nYou want to call who?\n", hostName)
			break
		case "WA":
			fmt.Printf("Wrong,again\n")
			break
		case "FM":
			connected <- true
			fmt.Printf("Connect from %s\n", string(buf[2:10]))
			_, _ = c.Write([]byte("OK" + string(buf[2:10])))
			return
		case "OK":
			connected <- true
			fmt.Printf("Connected %s\n", string(buf[2:10]))
			return
		default:
			return
		}
	}
}
func connHandler(c net.Conn) {
	connected := make(chan bool)
	stop := make(chan bool)
	go connWriter(c, stop)
	go connTryer(c, connected)
	for {
		select {
		case <-connected:
			go connReader(c, stop)
		case <-stop:
			return
		}
	}

}
func main() {
	conn, err := net.Dial("tcp", "localhost:1208")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	connHandler(conn)
}
