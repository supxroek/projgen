package config

// frameworks.go
// กำหนด mapping ของ frameworks, templates, และคำสั่งที่ใช้สร้าง

// ProjectType ประเภทของโปรเจค
type ProjectType string

const (
	Frontend  ProjectType = "Frontend"
	Backend   ProjectType = "Backend"
	Fullstack ProjectType = "Fullstack"
)

// FrameworkOption ตัวเลือก framework พร้อม metadata
type FrameworkOption struct {
	Name           string   // ชื่อ framework
	DisplayName    string   // ชื่อที่แสดงในเมนู
	Language       string   // ภาษาที่ใช้
	TemplatePath   string   // path ของ template
	Runtime        string   // runtime ที่ต้องการ
	InstallCmd     string   // คำสั่งติดตั้ง dependencies
	StartCmd       string   // คำสั่งรันโปรเจค
	BuildCmd       string   // คำสั่ง build (ถ้ามี)
	Description    string   // คำอธิบาย
	SupportedAddons []string // addons ที่รองรับ เช่น tailwind, prisma
}

// GetFrontendFrameworks คืนค่า frameworks สำหรับ Frontend
func GetFrontendFrameworks() []FrameworkOption {
	return []FrameworkOption{
		{
			Name:         "vite-react-ts",
			DisplayName:  "Vite + React + TypeScript",
			Language:     "TypeScript",
			TemplatePath: "templates/frontend/vite-react-ts",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm run dev",
			BuildCmd:     "npm run build",
			Description:  "Vite with React and TypeScript - Fast, modern frontend tooling",
			SupportedAddons: []string{"tailwindcss", "material-ui", "bootstrap", "eslint", "prettier"},
		},
		{
			Name:         "vite-vue-ts",
			DisplayName:  "Vite + Vue + TypeScript",
			Language:     "TypeScript",
			TemplatePath: "templates/frontend/vite-vue-ts",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm run dev",
			BuildCmd:     "npm run build",
			Description:  "Vite with Vue 3 and TypeScript - Progressive JavaScript framework",
			SupportedAddons: []string{"tailwindcss", "vuetify", "eslint", "prettier"},
		},
		{
			Name:         "vite-svelte-ts",
			DisplayName:  "Vite + Svelte + TypeScript",
			Language:     "TypeScript",
			TemplatePath: "templates/frontend/vite-svelte-ts",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm run dev",
			BuildCmd:     "npm run build",
			Description:  "Vite with Svelte and TypeScript - Cybernetically enhanced web apps",
			SupportedAddons: []string{"tailwindcss", "eslint", "prettier"},
		},
		{
			Name:         "nextjs-ts",
			DisplayName:  "Next.js + TypeScript + Tailwind",
			Language:     "TypeScript",
			TemplatePath: "templates/frontend/nextjs-ts",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm run dev",
			BuildCmd:     "npm run build",
			Description:  "Next.js with TypeScript and Tailwind CSS - React framework for production",
			SupportedAddons: []string{"prisma", "auth", "eslint", "prettier"},
		},
	}
}

// GetBackendFrameworks คืนค่า frameworks สำหรับ Backend
func GetBackendFrameworks() []FrameworkOption {
	return []FrameworkOption{
		{
			Name:         "nestjs-api",
			DisplayName:  "NestJS + TypeScript",
			Language:     "TypeScript",
			TemplatePath: "templates/backend/nestjs-api",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm run start:dev",
			BuildCmd:     "npm run build",
			Description:  "NestJS - Progressive Node.js framework for building efficient and scalable server-side applications",
			SupportedAddons: []string{"prisma", "typeorm", "mongoose", "passport", "swagger"},
		},
		{
			Name:         "express-api",
			DisplayName:  "Express.js + JavaScript",
			Language:     "JavaScript",
			TemplatePath: "templates/backend/express-api",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm start",
			BuildCmd:     "",
			Description:  "Express.js - Fast, unopinionated, minimalist web framework for Node.js",
			SupportedAddons: []string{"mongodb", "postgresql", "mysql", "jwt", "cors"},
		},
		{
			Name:         "go-fiber",
			DisplayName:  "Go + Fiber",
			Language:     "Go",
			TemplatePath: "templates/backend/go-fiber",
			Runtime:      "go",
			InstallCmd:   "go mod tidy",
			StartCmd:     "go run main.go",
			BuildCmd:     "go build -o app",
			Description:  "Fiber - Express-inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go",
			SupportedAddons: []string{"gorm", "postgresql", "mysql", "redis", "jwt"},
		},
	}
}

// GetFullstackFrameworks คืนค่า frameworks สำหรับ Fullstack
func GetFullstackFrameworks() []FrameworkOption {
	return []FrameworkOption{
		{
			Name:         "t3-stack",
			DisplayName:  "T3 Stack (Next.js + tRPC + Prisma + Tailwind)",
			Language:     "TypeScript",
			TemplatePath: "templates/fullstack/t3-stack",
			Runtime:      "node",
			InstallCmd:   "npm install",
			StartCmd:     "npm run dev",
			BuildCmd:     "npm run build",
			Description:  "T3 Stack - The best way to start a full-stack, typesafe Next.js app",
			SupportedAddons: []string{"auth", "trpc", "prisma"},
		},
		{
			Name:         "mern-stack",
			DisplayName:  "MERN Stack (MongoDB + Express + React + Node)",
			Language:     "JavaScript",
			TemplatePath: "templates/fullstack/mern-stack",
			Runtime:      "node",
			InstallCmd:   "npm run install-all",
			StartCmd:     "npm run dev",
			BuildCmd:     "npm run build",
			Description:  "MERN Stack - Full-stack JavaScript solution",
			SupportedAddons: []string{"redux", "tailwindcss", "jwt", "mongoose"},
		},
	}
}

// CSSFrameworkOption ตัวเลือก CSS frameworks
type CSSFrameworkOption struct {
	Name        string
	DisplayName string
	InstallCmd  string
	ConfigFiles []string // ไฟล์ config ที่ต้องสร้าง
}

// GetCSSFrameworks คืนค่า CSS frameworks
func GetCSSFrameworks() []CSSFrameworkOption {
	return []CSSFrameworkOption{
		{
			Name:        "tailwindcss",
			DisplayName: "Tailwind CSS",
			InstallCmd:  "npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p",
			ConfigFiles: []string{"tailwind.config.js", "postcss.config.js"},
		},
		{
			Name:        "bootstrap",
			DisplayName: "Bootstrap",
			InstallCmd:  "npm install bootstrap",
			ConfigFiles: []string{},
		},
		{
			Name:        "material-ui",
			DisplayName: "Material UI",
			InstallCmd:  "npm install @mui/material @emotion/react @emotion/styled",
			ConfigFiles: []string{},
		},
		{
			Name:        "none",
			DisplayName: "None (Skip CSS framework)",
			InstallCmd:  "",
			ConfigFiles: []string{},
		},
	}
}

// UILibraryOption ตัวเลือก UI libraries
type UILibraryOption struct {
	Name        string
	DisplayName string
	InstallCmd  string
}

// GetUILibraries คืนค่า UI libraries
func GetUILibraries() []UILibraryOption {
	return []UILibraryOption{
		{
			Name:        "shadcn",
			DisplayName: "shadcn/ui",
			InstallCmd:  "npx shadcn@latest init",
		},
		{
			Name:        "radix",
			DisplayName: "Radix UI",
			InstallCmd:  "npm install @radix-ui/react-*",
		},
		{
			Name:        "none",
			DisplayName: "None",
			InstallCmd:  "",
		},
	}
}

// ExtraOption ตัวเลือกเสริม
type ExtraOption struct {
	Name        string
	DisplayName string
	Action      string // action ที่ต้องทำ เช่น "create-file", "run-command"
	Value       string // ค่าของ action
}

// GetExtras คืนค่าตัวเลือกเสริม
func GetExtras() []ExtraOption {
	return []ExtraOption{
		{
			Name:        "dockerfile",
			DisplayName: "Dockerfile",
			Action:      "create-file",
			Value:       "Dockerfile",
		},
		{
			Name:        "docker-compose",
			DisplayName: "Docker Compose",
			Action:      "create-file",
			Value:       "docker-compose.yml",
		},
		{
			Name:        "eslint",
			DisplayName: "ESLint",
			Action:      "run-command",
			Value:       "npm install -D eslint && npx eslint --init",
		},
		{
			Name:        "prettier",
			DisplayName: "Prettier",
			Action:      "run-command",
			Value:       "npm install -D prettier && echo {} > .prettierrc",
		},
		{
			Name:        "github-actions",
			DisplayName: "GitHub Actions CI/CD",
			Action:      "create-file",
			Value:       ".github/workflows/ci.yml",
		},
		{
			Name:        "env",
			DisplayName: ".env file",
			Action:      "create-file",
			Value:       ".env",
		},
		{
			Name:        "gitignore",
			DisplayName: ".gitignore",
			Action:      "create-file",
			Value:       ".gitignore",
		},
	}
}
