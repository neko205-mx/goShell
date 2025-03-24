package goShell

import (
	"fmt"
	"testing"
)

func TestReverseShell(t *testing.T) {
	results := uploadInfo()
	for i, result := range results {
		fmt.Printf("\n=== 文件 %d ===\n", i+1)
		fmt.Println("路径:", result.Path)

		if result.Error != nil {
			fmt.Println("状态: 失败")
			fmt.Println("错误:", result.Error)
		} else {
			fmt.Println("状态: 成功")
			fmt.Printf("大小: %d bytes\n", len(result.Content))

			// 打印前 100 个字节（防止内容过大）
			preview := string(result.Content)
			if len(preview) > 100 {
				preview = preview[:100] + "...(截断)"
			}
			fmt.Println("内容预览:", preview)
		}
	}

}
