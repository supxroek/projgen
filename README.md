# 🚀 projgen - Project Generator CLI

> เครื่องมือสร้างโปรเจคอัตโนมัติที่ฉลาด ยืดหยุ่น และพร้อมใช้งานทันที

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)
[![GitHub](https://img.shields.io/badge/GitHub-supxroek%2Fprojgen-181717?style=flat&logo=github)](https://github.com/supxroek/projgen)

---

## ✨ Features

- 🎯 **Interactive CLI** - เมนูแบบโต้ตอบที่ใช้งานง่าย รองรับภาษาไทยเต็มรูปแบบ
- 🏗️ **Multi-tier Architecture** - เลือกสร้าง Frontend, Backend หรือ Fullstack
- 🔧 **Framework Flexibility** - รองรับ framework ยอดนิยมมากมาย
- 📦 **Auto Installation** - ติดตั้ง dependencies อัตโนมัติหลังสร้างโปรเจค
- 🎨 **Addon Support** - เลือก CSS framework, UI library, และเครื่องมือเสริม
- 🔍 **Runtime Detection** - ตรวจจับ runtime (Node, Bun, Deno, Go) อัตโนมัติ
- 📝 **Template Engine** - ใช้ template จริงจาก official CLI ของแต่ละ framework
- ⚡ **Fast & Efficient** - สร้างโปรเจคได้ภายในไม่กี่วินาที

---

## 📦 Installation

### Prerequisites

- Go 1.25 or higher
- Node.js 18+ (สำหรับ frontend/fullstack projects)
- Git

### From Source

```bash
# Clone repository
git clone https://github.com/supxroek/projgen.git
cd projgen

# Install dependencies
go mod tidy

# Build
go build -o projgen

# Run
./projgen create
```

### Using Go Install (Coming Soon)

```bash
go install github.com/yourusername/projgen@latest
```

---

## 🎮 Usage

### Basic Usage

```bash
# Start interactive wizard
projgen create

# หรือ
go run main.go create
```

### Interactive Flow

```
🚀 Project Generator (projgen)
───────────────────────────────────────

? คุณต้องการสร้างโปรเจคประเภทไหน?
  ▸ Frontend
    Backend
    Fullstack

? เลือก Framework/Stack:
  ▸ Vite + React + TypeScript
    Vite + Vue + TypeScript
    Vite + Svelte + TypeScript
    Next.js + TypeScript + Tailwind

? ต้องการเพิ่ม CSS Framework หรือไม่?
  ▸ Tailwind CSS
    Bootstrap
    Material UI
    None

? ต้องการเพิ่ม UI Library หรือไม่?
  ▸ shadcn/ui
    Radix UI
    None

✅ ตรวจพบรันไทม์: node

? ตั้งชื่อโปรเจ็กต์: my-awesome-app

? เลือกตัวเลือกเสริม:
  ◉ Dockerfile
  ◉ Docker Compose
  ◯ ESLint
  ◉ Prettier
  ◉ GitHub Actions CI/CD
  ◉ .env file
  ◉ .gitignore

? ต้องการติดตั้ง dependencies อัตโนมัติหลังสร้างโปรเจคหรือไม่? Yes

🧾 สรุปการตั้งค่า
┌─────────────────────┬──────────────────────────────────┐
│ รายการ              │ ค่า                              │
├─────────────────────┼──────────────────────────────────┤
│ ชื่อโปรเจ็กต์        │ my-awesome-app                   │
│ ประเภทโปรเจค         │ Frontend                         │
│ Framework           │ Vite + React + TypeScript        │
│ CSS Framework       │ Tailwind CSS                     │
│ UI Library          │ shadcn/ui                        │
│ รันไทม์              │ node                             │
│ ติดตั้งอัตโนมัติ     │ true                             │
└─────────────────────┴──────────────────────────────────┘

? เริ่มสร้างโปรเจ็กต์เลยไหม? Yes

⠋ กำลังสร้างโปรเจ็กต์และไฟล์ที่จำเป็น...
⠋ กำลังติดตั้ง dependencies...
⠋ กำลังติดตั้ง Tailwind CSS...
⠋ กำลังติดตั้ง shadcn/ui...

🎉 สร้างโปรเจ็กต์สำเร็จ!
ℹ️ โฟลเดอร์: my-awesome-app

👉 ขั้นตอนถัดไป
   cd my-awesome-app
   npm run dev
```

---

## 🏗️ Supported Frameworks

### 🎨 Frontend

| Framework                      | Description                      | Template         |
| ------------------------------ | -------------------------------- | ---------------- |
| **Vite + React + TypeScript**  | Fast, modern frontend with HMR   | `vite-react-ts`  |
| **Vite + Vue + TypeScript**    | Progressive JavaScript framework | `vite-vue-ts`    |
| **Vite + Svelte + TypeScript** | Cybernetically enhanced web apps | `vite-svelte-ts` |
| **Next.js + TypeScript**       | React framework for production   | `nextjs-ts`      |

### 🔧 Backend

| Framework               | Description                   | Template      |
| ----------------------- | ----------------------------- | ------------- |
| **NestJS + TypeScript** | Progressive Node.js framework | `nestjs-api`  |
| **Express.js**          | Minimalist web framework      | `express-api` |
| **Go + Fiber**          | Express-inspired Go framework | `go-fiber`    |

### 🌐 Fullstack

| Stack          | Description                        | Template     |
| -------------- | ---------------------------------- | ------------ |
| **T3 Stack**   | Next.js + tRPC + Prisma + Tailwind | `t3-stack`   |
| **MERN Stack** | MongoDB + Express + React + Node   | `mern-stack` |

---

## 🎨 CSS Frameworks

- **Tailwind CSS** - Utility-first CSS framework
- **Bootstrap** - Popular CSS framework
- **Material UI** - React component library (Material Design)

---

## 🧩 UI Libraries

- **shadcn/ui** - Re-usable components built with Radix UI and Tailwind
- **Radix UI** - Unstyled, accessible components
- **Headless UI** - Unstyled UI components

---

## ⚙️ Add-ons & Tools

- ✅ **Dockerfile** - Ready-to-use Docker configuration
- ✅ **Docker Compose** - Multi-container setup
- ✅ **ESLint** - Code linting
- ✅ **Prettier** - Code formatting
- ✅ **GitHub Actions** - CI/CD pipeline
- ✅ **.env** - Environment variables
- ✅ **.gitignore** - Git ignore rules

---

## 📁 Project Structure

```
projgen/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command
│   └── create.go          # Create command
├── internal/
│   ├── config/            # Configuration & framework definitions
│   │   ├── config.go
│   │   └── frameworks.go  # Framework mappings
│   ├── generator/         # Project generation logic
│   │   └── generator.go
│   ├── runtime/           # Runtime detection
│   │   └── runtime.go
│   ├── templates/         # Template management
│   │   └── templates.go
│   └── ui/                # Interactive UI
│       └── ui.go
├── templates/             # Project templates
│   ├── frontend/
│   │   ├── vite-react-ts/
│   │   ├── vite-vue-ts/
│   │   ├── vite-svelte-ts/
│   │   └── nextjs-ts/
│   ├── backend/
│   │   ├── nestjs-api/
│   │   ├── express-api/
│   │   └── go-fiber/
│   └── fullstack/
│       ├── t3-stack/
│       └── mern-stack/
├── go.mod
├── go.sum
├── main.go
├── README.md
└── TEMPLATES.md           # Template creation guide
```

---

## 🔧 Configuration

### Adding New Frameworks

1. สร้าง template ใน `templates/` (ดูรายละเอียดใน [TEMPLATES.md](TEMPLATES.md))
2. เพิ่ม framework config ใน `internal/config/frameworks.go`

```go
{
    Name:         "my-framework",
    DisplayName:  "My Framework + TypeScript",
    Language:     "TypeScript",
    TemplatePath: "templates/frontend/my-framework",
    Runtime:      "node",
    InstallCmd:   "npm install",
    StartCmd:     "npm run dev",
    BuildCmd:     "npm run build",
    Description:  "My awesome framework",
    SupportedAddons: []string{"tailwindcss", "eslint"},
}
```

3. Build และทดสอบ:

```bash
go build -o projgen
./projgen create
```

---

## 🧪 Development

### Running Tests

```bash
go test ./...
```

### Running with Hot Reload

```bash
# Install air
go install github.com/cosmtrek/air@latest

# Run
air
```

### Debugging

```bash
# Enable verbose logging
PROJGEN_DEBUG=1 go run main.go create
```

---

## 📝 Template Variables

Templates support Go template syntax with these variables:

| Variable         | Description          | Example        |
| ---------------- | -------------------- | -------------- |
| `{{.Name}}`      | Project name         | `my-app`       |
| `{{.KebabName}}` | Kebab-case name      | `my-app`       |
| `{{.Language}}`  | Programming language | `TypeScript`   |
| `{{.Framework}}` | Framework name       | `Vite + React` |
| `{{.Runtime}}`   | Runtime name         | `node`         |
| `{{.Port}}`      | Default port         | `3000`         |

Example:

```typescript
// package.json.tmpl
{
  "name": "{{.KebabName}}",
  "version": "0.1.0",
  "description": "{{.Name}} - Created with projgen"
}
```

---

## 🤝 Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) first.

### How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Survey](https://github.com/AlecAivazis/survey) - Interactive prompts
- [pterm](https://github.com/pterm/pterm) - Beautiful terminal output
- [Vite](https://vitejs.dev/) - Frontend tooling
- [Next.js](https://nextjs.org/) - React framework
- [NestJS](https://nestjs.com/) - Backend framework
- [Fiber](https://gofiber.io/) - Go web framework
- [T3 Stack](https://create.t3.gg/) - Fullstack template

---

## 📧 Contact

- **Author**: Your Name
- **Email**: suparoek.sm@gmail.com
- **GitHub**: [@suparoek](https://github.com/supxroek)
- **Twitter**: [@_supxroek](https://x.com/_supxroek)

---

## 🗺️ Roadmap

- [x] Frontend templates (Vite, Next.js)
- [x] Backend templates (NestJS, Express, Go Fiber)
- [x] Fullstack templates (T3, MERN)
- [x] CSS framework integration
- [x] UI library support
- [x] Auto dependency installation
- [ ] More backend frameworks (Django, Laravel, FastAPI)
- [ ] Database setup (PostgreSQL, MongoDB, MySQL)
- [ ] Authentication templates
- [ ] API documentation generation
- [ ] Testing setup (Jest, Vitest, Go test)
- [ ] Monorepo support (Turborepo, Nx)
- [ ] Cloud deployment helpers (Vercel, AWS, GCP)
- [ ] GUI version (Desktop app)

---

<p align="center">Made with ❤️ by supxroek, for developers</p>
