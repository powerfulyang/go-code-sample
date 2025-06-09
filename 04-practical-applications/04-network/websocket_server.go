package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// 注意：在实际项目中，你需要安装WebSocket库
// go get github.com/gorilla/websocket
// 这里我们提供一个简化的WebSocket实现概念

// WebSocketServer WebSocket服务器
type WebSocketServer struct {
	clients    map[string]*WebSocketClient
	register   chan *WebSocketClient
	unregister chan *WebSocketClient
	broadcast  chan []byte
	mutex      sync.RWMutex
	running    bool
}

// WebSocketClient WebSocket客户端
type WebSocketClient struct {
	ID       string
	Username string
	Room     string
	Send     chan []byte
	LastSeen time.Time
}

// Message 消息结构
type Message struct {
	Type      string    `json:"type"`
	From      string    `json:"from"`
	To        string    `json:"to,omitempty"`
	Room      string    `json:"room,omitempty"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// NewWebSocketServer 创建WebSocket服务器
func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients:    make(map[string]*WebSocketClient),
		register:   make(chan *WebSocketClient),
		unregister: make(chan *WebSocketClient),
		broadcast:  make(chan []byte),
		running:    false,
	}
}

// Start 启动WebSocket服务器
func (s *WebSocketServer) Start() {
	s.running = true

	// 启动消息处理协程
	go s.handleMessages()

	fmt.Println("🚀 WebSocket服务器已启动")
}

// Stop 停止WebSocket服务器
func (s *WebSocketServer) Stop() {
	s.running = false

	// 关闭所有客户端
	s.mutex.Lock()
	for _, client := range s.clients {
		close(client.Send)
	}
	s.clients = make(map[string]*WebSocketClient)
	s.mutex.Unlock()

	fmt.Println("🛑 WebSocket服务器已停止")
}

// handleMessages 处理消息
func (s *WebSocketServer) handleMessages() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for s.running {
		select {
		case client := <-s.register:
			s.registerClient(client)

		case client := <-s.unregister:
			s.unregisterClient(client)

		case message := <-s.broadcast:
			s.broadcastMessage(message)

		case <-ticker.C:
			s.cleanupClients()
		}
	}
}

// registerClient 注册客户端
func (s *WebSocketServer) registerClient(client *WebSocketClient) {
	s.mutex.Lock()
	s.clients[client.ID] = client
	s.mutex.Unlock()

	fmt.Printf("✅ 客户端连接: %s (用户: %s, 房间: %s)\n",
		client.ID, client.Username, client.Room)

	// 发送欢迎消息
	welcome := Message{
		Type:      "system",
		From:      "server",
		Content:   fmt.Sprintf("欢迎 %s 加入房间 %s", client.Username, client.Room),
		Timestamp: time.Now(),
	}

	if data, err := json.Marshal(welcome); err == nil {
		s.sendToRoom(client.Room, data, client.ID)
	}

	// 发送在线用户列表
	s.sendUserList(client.Room)
}

// unregisterClient 注销客户端
func (s *WebSocketServer) unregisterClient(client *WebSocketClient) {
	s.mutex.Lock()
	if _, ok := s.clients[client.ID]; ok {
		delete(s.clients, client.ID)
		close(client.Send)
	}
	s.mutex.Unlock()

	fmt.Printf("📤 客户端断开: %s (用户: %s)\n", client.ID, client.Username)

	// 发送离开消息
	leave := Message{
		Type:      "system",
		From:      "server",
		Content:   fmt.Sprintf("%s 离开了房间", client.Username),
		Timestamp: time.Now(),
	}

	if data, err := json.Marshal(leave); err == nil {
		s.sendToRoom(client.Room, data, "")
	}

	// 更新在线用户列表
	s.sendUserList(client.Room)
}

// broadcastMessage 广播消息
func (s *WebSocketServer) broadcastMessage(data []byte) {
	var msg Message
	if err := json.Unmarshal(data, &msg); err != nil {
		return
	}

	switch msg.Type {
	case "chat":
		// 聊天消息
		s.sendToRoom(msg.Room, data, msg.From)

	case "private":
		// 私聊消息
		s.sendToUser(msg.To, data)

	case "typing":
		// 输入状态
		s.sendToRoom(msg.Room, data, msg.From)

	default:
		fmt.Printf("未知消息类型: %s\n", msg.Type)
	}
}

// sendToRoom 发送消息到房间
func (s *WebSocketServer) sendToRoom(room string, data []byte, excludeID string) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, client := range s.clients {
		if client.Room == room && client.ID != excludeID {
			select {
			case client.Send <- data:
			default:
				// 客户端发送缓冲区满，关闭连接
				close(client.Send)
				delete(s.clients, client.ID)
			}
		}
	}
}

// sendToUser 发送消息给特定用户
func (s *WebSocketServer) sendToUser(username string, data []byte) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, client := range s.clients {
		if client.Username == username {
			select {
			case client.Send <- data:
			default:
				close(client.Send)
				delete(s.clients, client.ID)
			}
			break
		}
	}
}

// sendUserList 发送用户列表
func (s *WebSocketServer) sendUserList(room string) {
	s.mutex.RLock()
	var users []string
	for _, client := range s.clients {
		if client.Room == room {
			users = append(users, client.Username)
		}
	}
	s.mutex.RUnlock()

	userList := Message{
		Type:      "userlist",
		From:      "server",
		Content:   fmt.Sprintf("%v", users),
		Timestamp: time.Now(),
	}

	if data, err := json.Marshal(userList); err == nil {
		s.sendToRoom(room, data, "")
	}
}

// cleanupClients 清理不活跃的客户端
func (s *WebSocketServer) cleanupClients() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()
	for id, client := range s.clients {
		if now.Sub(client.LastSeen) > 5*time.Minute {
			fmt.Printf("🧹 清理不活跃客户端: %s\n", id)
			close(client.Send)
			delete(s.clients, id)
		}
	}
}

// GetStats 获取服务器统计信息
func (s *WebSocketServer) GetStats() map[string]interface{} {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	rooms := make(map[string]int)
	for _, client := range s.clients {
		rooms[client.Room]++
	}

	return map[string]interface{}{
		"total_clients": len(s.clients),
		"rooms":         rooms,
		"uptime":        time.Now().Format("2006-01-02 15:04:05"),
	}
}

// ChatRoom 聊天室
type ChatRoom struct {
	Name        string
	Description string
	MaxUsers    int
	Created     time.Time
}

// ChatRoomManager 聊天室管理器
type ChatRoomManager struct {
	rooms map[string]*ChatRoom
	mutex sync.RWMutex
}

// NewChatRoomManager 创建聊天室管理器
func NewChatRoomManager() *ChatRoomManager {
	manager := &ChatRoomManager{
		rooms: make(map[string]*ChatRoom),
	}

	// 创建默认房间
	manager.CreateRoom("general", "通用聊天室", 100)
	manager.CreateRoom("tech", "技术讨论", 50)
	manager.CreateRoom("random", "随机话题", 30)

	return manager
}

// CreateRoom 创建房间
func (m *ChatRoomManager) CreateRoom(name, description string, maxUsers int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.rooms[name] = &ChatRoom{
		Name:        name,
		Description: description,
		MaxUsers:    maxUsers,
		Created:     time.Now(),
	}
}

// GetRoom 获取房间
func (m *ChatRoomManager) GetRoom(name string) (*ChatRoom, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	room, exists := m.rooms[name]
	return room, exists
}

// ListRooms 列出所有房间
func (m *ChatRoomManager) ListRooms() []*ChatRoom {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var rooms []*ChatRoom
	for _, room := range m.rooms {
		rooms = append(rooms, room)
	}

	return rooms
}

// DeleteRoom 删除房间
func (m *ChatRoomManager) DeleteRoom(name string) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.rooms[name]; exists {
		delete(m.rooms, name)
		return true
	}

	return false
}

// WebSocketHandler HTTP处理器（模拟实现）
func WebSocketHandler(server *WebSocketServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 在实际实现中，这里会进行WebSocket握手
		// 并创建WebSocket连接

		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>WebSocket聊天室</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        #messages { border: 1px solid #ccc; height: 300px; overflow-y: scroll; padding: 10px; }
        #input { width: 70%%; }
        #send { width: 20%%; }
        .message { margin: 5px 0; }
        .system { color: #666; font-style: italic; }
        .user { color: #333; }
        .private { color: #0066cc; }
    </style>
</head>
<body>
    <h1>WebSocket聊天室演示</h1>
    <div id="messages"></div>
    <br>
    <input type="text" id="input" placeholder="输入消息..." />
    <button id="send">发送</button>
    
    <script>
        // 在实际实现中，这里会创建WebSocket连接
        // const ws = new WebSocket('ws://localhost:8080/ws');
        
        document.getElementById('messages').innerHTML = 
            '<div class="system">这是WebSocket聊天室的演示页面</div>' +
            '<div class="system">在实际实现中需要使用WebSocket库</div>' +
            '<div class="user">用户1: 大家好！</div>' +
            '<div class="user">用户2: 你好！欢迎来到聊天室</div>' +
            '<div class="private">私聊: 这是一条私聊消息</div>';
    </script>
</body>
</html>
		`)
	}
}

// WebSocketExamples WebSocket示例
func WebSocketExamples() {
	fmt.Println("=== WebSocket实时通信示例 ===")
	fmt.Println()
	fmt.Println("这个示例演示了WebSocket服务器的实现")
	fmt.Println()
	fmt.Println("WebSocket特点:")
	fmt.Println("- 全双工通信")
	fmt.Println("- 实时性强")
	fmt.Println("- 低延迟")
	fmt.Println("- 支持二进制和文本")
	fmt.Println("- 基于HTTP升级")
	fmt.Println()
	fmt.Println("功能特性:")
	fmt.Println("- 多房间聊天")
	fmt.Println("- 私聊功能")
	fmt.Println("- 用户管理")
	fmt.Println("- 输入状态")
	fmt.Println("- 在线用户列表")
	fmt.Println()
	fmt.Println("要使用WebSocket服务器:")
	fmt.Println("  server := NewWebSocketServer()")
	fmt.Println("  server.Start()")
	fmt.Println("  http.HandleFunc(\"/ws\", WebSocketHandler(server))")
	fmt.Println("  http.ListenAndServe(\":8080\", nil)")
	fmt.Println()

	// 演示API使用
	fmt.Println("🔹 API使用演示:")

	// 创建服务器
	server := NewWebSocketServer()
	server.Start()

	// 创建聊天室管理器
	roomManager := NewChatRoomManager()
	rooms := roomManager.ListRooms()

	fmt.Printf("WebSocket服务器已启动\n")
	fmt.Printf("可用聊天室数量: %d\n", len(rooms))
	for _, room := range rooms {
		fmt.Printf("  - %s: %s (最大用户: %d)\n",
			room.Name, room.Description, room.MaxUsers)
	}

	// 获取统计信息
	stats := server.GetStats()
	fmt.Printf("当前连接数: %v\n", stats["total_clients"])

	server.Stop()

	fmt.Println("\n💡 提示: 实际使用需要安装gorilla/websocket库")
	fmt.Println("💡 提示: go get github.com/gorilla/websocket")
	fmt.Println("💡 提示: 查看测试文件了解完整实现")
}
