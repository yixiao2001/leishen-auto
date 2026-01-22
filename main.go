package main

import (
	"fmt"
	"os"

	"leishen-auto/api"
	"leishen-auto/config"
)

func main() {
	fmt.Println("⌛️开始运行")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌错误: %v\n", err)
		os.Exit(1)
	}

	client := api.NewClient()

	totalAccounts := len(cfg.AccountTokens)
	successCount := 0
	failCount := 0

	fmt.Printf("📋共有 %d 个账户需要暂停\n", totalAccounts)

	for i, token := range cfg.AccountTokens {
		fmt.Printf("\n🔄正在处理账户 %d/%d...\n", i+1, totalAccounts)

		resp, err := client.Pause(token, cfg.Lang)
		if err != nil {
			fmt.Printf("❌账户 %d 暂停失败: %v\n", i+1, err)
			failCount++
			continue
		}

		if resp.Code != 0 {
			if resp.Code == 400803 { // 400803 - 账号已经停止加速，请不要重复操作
				fmt.Printf("👌账户 %d 已经暂停: %d - %s\n", i+1, resp.Code, resp.Msg)
				successCount++
				continue
			}
			fmt.Printf("❌账户 %d 暂停失败: %d - %s\n", i+1, resp.Code, resp.Msg)
			failCount++
			continue
		}

		fmt.Printf("✔️账户 %d 暂停成功: %d - %s\n", i+1, resp.Code, resp.Msg)
		successCount++
	}

	fmt.Printf("\n📊处理完成: 成功 %d 个, 失败 %d 个\n", successCount, failCount)
	fmt.Println("⌛️结束运行")

	if failCount > 0 {
		os.Exit(1)
	}
}
