// file包主要包含一些常用文件操作。
package file

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// 按行读取文件，返回行字符串切片，返回结果去除换行符和空白行。
func ReadFile(name string) (out []string, err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, isPrefix, e := reader.ReadLine()
		if e != nil {
			if e != io.EOF {
				err = e
			}
			break
		}
		if isPrefix {
			err = errors.New("The line is too long.")
			break
		}
		lineStr := string(line)
		if lineStr = strings.TrimSpace(lineStr); len(lineStr) > 0 {
			out = append(out, lineStr)
		}
	}
	return
}

// 把字符串数组写入到文件中，不自动添加换行符。
func WriteFile(in []string, name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return
	}
	defer file.Close()
	for _, line := range in {
		file.WriteString(line)
	}
	return
}
