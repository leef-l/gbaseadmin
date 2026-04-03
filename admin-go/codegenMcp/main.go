package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"gbaseadmin/codegenMcp/generator/menu"
	"gbaseadmin/codegenMcp/service"
)

func main() {
	configPath := "./codegen.yaml"
	if p := os.Getenv("CODEGEN_CONFIG"); p != "" {
		configPath = p
	}
	for i, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--config=") {
			configPath = strings.TrimPrefix(arg, "--config=")
		} else if arg == "--config" && i+2 < len(os.Args) {
			configPath = os.Args[i+2]
		}
	}

	cfg, err := LoadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[codegenMcp] 加载配置失败: %v\n", err)
		os.Exit(1)
	}

	execDir := executableDir()
	templateDir := filepath.Join(execDir, "templates")
	if _, err := os.Stat(templateDir); os.IsNotExist(err) {
		cwd, _ := os.Getwd()
		templateDir = filepath.Join(cwd, "templates")
	}

	menuApps := make(map[string]menu.MenuAppConfig, len(cfg.MenuApps))
	for k, v := range cfg.MenuApps {
		menuApps[k] = menu.MenuAppConfig{Title: v.Title, Icon: v.Icon, Sort: v.Sort}
	}
	menuModules := make(map[string]menu.MenuModuleConfig, len(cfg.MenuModules))
	for k, v := range cfg.MenuModules {
		menuModules[k] = menu.MenuModuleConfig{Sort: v.Sort, IsShow: v.IsShow}
	}

	engCfg := service.EngineConfig{
		TemplateDir: templateDir,
		BackendOut:  cfg.Backend.Output,
		FrontendOut: cfg.Frontend.Output,
		SkipFields:  cfg.SkipFields,
		EnableOpLog: cfg.OperationLog,
		MenuApps:    menuApps,
		MenuModules: menuModules,
	}

	eng, err := service.New(cfg.Database.DSN(), engCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[codegenMcp] 初始化引擎失败: %v\n", err)
		os.Exit(1)
	}
	defer eng.Close()

	s := server.NewMCPServer(
		"codegenMcp",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	inspectTool := mcp.NewTool("inspect_table",
		mcp.WithDescription("解析数据库表结构，返回字段元数据摘要（字段名、类型、组件、枚举等）"),
		mcp.WithString("table",
			mcp.Required(),
			mcp.Description("表名，如 system_dept"),
		),
	)
	s.AddTool(inspectTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tableName := req.GetString("table", "")
		if tableName == "" {
			return mcp.NewToolResultError("参数 table 不能为空"), nil
		}
		result, err := eng.InspectTable(tableName)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		b, _ := json.MarshalIndent(result, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	previewTool := mcp.NewTool("preview_generate",
		mcp.WithDescription("预览某张表将会生成的文件列表及内容，不实际写入磁盘"),
		mcp.WithString("table",
			mcp.Required(),
			mcp.Description("表名，如 system_dept"),
		),
		mcp.WithString("only",
			mcp.Description("只预览指定端: backend | frontend | menu（不传则全部预览）"),
		),
	)
	s.AddTool(previewTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tableName := req.GetString("table", "")
		if tableName == "" {
			return mcp.NewToolResultError("参数 table 不能为空"), nil
		}
		only := strings.ToLower(req.GetString("only", ""))
		if only != "" && only != "backend" && only != "frontend" {
			if only != "menu" {
				return mcp.NewToolResultError("参数 only 只能是 backend、frontend 或 menu"), nil
			}
		}

		result, err := eng.PreviewGenerate(tableName, only)
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}

		type fileInfo struct {
			Path  string `json:"path"`
			Bytes int    `json:"bytes"`
		}
		var summary struct {
			TableName string     `json:"table_name"`
			FileCount int        `json:"file_count"`
			Files     []fileInfo `json:"files"`
			Errors    []string   `json:"errors,omitempty"`
		}
		summary.TableName = result.TableName
		summary.Errors = result.Errors
		for path, content := range result.Files {
			summary.Files = append(summary.Files, fileInfo{Path: path, Bytes: len(content)})
		}
		summary.FileCount = len(summary.Files)
		b, _ := json.MarshalIndent(summary, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	generateTool := mcp.NewTool("generate_code",
		mcp.WithDescription("生成指定表的代码文件（后端/前端/菜单），实际写入磁盘"),
		mcp.WithString("table",
			mcp.Required(),
			mcp.Description("表名，多个用逗号分隔，如 system_dept 或 system_dept,system_role"),
		),
		mcp.WithString("only",
			mcp.Description("只生成指定端: backend | frontend | menu（不传则全部生成）"),
		),
		mcp.WithBoolean("force",
			mcp.Description("强制覆盖已存在文件（默认 false）"),
		),
		mcp.WithBoolean("with_menu",
			mcp.Description("同时生成菜单数据到数据库（默认 false）"),
		),
	)
	s.AddTool(generateTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tableArg := req.GetString("table", "")
		if tableArg == "" {
			return mcp.NewToolResultError("参数 table 不能为空"), nil
		}
		only := strings.ToLower(req.GetString("only", ""))
		if only != "" && only != "backend" && only != "frontend" && only != "menu" {
			return mcp.NewToolResultError("参数 only 只能是 backend、frontend 或 menu"), nil
		}

		force := req.GetBool("force", false)
		withMenu := req.GetBool("with_menu", false)
		eng.SetForce(force)

		tables := strings.Split(tableArg, ",")
		var allResults []*service.GenerateResult
		for _, t := range tables {
			t = strings.TrimSpace(t)
			if t == "" {
				continue
			}
			result, err := eng.GenerateCode(t, only, withMenu)
			if err != nil {
				result = &service.GenerateResult{
					TableName: t,
					Errors:    []string{err.Error()},
				}
			}
			allResults = append(allResults, result)
		}

		type summary struct {
			TableName string   `json:"table_name"`
			Files     []string `json:"files"`
			MenuCount int      `json:"menu_count,omitempty"`
			ElapsedMs int64    `json:"elapsed_ms"`
			Errors    []string `json:"errors,omitempty"`
		}
		var response struct {
			Results    []summary `json:"results"`
			TotalFiles int       `json:"total_files"`
			Note       string    `json:"note"`
		}
		response.Note = "不会执行 gf init / gf gen dao 等外部命令"
		for _, r := range allResults {
			response.TotalFiles += len(r.Files) + r.MenuCount
			response.Results = append(response.Results, summary{
				TableName: r.TableName,
				Files:     r.Files,
				MenuCount: r.MenuCount,
				ElapsedMs: r.Elapsed.Milliseconds(),
				Errors:    r.Errors,
			})
		}

		b, _ := json.MarshalIndent(response, "", "  ")
		return mcp.NewToolResultText(string(b)), nil
	})

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "[codegenMcp] MCP server 异常退出: %v\n", err)
		os.Exit(1)
	}
}

func executableDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}
