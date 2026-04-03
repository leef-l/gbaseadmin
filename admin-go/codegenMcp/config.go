package main

import (
	"fmt"
	"os"

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
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}
	// 支持环境变量语法 ${ENV_VAR}
	cfg.Database.Password = expandEnvVar(cfg.Database.Password)
	return &cfg, nil
}

// expandEnvVar 如果值形如 ${VAR_NAME}，则从环境变量读取替换
func expandEnvVar(val string) string {
	if len(val) > 3 && val[:2] == "${" && val[len(val)-1] == '}' {
		envName := val[2 : len(val)-1]
		if envVal := os.Getenv(envName); envVal != "" {
			return envVal
		}
	}
	return val
}

// DSN 生成标准数据库连接字符串
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
