package generator

// โมดูลนี้รับค่า ProjectOptions และทำหน้าที่สร้างโครงสร้างโปรเจ็กต์จริง ๆ
// ขั้นตอนหลัก: สร้างโฟลเดอร์, คัดลอกไฟล์เทมเพลต, เรนเดอร์ตัวแปร, และสร้างไฟล์พื้นฐาน

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/pterm/pterm"

	"projgen/internal/ui"
)

// Generate ประมวลผลการสร้างโครงสร้างโปรเจ็กต์จากตัวเลือกของผู้ใช้
func Generate(ctx context.Context, choices ui.ProjectOptions) error {
	destDir, err := projectDirFromChoices(choices)
	if err != nil {
		return err
	}

	// ตรวจสอบโฟลเดอร์ปลายทาง
	if err := ensureTargetDir(destDir); err != nil {
		return err
	}

	// แสดงสปินเนอร์ระหว่างสร้างไฟล์
	pterm.Println()
	spinner, _ := pterm.DefaultSpinner.Start("📦 กำลังสร้างโปรเจ็กต์และไฟล์ที่จำเป็น...")

	// 1) ค้นหาไดเรกทอรีเทมเพลต (ถ้ามี)
	tmplDir := resolveTemplateDir(choices)

	// 2) คัดลอก/เรนเดอร์ไฟล์จากเทมเพลต ถ้าพบ ไม่งั้นใช้ fallback
	if tmplDir != "" {
		if err := copyRenderTemplateDir(tmplDir, destDir, choices); err != nil {
			spinner.Fail("คัดลอกไฟล์จากเทมเพลตล้มเหลว")
			return fmt.Errorf("คัดลอกไฟล์จากเทมเพลตล้มเหลว: %w", err)
		}
	} else {
		if err := generateFallbackSkeleton(destDir, choices); err != nil {
			spinner.Fail("สร้างโครงสร้างพื้นฐานล้มเหลว")
			return fmt.Errorf("สร้างโครงสร้างพื้นฐานล้มเหลว: %w", err)
		}
	}
	spinner.Success("สร้างโครงสร้างโปรเจ็กต์เสร็จสิ้น")

	// 3) สร้างไฟล์เสริมตาม Extras เช่น .env, Dockerfile, README.md
	if err := generateExtras(destDir, choices); err != nil {
		return err
	}

	// 4) ติดตั้ง dependencies หากเลือกไว้
	if choices.AutoInstall && choices.Framework.InstallCmd != "" {
		spinner, _ = pterm.DefaultSpinner.Start("⬇️  กำลังติดตั้ง dependencies...")
		if err := installDependencies(ctx, destDir, choices); err != nil {
			spinner.Warning("ติดตั้ง dependencies ไม่สำเร็จ")
			pterm.Info.Printfln("   💡 คุณสามารถติดตั้งเองได้ด้วยคำสั่ง: %s", pterm.Cyan(choices.Framework.InstallCmd))
		} else {
			spinner.Success("ติดตั้ง dependencies สำเร็จ")
		}
	}

	// 5) ติดตั้ง CSS Framework หากเลือกไว้
	if choices.CSSFramework != nil && choices.CSSFramework.InstallCmd != "" {
		spinner, _ = pterm.DefaultSpinner.Start(fmt.Sprintf("🎨 กำลังติดตั้ง %s...", choices.CSSFramework.DisplayName))
		if err := installCSSFramework(ctx, destDir, choices); err != nil {
			spinner.Warning(fmt.Sprintf("ติดตั้ง %s ไม่สำเร็จ", choices.CSSFramework.DisplayName))
		} else {
			spinner.Success(fmt.Sprintf("ติดตั้ง %s สำเร็จ", choices.CSSFramework.DisplayName))
		}
	}

	// 6) ติดตั้ง UI Library หากเลือกไว้
	if choices.UILibrary != nil && choices.UILibrary.InstallCmd != "" {
		spinner, _ = pterm.DefaultSpinner.Start(fmt.Sprintf("🧩 กำลังติดตั้ง %s...", choices.UILibrary.DisplayName))
		if err := installUILibrary(ctx, destDir, choices); err != nil {
			spinner.Warning(fmt.Sprintf("ติดตั้ง %s ไม่สำเร็จ", choices.UILibrary.DisplayName))
		} else {
			spinner.Success(fmt.Sprintf("ติดตั้ง %s สำเร็จ", choices.UILibrary.DisplayName))
		}
	}

	printSuccessNextSteps(destDir, choices)
	return nil
}

// projectDirFromChoices กำหนดโฟลเดอร์ปลายทางจากชื่อโปรเจ็กต์ (ภายในโฟลเดอร์ปัจจุบัน)
func projectDirFromChoices(opts ui.ProjectOptions) (string, error) {
	if strings.TrimSpace(opts.Name) == "" {
		return "", errors.New("จำเป็นต้องระบุชื่อโปรเจ็กต์")
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// ใช้ชื่อโปรเจ็กต์เป็นโฟลเดอร์ปลายทาง
	return filepath.Join(wd, toKebab(opts.Name)), nil
}

// ensureTargetDir สร้างโฟลเดอร์ถ้ายังไม่มี และตรวจความว่างเปล่า
func ensureTargetDir(dir string) error {
	if fi, err := os.Stat(dir); err == nil {
		if !fi.IsDir() {
			return fmt.Errorf("ปลายทางชนไฟล์ที่มีอยู่แล้ว: %s", dir)
		}
		empty, err := isDirEmpty(dir)
		if err != nil {
			return err
		}
		if !empty {
			return fmt.Errorf("โฟลเดอร์ปลายทางไม่ว่างเปล่า: %s", dir)
		}
		return nil
	}
	return os.MkdirAll(dir, 0o755)
}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()
	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// resolveTemplateDir เลือกโฟลเดอร์เทมเพลตตามภาษา/เฟรมเวิร์กที่เลือก
// ค้นหาใน ./templates เป็นหลัก หากไม่พบจะคืนสตริงว่างเพื่อใช้ fallback
func resolveTemplateDir(opts ui.ProjectOptions) string {
	// ใช้ TemplatePath จาก framework option โดยตรง
	if opts.Framework.TemplatePath != "" {
		// ตรวจสอบว่า path มีอยู่จริง
		if _, err := os.Stat(opts.Framework.TemplatePath); err == nil {
			return opts.Framework.TemplatePath
		}
	}
	return ""
}

// copyRenderTemplateDir เดินสำรวจไดเรกทอรีเทมเพลตและเรนเดอร์ไฟล์ลงปลายทาง
func copyRenderTemplateDir(srcDir, destDir string, opts ui.ProjectOptions) error {
	data := map[string]any{
		"Name":         opts.Name,
		"Project":      opts.Name,
		"Language":     opts.Framework.Language,
		"Framework":    opts.Framework.DisplayName,
		"FrameworkName": opts.Framework.Name,
		"Runtime":      opts.Runtime,
		"Extras":       opts.Extras,
		"Port":         defaultPort(opts),
		"KebabName":    toKebab(opts.Name),
	}

	return filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(srcDir, path)
		target := filepath.Join(destDir, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		// รองรับไฟล์ .tmpl -> ตัดนามสกุลเมื่อเรนเดอร์ (ปลอดภัยแม้ไม่มีนามสกุลนี้)
		target = strings.TrimSuffix(target, ".tmpl")
		b, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		// พยายามเรนเดอร์เป็น text/template เสมอ (เหมาะกับไฟล์ข้อความ)
		if err := renderToFile(string(b), target, data); err != nil {
			return err
		}
		return nil
	})
}

// generateFallbackSkeleton กรณีไม่มีเทมเพลต ให้สร้างไฟล์พื้นฐานขั้นต่ำ
func generateFallbackSkeleton(destDir string, opts ui.ProjectOptions) error {
	// README.md
	readme := fmt.Sprintf("# %s\n\nโปรเจ็กต์ที่สร้างด้วย projgen (โหมดพื้นฐาน)\n\nภาษา: %s\nเฟรมเวิร์ก: %s\nรันไทม์: %s\n", 
		opts.Name, opts.Framework.Language, opts.Framework.DisplayName, opts.Runtime)
	if err := os.WriteFile(filepath.Join(destDir, "README.md"), []byte(readme), 0o644); err != nil {
		return err
	}
	// โครงสร้าง src ง่าย ๆ
	if err := os.MkdirAll(filepath.Join(destDir, "src"), 0o755); err != nil {
		return err
	}
	mainFile := filepath.Join(destDir, "src", "main.txt")
	content := fmt.Sprintf("โปรเจ็กต์ %s สร้างเมื่อ %s", opts.Name, time.Now().Format(time.RFC3339))
	if err := os.WriteFile(mainFile, []byte(content), 0o644); err != nil {
		return err
	}
	return nil
}

// generateExtras สร้างไฟล์เสริมตามตัวเลือก
func generateExtras(destDir string, opts ui.ProjectOptions) error {
	// .env เสมอถ้าเลือก
	if contains(opts.Extras, ".env") {
		env := fmt.Sprintf("PORT=%d\nAPP_NAME=%s\n", defaultPort(opts), toKebab(opts.Name))
		if err := os.WriteFile(filepath.Join(destDir, ".env"), []byte(env), 0o644); err != nil {
			return err
		}
	}
	// Dockerfile
	if contains(opts.Extras, "Dockerfile") {
		df := dockerfileFor(opts)
		if err := os.WriteFile(filepath.Join(destDir, "Dockerfile"), []byte(df), 0o644); err != nil {
			return err
		}
	}
	// README.md เสริม (ถ้ายังไม่มี)
	readmePath := filepath.Join(destDir, "README.md")
	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		content := fmt.Sprintf("# %s\n\nสร้างด้วย projgen\n", opts.Name)
		if err := os.WriteFile(readmePath, []byte(content), 0o644); err != nil {
			return err
		}
	}
	return nil
}

func renderToFile(tpl string, dest string, data any) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
		"ToUpper": strings.ToUpper,
		"Kebab":   toKebab,
	}
	t, err := template.New("file").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return err
	}
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := t.Execute(f, data); err != nil {
		return err
	}
	return nil
}

func defaultPort(opts ui.ProjectOptions) int {
	if strings.EqualFold(opts.Framework.Language, "Go") {
		return 8080
	}
	// ค่าปกติสำหรับเว็บ JS/TS
	return 3000
}

func dockerfileFor(opts ui.ProjectOptions) string {
	if strings.EqualFold(opts.Framework.Language, "Go") || strings.EqualFold(opts.Runtime, "go") {
		return fmt.Sprintf(`FROM golang:1.25.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app ./...

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app /usr/local/bin/app
EXPOSE %d
CMD ["/usr/local/bin/app"]
`, defaultPort(opts))
	}
	// Node/Bun/Deno (ใช้ Node เป็นค่าปกติ)
	return fmt.Sprintf(`FROM node:20-alpine
WORKDIR /app
COPY package*.json ./
RUN npm ci || npm install
COPY . .
EXPOSE %d
CMD ["npm","run","start"]
`, defaultPort(opts))
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}

func toKebab(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ToLower(s)
	// ลบอักขระซ้ำ ๆ ที่ไม่จำเป็น
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return s
}

// แสดงข้อความสำเร็จและขั้นตอนถัดไปแบบเป็นมิตร
func printSuccessNextSteps(destDir string, opts ui.ProjectOptions) {
	projectDirName := filepath.Base(destDir)
	
	pterm.Println()
	pterm.Println(pterm.LightCyan("═══════════════════════════════════════════════════════════"))
	
	// Success message with style
	successBox := pterm.DefaultBox.
		WithTitle("🎉 สำเร็จ!").
		WithTitleTopCenter().
		WithBoxStyle(pterm.NewStyle(pterm.FgLightGreen)).
		Sprintln(fmt.Sprintf("โปรเจ็กต์ %s สร้างเสร็จเรียบร้อยแล้ว", pterm.LightGreen(projectDirName)))
	
	pterm.Println(successBox)

	// Next steps
	cmds := nextCommands(projectDirName, opts)
	pterm.Println()
	pterm.DefaultSection.WithStyle(pterm.NewStyle(pterm.FgLightCyan)).Println("� ขั้นตอนถัดไป")
	
	for i, c := range cmds {
		prefix := pterm.LightMagenta(fmt.Sprintf("%d.", i+1))
		pterm.Printfln("   %s %s", prefix, pterm.Cyan(c))
	}
	
	pterm.Println()
	pterm.Println(pterm.LightCyan("═══════════════════════════════════════════════════════════"))
	pterm.Println()
	pterm.DefaultCenter.Println(pterm.LightGreen("✨ Happy Coding! ✨"))
	pterm.Println()
}

func nextCommands(dir string, opts ui.ProjectOptions) []string {
	cmds := []string{fmt.Sprintf("cd %s", dir)}
	// แนะนำคำสั่งรันเริ่มต้นตามภาษาหรือรันไทม์
	if strings.EqualFold(opts.Framework.Language, "Go") || strings.EqualFold(opts.Runtime, "go") {
		if opts.Framework.StartCmd != "" {
			cmds = append(cmds, opts.Framework.StartCmd)
		} else {
			cmds = append(cmds, "go run ./...")
		}
		return cmds
	}
	
	// ถ้ามีคำสั่งกำหนดไว้ใน framework config ให้ใช้เลย
	if opts.AutoInstall && opts.Framework.InstallCmd != "" {
		cmds = append(cmds, opts.Framework.InstallCmd)
	}
	if opts.Framework.StartCmd != "" {
		cmds = append(cmds, opts.Framework.StartCmd)
		return cmds
	}
	
	// ค่าปกติฝั่งเว็บ JS/TS
	switch strings.ToLower(opts.Runtime) {
	case "bun":
		if !opts.AutoInstall {
			cmds = append(cmds, "bun install")
		}
		cmds = append(cmds, "bun run dev")
	case "deno":
		// สมมติว่ามี task ชื่อ dev ใน deno.json (อาจต้องแก้ไขตามเทมเพลตจริง)
		cmds = append(cmds, "deno task dev")
	default:
		if !opts.AutoInstall {
			cmds = append(cmds, "npm install")
		}
		cmds = append(cmds, "npm run dev")
	}
	return cmds
}

// installDependencies ติดตั้ง dependencies หลัก
func installDependencies(ctx context.Context, destDir string, opts ui.ProjectOptions) error {
	if opts.Framework.InstallCmd == "" {
		return nil
	}
	return runCommandInDir(ctx, destDir, opts.Framework.InstallCmd)
}

// installCSSFramework ติดตั้ง CSS framework
func installCSSFramework(ctx context.Context, destDir string, opts ui.ProjectOptions) error {
	if opts.CSSFramework == nil || opts.CSSFramework.InstallCmd == "" {
		return nil
	}
	return runCommandInDir(ctx, destDir, opts.CSSFramework.InstallCmd)
}

// installUILibrary ติดตั้ง UI library
func installUILibrary(ctx context.Context, destDir string, opts ui.ProjectOptions) error {
	if opts.UILibrary == nil || opts.UILibrary.InstallCmd == "" {
		return nil
	}
	return runCommandInDir(ctx, destDir, opts.UILibrary.InstallCmd)
}

// runCommandInDir รันคำสั่งใน directory ที่ระบุ
func runCommandInDir(ctx context.Context, dir string, cmdStr string) error {
	// แยกคำสั่งและ arguments
	parts := strings.Fields(cmdStr)
	if len(parts) == 0 {
		return fmt.Errorf("คำสั่งว่างเปล่า")
	}
	
	cmd := exec.CommandContext(ctx, parts[0], parts[1:]...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

