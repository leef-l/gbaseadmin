package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gbaseadmin/codegenMcp/generator/backend"
	"gbaseadmin/codegenMcp/generator/frontend"
	"gbaseadmin/codegenMcp/generator/menu"
	"gbaseadmin/codegenMcp/generator/util"
	"gbaseadmin/codegenMcp/parser"
)

type DBConfig struct {
	DSN string
}

// EngineConfig 引擎配置
type EngineConfig struct {
	TemplateDir string
	BackendOut  string
	FrontendOut string
	SkipFields  []string
	EnableOpLog bool
	MenuApps    map[string]menu.MenuAppConfig
	MenuModules map[string]menu.MenuModuleConfig
	Force       bool
	RunPostGen  bool
}

type GenerateResult struct {
	TableName string        `json:"table_name"`
	Files     []string      `json:"files"`
	MenuCount int           `json:"menu_count"`
	Elapsed   time.Duration `json:"-"`
	Errors    []string      `json:"errors,omitempty"`
}

type PreviewResult struct {
	TableName string            `json:"table_name"`
	Files     map[string][]byte `json:"files"`
	Errors    []string          `json:"errors,omitempty"`
}

type InspectResult struct {
	TableName   string         `json:"table_name"`
	AppName     string         `json:"app_name"`
	ModuleName  string         `json:"module_name"`
	Comment     string         `json:"comment"`
	FieldCount  int            `json:"field_count"`
	Fields      []FieldSummary `json:"fields"`
	HasParentID bool           `json:"has_parent_id"`
	HasStatus   bool           `json:"has_status"`
	HasSort     bool           `json:"has_sort"`
}

type FieldSummary struct {
	Name       string   `json:"name"`
	GoType     string   `json:"go_type"`
	TSType     string   `json:"ts_type"`
	Component  string   `json:"component"`
	Label      string   `json:"label"`
	IsHidden   bool     `json:"is_hidden"`
	IsEnum     bool     `json:"is_enum"`
	IsRequired bool     `json:"is_required"`
	EnumValues []string `json:"enum_values,omitempty"`
}

// Engine 代码生成引擎，封装 parser + backend/frontend/menu generator
type Engine struct {
	cfg      EngineConfig
	dsn      string
	parser   *parser.Parser
	tplCache *util.TemplateCache
}

// New 创建引擎实例并初始化数据库连接
func New(dsn string, cfg EngineConfig) (*Engine, error) {
	p, err := parser.New(dsn, cfg.SkipFields)
	if err != nil {
		return nil, fmt.Errorf("初始化解析器失败: %w", err)
	}
	return &Engine{
		cfg:      cfg,
		dsn:      dsn,
		parser:   p,
		tplCache: util.NewTemplateCache(),
	}, nil
}

// Close 释放数据库连接
func (e *Engine) Close() {
	e.parser.Close()
}

// SetForce 设置是否强制覆盖
func (e *Engine) SetForce(force bool) {
	e.cfg.Force = force
}

// InspectTable 解析表结构，返回简化元数据摘要
func (e *Engine) InspectTable(tableName string) (*InspectResult, error) {
	meta, err := e.parser.ParseTable(tableName)
	if err != nil {
		return nil, err
	}

	result := &InspectResult{
		TableName:   meta.TableName,
		AppName:     meta.AppName,
		ModuleName:  meta.ModuleName,
		Comment:     meta.Comment,
		FieldCount:  len(meta.Fields),
		HasParentID: meta.HasParentID,
		HasStatus:   meta.HasStatus,
		HasSort:     meta.HasSort,
	}

	for _, f := range meta.Fields {
		fs := FieldSummary{
			Name:       f.Name,
			GoType:     f.GoType,
			TSType:     f.TSType,
			Component:  f.Component,
			Label:      f.Label,
			IsHidden:   f.IsHidden,
			IsEnum:     f.IsEnum,
			IsRequired: f.IsRequired,
		}
		for _, ev := range f.EnumValues {
			fs.EnumValues = append(fs.EnumValues, fmt.Sprintf("%s=%s", ev.Value, ev.Label))
		}
		result.Fields = append(result.Fields, fs)
	}

	return result, nil
}

// PreviewGenerate 预览将生成的文件内容，不写磁盘
func (e *Engine) PreviewGenerate(tableName string, only string) (*PreviewResult, error) {
	meta, err := e.parser.ParseTable(tableName)
	if err != nil {
		return nil, err
	}
	meta.EnableOpLog = e.cfg.EnableOpLog

	result := &PreviewResult{
		TableName: tableName,
		Files:     make(map[string][]byte),
	}

	if only != "frontend" && only != "menu" {
		gen := backend.New(backend.Config{
			TemplateDir: filepath.Join(e.cfg.TemplateDir, "backend"),
			OutputDir:   filepath.Join(e.cfg.BackendOut, meta.AppName),
			Force:       e.cfg.Force,
			Cache:       e.tplCache,
		})
		files, err := gen.GenerateToMemory(meta)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("backend: %v", err))
		} else {
			for k, v := range files {
				result.Files[k] = v
			}
		}
	}

	if only != "backend" && only != "menu" {
		gen := frontend.New(frontend.Config{
			TemplateDir: filepath.Join(e.cfg.TemplateDir, "frontend"),
			OutputDir:   e.cfg.FrontendOut,
			Force:       e.cfg.Force,
			Cache:       e.tplCache,
		})
		files, err := gen.GenerateToMemory(meta)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("frontend: %v", err))
		} else {
			for k, v := range files {
				result.Files[k] = v
			}
		}
	}

	return result, nil
}

// GenerateCode 生成代码到磁盘
func (e *Engine) GenerateCode(tableName string, only string, withMenu bool) (*GenerateResult, error) {
	start := time.Now()

	meta, err := e.parser.ParseTable(tableName)
	if err != nil {
		return nil, err
	}
	meta.EnableOpLog = e.cfg.EnableOpLog

	result := &GenerateResult{TableName: tableName}

	if only != "frontend" && only != "menu" {
		appDir := filepath.Join(e.cfg.BackendOut, meta.AppName)
		if err := os.MkdirAll(appDir, 0o755); err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("创建应用目录失败: %v", err))
		}
	}

	if only != "frontend" && only != "menu" {
		gen := backend.New(backend.Config{
			TemplateDir: filepath.Join(e.cfg.TemplateDir, "backend"),
			OutputDir:   filepath.Join(e.cfg.BackendOut, meta.AppName),
			Force:       e.cfg.Force,
			Cache:       e.tplCache,
		})
		files, err := gen.Generate(meta)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("backend: %v", err))
		} else {
			result.Files = append(result.Files, files...)
		}
	}

	if only != "backend" && only != "menu" {
		gen := frontend.New(frontend.Config{
			TemplateDir: filepath.Join(e.cfg.TemplateDir, "frontend"),
			OutputDir:   e.cfg.FrontendOut,
			Force:       e.cfg.Force,
			Cache:       e.tplCache,
		})
		files, err := gen.Generate(meta)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("frontend: %v", err))
		} else {
			result.Files = append(result.Files, files...)
		}
	}

	if only == "menu" || withMenu {
		menuGen := menu.New(menu.Config{
			DSN:         e.dsn,
			Force:       e.cfg.Force,
			MenuApps:    e.cfg.MenuApps,
			MenuModules: e.cfg.MenuModules,
		})
		count, err := menuGen.Generate(meta)
		if err != nil {
			result.Errors = append(result.Errors, fmt.Sprintf("menu: %v", err))
		} else {
			result.MenuCount = count
		}
	}

	result.Elapsed = time.Since(start)
	return result, nil
}
