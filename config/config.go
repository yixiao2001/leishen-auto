package config

import (
	"fmt"
	"os"
	"strings"
)

// Config 配置结构体
type Config struct {
	AccountTokens []string // 支持多个账户令牌
	Lang          string
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	tokenStr := os.Getenv("TOKEN")
	if tokenStr == "" {
		return nil, fmt.Errorf("TOKEN 环境变量未设置")
	}

	// 支持多个token，以逗号分隔
	tokens := strings.Split(tokenStr, ",")
	var accountTokens []string
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token != "" {
			accountTokens = append(accountTokens, token)
		}
	}

	if len(accountTokens) == 0 {
		return nil, fmt.Errorf("TOKEN 环境变量未包含有效的令牌")
	}

	return &Config{
		AccountTokens: accountTokens,
		Lang:          "zh_CN",
	}, nil
}
