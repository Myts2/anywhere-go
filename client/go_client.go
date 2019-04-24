package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	p "github.com/Myts2/anywhere-go/common/proto"
	"github.com/armon/go-socks5"
	"github.com/golang/protobuf/proto"
	"net"
)

var hostName = ""
var nowConnection = make(map[string]map[int32]portLink)
var portList = make(map[string]map[int32]portLink)
var portStop = make(map[string]struct{
	L net.Listener
	stop chan struct{}
})
var socksPort int

type portLink struct {
	Ws mid
}

type mid struct {
	DataToWrite chan []byte
	DataToSend  chan []byte
	Stop        chan struct{}
	BeStop       bool
}

func initMid(c net.Conn) mid {
	m := mid{
		DataToWrite: make(chan []byte),
		DataToSend:  make(chan []byte),
		Stop:        make(chan struct{}),
		BeStop: false,
	}
	go func() {
		for !m.BeStop {
			buf := <-m.DataToWrite
			_, err := c.Write(buf)
			if err != nil {
				if !m.BeStop{
					close(m.Stop)
				}
			}
		}
		return
	}()
	go func() {
		r := bufio.NewReader(c)
		for !m.BeStop {
			buf := make([]byte, 40960)
			l, err := r.Read(buf)
			if err != nil {
				return
			}
			m.DataToSend <- buf[:l]
		}
		return
	}()
	go func() {
		<-m.Stop
		m.BeStop = true
		_ = c.Close()
		return
	}()
	return m
}

func initLink(c net.Conn, r string, ch int32, m mid, f string) portLink {
	link := portLink{
		Ws: m,
	}
	var stop = false
	go func() {
		for !stop {
			buf := <-m.DataToSend
			msgToSend := &p.Msg{
				Chan: ch,
				Data: buf,
				From: f,
			}
			bs, _ := msgToSend.Marshal()
			pkgToSend := &p.Pkg{
			From: hostName,
				To:   r,
					Type: p.Type_PassWay,
					Data: bs,
			}
			b, _ := pkgToSend.Marshal()
			_, _ = c.Write(b)
		}
	}()
	go func() {
		<-m.Stop
		stop = true
		msgToSend := &p.Msg{
			Chan:  ch,
			Data:  nil,
			From:  f,
			Close: true,
		}
		bs, _ := msgToSend.Marshal()
		pkgToSend := &p.Pkg{
			From: hostName,
			To:   r,
			Type: p.Type_PassWay,
			Data: bs,
		}
		b, _ := pkgToSend.Marshal()
		_, _ = c.Write(b)
	}()
	return link
}

func CreateSocks5Server(finish chan int) {
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	normal := 8000
	for {
		l, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d",normal))
		if err == nil {
			finish <- normal
			close(finish)
			_ = server.Serve(l)
			return
		}
		normal += 1
	}
}

func createConnection(c net.Conn, pkg *p.Pkg, Chan int32) {
	con, _ := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d",socksPort))
	nowConnection[pkg.From][Chan] = initLink(c, pkg.From, Chan, initMid(con), pkg.From)
	go func() {
		<-nowConnection[pkg.From][Chan].Ws.Stop
		delete(nowConnection[pkg.From], Chan)
	}()
}

func handler(c net.Conn, OK chan p.Type) {
	r := bufio.NewReader(c)
	for {
		buf := make([]byte, 40960)
		l, err := r.Read(buf)
		if err != nil {
			fmt.Printf("Fail to read data, %s\n", err)
			break
		}
		pkg := &p.Pkg{}
		err = proto.Unmarshal(buf[:l], pkg)
		if err != nil {
			fmt.Printf("Fail to Unmarshal data, %s\n", err)
			continue
		}
		switch pkg.Type {
		case p.Type_Hello:
			hostName = pkg.To
			OK <- p.Type_Hello
			break
		case p.Type_CantFind:
			OK <- p.Type_CantFind
			break
		case p.Type_ConnectionFrom:
			fmt.Printf("\nConnect from %s\n", pkg.From)
			go func() {
				nowConnection[pkg.From] = map[int32]portLink{}
				pkgToSend := p.Pkg{
					From: hostName,
					To:   pkg.From,
					Type: p.Type_OK,
					Data: nil,
				}
				b, _ := pkgToSend.Marshal()
				_, _ = c.Write(b)
			}()
			// OK <- int(p.Type_ConnectionFrom)
			break
		case p.Type_OK:
			OK <- p.Type_OK
			break
		case p.Type_CantConnectSocks:
			OK <- p.Type_CantConnectSocks
			break
		case p.Type_ConnectClosed:
			fmt.Printf("\nConnection closed %s\n", pkg.From)
			go func() {
				if l, ok := nowConnection[pkg.From]; ok {
					for _, ws := range l {
						if !ws.Ws.BeStop {
							close(ws.Ws.Stop)
						}
					}
					delete(portList, pkg.From)
				} else if l, ok := portList[pkg.From]; ok {
					for _, ws := range l {
						if !ws.Ws.BeStop {
							close(ws.Ws.Stop)
						}
					}
					delete(portList, pkg.From)
					close(portStop[pkg.From].stop)
				}
			}()
			break
		case p.Type_Ping:
			go func() {
				pkgToSend := &p.Pkg{
					From: hostName,
					To:   "",
					Type: p.Type_Ping,
					Data: nil,
				}
				b, _ := pkgToSend.Marshal()
				_, _ = c.Write(b)
			}()
			break
		case p.Type_PassWay:
			go func() {
				msg := &p.Msg{}
				_ = proto.Unmarshal(pkg.Data, msg)
				var aims *map[string]map[int32]portLink
				if msg.From == hostName {
					aims = &portList
				} else {
					aims = &nowConnection
				}
				if l, ok := (*aims)[pkg.From][msg.Chan]; ok {
					if msg.Close {
						if !l.Ws.BeStop {
							close(l.Ws.Stop)
						}
						return
					}
					l.Ws.DataToWrite <- msg.Data
				} else if msg.From != hostName {
					createConnection(c, pkg, msg.Chan)
					(*aims)[pkg.From][msg.Chan].Ws.DataToWrite <- msg.Data
				}
			}()
			break
		default:
			return
		}
	}
}

func main() {
	// Parse some flags
	remoteAddr := flag.String("r", "123.206.25.54", "远程服务器IP")
	port := flag.Int("p", 1208, "远程服务器端口")
	flag.Parse()
	socks5Creat := make(chan int)
	go CreateSocks5Server(socks5Creat)
	remote := fmt.Sprintf("%s:%d",*remoteAddr,*port)
	conn, err := net.Dial("tcp", remote)
	defer conn.Close()
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	socksPort = <-socks5Creat
	Cok := make(chan p.Type)
	go handler(conn, Cok)
	pkgToSend := &p.Pkg{From:"???",To:"???",Type: p.Type_Hello}
	b,_ := pkgToSend.Marshal()
	_,err = conn.Write(b)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	<-Cok
	for {
		fmt.Printf(`
Your HostName is %s
What do you want to do now?
1. Connected to a client
2. List NowConnection
3. List PortLink
Please enter the label > `, hostName)
		var chance int
		_, err := fmt.Scanf("%d", &chance)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		switch chance {
		case 1:
			fmt.Printf("The host name you want to connect to > ")
			var rHostName string
			_, _ = fmt.Scanf("%s", &rHostName)
			fmt.Println("Try to connect. Wait please.")
			pkgToSend := &p.Pkg{
				From: hostName,
				To:   rHostName,
				Type: p.Type_WantToConnection,
				Data: nil,
			}
			b, _ := pkgToSend.Marshal()
			_, _ = conn.Write(b)
			switch <-Cok {
			case p.Type_CantFind:
				fmt.Printf("Cant find it.\n")
				continue
			case p.Type_CantConnectSocks:
				fmt.Printf(" Cant connect %s\n", rHostName)
				continue
			case p.Type_OK:
				fmt.Printf("Connected %s\n", rHostName)
				break
			default:
				continue
			}
			err = errors.New("no error")
			for err != nil {
				var addr string
				fmt.Printf("What address do you want to listen (like: 127.0.0.1:8001) > ")
				_, _ = fmt.Scanf("%s",&addr)
				l, err := net.Listen("tcp", addr)
				if err != nil {
					fmt.Println(err.Error() + "\nTry another.")
					continue
				}
				portStop[rHostName] = struct {
					L    net.Listener
					stop chan struct{}
				}{L: l, stop: make(chan struct{})}
				go func() {
					<-portStop[rHostName].stop
					_ = l.Close()
					delete(portStop, rHostName)
				}()
				go func() {
					var ch int32 = 0
					portList[rHostName] = make(map[int32]portLink)
					for {
						c, _ := l.Accept()
						if portList[rHostName] != nil {
							portList[rHostName][ch] = initLink(conn, rHostName, ch, initMid(c), hostName)
						}
						ch++
					}
				}()
				fmt.Println("OK!")
				break
			}
			continue
		case 2:
			println("The nodes currently connected to this computer are:")
			for index := range nowConnection {
				fmt.Printf("%s\n", index)
			}
			var removeHostName string
			fmt.Printf("Enter the host name to disconnect (q exit) > ")
			_,_ = fmt.Scanf("%s", &removeHostName)
			if i, ok := nowConnection[removeHostName]; ok {
				for _, ws := range i {
					if !ws.Ws.BeStop {
						close(ws.Ws.Stop)
					}
				}
				pkgToSend := &p.Pkg{
					From:                 hostName,
					To:                   removeHostName,
					Type:                 p.Type_ConnectClosed,
					Data:                 nil,
				}
				b,_:= pkgToSend.Marshal()
				_,_ = conn.Write(b)
				delete(nowConnection, removeHostName)
				fmt.Println("Close success.")
			}
			continue
		case 3:
			println("The nodes currently connected to this computer are:")
			for index,v := range portStop {
				fmt.Printf("%s -- %s\n", index, v.L.Addr().String())
			}
			var removeHostName string
			fmt.Printf("Enter the host name to disconnect (q exit) > ")
			_,_ = fmt.Scanf("%s", &removeHostName)
			if i, ok := portList[removeHostName]; ok {
				for _, ws := range i {
					if !ws.Ws.BeStop {
						close(ws.Ws.Stop)
					}
				}
				pkgToSend := &p.Pkg{
					From:                 hostName,
					To:                   removeHostName,
					Type:                 p.Type_ConnectClosed,
					Data:                 nil,
				}
				b,_:= pkgToSend.Marshal()
				_,_ = conn.Write(b)
				delete(portList, removeHostName)
				close(portStop[removeHostName].stop)
				fmt.Println("Close success.")
			}
			continue
		default:
			fmt.Println("Wrong input")
			continue
		}
	}
}
