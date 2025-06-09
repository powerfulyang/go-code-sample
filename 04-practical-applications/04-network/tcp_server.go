package network

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

// TCPServer TCP服务器
type TCPServer struct {
	address  string
	listener net.Listener
	clients  map[string]*Client
	mutex    sync.RWMutex
	running  bool
}

// Client 客户端连接
type Client struct {
	ID       string
	Conn     net.Conn
	Reader   *bufio.Reader
	Writer   *bufio.Writer
	LastSeen time.Time
}

// NewTCPServer 创建TCP服务器
func NewTCPServer(address string) *TCPServer {
	return &TCPServer{
		address: address,
		clients: make(map[string]*Client),
		running: false,
	}
}

// Start 启动服务器
func (s *TCPServer) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("启动TCP服务器失败: %v", err)
	}

	s.listener = listener
	s.running = true

	fmt.Printf("🚀 TCP服务器启动在 %s\n", s.address)

	// 启动客户端清理协程
	go s.cleanupClients()

	// 接受连接
	for s.running {
		conn, err := listener.Accept()
		if err != nil {
			if s.running {
				fmt.Printf("接受连接失败: %v\n", err)
			}
			continue
		}

		// 处理新连接
		go s.handleConnection(conn)
	}

	return nil
}

// Stop 停止服务器
func (s *TCPServer) Stop() error {
	s.running = false

	if s.listener != nil {
		s.listener.Close()
	}

	// 关闭所有客户端连接
	s.mutex.Lock()
	for _, client := range s.clients {
		client.Conn.Close()
	}
	s.clients = make(map[string]*Client)
	s.mutex.Unlock()

	fmt.Println("🛑 TCP服务器已停止")
	return nil
}

// handleConnection 处理客户端连接
func (s *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	// 创建客户端
	clientID := fmt.Sprintf("%s_%d", conn.RemoteAddr().String(), time.Now().Unix())
	client := &Client{
		ID:       clientID,
		Conn:     conn,
		Reader:   bufio.NewReader(conn),
		Writer:   bufio.NewWriter(conn),
		LastSeen: time.Now(),
	}

	// 注册客户端
	s.mutex.Lock()
	s.clients[clientID] = client
	s.mutex.Unlock()

	fmt.Printf("✅ 客户端连接: %s\n", clientID)

	// 发送欢迎消息
	s.sendMessage(client, "欢迎连接到TCP服务器! 输入 'help' 查看可用命令\n")

	// 处理客户端消息
	for {
		// 设置读取超时
		conn.SetReadDeadline(time.Now().Add(30 * time.Second))

		message, err := client.Reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("📤 客户端断开连接: %s\n", clientID)
			} else {
				fmt.Printf("读取消息失败: %v\n", err)
			}
			break
		}

		// 更新最后活跃时间
		client.LastSeen = time.Now()

		// 处理消息
		s.processMessage(client, strings.TrimSpace(message))
	}

	// 移除客户端
	s.mutex.Lock()
	delete(s.clients, clientID)
	s.mutex.Unlock()
}

// processMessage 处理客户端消息
func (s *TCPServer) processMessage(client *Client, message string) {
	fmt.Printf("📨 收到消息 [%s]: %s\n", client.ID, message)

	switch {
	case message == "help":
		help := `可用命令:
- help: 显示帮助信息
- time: 获取服务器时间
- echo <message>: 回显消息
- clients: 查看在线客户端
- broadcast <message>: 广播消息给所有客户端
- quit: 断开连接
`
		s.sendMessage(client, help)

	case message == "time":
		timeStr := fmt.Sprintf("服务器时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		s.sendMessage(client, timeStr)

	case strings.HasPrefix(message, "echo "):
		echoMsg := strings.TrimPrefix(message, "echo ")
		response := fmt.Sprintf("回显: %s\n", echoMsg)
		s.sendMessage(client, response)

	case message == "clients":
		s.mutex.RLock()
		clientList := fmt.Sprintf("在线客户端数量: %d\n", len(s.clients))
		for id, c := range s.clients {
			clientList += fmt.Sprintf("- %s (最后活跃: %s)\n",
				id, c.LastSeen.Format("15:04:05"))
		}
		s.mutex.RUnlock()
		s.sendMessage(client, clientList)

	case strings.HasPrefix(message, "broadcast "):
		broadcastMsg := strings.TrimPrefix(message, "broadcast ")
		s.broadcastMessage(client.ID, broadcastMsg)
		s.sendMessage(client, "消息已广播\n")

	case message == "quit":
		s.sendMessage(client, "再见!\n")
		client.Conn.Close()

	default:
		response := fmt.Sprintf("未知命令: %s (输入 'help' 查看可用命令)\n", message)
		s.sendMessage(client, response)
	}
}

// sendMessage 发送消息给客户端
func (s *TCPServer) sendMessage(client *Client, message string) {
	client.Writer.WriteString(message)
	client.Writer.Flush()
}

// broadcastMessage 广播消息给所有客户端
func (s *TCPServer) broadcastMessage(senderID, message string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	broadcastMsg := fmt.Sprintf("[广播 from %s]: %s\n", senderID, message)

	for id, client := range s.clients {
		if id != senderID { // 不发送给发送者自己
			s.sendMessage(client, broadcastMsg)
		}
	}
}

// cleanupClients 清理不活跃的客户端
func (s *TCPServer) cleanupClients() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for s.running {
		select {
		case <-ticker.C:
			s.mutex.Lock()
			now := time.Now()
			for id, client := range s.clients {
				// 如果客户端超过5分钟没有活动，则断开连接
				if now.Sub(client.LastSeen) > 5*time.Minute {
					fmt.Printf("🧹 清理不活跃客户端: %s\n", id)
					client.Conn.Close()
					delete(s.clients, id)
				}
			}
			s.mutex.Unlock()
		}
	}
}

// GetClientCount 获取客户端数量
func (s *TCPServer) GetClientCount() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.clients)
}

// TCPClient TCP客户端
type TCPClient struct {
	address string
	conn    net.Conn
	reader  *bufio.Reader
	writer  *bufio.Writer
}

// NewTCPClient 创建TCP客户端
func NewTCPClient(address string) *TCPClient {
	return &TCPClient{
		address: address,
	}
}

// Connect 连接到服务器
func (c *TCPClient) Connect() error {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return fmt.Errorf("连接服务器失败: %v", err)
	}

	c.conn = conn
	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)

	fmt.Printf("✅ 已连接到服务器: %s\n", c.address)
	return nil
}

// Disconnect 断开连接
func (c *TCPClient) Disconnect() error {
	if c.conn != nil {
		c.conn.Close()
		fmt.Println("📤 已断开连接")
	}
	return nil
}

// SendMessage 发送消息
func (c *TCPClient) SendMessage(message string) error {
	if c.writer == nil {
		return fmt.Errorf("未连接到服务器")
	}

	_, err := c.writer.WriteString(message + "\n")
	if err != nil {
		return err
	}

	return c.writer.Flush()
}

// ReadMessage 读取消息
func (c *TCPClient) ReadMessage() (string, error) {
	if c.reader == nil {
		return "", fmt.Errorf("未连接到服务器")
	}

	message, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(message), nil
}

// TCPExamples TCP网络编程示例
func TCPExamples() {
	fmt.Println("=== TCP网络编程示例 ===")
	fmt.Println()
	fmt.Println("这个示例演示了TCP服务器和客户端的实现")
	fmt.Println()
	fmt.Println("功能特性:")
	fmt.Println("- 多客户端并发连接")
	fmt.Println("- 消息广播")
	fmt.Println("- 客户端管理")
	fmt.Println("- 连接超时处理")
	fmt.Println("- 优雅关闭")
	fmt.Println()
	fmt.Println("要运行TCP服务器:")
	fmt.Println("  server := NewTCPServer(\":8080\")")
	fmt.Println("  go server.Start()")
	fmt.Println()
	fmt.Println("要连接TCP客户端:")
	fmt.Println("  client := NewTCPClient(\"localhost:8080\")")
	fmt.Println("  client.Connect()")
	fmt.Println("  client.SendMessage(\"hello\")")
	fmt.Println()
	fmt.Println("可以使用telnet测试:")
	fmt.Println("  telnet localhost 8080")
	fmt.Println()

	// 演示API使用
	fmt.Println("🔹 API使用演示:")

	// 创建服务器
	server := NewTCPServer(":0") // 使用随机端口

	// 模拟启动（实际使用中需要在goroutine中启动）
	fmt.Println("创建TCP服务器成功")
	fmt.Printf("当前客户端数量: %d\n", server.GetClientCount())

	// 创建客户端
	_ = NewTCPClient("localhost:8080")
	fmt.Println("创建TCP客户端成功")

	fmt.Println("\n💡 提示: 在实际使用中，服务器需要在单独的goroutine中启动")
	fmt.Println("💡 提示: 可以使用测试文件查看完整的使用示例")
}
