# Contributing to Storage Optimizer

Thank you for your interest in contributing to Storage Optimizer! We welcome contributions from everyone.

## Code of Conduct

Please be respectful and professional in all interactions.

## How Can I Contribute?

###  Reporting Bugs

If you find a bug, please create an issue with:
- Clear description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version, etc.)

###  Suggesting Enhancements

We'd love to hear your ideas! Create an issue with:
- Clear description of the feature
- Use cases and benefits
- Possible implementation approach

###  Code Contributions

1. **Fork the repository**
   ```bash
   git clone https://github.com/hamdyelbatal122/storage-optimizer.git
   cd storage-optimizer
   ```

2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-amazing-feature
   ```

3. **Make your changes**
   - Follow Go conventions
   - Keep code clean and well-documented
   - Add comments for complex logic

4. **Test your changes**
   ```bash
   make test
   make lint
   make build
   ```

5. **Commit with clear messages**
   ```bash
   git commit -m "Add your feature description"
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/your-amazing-feature
   ```

7. **Create a Pull Request**
   - Link related issues
   - Describe your changes
   - Explain the benefits

## Development Setup

### Prerequisites
- Go 1.21 or later
- Make (optional but recommended)

### Build
```bash
make build
```

### Run Tests
```bash
make test
```

### Format Code
```bash
make fmt
```

### Lint
```bash
make lint
```

### Full Check
```bash
make all
```

## Project Structure

```
storage-optimizer/
├── main.go                    # Entry point
├── cmd/                       # CLI commands
│   ├── analyze.go
│   ├── duplicates.go
│   ├── large.go
│   ├── cleanup.go
│   └── root.go
├── internal/analyzer/         # Core analysis logic
│   ├── analyzer.go
│   └── report.go
├── Makefile                   # Build automation
├── README.md                  # Documentation
└── ... (other files)
```

## Code Style

- Follow [Go Code Review Comments](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Keep functions small and focused
- Add comments for exported functions
- Use meaningful variable names

## Testing

- Write tests for new features
- Ensure existing tests pass
- Test on multiple platforms if possible

## Documentation

- Update README.md for user-facing changes
- Update CHANGELOG.md with new features
- Add comments to complex code
- Provide clear examples

## Commit Messages

- Use clear, descriptive messages
- Start with present tense ("Add feature" not "Added feature")
- Reference issues when relevant: "Fixes #123"

Example:
```
Add CSV export format

- Implement CSV writer in report.go
- Add format flag to all commands
- Update documentation with examples

Fixes #42
```

## Pull Request Process

1. Update documentation as needed
2. Update CHANGELOG.md
3. Ensure tests pass: `make all`
4. Request reviews from maintainers
5. Address review feedback

## Areas We Need Help

- [ ] GUI/TUI interface
- [ ] Windows batch scripts
- [ ] Performance optimizations
- [ ] Additional export formats
- [ ] Documentation translations
- [ ] CI/CD pipeline setup
- [ ] Docker containerization

## Release Process

Maintainers handle releases, but you'll see:
1. Version bump in appropriate files
2. CHANGELOG update
3. GitHub release creation
4. Tag pushed

## Questions?

Create an issue or contact the maintainers.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

**Thank you for contributing to Storage Optimizer!** 
