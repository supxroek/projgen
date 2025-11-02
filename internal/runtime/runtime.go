package runtime

// เครื่องมือตรวจจับสภาพแวดล้อมรันไทม์ (Node/Bun/Deno/Go/Python/npm/pip) และตรวจสอบเวอร์ชัน
// รวมถึงยูทิลิตี้สำหรับเรียกคำสั่งแบบข้ามแพลตฟอร์ม พร้อมข้อความเตือนแบบมีสีสัน

import (
	"bytes"
	"context"
	"os/exec"
	"regexp"
	"runtime"
	"strings"

	"github.com/pterm/pterm"
)

// RuntimeStatus สถานะของรันไทม์ที่ตรวจพบ
type RuntimeStatus struct {
	Name    string // ชื่อรันไทม์ เช่น node, npm, go, python, bun, deno, pip
	Found   bool   // พบหรือไม่
	Version string // เวอร์ชันที่ตรวจพบ (เช่น 20.11.1)
}

// Detect ค้นหารันไทม์ที่ใช้งานได้จากเครื่อง โดยเรียงลำดับความนิยม
// คืนค่าเป็นชื่อรันไทม์เช่น "node", "bun", "deno", "go" หรือ "unknown"
func Detect(ctx context.Context) string {
	// ลองตรวจจากคำสั่งที่นิยมใช้ในการพัฒนาเว็บ/แอปก่อน
	if has("node") {
		return "node"
	}
	if has("bun") {
		return "bun"
	}
	if has("deno") {
		return "deno"
	}
	if has("go") {
		return "go"
	}
	return "unknown"
}

// InspectAll ตรวจสอบรันไทม์ยอดนิยมและคืนผลลัพธ์ทั้งหมด
func InspectAll(ctx context.Context) []RuntimeStatus {
	names := []string{"node", "npm", "go", "python", "pip", "bun", "deno"}
	out := make([]RuntimeStatus, 0, len(names))
	for _, n := range names {
		out = append(out, CheckRuntime(ctx, n))
	}
	return out
}

// CheckRuntime ตรวจสอบรันไทม์เดี่ยว ๆ โดยพยายามเรียกคำสั่งเวอร์ชันและแยกเลขเวอร์ชัน
func CheckRuntime(ctx context.Context, name string) RuntimeStatus {
	bin, args := commandFor(name)
	if bin == "" {
		return RuntimeStatus{Name: name, Found: false}
	}
	// บางตัวมีตัวเลือกหลายไบนารี (เช่น python/python3, pip/pip3)
	bins := strings.Split(bin, "|")
	for _, candidate := range bins {
		if !has(strings.TrimSpace(candidate)) {
			continue
		}
		ver := runVersion(ctx, strings.TrimSpace(candidate), args)
		if ver != "" {
			return RuntimeStatus{Name: name, Found: true, Version: ver}
		}
		// แม้จะรันได้แต่แยกเวอร์ชันไม่ได้
		return RuntimeStatus{Name: name, Found: true, Version: "unknown"}
	}
	return RuntimeStatus{Name: name, Found: false}
}

// PrintReport แสดงผลสรุปการตรวจสอบด้วยสี และแนะนำวิธีติดตั้งเมื่อไม่พบ
func PrintReport(statuses []RuntimeStatus) {
	// แสดงตารางย่อ
	pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(1).Println("สรุประบบรันไทม์บนเครื่อง")
	for _, s := range statuses {
		if s.Found {
			pterm.Success.Printfln("พบ %s เวอร์ชัน %s", s.Name, s.Version)
		} else {
			pterm.Warning.Printfln("ไม่พบ %s บนเครื่องนี้", s.Name)
			for _, tip := range suggestInstall(s.Name) {
				pterm.Info.Println(tip)
			}
		}
	}
}

// ภายใน
func has(bin string) bool {
	_, err := exec.LookPath(bin)
	return err == nil
}

func commandFor(name string) (bin string, args []string) {
	switch strings.ToLower(name) {
	case "node":
		return "node", []string{"--version"}
	case "npm":
		return "npm", []string{"--version"}
	case "go":
		return "go", []string{"version"}
	case "python":
		// ลองทั้ง python และ python3
		return "python|python3", []string{"--version"}
	case "pip":
		return "pip|pip3", []string{"--version"}
	case "bun":
		return "bun", []string{"--version"}
	case "deno":
		return "deno", []string{"--version"}
	default:
		return "", nil
	}
}

func runVersion(ctx context.Context, bin string, args []string) string {
	cmd := exec.CommandContext(ctx, bin, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return ""
	}
	out := strings.TrimSpace(stdout.String())
	if out == "" {
		out = strings.TrimSpace(stderr.String())
	}
	return parseVersion(out)
}

var verRe = regexp.MustCompile(`\d+\.\d+(\.\d+)?`)

func parseVersion(s string) string {
	m := verRe.FindString(s)
	return m
}

func suggestInstall(name string) []string {
	os := runtime.GOOS
	n := strings.ToLower(name)
	tips := []string{"คำแนะนำการติดตั้ง (ตรวจสอบความเหมาะสมกับเครื่องของคุณก่อนรัน):"}
	switch n {
	case "node":
		if os == "windows" {
			tips = append(tips,
				"winget install OpenJS.NodeJS.LTS",
				"หรือ: choco install nodejs-lts",
				"หรือดาวน์โหลดจาก https://nodejs.org/")
		} else {
			tips = append(tips,
				"nvm (แนะนำ) หรือแพ็กเกจเมเนเจอร์ของระบบ",
				"ดูเพิ่มเติม: https://nodejs.org/")
		}
	case "npm":
		tips = append(tips, "ติดมากับ Node.js โดยปกติ — ติดตั้ง Node.js แล้วจะมี npm")
	case "go":
		if os == "windows" {
			tips = append(tips,
				"winget install GoLang.Go",
				"หรือ: choco install golang",
				"หรือดาวน์โหลดจาก https://go.dev/dl/")
		} else {
			tips = append(tips, "ดาวน์โหลดจาก https://go.dev/dl/ หรือใช้แพ็กเกจเมเนเจอร์ของระบบ")
		}
	case "python":
		if os == "windows" {
			tips = append(tips,
				"winget install Python.Python.3",
				"หรือ: choco install python",
				"หรือดาวน์โหลดจาก https://www.python.org/downloads/")
		} else {
			tips = append(tips, "ใช้แพ็กเกจเมเนเจอร์ของระบบ หรือ https://www.python.org/downloads/")
		}
	case "pip":
		tips = append(tips, "ติดมากับ Python หรือใช้ get-pip.py ตามเอกสารทางการ")
	case "bun":
		if os == "windows" {
			tips = append(tips, "ดูวิธีติดตั้ง Windows: https://bun.sh/docs/installation#windows")
		} else {
			tips = append(tips, "curl -fsSL https://bun.sh/install | bash (โปรดตรวจสอบสคริปต์ก่อนรัน)")
		}
	case "deno":
		if os == "windows" {
			tips = append(tips, "iwr https://deno.land/install.ps1 -useb | iex (โปรดตรวจสอบสคริปต์ก่อนรัน)")
		} else {
			tips = append(tips, "curl -fsSL https://deno.land/x/install/install.sh | sh (โปรดตรวจสอบสคริปต์ก่อนรัน)")
		}
	default:
		tips = append(tips, "ค้นหาวิธีติดตั้งจากเอกสารทางการ")
	}
	return tips
}
