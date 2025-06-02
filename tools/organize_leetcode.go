package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// go run .\tools\organize_leetcode.go .\leetcode
	if len(os.Args) < 2 {
		fmt.Println("请指定文件夹路径作为参数")
		return
	}

	sourceDir := os.Args[1]
	parentDir := filepath.Dir(sourceDir)

	files, err := os.ReadDir(sourceDir)
	if err != nil {
		fmt.Printf("读取目录失败: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()
		// 获取下划线前的数字部分
		parts := strings.SplitN(filename, "_", 2)
		if len(parts) < 2 {
			fmt.Printf("跳过不符合格式的文件: %s\n", filename)
			continue
		}

		numStr := parts[0]
		nums, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("无效的数字格式: %s (%v)\n", numStr, err)
			continue
		}

		// 创建目标目录
		targetDir := filepath.Join(parentDir, fmt.Sprintf("leetcode%d", nums))
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			fmt.Printf("创建目录失败: %s (%v)\n", targetDir, err)
			continue
		}

		// 处理文件名冲突
		targetPath := filepath.Join(targetDir, filename)
		counter := 1
		ext := filepath.Ext(filename)
		nameWithoutExt := filename[:len(filename)-len(ext)]
		for {
			if _, err := os.Stat(targetPath); os.IsNotExist(err) {
				break
			}
			newFilename := fmt.Sprintf("%s_%d%s", nameWithoutExt, counter, ext)
			targetPath = filepath.Join(targetDir, newFilename)
			counter++
		}

		// 复制文件
		if err := copyFile(
			filepath.Join(sourceDir, filename),
			targetPath,
		); err != nil {
			fmt.Printf("复制文件失败: %s -> %s (%v)\n", filename, targetPath, err)
		} else {
			fmt.Printf("成功复制: %s -> %s\n", filename, targetPath)
		}
	}
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
