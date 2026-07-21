# RVPA

```
    ____ _    ______  ___
   / __ \ |  / / __ \/   |
  / /_/ / | / / /_/ / /| |
 / _, _/| |/ / ____/ ___ |
/_/ |_| |___/_/   /_/  |_|
```

Fast Web Scanning & Security Audit Framework written in Go.

## Features

- ✅ Website fingerprinting
- ✅ Technology detection
- ✅ HTTP header analysis
- ✅ Security header inspection
- ✅ TLS/SSL information
- ✅ JavaScript discovery
- ✅ robots.txt discovery
- ✅ sitemap.xml discovery
- ✅ API endpoint discovery
- ✅ Directory enumeration
- ✅ Subdomain enumeration
- ✅ WAF detection
- ✅ Redirect analysis
- ✅ Cookie inspection
- ✅ CORS inspection
- ✅ CSP inspection
- ✅ DNS information
- ✅ WHOIS lookup
- ✅ HTML report
- ✅ JSON report
- ✅ Multi-threaded scanning

## Installation

### Linux / macOS / Windows
```bash
go install github.com/mrrvpa-gemini/rvpa/cmd/rvpa@latest
```

### Termux (Android)
```bash
pkg install git golang
git clone https://github.com/mrrvpa-gemini/rvpa
cd rvpa
go build -o rvpa ./cmd/rvpa
```

## Quick Start

### Basic scan
```bash
rvpa -u https://example.com
```

### Deep scan (slow but thorough)
```bash
rvpa -u https://example.com -deep
```

### Save JSON report
```bash
rvpa -u https://example.com -json -o report.json
```

### Generate HTML report
```bash
rvpa -u https://example.com -html -o report.html
```

### Increase threads (faster scanning)
```bash
rvpa -u https://example.com -threads 100
```

### Scan multiple targets from file
```bash
rvpa -list targets.txt -json
```

### Check version
```bash
rvpa -version
```

## Usage

```
rvpa - Fast Web Scanning & Security Audit Framework

Usage:
  rvpa [flags]

Flags:
  -u, --url string              Target URL to scan
  -list string                  File containing list of targets (one per line)
  -deep                         Enable deep scanning (slower but comprehensive)
  -json                         Output results in JSON format
  -html                         Generate HTML report
  -o, --output string           Output file path
  -threads int                  Number of concurrent threads (default: 20)
  -timeout int                  Request timeout in seconds (default: 30)
  -proxy string                 HTTP proxy URL
  -ua, --user-agent string      Custom user agent
  -h, --help                    Display this help message
  -version                      Show version information
  -update                       Update to latest version
```

## Examples

### Example 1: Quick security audit
```bash
rvpa -u https://example.com -json -o audit.json
```

### Example 2: Deep scan with HTML report
```bash
rvpa -u https://example.com -deep -html -o deep_scan.html
```

### Example 3: Batch scanning
```bash
rvpa -list targets.txt -json -threads 50
```

### Example 4: Scan with proxy
```bash
rvpa -u https://example.com -proxy http://127.0.0.1:8080
```

## Requirements

- **Go 1.24+**
- Linux
- macOS
- Windows
- Android (Termux)

## Termux Installation

1. Install Termux from F-Droid
2. Install dependencies:
```bash
pkg update && pkg upgrade
pkg install git golang
```

3. Clone and build:
```bash
git clone https://github.com/mrrvpa-gemini/rvpa
cd rvpa
go build -o rvpa ./cmd/rvpa
```

4. Make it available globally:
```bash
mv rvpa $HOME/go/bin/
export PATH=$PATH:$HOME/go/bin
```

## Configuration

Configuration file: `~/.rvpa/config.yaml`

```yaml
# HTTP client settings
http:
  timeout: 30
  retries: 3
  user_agent: "RVPA/1.0"

# Scanning preferences
scan:
  threads: 20
  deep_scan: false
  timeout: 30

# Output preferences
output:
  json: false
  html: false
  format: text
```

## License

MIT License - See LICENSE file for details

## Author

mrrvpa-gemini - [@mrrvpa](https://github.com/mrrvpa-gemini)

## Disclaimer

This tool is designed for authorized security testing only. Unauthorized access to computer systems is illegal. Always get proper authorization before testing.
