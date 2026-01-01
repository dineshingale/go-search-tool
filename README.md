# 🔍 Go Search Tool

A high-performance file search utility built in **Go**, featuring both a blazing fast CLI and a modern React Dashboard.

## 🚀 Features
- **Ultra-fast Search**: Powered by Go's efficient file walking.
- **Dual Interface**: Use the terminal or a beautiful Web UI.
- **REST API**: Exposes search functionality via JSON API.
- **Smart Filtering**: Filter by path, filename, and extension.

## 🛠️ Installation & Setup

### 1. Backend (Go)
Ensure you have **Go 1.21+** installed.

```bash
# Clone the repository
git clone https://github.com/dineshingale/go-search-tool.git
cd go-search-tool

# Run the API Server
go run cmd/server/main.go
```
The server will start on `http://localhost:8080`.

### 2. Frontend (React + Vite)
Ensure you have **Node.js** installed.

```bash
# Open a new terminal
cd dashboard

# Install dependencies
npm install

# Start the dashboard
npm run dev
```
The UI will be available at `http://localhost:5173`.

## 📈 CLI Usage
If you prefer the command line, you can build the CLI tool:

```bash
go build -o findit cmd/findit/main.go
```

**Examples:**
```bash
# Search for PDFs
./findit -ext .pdf

# Search in a specific path
./findit -path C:/Users/Downloads -name invoice
```

## 🏗️ Architecture
- **cmd/**: Application entry points.
  - `findit/`: CLI tool entry point.
  - `server/`: REST API server entry point.
- **internal/engine/**: Core search logic. Uses `filepath.Walk` for efficient recursion.
- **dashboard/**: React + Vite frontend application.
- **Makefile**: Automates the build and run process.

## 🤝 Contributing
1. Fork the repo
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---
**License**: MIT
