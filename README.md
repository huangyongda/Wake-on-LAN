# 🌐 网络唤醒工具（Wake-on-LAN）使用说明书

> **工具名称**：`wol`  
> **功能**：通过发送 Magic Packet 唤醒局域网内的设备  
> **支持平台**：Windows / macOS（M1 & Intel）/ Linux  
> **特点**：绿色单文件、无需安装、配置灵活

---

## 📁 文件组成

每个版本包含以下两个文件：
wol_xxx/
├── wol （或 wol.exe） # 可执行程序
└── config.json # 配置文件（必须与程序同目录）

## 🛠️ `config.json` 配置文件说明

```json
{
  "macAddress": "04:7c:16:6e:18:67",
  "broadcast": "255.255.255.255"
}
```
macAddress
目标设备的 MAC 地址（物理地址），必须正确填写
broadcast
广播地址，推荐使用
255.255.255.255
或子网广播（如
192.168.1.255
）

各平台使用方法
✅ 1. Windows 版本（wol.exe）
🔧 环境要求
Windows 7 / 10 / 11（64位）
建议以管理员身份运行命令行（避免网络权限限制）
📦 文件名
wol.exe
🚀 使用步骤
将 wol.exe 和 config.json 放在同一文件夹中。
按下 Win + R，输入 cmd 打开命令提示符。
切换到程序所在目录：
cmd
cd C:\path\to\your\folder
运行程序：
cmd
wol.exe
📝 输出示例
2025/04/05 10:00:00 发送网络唤醒包 -> MAC: 04:7c:16:6e:18:67, 广播: 255.255.255.255:9
2025/04/05 10:00:00 ✅ Magic Packet 发送成功！设备应开始唤醒。
⚠️ 安全提示：若提示“无法识别的程序”，请右键 wol.exe → 属性 → 勾选“解除锁定”。 

✅ 2. macOS 版本
🔧 环境要求
macOS  
支持 Apple Silicon（M1/M2/M3）和 Intel 芯片
📦 文件名
Apple Silicon（M系列）：wol-macos-arm64
Intel 芯片：wol-macos-amd64
🚀 使用步骤
将可执行文件和 config.json 放在同一目录。
打开 终端（Terminal），进入目录：
bash
 
cd /Users/你的用户名/Downloads/wol-macos
赋予执行权限（首次运行）：
bash
 
chmod +x wol-macos-arm64
# 或
chmod +x wol-macos-amd64
运行程序：
bash ./wol-macos-arm64

✅ 3. Linux 版本
🔧 环境要求
Ubuntu / Debian / CentOS / Fedora 等主流发行版
架构：amd64（64位 Intel/AMD）
📦 文件名 wol-linux-amd64
🚀 使用步骤
将 wol-linux-amd64 和 config.json 放在同一目录。
打开终端，进入目录：
bash 
cd ~/wol-linux
赋予执行权限：
bash
 
chmod +x wol-linux-amd64
运行程序：
bash
 
./wol-linux-amd64
📝 输出示例

 
2025/04/05 10:00:00 发送网络唤醒包 -> MAC: 04:7c:16:6e:18:67, 广播: 255.255.255.255:9
2025/04/05 10:00:00 ✅ Magic Packet 发送成功！设备应开始唤醒。
💡 提示：某些系统需 root 权限发送广播包，可尝试： 

bash 
sudo ./wol-linux-amd64

🌐 网络要求
发送设备与目标设备必须在 同一局域网。
目标设备需在 BIOS/UEFI 和操作系统中开启 Wake-on-LAN。
路由器或交换机未屏蔽广播包（如 255.255.255.255）。