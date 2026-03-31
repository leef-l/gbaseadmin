package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type BackendConfig struct {
	Output string `yaml:"output"`
}

type FrontendConfig struct {
	Output string `yaml:"output"`
}

type Config struct {
	Database   DatabaseConfig `yaml:"database"`
	Backend    BackendConfig  `yaml:"backend"`
	Frontend   FrontendConfig `yaml:"frontend"`
	SkipFields []string       `yaml:"skip_fields"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}
	return &cfg, nil
}

// DSN 生成数据库连接字符串
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}

// DSNForHack 生成 hack/config.yaml 中 gf gen dao 使用的 link 格式
// 格式: mysql:user:password@tcp(host:port)/dbname
func (c *DatabaseConfig) DSNForHack() string {
	return fmt.Sprintf("mysql:%s:%s@tcp(%s:%d)/%s",
		c.User, c.Password, c.Host, c.Port, c.DBName)
}
