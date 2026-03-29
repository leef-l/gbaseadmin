package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gbaseadmin/codegen/generator/backend"
	"gbaseadmin/codegen/generator/frontend"
	"gbaseadmin/codegen/parser"
)

func main() {
	// 命令行参数
	var (
		table  string // 表名，逗号分隔
		only   string // backend | frontend | 空=都生成
		force  bool   // 强制覆盖
		config string // 配置文件路径
		dryRun bool   // 只打印不写入
	)

	flag.StringVar(&table, "table", "", "要生成的表名，多个用逗号分隔 (required)")
	flag.StringVar(&only, "only", "", "只生成指定端: backend | frontend")
	flag.BoolVar(&force, "force", false, "强制覆盖已存在文件")
	flag.StringVar(&config, "config", "./codegen.yaml", "配置文件路径")
	flag.BoolVar(&dryRun, "dry-run", false, "只打印将生成的文件列表")
	flag.Parse()

	if table == "" {
		fmt.Println("错误: --table 参数不能为空")
		flag.Usage()
		os.Exit(1)
	}

	// 加载配置
	cfg, err := LoadConfig(config)
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		os.Exit(1)
	}

	// 创建解析器
	p := parser.New(cfg.Database.DSN())

	// 解析表名列表
	tableNames := strings.Split(table, ",")
	for i := range tableNames {
		tableNames[i] = strings.TrimSpace(tableNames[i])
	}

	start := time.Now()
	totalFiles := 0

	// 获取当前工作目录（用于计算模板路径）
	cwd, _ := os.Getwd()
	templateDir := filepath.Join(cwd, "templates")

	for _, tableName := range tableNames {
		fmt.Printf("\n[codegen] 开始生成表: %s\n", tableName)

		// 解析表结构
		meta, err := p.ParseTable(tableName)
		if err != nil {
			fmt.Printf("[codegen] ✗ 解析表 %s 失败: %v\n", tableName, err)
			continue
		}

		var files []string

		// 生成后端代码
		if only != "frontend" {
			backendGen := backend.New(backend.Config{
				TemplateDir: filepath.Join(templateDir, "backend"),
				OutputDir:   cfg.Backend.Output,
				Force:       force,
			})
			if dryRun {
				fmt.Println("[codegen] [dry-run] 后端文件将生成到:", cfg.Backend.Output)
			} else {
				generated, err := backendGen.Generate(meta)
				if err != nil {
					fmt.Printf("[codegen] ✗ 后端生成失败: %v\n", err)
				} else {
					for _, f := range generated {
						fmt.Printf("[codegen] ✓ 后端: %s\n", f)
					}
					files = append(files, generated...)
				}
			}
		}

		// 生成前端代码
		if only != "backend" {
			frontendGen := frontend.New(frontend.Config{
				TemplateDir: filepath.Join(templateDir, "frontend"),
				OutputDir:   cfg.Frontend.Output,
				Force:       force,
			})
			if dryRun {
				fmt.Println("[codegen] [dry-run] 前端文件将生成到:", cfg.Frontend.Output)
			} else {
				generated, err := frontendGen.Generate(meta)
				if err != nil {
					fmt.Printf("[codegen] ✗ 前端生成失败: %v\n", err)
				} else {
					for _, f := range generated {
						fmt.Printf("[codegen] ✓ 前端: %s\n", f)
					}
					files = append(files, generated...)
				}
			}
		}

		fmt.Printf("[codegen] 表 %s 生成完成，共 %d 个文件\n", tableName, len(files))
		totalFiles += len(files)
	}

	elapsed := time.Since(start)
	fmt.Printf("\n[codegen] 全部完成！共生成 %d 个文件，耗时 %.1fs\n", totalFiles, elapsed.Seconds())
}