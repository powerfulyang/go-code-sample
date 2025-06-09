package network

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// UDPServer UDP服务器
type UDPServer struct {
	address string
	conn    *net.UDPConn
	clients map[string]*UDPClient
	mutex   sync.RWMutex
	running bool
}

// UDPClient UDP客户端信息
type UDPClient struct {
	Addr     *net.UDPAddr
	LastSeen time.Time
}

// NewUDPServer 创建UDP服务器
func NewUDPServer(address string) *UDPServer {
	return &UDPServer{
		address: address,
		clients: make(map[string]*UDPClient),
		running: false,
	}
}

// Start 启动UDP服务器
func (s *UDPServer) Start() error {
	addr, err := net.ResolveUDPAddr("udp", s.address)
	if err != nil {
		return fmt.Errorf("解析UDP地址失败: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("启动UDP服务器失败: %v", err)
	}

	s.conn = conn
	s.running = true

	fmt.Printf("🚀 UDP服务器启动在 %s\n", s.address)

	// 启动客户端清理协程
	go s.cleanupClients()

	// 处理UDP消息
	buffer := make([]byte, 1024)
	for s.running {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			if s.running {
				fmt.Printf("读取UDP消息失败: %v\n", err)
			}
			continue
		}

		// 处理消息
		go s.handleMessage(clientAddr, string(buffer[:n]))
	}

	return nil
}

// Stop 停止UDP服务器
func (s *UDPServer) Stop() error {
	s.running = false

	if s.conn != nil {
		s.conn.Close()
	}

	// 清空客户端列表
	s.mutex.Lock()
	s.clients = make(map[string]*UDPClient)
	s.mutex.Unlock()

	fmt.Println("🛑 UDP服务器已停止")
	return nil
}

// handleMessage 处理UDP消息
func (s *UDPServer) handleMessage(clientAddr *net.UDPAddr, message string) {
	clientKey := clientAddr.String()
	message = strings.TrimSpace(message)

	// 更新客户端信息
	s.mutex.Lock()
	s.clients[clientKey] = &UDPClient{
		Addr:     clientAddr,
		LastSeen: time.Now(),
	}
	s.mutex.Unlock()

	fmt.Printf("📨 收到UDP消息 [%s]: %s\n", clientKey, message)

	// 处理不同的命令
	switch {
	case message == "ping":
		s.sendMessage(clientAddr, "pong")

	case message == "time":
		timeStr := fmt.Sprintf("服务器时间: %s", time.Now().Format("2006-01-02 15:04:05"))
		s.sendMessage(clientAddr, timeStr)

	case strings.HasPrefix(message, "echo "):
		echoMsg := strings.TrimPrefix(message, "echo ")
		response := fmt.Sprintf("回显: %s", echoMsg)
		s.sendMessage(clientAddr, response)

	case message == "clients":
		s.mutex.RLock()
		clientList := fmt.Sprintf("在线客户端数量: %d", len(s.clients))
		for addr, client := range s.clients {
			clientList += fmt.Sprintf("\n- %s (最后活跃: %s)",
				addr, client.LastSeen.Format("15:04:05"))
		}
		s.mutex.RUnlock()
		s.sendMessage(clientAddr, clientList)

	case strings.HasPrefix(message, "broadcast "):
		broadcastMsg := strings.TrimPrefix(message, "broadcast ")
		count := s.broadcastMessage(clientAddr, broadcastMsg)
		response := fmt.Sprintf("消息已广播给 %d 个客户端", count)
		s.sendMessage(clientAddr, response)

	case message == "help":
		help := `UDP服务器可用命令:
- ping: 测试连接
- time: 获取服务器时间
- echo <message>: 回显消息
- clients: 查看在线客户端
- broadcast <message>: 广播消息
- help: 显示帮助信息`
		s.sendMessage(clientAddr, help)

	default:
		response := fmt.Sprintf("未知命令: %s (发送 'help' 查看可用命令)", message)
		s.sendMessage(clientAddr, response)
	}
}

// sendMessage 发送消息给指定客户端
func (s *UDPServer) sendMessage(clientAddr *net.UDPAddr, message string) {
	if s.conn == nil {
		return
	}

	_, err := s.conn.WriteToUDP([]byte(message), clientAddr)
	if err != nil {
		fmt.Printf("发送UDP消息失败: %v\n", err)
	}
}

// broadcastMessage 广播消息给所有客户端
func (s *UDPServer) broadcastMessage(senderAddr *net.UDPAddr, message string) int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	broadcastMsg := fmt.Sprintf("[广播 from %s]: %s", senderAddr.String(), message)
	count := 0

	for _, client := range s.clients {
		if !client.Addr.IP.Equal(senderAddr.IP) || client.Addr.Port != senderAddr.Port {
			s.sendMessage(client.Addr, broadcastMsg)
			count++
		}
	}

	return count
}

// cleanupClients 清理不活跃的客户端
func (s *UDPServer) cleanupClients() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for s.running {
		select {
		case <-ticker.C:
			s.mutex.Lock()
			now := time.Now()
			for addr, client := range s.clients {
				// 如果客户端超过5分钟没有活动，则移除
				if now.Sub(client.LastSeen) > 5*time.Minute {
					fmt.Printf("🧹 清理不活跃UDP客户端: %s\n", addr)
					delete(s.clients, addr)
				}
			}
			s.mutex.Unlock()
		}
	}
}

// GetClientCount 获取客户端数量
func (s *UDPServer) GetClientCount() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.clients)
}

// SimpleUDPClient 简单UDP客户端
type SimpleUDPClient struct {
	serverAddr string
	conn       *net.UDPConn
}

// NewSimpleUDPClient 创建简单UDP客户端
func NewSimpleUDPClient(serverAddr string) *SimpleUDPClient {
	return &SimpleUDPClient{
		serverAddr: serverAddr,
	}
}

// Connect 连接到UDP服务器
func (c *SimpleUDPClient) Connect() error {
	serverAddr, err := net.ResolveUDPAddr("udp", c.serverAddr)
	if err != nil {
		return fmt.Errorf("解析服务器地址失败: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return fmt.Errorf("连接UDP服务器失败: %v", err)
	}

	c.conn = conn
	fmt.Printf("✅ 已连接到UDP服务器: %s\n", c.serverAddr)
	return nil
}

// Disconnect 断开连接
func (c *SimpleUDPClient) Disconnect() error {
	if c.conn != nil {
		c.conn.Close()
		fmt.Println("📤 已断开UDP连接")
	}
	return nil
}

// SendMessage 发送消息
func (c *SimpleUDPClient) SendMessage(message string) error {
	if c.conn == nil {
		return fmt.Errorf("未连接到服务器")
	}

	_, err := c.conn.Write([]byte(message))
	return err
}

// ReadMessage 读取消息
func (c *SimpleUDPClient) ReadMessage() (string, error) {
	if c.conn == nil {
		return "", fmt.Errorf("未连接到服务器")
	}

	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

// SendAndReceive 发送消息并接收响应
func (c *SimpleUDPClient) SendAndReceive(message string) (string, error) {
	if err := c.SendMessage(message); err != nil {
		return "", err
	}

	// 设置读取超时
	c.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	response, err := c.ReadMessage()
	c.conn.SetReadDeadline(time.Time{}) // 清除超时

	return response, err
}

// UDPExamples UDP网络编程示例
func UDPExamples() {
	fmt.Println("=== UDP网络编程示例 ===")
	fmt.Println()
	fmt.Println("这个示例演示了UDP服务器和客户端的实现")
	fmt.Println()
	fmt.Println("UDP特点:")
	fmt.Println("- 无连接协议")
	fmt.Println("- 不保证消息顺序")
	fmt.Println("- 不保证消息到达")
	fmt.Println("- 低延迟，高效率")
	fmt.Println("- 适合实时应用")
	fmt.Println()
	fmt.Println("功能特性:")
	fmt.Println("- 多客户端支持")
	fmt.Println("- 消息广播")
	fmt.Println("- 客户端管理")
	fmt.Println("- 超时处理")
	fmt.Println()
	fmt.Println("要运行UDP服务器:")
	fmt.Println("  server := NewUDPServer(\":8081\")")
	fmt.Println("  go server.Start()")
	fmt.Println()
	fmt.Println("要使用UDP客户端:")
	fmt.Println("  client := NewSimpleUDPClient(\"localhost:8081\")")
	fmt.Println("  client.Connect()")
	fmt.Println("  response, _ := client.SendAndReceive(\"ping\")")
	fmt.Println()

	// 演示API使用
	fmt.Println("🔹 API使用演示:")

	// 创建服务器
	server := NewUDPServer(":0") // 使用随机端口
	fmt.Println("创建UDP服务器成功")
	fmt.Printf("当前客户端数量: %d\n", server.GetClientCount())

	// 创建客户端
	_ = NewSimpleUDPClient("localhost:8081")
	fmt.Println("创建UDP客户端成功")

	fmt.Println("\n💡 提示: UDP是无连接协议，适合实时通信")
	fmt.Println("💡 提示: 可以使用nc命令测试: nc -u localhost 8081")
	fmt.Println("💡 提示: 查看测试文件了解完整使用示例")
}
