# TamperX - Verb Tampering Vulnerability Checker

A high-performance Go tool for testing HTTP method vulnerabilities with concurrent requests, proxy support, and advanced error handling.

## 🚀 Features

- **Concurrent Processing**: Configurable concurrency for simultaneous HTTP requests
- **Proxy Support**: Built-in proxy support with SSL verification disabled
- **Random User-Agent**: Randomize User-Agent headers to avoid detection
- **Custom Headers**: Add custom HTTP headers with `-H` flag
- **Advanced Error Handling**: Detailed error categorization and reporting
- **Multiple HTTP Methods**: Tests GET, HEAD, POST, PUT, DELETE, CONNECT, TRACE, and PATCH
- **Configurable Timeout**: Customizable request timeout
- **Clean Output**: Professional formatted output with status codes and content lengths
- **No Dependencies**: Uses only Go standard library

## 🛠️ Installation

### Quick Install (Recommended)
```bash
go install github.com/TheWation/GoTamperX/cmd/tamperx@latest
```

### Prerequisites
- Go 1.20 or later

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/yourusername/tamperx.git
cd tamperx
```

2. Build the application:
```bash
go build -o tamperx ./cmd/tamperx
```

3. (Optional) Install globally:
```bash
go install ./cmd/tamperx
```

## 📖 Usage

### Basic Usage

```bash
# Test a single URL
./tamperx -u https://example.com

# Test with custom timeout
./tamperx -u https://example.com -t 30

# Test with custom concurrency
./tamperx -u https://example.com -c 16

# Test through a proxy
./tamperx -u https://example.com -p http://proxy:8080

# Test with random User-Agent headers
./tamperx -u https://example.com --random-agent

# Test with custom headers
./tamperx -u https://example.com -H "Authorization: Bearer token123" -H "X-Custom-Header: value"
```

### Command Line Options

| Flag | Description | Default | Example |
|------|-------------|---------|---------|
| `-u` | Target URL to test | Required | `-u https://example.com` |
| `-t` | Timeout in seconds | 10 | `-t 30` |
| `-c` | Number of concurrent requests | 8 | `-c 16` |
| `-p` | Proxy URL | None | `-p http://proxy:8080` |
| `--random-agent` | Use random User-Agent for each request | false | `--random-agent` |
| `-H` | Add custom header (can be used multiple times) | None | `-H "Header: Value"` |

### Example Output

```
  __     __     ______     ______   __     ______     __   __    
 /\ \  _ \ \   /\  __ \   /\__  _\ /\ \   /\  __ \   /\ "-.\ \   
 \ \ \/ ".\ \  \ \  __ \  \/_/\ \/ \ \ \  \ \ \/\ \  \ \ \-.  \  
  \ \__/".~\_\  \ \_\ \_\    \ \_\  \ \_\  \ \_____\  \ \_\"\_ \ 
   \/_/   \/_/   \/_/\/_/     \/_/   \/_/   \/_____/   \/_/ \/_/ 
        TamperX V2.0: Verb Tampering Vulnerability Checker

[+] Target Url: https://example.com/admin/restricted
[+] Concurrency: 8
[+] Timeout: 10s

Method     Status     Content    
--------------------------------
GET        200        4416       
HEAD       200        0          
POST       405        178        
PUT        405        178        
DELETE     405        178        
CONNECT    405        150        
TRACE      405        150        
PATCH      405        178       
```

## 🔧 Advanced Features

### Concurrency Control

Control the number of simultaneous requests:

```bash
# Low concurrency for rate-limited targets
./tamperx -u https://example.com -c 2

# High concurrency for fast testing
./tamperx -u https://example.com -c 32
```

### Proxy Support

Test through various proxy types:

```bash
# HTTP proxy
./tamperx -u https://example.com -p http://proxy:8080

# HTTPS proxy
./tamperx -u https://example.com -p https://proxy:8080

# SOCKS proxy (if supported by Go)
./tamperx -u https://example.com -p socks5://proxy:1080

# Random User-Agent for stealth testing
./tamperx -u https://example.com --random-agent

# Combine multiple features
./tamperx -u https://example.com -p http://proxy:8080 --random-agent -c 16 -t 30
```

**Note**: SSL verification is disabled when using proxies for maximum compatibility.

### Random User-Agent

Randomize User-Agent headers to avoid detection and test different browser behaviors:

```bash
# Enable random User-Agent
./tamperx -u https://example.com --random-agent

# Random User-Agent with custom concurrency
./tamperx -u https://example.com --random-agent -c 4
```

**Available User-Agents**: Chrome, Firefox, Safari, Edge on Windows, macOS, and Linux.

### Custom Headers

Add custom HTTP headers to test specific scenarios:

```bash
# Single custom header
./tamperx -u https://example.com -H "Authorization: Bearer token123"

# Multiple custom headers
./tamperx -u https://example.com -H "X-API-Key: abc123" -H "Content-Type: application/json"

# Custom headers with other features
./tamperx -u https://example.com -H "X-Custom: value" --random-agent -c 4
```

**Common Use Cases**:
- Authentication testing
- API key validation
- Content-Type testing
- Custom business logic headers

### Error Handling

The tool provides detailed error information:

| Error Type | Description | Common Causes |
|------------|-------------|---------------|
| `domain not found` | DNS resolution failed | Invalid hostname |
| `connection refused` | Server refused connection | Port closed, firewall |
| `connection reset` | Connection reset by peer | Network issues |
| `timeout` | Request timed out | Slow server, network latency |
| `host not found` | Hostname resolution failed | Invalid hostname |

## 🏗️ Architecture

### Design Principles

- **Modular Structure**: Clean separation of concerns
- **Concurrent by Design**: Built for high-performance testing
- **Error Resilient**: Comprehensive error handling
- **Extensible**: Easy to add new features

### Package Organization

- **`cmd/tamperx`**: Main application entry point
- **`internal/httpclient`**: HTTP client implementation with concurrency
- **`pkg/errors`**: Error handling utilities

## 🚀 Performance

### Benchmarks

Compared to sequential testing, TamperX provides:

- **8x faster** with default concurrency (8)
- **16x faster** with high concurrency (16)
- **Linear scaling** with concurrency up to network limits

### Resource Usage

- **Memory**: ~2MB base usage
- **CPU**: Efficient goroutine management
- **Network**: Configurable concurrency prevents overwhelming targets

## 🔒 Security Considerations

- **SSL Verification**: Disabled when using proxies for compatibility
- **Rate Limiting**: Built-in concurrency control prevents DoS
- **User Agent**: Identifies as TamperX/2.0
- **Timeout Protection**: Prevents hanging requests

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

**For educational and authorized testing purposes only.**

- Do not use for illegal activities
- Always obtain proper authorization before testing
- Use at your own risk
- The authors are not responsible for any misuse

## 🙏 Acknowledgments

- Inspired by the original TamperX Python implementation
- Built with Go's excellent concurrency primitives
- Community feedback and contributions

---

**TamperX V2.0** - Made with ❤️ by [Wation](https://github.com/TheWation)