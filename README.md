# Storage Optimizer

A professional Go tool for analyzing and optimizing disk storage usage efficiently with advanced filtering and reporting capabilities.

## Features

- **Find Duplicate Files** using secure SHA256 hash-based detection
- **Comprehensive Analysis** of storage space usage and file distribution
- **Detect Large Files** that consume excessive storage
- **Safe Cleanup** with preview capability (Dry Run)
- **Optional Backup** before deletion operations
- **Detailed Reports** in table, JSON, and CSV formats
- **Smart Filtering** by extension, minimum size, and file age
- **Exclude Patterns** to skip directories and files
- **Cross-platform** support (Windows, macOS, Linux)

## Requirements

- Go 1.21 or later
- Linux, macOS, or Windows

## Installation

```bash
# Download dependencies
go mod download

# Build the binary
go build -o storage-optimizer main.go

# Or use make
make build
```

## Usage

### Analyze Directory

```bash
./storage-optimizer analyze /path/to/folder
```

Options:
- `-f, --format` - Output format: `table`, `json`, `csv` (default: table)
- `-o, --output` - Save output to file
- `--exclude` - Exclude patterns (comma-separated, e.g., node_modules,vendor,.git)
- `--extensions` - Include only specific file types (e.g., .go,.js,.txt)
- `--min-size` - Minimum file size in bytes
- `--max-age` - Maximum file age in days

**Examples:**
```bash
# Analyze with excluded directories
./storage-optimizer analyze . --exclude node_modules,vendor

# Analyze only Go and JS files
./storage-optimizer analyze . --extensions .go,.js

# Find files larger than 1MB
./storage-optimizer analyze . --min-size 1048576

# Export to CSV
./storage-optimizer analyze . -f csv -o report.csv
```

### Find Duplicate Files

```bash
./storage-optimizer duplicates /path/to/folder
```

Options:
- `-f, --format` - Output format: `table`, `json`, `csv`
- `-o, --output` - Save output to file
- `--exclude` - Exclude patterns
- `--extensions` - Include only specific file types
- `--min-size` - Minimum file size in bytes

**Examples:**
```bash
# Find duplicates, export to JSON
./storage-optimizer duplicates . -f json -o duplicates.json

# Find duplicate videos only
./storage-optimizer duplicates . --extensions .mp4,.mkv

# Find large duplicate files (>10MB)
./storage-optimizer duplicates . --min-size 10485760
```

### Find Large Files

```bash
./storage-optimizer large /path/to/folder
```

Options:
- `-l, --limit` - Number of files to display (default: 50)
- `-f, --format` - Output format: `table`, `json`, `csv`
- `-o, --output` - Save output to file
- `--min-size` - Minimum file size in bytes
- `--exclude` - Exclude patterns

**Examples:**
```bash
# Show top 100 largest files
./storage-optimizer large . --limit 100

# Find files >100MB, export to CSV
./storage-optimizer large . --min-size 104857600 -f csv -o large.csv

# Exclude cache directories
./storage-optimizer large . --exclude .cache,tmp,temp
```

### Safe Cleanup (Delete Duplicates)

```bash
./storage-optimizer cleanup /path/to/folder
```

Options:
- `--dry-run` - Preview without actual deletion (default: true)
- `--backup` - Create backup before deletion
- `--backup-dir` - Custom backup directory
- `--exclude` - Exclude patterns
- `--extensions` - Include only specific file types

**Examples:**
```bash
# Preview cleanup without deleting
./storage-optimizer cleanup . --dry-run=true

# Delete duplicates with backup
./storage-optimizer cleanup . --dry-run=false --backup=true

# Delete only duplicate .log files
./storage-optimizer cleanup . --dry-run=false --extensions .log
```

## Real-World Examples

### Free up space (remove duplicates)
```bash
./storage-optimizer duplicates ~/Downloads --min-size 1048576
./storage-optimizer cleanup ~/Downloads --dry-run=false --backup=true
```

### Analyze project excluding dependencies
```bash
./storage-optimizer analyze ./myproject --exclude node_modules,vendor,.git -f csv -o analysis.csv
```

### Find large video files
```bash
./storage-optimizer large ~/Videos --extensions .mp4,.mkv,.avi --limit 20
```

### Track old files
```bash
./storage-optimizer analyze . --max-age 365 -f json
```

### Clean specific file types
```bash
./storage-optimizer cleanup . --extensions .tmp,.log --dry-run=false
```

## Export Formats

### Table Format (Default)
Human-readable console output with colors and formatting.

### JSON Format
Machine-readable JSON for automation and integration:
```json
{
  "totalFiles": 100,
  "totalSize": 5242880,
  "filesByType": {
    ".go": 10,
    ".txt": 5
  },
  "sizeByType": {
    ".go": 102400
  }
}
```

### CSV Format
Spreadsheet-compatible format for Excel, Google Sheets, etc.:
```csv
File Type,Count,Total Size
.go,10,102400
.txt,5,51200
```

## Project Structure

```
storage-optimizer/
├── main.go                      # Entry point
├── cmd/                         # CLI commands
│   ├── root.go                  # Root command
│   ├── analyze.go               # Analyze command
│   ├── duplicates.go            # Duplicates command
│   ├── large.go                 # Large files command
│   └── cleanup.go               # Cleanup command
├── internal/                    # Internal packages
│   └── analyzer/                # Analysis engine
│       ├── analyzer.go          # Core analysis logic
│       └── report.go            # Report formatting
├── go.mod                       # Dependency management
└── README.md                    # Documentation
```

## Performance

- Process thousands of files easily
- Real-time progress tracking
- Optimized memory usage with efficient hashing (SHA256)
- Scales to large directories

## Security & Privacy

- **Local analysis only** - No data sent to internet
- **SHA256 hashing** - More reliable duplicate detection than MD5
- **No root privileges required** - Regular user analysis
- **Optional backups** before any file deletion
- **Dry Run mode** for safe preview before actual deletion

## Troubleshooting

### Issue: "Permission denied"
```bash
# Make executable
chmod +x storage-optimizer
```

### Issue: Invalid path
```bash
# Use absolute path
./storage-optimizer analyze /absolute/path/to/folder
```

### Issue: Out of memory with very large directories
Use filters to narrow scope:
```bash
./storage-optimizer analyze . --exclude lazy_backup,archive --extensions .go
```

## Build and Development

```bash
make help        # Show all commands
make build       # Build the binary
make test        # Run tests
make clean       # Clean build artifacts
make fmt         # Format code
make lint        # Lint code quality
make start       # Build and show help
make all         # Full build, test, and check
make install     # Install globally
```

## Contributing

Contributions welcome! Please:
1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Submit a pull request

## License

MIT License - You are free to use, modify, and distribute

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI Framework
- [Color](https://github.com/fatih/color) - Terminal Colors
- [Progressbar](https://github.com/schollz/progressbar) - Progress Bars

## Changelog

### v1.1.0 (Current)
- ✅ Enhanced filtering (extension, size, age, exclude patterns)
- ✅ CSV export format
- ✅ SHA256 hashing for duplicates (more reliable)
- ✅ Custom backup directory option
- ✅ Real-time progress tracking

### v1.0.0
- ✅ Comprehensive directory analysis
- ✅ Duplicate file detection
- ✅ Large file detection
- ✅ Safe cleanup with backups
- ✅ JSON and table output formats
- ✅ Cross-platform support

