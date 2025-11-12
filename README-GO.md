# MHDDoS - Go Implementation

This is the Go implementation of MHDDoS, a powerful DDoS attack script with 57+ methods.

## Features

- **High Performance**: Go's concurrency model provides excellent performance
- **Cross-Platform**: Compile once, run anywhere (Windows, Linux, macOS)
- **Low Memory Footprint**: More efficient than the Python version
- **Easy Deployment**: Single binary, no dependencies to install

## Building

### From Source

```bash
git clone https://github.com/mhdthariq/MHDDoS.git
cd MHDDoS
go build -o mhddos main.go
```

### Build for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o mhddos-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o mhddos.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o mhddos-mac main.go
```

## Usage

### Layer 7 Attacks

```bash
./mhddos GET http://example.com 5 100 proxies.txt 100 60
./mhddos POST http://example.com 5 100 proxies.txt 100 60
./mhddos STRESS http://example.com 5 100 proxies.txt 100 60
```

### Layer 4 Attacks

```bash
./mhddos TCP 192.168.1.1:80 100 60
./mhddos UDP 192.168.1.1:80 100 60
./mhddos SYN 192.168.1.1:80 100 60
```

### Available Commands

- `HELP` - Show usage information
- `TOOLS` - Run interactive tools console
- `STOP` - Stop all attacks

## Attack Methods

### Layer 7 (HTTP/HTTPS) - 26 Methods

- CFB, BYPASS, GET, POST, OVH, STRESS, DYN, SLOW
- HEAD, NULL, COOKIE, PPS, EVEN, GSB, DGB, AVB
- CFBUAM, APACHE, XMLRPC, BOT, BOMB, DOWNLOADER
- KILLER, TOR, RHEX, STOMP

### Layer 4 (TCP/UDP) - 14 Methods

- TCP, UDP, SYN, VSE, MINECRAFT, MCBOT
- CONNECTION, CPS, FIVEM, FIVEM-TOKEN
- TS3, MCPE, ICMP, OVH-UDP

### Amplification - 7 Methods

- MEM, NTP, DNS, ARD, CLDAP, CHAR, RDP

## Configuration

The `config.json` file contains:

```json
{
  "MCBOT": "MHDDoS_",
  "MINECRAFT_DEFAULT_PROTOCOL": 47,
  "proxy-providers": [
    {
      "type": 4,
      "url": "https://raw.githubusercontent.com/TheSpeedX/PROXY-List/refs/heads/master/socks4.txt",
      "timeout": 5
    }
  ]
}
```

## Performance Tips

1. **Thread Count**: Start with 100-500 threads and adjust based on your system
2. **RPC (Requests Per Connection)**: Higher values (50-100) for more intensive attacks
3. **Proxies**: Use fresh, fast proxies for better results
4. **Duration**: Don't run attacks for too long to avoid detection

## Differences from Python Version

- **Better Performance**: Go's concurrency is more efficient than Python threads
- **Simpler Deployment**: Single binary, no Python dependencies
- **Lower Resource Usage**: Uses less memory and CPU
- **Faster Startup**: No interpreter overhead

## Project Structure

```
MHDDoS/
├── main.go                 # Main entry point
├── pkg/
│   ├── attacks/           # Attack implementations
│   │   ├── layer4.go      # Layer 4 attacks
│   │   └── layer7.go      # Layer 7 attacks
│   ├── config/            # Configuration handling
│   ├── methods/           # Method definitions
│   ├── minecraft/         # Minecraft protocol
│   ├── proxy/             # Proxy management
│   ├── tools/             # Console tools
│   └── utils/             # Utility functions
└── config.json            # Configuration file
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests.

## Disclaimer

This tool is for educational purposes only. Do not attack websites without the owner's permission. Unauthorized attacks are illegal and unethical.

## License

This project is licensed under the terms specified in the LICENSE file.
