# Template Creation Commands

เอกสารนี้รวมคำสั่งที่ใช้สร้าง template ต่างๆ ใน projgen

## 📋 Frontend Templates

### Vite + React + TypeScript

```bash
cd templates/frontend
npm create vite@latest vite-react-ts -- --template react-ts
```

### Vite + Vue + TypeScript

```bash
cd templates/frontend
npm create vite@latest vite-vue-ts -- --template vue-ts
```

### Vite + Svelte + TypeScript

```bash
cd templates/frontend
npm create vite@latest vite-svelte-ts -- --template svelte-ts
```

### Next.js + TypeScript + Tailwind

```bash
cd templates/frontend
npx create-next-app@latest nextjs-ts --typescript --tailwind --app --no-src-dir --import-alias "@/*" --turbopack --eslint --no-git
```

### React (CRA) - Optional

```bash
cd templates/frontend
npx create-react-app react-cra-ts --template typescript
```

### Vue 3

```bash
cd templates/frontend
npm create vue@latest vue3-ts
# เลือก: TypeScript, Router, Pinia, ESLint ตามต้องการ
```

### Angular

```bash
cd templates/frontend
npm install -g @angular/cli
ng new angular-ts --routing --style=scss --skip-git
```

---

## 🔧 Backend Templates

### NestJS + TypeScript

```bash
cd templates/backend
npm install -g @nestjs/cli
nest new nestjs-api --package-manager npm --skip-git
```

### Express.js

```bash
cd templates/backend
npx express-generator express-api --no-view --git
```

### Go + Fiber (Manual Setup)

```bash
# สร้างโฟลเดอร์และไฟล์เอง (ดูใน templates/backend/go-fiber)
cd templates/backend/go-fiber
# มี go.mod.tmpl, main.go.tmpl, README.md.tmpl
```

### Django

```bash
cd templates/backend
python -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate
pip install django
django-admin startproject django_api .
```

### Laravel

```bash
cd templates/backend
composer create-project laravel/laravel laravel-api
```

### FastAPI

```bash
cd templates/backend
mkdir fastapi-api && cd fastapi-api
python -m venv venv
source venv/bin/activate  # Windows: venv\Scripts\activate
pip install fastapi uvicorn
# สร้าง main.py, requirements.txt
```

---

## 🌐 Fullstack Templates

### T3 Stack (Next.js + tRPC + Prisma + Tailwind)

```bash
cd templates/fullstack
npx create-t3-app@latest t3-stack --noGit --CI
```

### MERN Stack (MongoDB + Express + React + Node)

```bash
cd templates/fullstack
mkdir mern-stack && cd mern-stack

# Backend
mkdir backend && cd backend
npm init -y
npm install express mongoose cors dotenv
# สร้าง server.js, models/, routes/

# Frontend
cd ..
npm create vite@latest frontend -- --template react-ts
cd frontend
npm install axios react-router-dom
```

### Next.js + NestJS

```bash
cd templates/fullstack
mkdir next-nest && cd next-nest

# Frontend
npx create-next-app@latest frontend --typescript --tailwind --app
# Backend
nest new backend --package-manager npm --skip-git
```

---

## 🎨 CSS Frameworks (Add-ons)

### Tailwind CSS

```bash
npm install -D tailwindcss postcss autoprefixer
npx tailwindcss init -p
```

### Bootstrap

```bash
npm install bootstrap
```

### Material UI (React)

```bash
npm install @mui/material @emotion/react @emotion/styled
```

### Vuetify (Vue)

```bash
npm install vuetify@^3.0.0
```

---

## 🧩 UI Libraries (Add-ons)

### shadcn/ui

```bash
npx shadcn@latest init
```

### Radix UI

```bash
npm install @radix-ui/react-dialog @radix-ui/react-dropdown-menu
```

### Headless UI

```bash
npm install @headlessui/react
```

---

## ⚙️ Development Tools (Add-ons)

### ESLint

```bash
npm install -D eslint
npx eslint --init
```

### Prettier

```bash
npm install -D prettier
echo {} > .prettierrc
```

### Docker

สร้างไฟล์ `Dockerfile` และ `docker-compose.yml` แบบ manual

### GitHub Actions

สร้างไฟล์ `.github/workflows/ci.yml` แบบ manual

---

## 📝 หมายเหตุ

1. **ติดตั้ง Dependencies**: หลังสร้าง template แล้ว ให้รันคำสั่ง:

   - Node.js: `npm install`
   - Go: `go mod tidy`
   - Python: `pip install -r requirements.txt`
   - PHP: `composer install`

2. **Template Variables**: ไฟล์ที่มีนามสกุล `.tmpl` จะถูกแปลงด้วย Go template engine:

   - `{{.Name}}` - ชื่อโปรเจค
   - `{{.KebabName}}` - ชื่อโปรเจคแบบ kebab-case
   - `{{.Port}}` - Port number
   - `{{.Language}}` - ภาษาที่ใช้
   - `{{.Framework}}` - Framework ที่เลือก

3. **Custom Templates**: คุณสามารถเพิ่ม template ของคุณเองได้โดย:

   - สร้างโฟลเดอร์ใน `templates/frontend/`, `templates/backend/`, หรือ `templates/fullstack/`
   - เพิ่ม config ใน `internal/config/frameworks.go`
   - ใช้ `.tmpl` suffix สำหรับไฟล์ที่ต้องการ template rendering

4. **Testing Templates**: หลังสร้าง template ใหม่ ให้ทดสอบด้วย:
   ```bash
   go run main.go create
   ```

---

## 🔄 Update Templates

เมื่อต้องการอัปเดต template ให้ใช้คำสั่งเดิม แล้วแทนที่ template เก่า:

```bash
# ลบ template เก่า
rm -rf templates/frontend/vite-react-ts

# สร้างใหม่
cd templates/frontend
npm create vite@latest vite-react-ts -- --template react-ts
```

---

## 📦 Auto-Install Scripts

CLI จะรันคำสั่งเหล่านี้โดยอัตโนมัติหากเลือก "Auto Install":

- **Node.js**: `npm install`
- **Go**: `go mod tidy`
- **Python**: `pip install -r requirements.txt` (if exists)
- **PHP**: `composer install` (if exists)

---

## 🚀 Quick Start

สร้างโปรเจคใหม่ด้วย:

```bash
# Build CLI
go build -o projgen

# Run
./projgen create

# หรือรันโดยตรง
go run main.go create
```

---

## 📚 Resources

- [Vite](https://vitejs.dev/)
- [Next.js](https://nextjs.org/)
- [NestJS](https://nestjs.com/)
- [Fiber](https://gofiber.io/)
- [T3 Stack](https://create.t3.gg/)
- [Tailwind CSS](https://tailwindcss.com/)
- [shadcn/ui](https://ui.shadcn.com/)
