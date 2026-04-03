package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// DatabaseConfig 数据库连接配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

// BackendConfig 后端输出配置
type BackendConfig struct {
	Output string `yaml:"output"`
}

// FrontendConfig 前端输出配置
type FrontendConfig struct {
	Output string `yaml:"output"`
}

// MenuAppConfig 菜单应用目录配置
type MenuAppConfig struct {
	Title string `yaml:"title"`
	Icon  string `yaml:"icon"`
	Sort  int    `yaml:"sort"`
}

// MenuModuleConfig 模块级菜单配置
type MenuModuleConfig struct {
	Sort   int  `yaml:"sort"`
	IsShow *int `yaml:"is_show"`
}

// PostGenConfig 后置生成配置。
// MCP 模式默认不执行任何外部命令，这两个字段默认值均为 false。
type PostGenConfig struct {
	RunGfGenDao bool `yaml:"run_gf_gen_dao"`
	RunGfInit   bool `yaml:"run_gf_init"`
}

// Config 完整配置
type Config struct {
	Database     DatabaseConfig              `yaml:"database"`
	Backend      BackendConfig               `yaml:"backend"`
	Frontend     FrontendConfig              `yaml:"frontend"`
	SkipFields   []string                    `yaml:"skip_fields"`
	MenuApps     map[string]MenuAppConfig    `yaml:"menu_apps"`
	MenuModules  map[string]MenuModuleConfig `yaml:"menu_modules"` // key: "appName/moduleName"
	OperationLog bool                        `yaml:"operation_log"`
	PostGen      PostGenConfig               `yaml:"post_generation"`
}

// LoadConfig 从 YAML 文件加载配置
func LoadConfig(path string) (*Config, error) {
	loadEnvFileIfExists(filepath.Join(filepath.Dir(path), "..", ".env"))

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	expanded := expandEnvPlaceholders(string(data))
	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}
	return &cfg, nil
}

var envPattern = regexp.MustCompile(`\$\{([A-Za-z_][A-Za-z0-9_]*)\}`)

func expandEnvPlaceholders(input string) string {
	return envPattern.ReplaceAllStringFunc(input, func(match string) string {
		sub := envPattern.FindStringSubmatch(match)
		if len(sub) != 2 {
			return match
		}
		if value, ok := os.LookupEnv(sub[1]); ok {
			return value
		}
		return match
	})
}

func loadEnvFileIfExists(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}
		if _, exists := os.LookupEnv(key); exists {
			continue
		}
		value = strings.TrimSpace(value)
		value = strings.Trim(value, `"'`)
		_ = os.Setenv(key, value)
	}
}

// DSN 生成标准数据库连接字符串
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
