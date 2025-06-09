package network

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	// 使用随机端口进行测试
	server := NewTCPServer(":0")

	// 启动服务器
	go func() {
		if err := server.Start(); err != nil {
			t.Logf("服务器启动失败: %v", err)
		}
	}()

	// 等待服务器启动
	time.Sleep(100 * time.Millisecond)

	// 获取实际监听的地址
	if server.listener == nil {
		t.Fatal("服务器未启动")
	}
	serverAddr := server.listener.Addr().String()

	t.Run("ClientConnection", func(t *testing.T) {
		// 创建客户端连接
		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			t.Fatalf("连接服务器失败: %v", err)
		}
		defer conn.Close()

		// 发送消息
		_, err = conn.Write([]byte("help\n"))
		if err != nil {
			t.Fatalf("发送消息失败: %v", err)
		}

		// 读取响应
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			t.Fatalf("读取响应失败: %v", err)
		}

		response := string(buffer[:n])
		if len(response) == 0 {
			t.Error("应该收到帮助信息")
		}

		t.Log("TCP客户端连接测试通过")
	})

	t.Run("EchoCommand", func(t *testing.T) {
		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			t.Fatalf("连接服务器失败: %v", err)
		}
		defer conn.Close()

		// 发送echo命令
		testMessage := "Hello TCP Server"
		_, err = conn.Write([]byte(fmt.Sprintf("echo %s\n", testMessage)))
		if err != nil {
			t.Fatalf("发送echo命令失败: %v", err)
		}

		// 先读取欢迎消息
		buffer := make([]byte, 1024)
		conn.Read(buffer) // 读取并丢弃欢迎消息

		// 读取echo响应
		n, err := conn.Read(buffer)
		if err != nil {
			t.Fatalf("读取echo响应失败: %v", err)
		}

		response := strings.TrimSpace(string(buffer[:n]))
		expectedResponse := fmt.Sprintf("回显: %s", testMessage)
		if response != expectedResponse {
			t.Errorf("Echo响应不正确: 期望 %q, 实际 %q", expectedResponse, response)
		}

		t.Log("TCP Echo命令测试通过")
	})

	t.Run("ClientCount", func(t *testing.T) {
		// 创建多个连接
		var conns []net.Conn
		for i := 0; i < 3; i++ {
			conn, err := net.Dial("tcp", serverAddr)
			if err != nil {
				t.Fatalf("创建连接%d失败: %v", i, err)
			}
			conns = append(conns, conn)
		}

		// 等待连接注册
		time.Sleep(50 * time.Millisecond)

		// 检查客户端数量
		if count := server.GetClientCount(); count < 3 {
			t.Errorf("客户端数量不正确: 期望至少3, 实际 %d", count)
		}

		// 关闭连接
		for _, conn := range conns {
			conn.Close()
		}

		t.Log("TCP客户端计数测试通过")
	})

	// 停止服务器
	server.Stop()
}

func TestUDPServer(t *testing.T) {
	// 使用随机端口进行测试
	server := NewUDPServer(":0")

	// 启动服务器
	go func() {
		if err := server.Start(); err != nil {
			t.Logf("UDP服务器启动失败: %v", err)
		}
	}()

	// 等待服务器启动
	time.Sleep(100 * time.Millisecond)

	// 获取实际监听的地址
	if server.conn == nil {
		t.Fatal("UDP服务器未启动")
	}
	serverAddr := server.conn.LocalAddr().String()

	t.Run("PingPong", func(t *testing.T) {
		// 创建UDP客户端
		client := NewSimpleUDPClient(serverAddr)
		err := client.Connect()
		if err != nil {
			t.Fatalf("连接UDP服务器失败: %v", err)
		}
		defer client.Disconnect()

		// 发送ping消息
		response, err := client.SendAndReceive("ping")
		if err != nil {
			t.Fatalf("发送ping失败: %v", err)
		}

		if response != "pong" {
			t.Errorf("Ping响应不正确: 期望 'pong', 实际 '%s'", response)
		}

		t.Log("UDP Ping-Pong测试通过")
	})

	t.Run("EchoCommand", func(t *testing.T) {
		client := NewSimpleUDPClient(serverAddr)
		err := client.Connect()
		if err != nil {
			t.Fatalf("连接UDP服务器失败: %v", err)
		}
		defer client.Disconnect()

		// 发送echo命令
		testMessage := "Hello UDP Server"
		response, err := client.SendAndReceive(fmt.Sprintf("echo %s", testMessage))
		if err != nil {
			t.Fatalf("发送echo命令失败: %v", err)
		}

		expectedResponse := fmt.Sprintf("回显: %s", testMessage)
		if response != expectedResponse {
			t.Errorf("Echo响应不正确: 期望 %q, 实际 %q", expectedResponse, response)
		}

		t.Log("UDP Echo命令测试通过")
	})

	t.Run("ClientTracking", func(t *testing.T) {
		// 创建多个UDP客户端
		var clients []*SimpleUDPClient
		for i := 0; i < 3; i++ {
			client := NewSimpleUDPClient(serverAddr)
			err := client.Connect()
			if err != nil {
				t.Fatalf("创建UDP客户端%d失败: %v", i, err)
			}
			clients = append(clients, client)

			// 发送消息以注册客户端
			_, err = client.SendAndReceive("ping")
			if err != nil {
				t.Fatalf("注册UDP客户端%d失败: %v", i, err)
			}
		}

		// 等待客户端注册
		time.Sleep(50 * time.Millisecond)

		// 检查客户端数量
		if count := server.GetClientCount(); count < 3 {
			t.Errorf("UDP客户端数量不正确: 期望至少3, 实际 %d", count)
		}

		// 关闭客户端
		for _, client := range clients {
			client.Disconnect()
		}

		t.Log("UDP客户端跟踪测试通过")
	})

	// 停止服务器
	server.Stop()
}

func TestWebSocketServer(t *testing.T) {
	server := NewWebSocketServer()
	server.Start()
	defer server.Stop()

	t.Run("ServerStartStop", func(t *testing.T) {
		// 测试服务器启动和停止
		if !server.running {
			t.Error("服务器应该在运行状态")
		}

		stats := server.GetStats()
		if stats["total_clients"] != 0 {
			t.Errorf("初始客户端数量应该为0, 实际 %v", stats["total_clients"])
		}

		t.Log("WebSocket服务器启动停止测试通过")
	})

	t.Run("ClientRegistration", func(t *testing.T) {
		// 模拟客户端注册
		client := &WebSocketClient{
			ID:       "test_client_1",
			Username: "testuser",
			Room:     "general",
			Send:     make(chan []byte, 256),
			LastSeen: time.Now(),
		}

		// 注册客户端
		server.register <- client

		// 等待注册完成
		time.Sleep(50 * time.Millisecond)

		stats := server.GetStats()
		if stats["total_clients"] != 1 {
			t.Errorf("客户端数量应该为1, 实际 %v", stats["total_clients"])
		}

		// 注销客户端
		server.unregister <- client

		// 等待注销完成
		time.Sleep(50 * time.Millisecond)

		stats = server.GetStats()
		if stats["total_clients"] != 0 {
			t.Errorf("注销后客户端数量应该为0, 实际 %v", stats["total_clients"])
		}

		t.Log("WebSocket客户端注册测试通过")
	})
}

func TestChatRoomManager(t *testing.T) {
	manager := NewChatRoomManager()

	t.Run("DefaultRooms", func(t *testing.T) {
		rooms := manager.ListRooms()
		if len(rooms) < 3 {
			t.Errorf("应该有至少3个默认房间, 实际 %d", len(rooms))
		}

		// 检查默认房间是否存在
		if _, exists := manager.GetRoom("general"); !exists {
			t.Error("应该存在general房间")
		}

		t.Log("默认聊天室测试通过")
	})

	t.Run("CreateRoom", func(t *testing.T) {
		roomName := "test_room"
		manager.CreateRoom(roomName, "测试房间", 10)

		room, exists := manager.GetRoom(roomName)
		if !exists {
			t.Error("创建的房间应该存在")
		}

		if room.Name != roomName {
			t.Errorf("房间名称不正确: 期望 %s, 实际 %s", roomName, room.Name)
		}

		if room.MaxUsers != 10 {
			t.Errorf("房间最大用户数不正确: 期望 10, 实际 %d", room.MaxUsers)
		}

		t.Log("创建聊天室测试通过")
	})

	t.Run("DeleteRoom", func(t *testing.T) {
		roomName := "temp_room"
		manager.CreateRoom(roomName, "临时房间", 5)

		// 确认房间存在
		if _, exists := manager.GetRoom(roomName); !exists {
			t.Fatal("临时房间应该存在")
		}

		// 删除房间
		if !manager.DeleteRoom(roomName) {
			t.Error("删除房间应该成功")
		}

		// 确认房间已删除
		if _, exists := manager.GetRoom(roomName); exists {
			t.Error("删除后房间不应该存在")
		}

		t.Log("删除聊天室测试通过")
	})
}

// 基准测试
func BenchmarkTCPServerConnection(b *testing.B) {
	server := NewTCPServer(":0")
	go server.Start()
	defer server.Stop()

	// 等待服务器启动
	time.Sleep(100 * time.Millisecond)

	if server.listener == nil {
		b.Fatal("服务器未启动")
	}
	serverAddr := server.listener.Addr().String()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			conn, err := net.Dial("tcp", serverAddr)
			if err != nil {
				b.Fatalf("连接失败: %v", err)
			}
			conn.Write([]byte("ping\n"))
			buffer := make([]byte, 1024)
			conn.Read(buffer)
			conn.Close()
		}
	})
}

func BenchmarkUDPServerMessage(b *testing.B) {
	server := NewUDPServer(":0")
	go server.Start()
	defer server.Stop()

	// 等待服务器启动
	time.Sleep(100 * time.Millisecond)

	if server.conn == nil {
		b.Fatal("UDP服务器未启动")
	}
	serverAddr := server.conn.LocalAddr().String()

	client := NewSimpleUDPClient(serverAddr)
	client.Connect()
	defer client.Disconnect()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.SendAndReceive("ping")
	}
}

func BenchmarkWebSocketClientRegistration(b *testing.B) {
	server := NewWebSocketServer()
	server.Start()
	defer server.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client := &WebSocketClient{
			ID:       fmt.Sprintf("bench_client_%d", i),
			Username: fmt.Sprintf("user%d", i),
			Room:     "general",
			Send:     make(chan []byte, 256),
			LastSeen: time.Now(),
		}

		server.register <- client
		server.unregister <- client
	}
}
