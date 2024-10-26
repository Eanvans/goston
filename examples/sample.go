package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/net/proxy"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建一个 SOCKS5 代理拨号器
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
	if err != nil {
		log.Fatalf("Failed to create SOCKS5 dialer: %v", err)
	}

	// 读取客户端请求的目标地址
	br := bufio.NewReader(conn)
	req, err := proxy.SOCKS5RequestFrom(br)
	if err != nil {
		log.Printf("Failed to read SOCKS5 request: %v", err)
		return
	}

	// 通过 SOCKS5 代理连接到目标地址
	targetConn, err := dialer.Dial("tcp", req.Address())
	if err != nil {
		log.Printf("Failed to connect to target: %v", err)
		return
	}
	defer targetConn.Close()

	// 回复客户端表示连接成功
	if err := req.Reply(nil, br); err != nil {
		log.Printf("Failed to reply to client: %v", err)
		return
	}

	// 双向转发数据
	go func() {
		io.Copy(targetConn, conn)
	}()

	// 统计转发的数据包大小
	totalBytes := 0
	reader := bufio.NewReader(targetConn)
	for {
		buf := make([]byte, 4096)
		n, err := reader.Read(buf)
		if n > 0 {
			totalBytes += n
			fmt.Printf("Data packet size: %d bytes, Total bytes: %d\n", n, totalBytes)
		}
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed")
			} else {
				log.Printf("Error reading from target: %v", err)
			}
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":1081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()
	log.Println("Listening on :1081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
