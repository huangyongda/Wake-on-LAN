package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
)

// Config 结构体定义配置文件内容
type Config struct {
	MACAddress string `json:"macAddress"`
	Broadcast  string `json:"broadcast"`
}

// 从文件加载配置
func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// 创建 Magic Packet
func createMagicPacket(macAddr string) ([]byte, error) {
	mac, err := net.ParseMAC(macAddr)
	if err != nil {
		return nil, err
	}

	packet := make([]byte, 102)
	// 前 6 个字节是 0xFF
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}
	// 16 次重复 MAC
	for i := 0; i < 16; i++ {
		copy(packet[6+i*6:], mac)
	}
	return packet, nil
}

// 发送 Magic Packet
func sendMagicPacket(macAddr, broadcastAddr string, port int) error {
	udpAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(broadcastAddr, "9"))
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet, err := createMagicPacket(macAddr)
	if err != nil {
		return err
	}

	_, err = conn.Write(packet)
	return err
}

func main() {
	// 从 config.json 加载配置
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	log.Printf("发送网络唤醒包 -> MAC: %s, 广播: %s:%d\n", config.MACAddress, config.Broadcast, 9)
	err = sendMagicPacket(config.MACAddress, config.Broadcast, 9)
	if err != nil {
		log.Printf("发送失败: %v", err)
	} else {
		log.Println("✅ Magic Packet 发送成功！设备应开始唤醒。")
	}
}
