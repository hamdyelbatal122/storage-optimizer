# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2026-04-29

### Added
- **CSV Export Format** - Export analysis results to CSV for Excel/Google Sheets
- **Advanced Filtering**
  - `--extensions` - Include only specific file types
  - `--exclude` - Exclude directories and patterns
  - `--min-size` - Filter files by minimum size
  - `--max-age` - Filter files by modification age (days)
- **SHA256 Hashing** - Upgraded from MD5 for more reliable duplicate detection
- **Backup Directory Option** - `--backup-dir` flag for custom backup locations
- **Builder Pattern** - Chainable API for Analyzer configuration
- **Real-time Progress Display** - Visual feedback during file scanning
- Comprehensive README with examples and real-world use cases
- MIT License

### Features
- Duplicate file detection with SHA256 hashing
- Large file discovery with adjustable limits
- Directory analysis with file distribution
- Safe cleanup with backup capabilities
- Multiple output formats (Table, JSON, CSV)
- Cross-platform support (Linux, macOS, Windows)
- Professional CLI with Cobra framework

### Technical Details
- Built with Go 1.21+
- Modular architecture with separation of concerns
- Optimized for large directory traversal
- Memory-efficient file processing

---

## Version History

### v1.0.0 (Theoretical - Foundation)
- Initial concept with core features planned
- Basic duplicate detection
- File size analysis
- Safe cleanup mechanics
