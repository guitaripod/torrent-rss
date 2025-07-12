# 🌊 Torrent RSS Downloader

[![Go Version](https://img.shields.io/badge/Go-1.22.5-00ADD8?style=flat-square&logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)
[![Maintenance](https://img.shields.io/badge/maintained%3F-yes-green.svg?style=flat-square)](https://github.com/guitaripod/torrent-rss/graphs/commit-activity)
[![Docker Support](https://img.shields.io/badge/Docker-ready-2496ED?style=flat-square&logo=docker)](https://www.docker.com/)
[![Platform Support](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey?style=flat-square)](https://github.com/guitaripod/torrent-rss)
[![Static Badge](https://img.shields.io/badge/RSS-feed-orange?style=flat-square&logo=rss)](https://github.com/guitaripod/torrent-rss)

![torrent-rss-demo](https://github.com/user-attachments/assets/af50f8f9-c780-4f67-b286-a680a38f16b3)



A sleek, automated torrent RSS feed monitor and downloader written in Go. This tool automatically checks your private tracker's RSS feed for new torrents matching your search terms and downloads them.

## ✨ Features

- 🔄 Automated RSS feed monitoring
- 🔍 Configurable search terms
- ⚡️ Fast and lightweight
- 🎯 File quality filters
- 🔐 Secure authentication handling
- 📁 Customizable download directory
- ⏰ Configurable check intervals
- 🐳 Docker support

## 📋 Prerequisites

- Support is only for **Anime, TV/x264**files (for now)
- Docker (recommended) or Go 1.22.5+
- Access to TorrentDay.com (invite-only website)
- RSS feed access on TorrentDay

## 🚀 Quick Start

### ⚡️ Docker Method (Recommended)

```bash
# Clone the repository
git clone https://github.com/username/torrent-rss.git

# Navigate to the project directory
cd torrent-rss

# Copy the example environment file
cp .env.example .env

# Edit your .env file with your credentials
nvim .env

# Create downloads directory
mkdir downloads

# Start the container
docker compose up -d
```

### 🔧 Traditional Method

```bash
# Install dependencies
go mod download

# Run the program
go run cmd/torrent-rss/main.go
```

## ⚙️ Configuration

1. Visit TorrentDay's RSS setup page at `https://www.torrentday.com/rss`
2. Generate your RSS feed and get the following information from the generated RSS URL:
   - User ID
   - RSS Token
   - Download Token
3. Edit your `.env` file with your details:

```env
# Required configuration
TD_BASE_URL=https://www.torrentday.com
TD_USER_ID=your_user_id_here
TD_TOKEN=your_download_token_here
TD_RSS_TOKEN=your_rss_token_here
TD_SEARCH_TERMS=Formula1,UFC

# Optional configuration
TD_CHECK_INTERVAL=0 */12 * 
TD_DOWNLOAD_PATH=/downloads  # Default path in Docker
```

### 🔮 Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `TD_BASE_URL` | TorrentDay base URL | Yes | https://www.torrentday.com |
| `TD_USER_ID` | Your user ID | Yes | - |
| `TD_TOKEN` | Download token | Yes | - |
| `TD_RSS_TOKEN` | RSS feed token | Yes | - |
| `TD_SEARCH_TERMS` | Search terms (comma-separated) | Yes | - |
| `TD_CHECK_INTERVAL` | Check interval (cron format) | No | `0 */12 * * *` |
| `TD_DOWNLOAD_PATH` | Download directory | No | `/downloads` |

## 🐳 Docker Configuration

The application comes with a pre-configured `compose.yml` file for easy deployment. The container:

- Automatically restarts unless stopped
- Mounts a local `downloads` directory
- Uses environment variables from `.env`
- Runs in a lightweight Alpine Linux container

### 📦 Container Management

```bash
# Start the container
docker compose up -d

# View logs
docker compose logs -f

# Stop the container
docker compose down

# Rebuild and restart
docker compose up -d --build
```

## 🔒 Security Notes

- Keep your `.env` file secure and never commit it to version control
- Your RSS feed URL contains private tokens - never share it
- The program stores sensitive data only in environment variables
- Docker containers provide isolation and security by default

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

This tool is for personal use only. Ensure you comply with your tracker's rules and regulations regarding automated downloads and RSS usage.
