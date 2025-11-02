package ui

// ส่วนติดต่อผู้ใช้แบบเทอร์มินัล (Interactive UI) ด้วย survey/pterm/spinner

import (
	"context"
	"fmt"

	"projgen/internal/config"
	uiRuntime "projgen/internal/runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/pterm/pterm"
)

// ProjectOptions แทนตัวเลือกที่ผู้ใช้ระบุผ่านวิซาร์ด
type ProjectOptions struct {
	Name          string                  // ชื่อโปรเจ็กต์
	ProjectType   config.ProjectType      // ประเภทโปรเจค (Frontend/Backend/Fullstack)
	Framework     config.FrameworkOption  // framework ที่เลือก
	CSSFramework  *config.CSSFrameworkOption // CSS framework (optional)
	UILibrary     *config.UILibraryOption    // UI library (optional)
	Language      string                  // ภาษา (สำหรับ fallback)
	Runtime       string                  // รันไทม์ เช่น node, bun, deno, go
	Extras        []string                // ตัวเลือกเสริม เช่น Dockerfile, ESLint
	AutoInstall   bool                    // ติดตั้ง dependencies อัตโนมัติหรือไม่
}

// RunWizard เรียกใช้งานวิซาร์ดแบบโต้ตอบเพื่อเก็บตัวเลือกจากผู้ใช้ (ภาษาไทยทั้งหมด)
func RunWizard(ctx context.Context) (ProjectOptions, error) {
	var opts ProjectOptions
	opts.AutoInstall = true // default ให้ติดตั้งอัตโนมัติ

	// สร้าง header แบบสวยงามด้วยสีสันแต่ไม่มีพื้นหลัง
	titleStyle := pterm.NewStyle(pterm.FgCyan, pterm.Bold)
	subtitleStyle := pterm.NewStyle(pterm.FgLightMagenta, pterm.Bold)
	
	pterm.Println()
	pterm.DefaultCenter.Println(titleStyle.Sprint("╔══════════════════════════════════════════════════════════╗"))
	pterm.DefaultCenter.Println(titleStyle.Sprint("║") + "  " + titleStyle.Sprint("🚀 projgen") + subtitleStyle.Sprint(" - Project Generator CLI v2.0") + "  " + titleStyle.Sprint("║"))
	pterm.DefaultCenter.Println(titleStyle.Sprint("╚══════════════════════════════════════════════════════════╝"))
	pterm.Println()
	
	pterm.DefaultCenter.Println(pterm.LightGreen("✨ สร้างโปรเจ็กต์ของคุณในไม่กี่วินาที ✨"))
	pterm.Println()

	// 1) เลือกประเภทโปรเจค
	projectTypePrompt := &survey.Select{
		Message: "🎯 คุณต้องการสร้างโปรเจคประเภทไหน?",
		Options: []string{
			string(config.Frontend),
			string(config.Backend),
			string(config.Fullstack),
		},
		Default: string(config.Frontend),
	}
	var projectTypeStr string
	if err := survey.AskOne(projectTypePrompt, &projectTypeStr, survey.WithValidator(survey.Required)); err != nil {
		return ProjectOptions{}, err
	}
	opts.ProjectType = config.ProjectType(projectTypeStr)

	// 2) เลือก Framework ตามประเภทโปรเจค
	var frameworks []config.FrameworkOption
	switch opts.ProjectType {
	case config.Frontend:
		frameworks = config.GetFrontendFrameworks()
	case config.Backend:
		frameworks = config.GetBackendFrameworks()
	case config.Fullstack:
		frameworks = config.GetFullstackFrameworks()
	}

	frameworkOptions := make([]string, len(frameworks))
	for i, fw := range frameworks {
		frameworkOptions[i] = fw.DisplayName
	}

	frameworkPrompt := &survey.Select{
		Message: "🛠️  เลือก Framework/Stack:",
		Options: frameworkOptions,
		Description: func(value string, index int) string {
			if index < len(frameworks) {
				return frameworks[index].Description
			}
			return ""
		},
	}
	var selectedFrameworkName string
	if err := survey.AskOne(frameworkPrompt, &selectedFrameworkName, survey.WithValidator(survey.Required)); err != nil {
		return ProjectOptions{}, err
	}

	// หา framework ที่ถูกเลือก
	for _, fw := range frameworks {
		if fw.DisplayName == selectedFrameworkName {
			opts.Framework = fw
			break
		}
	}

	// 3) ถ้าเป็น Frontend ให้เลือก CSS Framework (ถ้า framework รองรับ)
	if opts.ProjectType == config.Frontend && len(opts.Framework.SupportedAddons) > 0 {
		// ตรวจสอบว่ารองรับ CSS framework หรือไม่
		supportsCSSFramework := false
		for _, addon := range opts.Framework.SupportedAddons {
			if addon == "tailwindcss" || addon == "bootstrap" {
				supportsCSSFramework = true
				break
			}
		}

		if supportsCSSFramework {
			cssFrameworks := config.GetCSSFrameworks()
			cssOptions := make([]string, len(cssFrameworks))
			for i, css := range cssFrameworks {
				cssOptions[i] = css.DisplayName
			}

			cssPrompt := &survey.Select{
				Message: "🎨 ต้องการเพิ่ม CSS Framework หรือไม่?",
				Options: cssOptions,
				Default: "None (Skip CSS framework)",
			}
			var selectedCSS string
			if err := survey.AskOne(cssPrompt, &selectedCSS); err != nil {
				return ProjectOptions{}, err
			}

			// หา CSS framework ที่เลือก
			for _, css := range cssFrameworks {
				if css.DisplayName == selectedCSS && css.Name != "none" {
					opts.CSSFramework = &css
					break
				}
			}
		}

		// 4) ถามเรื่อง UI Library (สำหรับบาง framework)
		if opts.Framework.Name == "vite-react-ts" || opts.Framework.Name == "nextjs-ts" {
			uiLibs := config.GetUILibraries()
			uiOptions := make([]string, len(uiLibs))
			for i, lib := range uiLibs {
				uiOptions[i] = lib.DisplayName
			}

			uiPrompt := &survey.Select{
				Message: "🧩 ต้องการเพิ่ม UI Library หรือไม่?",
				Options: uiOptions,
				Default: "None",
			}
			var selectedUI string
			if err := survey.AskOne(uiPrompt, &selectedUI); err != nil {
				return ProjectOptions{}, err
			}

			// หา UI library ที่เลือก
			for _, lib := range uiLibs {
				if lib.DisplayName == selectedUI && lib.Name != "none" {
					opts.UILibrary = &lib
					break
				}
			}
		}
	}

	// 5) ตรวจจับรันไทม์ (แสดงสปินเนอร์ระหว่างตรวจสอบ)
	pterm.Println()
	spinner, _ := pterm.DefaultSpinner.Start("🔍 กำลังตรวจสอบสภาพแวดล้อมรันไทม์...")
	runtimeDetected := uiRuntime.Detect(ctx)
	spinner.Stop()

	if runtimeDetected == "unknown" {
		pterm.Warning.WithPrefix(pterm.Prefix{
			Text:  " WARNING ",
			Style: pterm.NewStyle(pterm.FgBlack, pterm.BgYellow),
		}).Println("ไม่พบรันไทม์ที่รองรับ (Node, Bun, Deno หรือ Go) บนเครื่อง")
	} else {
		pterm.Success.WithPrefix(pterm.Prefix{
			Text:  " SUCCESS ",
			Style: pterm.NewStyle(pterm.FgBlack, pterm.BgGreen),
		}).Printfln("ตรวจพบรันไทม์: %s", pterm.Cyan(runtimeDetected))
	}
	opts.Runtime = runtimeDetected

	// แสดงรายงานรันไทม์ทั้งหมดที่พบบนเครื่อง
	statuses := uiRuntime.InspectAll(ctx)
	uiRuntime.PrintReport(statuses)

	// 6) ตั้งชื่อโปรเจ็กต์
	namePrompt := &survey.Input{
		Message: "📝 ตั้งชื่อโปรเจ็กต์:",
		Default: "my-app",
	}
	if err := survey.AskOne(namePrompt, &opts.Name, survey.WithValidator(survey.Required)); err != nil {
		return ProjectOptions{}, err
	}

	// 7) เลือกตัวเลือกเสริม
	extras := config.GetExtras()
	extraOptions := make([]string, len(extras))
	for i, ex := range extras {
		extraOptions[i] = ex.DisplayName
	}

	extrasPrompt := &survey.MultiSelect{
		Message: "⚙️  เลือกตัวเลือกเสริม (เลือกได้หลายข้อ):",
		Options: extraOptions,
	}
	var selectedExtras []string
	if err := survey.AskOne(extrasPrompt, &selectedExtras); err != nil {
		return ProjectOptions{}, err
	}
	opts.Extras = selectedExtras

	// 8) ถามว่าต้องการติดตั้ง dependencies อัตโนมัติหรือไม่
	autoInstallPrompt := &survey.Confirm{
		Message: "📦 ต้องการติดตั้ง dependencies อัตโนมัติหลังสร้างโปรเจคหรือไม่?",
		Default: true,
	}
	if err := survey.AskOne(autoInstallPrompt, &opts.AutoInstall); err != nil {
		return ProjectOptions{}, err
	}

	// 9) แสดงสรุปก่อนสร้าง
	pterm.Println()
	pterm.Println(pterm.LightCyan("─────────────────────────────────────────────────────────────"))
	pterm.DefaultSection.WithStyle(pterm.NewStyle(pterm.FgLightCyan)).Println("📋 สรุปการตั้งค่า")
	
	// สร้างตารางสวยๆ ด้วยสี
	tableData := pterm.TableData{
		{pterm.LightMagenta("รายการ"), pterm.LightMagenta("ค่า")},
		{pterm.Cyan("ชื่อโปรเจ็กต์"), pterm.LightGreen(opts.Name)},
		{pterm.Cyan("ประเภทโปรเจค"), pterm.LightYellow(string(opts.ProjectType))},
		{pterm.Cyan("Framework"), pterm.LightBlue(opts.Framework.DisplayName)},
		{pterm.Cyan("ภาษา"), pterm.White(opts.Framework.Language)},
		{pterm.Cyan("รันไทม์"), pterm.LightGreen(opts.Runtime)},
	}

	if opts.CSSFramework != nil {
		tableData = append(tableData, []string{pterm.Cyan("CSS Framework"), pterm.LightMagenta(opts.CSSFramework.DisplayName)})
	}
	if opts.UILibrary != nil {
		tableData = append(tableData, []string{pterm.Cyan("UI Library"), pterm.LightBlue(opts.UILibrary.DisplayName)})
	}
	if len(opts.Extras) > 0 {
		tableData = append(tableData, []string{pterm.Cyan("ตัวเลือกเสริม"), pterm.Yellow(fmt.Sprintf("%d รายการ", len(opts.Extras)))})
	}
	autoInstallText := "❌ ไม่"
	if opts.AutoInstall {
		autoInstallText = "✅ ใช่"
	}
	tableData = append(tableData, []string{pterm.Cyan("ติดตั้งอัตโนมัติ"), autoInstallText})

	// แสดงตารางแบบสวยงาม
	pterm.DefaultTable.
		WithHasHeader().
		WithHeaderRowSeparator("─").
		WithBoxed().
		WithData(tableData).
		Render()
	
	pterm.Println()
	pterm.Println(pterm.LightCyan("─────────────────────────────────────────────────────────────"))

	// 10) ยืนยันก่อนเริ่มสร้าง
	confirmPrompt := &survey.Confirm{
		Message: "🚀 เริ่มสร้างโปรเจ็กต์เลยไหม?",
		Default: true,
	}
	var confirm bool
	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		return ProjectOptions{}, err
	}
	if !confirm {
		pterm.Info.Println("ยกเลิกการสร้างโปรเจ็กต์")
		return ProjectOptions{}, fmt.Errorf("ยกเลิกโดยผู้ใช้")
	}

	return opts, nil
}
