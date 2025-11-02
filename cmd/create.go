package cmd

import (
	"context"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"

	"projgen/internal/generator"
	"projgen/internal/ui"
)

// createCmd defines the "create" subcommand which triggers the interactive setup wizard.
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project via an interactive wizard",
	Long:  "Launches an interactive prompt to choose language, framework, runtime, and features, then scaffolds the project.",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Use the command's context for cancellation and deadlines if provided.
		var ctx context.Context = cmd.Context()

		// 1) Trigger interactive prompt sequence in the UI layer.
		choices, err := ui.RunWizard(ctx)
		if err != nil {
			// แสดงผลแบบเป็นมิตรและออกอย่างนุ่มนวล
			if strings.Contains(err.Error(), "ยกเลิกโดยผู้ใช้") {
				pterm.Warning.Println("ยกเลิกการสร้างโปรเจ็กต์")
				return nil
			}
			pterm.Error.Printfln("เกิดข้อผิดพลาดระหว่างขั้นตอนวิซาร์ด: %v", err)
			return err
		}

		// 2) Pass collected data to the generator to scaffold the project.
		if err := generator.Generate(ctx, choices); err != nil {
			pterm.Error.Printfln("สร้างโปรเจ็กต์ไม่สำเร็จ: %v", err)
			return err
		}

		return nil
	},
}

func init() {
	// Register the create subcommand under the root command.
	rootCmd.AddCommand(createCmd)
}

