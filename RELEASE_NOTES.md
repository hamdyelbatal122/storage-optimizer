# Release Notes - Storage Optimizer v1.1.0

**Release Date:** April 29, 2026  
**Status:** Production Ready  
**License:** MIT

---

## 🎉 Welcome to Storage Optimizer v1.1.0

This is the **first public release** of Storage Optimizer - a professional-grade disk storage analysis and optimization tool built with Go.

### What is Storage Optimizer?

A powerful command-line tool that helps you:
- ✅ Find and analyze duplicate files
- ✅ Discover large files consuming storage
- ✅ Safely clean up storage with backup options
- ✅ Export analysis in multiple formats (Table, JSON, CSV)
- ✅ Filter and exclude files intelligently

---

## 🚀 Key Features

### 1. **Duplicate File Detection**
Finds all duplicate files using SHA256 hashing for accuracy
```bash
./storage-optimizer duplicates /path --extensions .mp4,.zip
```

### 2. **Large File Discovery**
Identifies the largest files consuming your storage
```bash
./storage-optimizer large /path --limit 100 --min-size 104857600
```

### 3. **Comprehensive Analysis**
Get detailed storage breakdown by file type
```bash
./storage-optimizer analyze /path --exclude node_modules,vendor
```

### 4. **Safe Cleanup**
Safely remove duplicates with optional backup
```bash
./storage-optimizer cleanup /path --dry-run=false --backup=true
```

### 5. **Smart Filtering**
Filter by:
- File extension (`--extensions .go,.js,.txt`)
- Minimum size (`--min-size 1000000`)
- File age (`--max-age 365`)
- Exclude patterns (`--exclude cache,tmp,vendor`)

### 6. **Multiple Export Formats**
Export results as:
- **Table** - Human-readable console output
- **JSON** - Machine-readable for automation
- **CSV** - Spreadsheet-compatible format
```bash
./storage-optimizer analyze /path -f csv -o report.csv
```

---

## 📦 Installation

### Requirements
- Go 1.21 or later
- Linux, macOS, or Windows

### Build from Source
```bash
git clone https://github.com/hamdyelbatal122/storage-optimizer.git
cd storage-optimizer
go mod download
go build -o storage-optimizer main.go
```

### Using Make
```bash
make build
make test
make install
```

---

## 🎯 Quick Start

### 1. Analyze Your Directory
```bash
./storage-optimizer analyze /home/user
```

### 2. Find Duplicates
```bash
./storage-optimizer duplicates /home/user -f json -o duplicates.json
```

### 3. Find Large Files
```bash
./storage-optimizer large /home/user --limit 50
```

### 4. Safe Cleanup (Preview First)
```bash
# Preview mode (no deletion)
./storage-optimizer cleanup /home/user --dry-run=true

# Execute cleanup
./storage-optimizer cleanup /home/user --dry-run=false --backup=true
```

---

## 💡 Real-World Examples

### Free Up Space in Downloads
```bash
./storage-optimizer duplicates ~/Downloads --min-size 1048576
./storage-optimizer cleanup ~/Downloads --dry-run=false --backup=true
```

### Analyze Project Excluding Dependencies
```bash
./storage-optimizer analyze ./myproject --exclude node_modules,vendor,.git -f csv -o analysis.csv
```

### Find Large Video Files
```bash
./storage-optimizer large ~/Videos --extensions .mp4,.mkv,.avi --limit 20
```

### Track Files Modified in Last Year
```bash
./storage-optimizer analyze . --max-age 365 -f json
```

### Clean Specific File Types
```bash
./storage-optimizer cleanup . --extensions .tmp,.log --dry-run=false
```

---

## 🔍 Command Reference

### Global Commands
- `analyze [path]` - Analyze storage usage
- `duplicates [path]` - Find duplicate files
- `large [path]` - Find large files
- `cleanup [path]` - Safe cleanup with backup
- `help` - Show help information

### Common Flags
- `-f, --format` - Output format: `table`, `json`, `csv`
- `-o, --output` - Save output to file
- `--exclude` - Exclude patterns (comma-separated)
- `--extensions` - Include only these file types
- `--min-size` - Minimum file size in bytes
- `--max-age` - Maximum file age in days

### Cleanup Flags
- `--dry-run` - Preview without deletion (default: true)
- `--backup` - Create backup before deletion
- `--backup-dir` - Custom backup directory

---

## 🛠️ Development

### Build & Test
```bash
make all          # Full build and test
make build        # Build binary
make test         # Run tests
make clean        # Clean artifacts
make fmt          # Format code
make lint         # Lint code
```

### Directory Structure
```
storage-optimizer/
├── main.go                    # Entry point
├── cmd/                       # CLI commands
│   ├── analyze.go
│   ├── duplicates.go
│   ├── large.go
│   ├── cleanup.go
│   └── root.go
├── internal/analyzer/         # Core logic
│   ├── analyzer.go
│   └── report.go
├── Makefile
├── README.md
├── CHANGELOG.md
└── LICENSE
```

---

## 🔐 Security & Privacy

✅ **Local Analysis Only** - No data sent to internet  
✅ **SHA256 Hashing** - Cryptographically secure duplicate detection  
✅ **No Root Required** - Runs as regular user  
✅ **Safe Deletion** - Dry-run preview mode by default  
✅ **Optional Backups** - Keep copies before deletion  

---

## 📊 Performance

- Process **thousands of files** efficiently
- **Optimized memory usage** with streaming analysis
- **Real-time progress** feedback
- Scales to **large directories** (100GB+)

---

## 🤝 Contributing

Contributions welcome! Please:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📝 License

This project is licensed under the MIT License - see [LICENSE](LICENSE) file for details.

---

## 🎓 Use Cases

### DevOps & System Administration
- Monitor server storage usage
- Identify space hogs in large filesystems
- Automated storage cleanup scripts

### Development Teams
- Clean project directories from build artifacts
- Identify large dependencies
- Manage backup storage

### Personal Use
- Free up disk space
- Find and remove duplicates
- Analyze storage distribution

---

## 🙏 Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Color](https://github.com/fatih/color) - Terminal colors
- [Progressbar](https://github.com/schollz/progressbar) - Progress visualization

---

## 📞 Support

For issues, feature requests, or questions:
- GitHub Issues: [storage-optimizer/issues](https://github.com/hamdyelbatal122/storage-optimizer/issues)
- Check [README.md](README.md) for full documentation

---

## 🚀 What's Next?

Planned features for future releases:
- [ ] Watch mode for continuous monitoring
- [ ] Age-based automatic cleanup
- [ ] Performance metrics and benchmarks
- [ ] Web UI dashboard
- [ ] Scheduled cleanup tasks
- [ ] Network disk support
- [ ] Docker/Kubernetes integration

---

**Enjoy Storage Optimizer!** 🎉

Questions? Check the [README.md](README.md) and [CHANGELOG.md](CHANGELOG.md) for more information.
