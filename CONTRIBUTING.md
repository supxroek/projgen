# 🎉 Contributing to projgen

Thank you for considering contributing to projgen! We welcome contributions from everyone.

## 🚀 Getting Started

### Prerequisites

- Go 1.22 or higher
- Node.js 18+ (for testing templates)
- Git
- Basic knowledge of Go and CLI development

### Setting Up Development Environment

1. **Fork and Clone**

```bash
git clone https://github.com/yourusername/projgen.git
cd projgen
```

2. **Install Dependencies**

```bash
go mod download
```

3. **Build**

```bash
go build -o projgen
```

4. **Run**

```bash
./projgen create
```

---

## 📝 Development Workflow

### Making Changes

1. **Create a Branch**

```bash
git checkout -b feature/your-feature-name
```

2. **Make Your Changes**

- Write clean, readable code
- Follow Go conventions
- Add comments for complex logic
- Update documentation if needed

3. **Test Your Changes**

```bash
# Run tests
go test ./...

# Build and test CLI
go build -o projgen
./projgen create
```

4. **Commit**

```bash
git add .
git commit -m "feat: add awesome feature"
```

Use conventional commits:

- `feat:` - New feature
- `fix:` - Bug fix
- `docs:` - Documentation changes
- `refactor:` - Code refactoring
- `test:` - Adding tests
- `chore:` - Maintenance tasks

5. **Push and Create PR**

```bash
git push origin feature/your-feature-name
```

---

## 🏗️ Project Structure

```
projgen/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command setup
│   └── create.go          # Create command logic
├── internal/              # Internal packages
│   ├── config/           # Framework configurations
│   ├── generator/        # Project generation logic
│   ├── runtime/          # Runtime detection
│   ├── templates/        # Template management
│   └── ui/               # Interactive UI
├── templates/            # Project templates
└── main.go               # Entry point
```

---

## 🎯 What to Contribute

### 🌟 High Priority

- **New Framework Templates** - Add support for popular frameworks
- **Bug Fixes** - Fix issues and improve stability
- **Documentation** - Improve guides and examples
- **Tests** - Add unit and integration tests

### 💡 Good Ideas

- **UI Improvements** - Better interactive experience
- **Error Handling** - Better error messages
- **Performance** - Optimize generation speed
- **New Features** - Database setup, auth templates, etc.

### 📋 Areas of Contribution

1. **Adding New Templates**

   - Create framework templates (see TEMPLATES.md)
   - Add configuration in `internal/config/frameworks.go`
   - Test template generation

2. **Improving Generator**

   - Better file copying logic
   - Template variable expansion
   - Post-generation hooks

3. **UI Enhancements**

   - More interactive prompts
   - Better progress indicators
   - Colorful output

4. **Documentation**
   - Usage examples
   - Tutorial videos
   - Architecture docs

---

## 📚 Adding New Frameworks

### Step 1: Create Template

```bash
# Navigate to appropriate directory
cd templates/frontend  # or backend, fullstack

# Create template using official CLI
npm create vite@latest my-framework -- --template react
```

### Step 2: Add Configuration

Edit `internal/config/frameworks.go`:

```go
{
    Name:           "my-framework",
    DisplayName:    "My Framework + TypeScript",
    Language:       "TypeScript",
    TemplatePath:   "templates/frontend/my-framework",
    Runtime:        "node",
    InstallCmd:     "npm install",
    StartCmd:       "npm run dev",
    BuildCmd:       "npm run build",
    Description:    "Description of the framework",
    SupportedAddons: []string{"tailwindcss", "eslint"},
}
```

### Step 3: Test

```bash
go build -o projgen
./projgen create
# Select your new framework and test
```

### Step 4: Update Documentation

- Add to README.md
- Update TEMPLATES.md with creation commands
- Add examples

---

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run specific package
go test ./internal/generator

# Run with coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

### Manual Testing

```bash
# Build
go build -o projgen

# Test various scenarios
./projgen create
# - Test different frameworks
# - Test with/without addons
# - Test error cases
# - Test cancellation
```

---

## 📖 Code Style

### Go Code

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Keep functions small and focused
- Add comments for exported functions
- Use meaningful variable names

Example:

```go
// GenerateProject creates a new project based on user options
func GenerateProject(ctx context.Context, opts ProjectOptions) error {
    // Validate options
    if err := validateOptions(opts); err != nil {
        return fmt.Errorf("invalid options: %w", err)
    }

    // Create project directory
    if err := createDirectory(opts.Name); err != nil {
        return fmt.Errorf("failed to create directory: %w", err)
    }

    return nil
}
```

### Thai Language Support

- Use Thai in UI prompts for better UX
- Keep internal code/comments in English
- Thai messages should be clear and friendly

---

## 🐛 Reporting Bugs

### Before Submitting

1. Check if the bug was already reported
2. Try to reproduce with latest version
3. Gather relevant information:
   - OS and version
   - Go version
   - Command used
   - Error messages
   - Steps to reproduce

### Bug Report Template

```markdown
**Describe the bug**
A clear description of what the bug is.

**To Reproduce**
Steps to reproduce:

1. Run command '...'
2. Select option '...'
3. See error

**Expected behavior**
What you expected to happen.

**Screenshots**
If applicable, add screenshots.

**Environment:**

- OS: [e.g., Windows 11]
- Go Version: [e.g., 1.22]
- projgen Version: [e.g., 1.0.0]

**Additional context**
Any other context about the problem.
```

---

## 💡 Feature Requests

We love feature ideas! Please:

1. Check if it's already requested
2. Describe the use case
3. Explain why it's useful
4. Provide examples if possible

### Feature Request Template

```markdown
**Is your feature related to a problem?**
A clear description of the problem.

**Describe the solution you'd like**
What you want to happen.

**Describe alternatives considered**
Other solutions you've thought about.

**Additional context**
Any other context, screenshots, or examples.
```

---

## 📜 Code of Conduct

### Our Pledge

We pledge to make participation in our project a harassment-free experience for everyone.

### Our Standards

✅ **Positive behaviors:**

- Being respectful and inclusive
- Accepting constructive criticism
- Focusing on what's best for the community
- Showing empathy

❌ **Unacceptable behaviors:**

- Harassment or discrimination
- Trolling or insulting comments
- Public or private harassment
- Publishing others' private information

### Enforcement

Report violations to [your.email@example.com](mailto:your.email@example.com).

---

## 📞 Getting Help

- 💬 **Discussions**: Ask questions in GitHub Discussions
- 🐛 **Issues**: Report bugs in GitHub Issues
- 📧 **Email**: Contact maintainers directly
- 📚 **Docs**: Check README.md and TEMPLATES.md

---

## 🎁 Recognition

Contributors will be:

- Added to CONTRIBUTORS.md
- Mentioned in release notes
- Thanked in README.md

---

## ✅ Checklist Before Submitting PR

- [ ] Code follows project style guidelines
- [ ] Self-review of code done
- [ ] Comments added for complex parts
- [ ] Documentation updated (if needed)
- [ ] Tests added/updated (if applicable)
- [ ] All tests pass
- [ ] No new warnings
- [ ] Commit messages follow conventions

---

## 🌟 Thank You!

Your contributions make projgen better for everyone. We appreciate your time and effort! 🙏

---

**Happy Coding! 🚀**
