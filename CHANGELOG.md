# Changelog

All notable changes to projgen will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Initial release
- Interactive CLI with Thai language support
- Multi-tier project selection (Frontend/Backend/Fullstack)
- Frontend frameworks: Vite + React/Vue/Svelte, Next.js
- Backend frameworks: NestJS, Express, Go Fiber
- Fullstack templates: T3 Stack, MERN
- CSS framework integration (Tailwind, Bootstrap, Material UI)
- UI library support (shadcn/ui, Radix UI)
- Auto dependency installation
- Runtime detection (Node, Bun, Deno, Go)
- Addon support (Dockerfile, ESLint, Prettier, GitHub Actions)
- Template engine with variable substitution
- Beautiful terminal output with pterm
- Progress indicators with spinner

### Features

#### Core Features

- 🎯 Interactive wizard with step-by-step guidance
- 🏗️ Multi-tier architecture support
- 📦 Automatic dependency installation
- 🎨 CSS framework integration
- 🧩 UI library support
- 🔍 Runtime auto-detection
- ⚙️ Configurable add-ons

#### Supported Templates

- **Frontend**: Vite + React/Vue/Svelte, Next.js
- **Backend**: NestJS, Express, Go Fiber
- **Fullstack**: T3 Stack, MERN Stack

#### Developer Experience

- Thai language support in UI
- Beautiful colored terminal output
- Progress indicators
- Comprehensive error messages
- Detailed success messages with next steps

### Technical

#### Architecture

- Clean modular structure
- Separate concerns (UI, Generator, Config, Runtime)
- Template-based generation
- Extensible framework system

#### Dependencies

- `github.com/spf13/cobra` - CLI framework
- `github.com/AlecAivazis/survey/v2` - Interactive prompts
- `github.com/pterm/pterm` - Terminal styling
- `github.com/briandowns/spinner` - Loading indicators

## [0.1.0] - 2025-11-02

### Added

- Initial scaffolding
- Basic CLI structure
- Template system foundation

---

## Future Releases

### [0.2.0] - Planned

- [ ] Database setup support (PostgreSQL, MongoDB, MySQL)
- [ ] Authentication templates
- [ ] More backend frameworks (Django, Laravel, FastAPI)
- [ ] Testing setup (Jest, Vitest, Go test)
- [ ] API documentation generation

### [0.3.0] - Planned

- [ ] Monorepo support (Turborepo, Nx)
- [ ] Docker Compose advanced configurations
- [ ] Kubernetes manifests
- [ ] Cloud deployment helpers (Vercel, AWS, GCP)

### [1.0.0] - Planned

- [ ] Stable API
- [ ] Complete documentation
- [ ] Video tutorials
- [ ] GUI version (Desktop app)
- [ ] Plugin system
- [ ] Template marketplace

---

## Release Notes Template

### [Version] - YYYY-MM-DD

#### Added

- New features

#### Changed

- Changes in existing functionality

#### Deprecated

- Soon-to-be removed features

#### Removed

- Removed features

#### Fixed

- Bug fixes

#### Security

- Security fixes

---

[Unreleased]: https://github.com/yourusername/projgen/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/yourusername/projgen/releases/tag/v0.1.0
