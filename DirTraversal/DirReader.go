package main

/**
 * 读取文件夹
 *
 * @author itaken <regelhh@gmail.com>
 * @since 2014-6-13
 */

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
GetFileList 获取path目录下的文件列表
*/
func GetFileList(path string) (slice []string) {
	var files []string // slice
	if path == "" {
		return files
	}
	_, err := os.Stat(path)               // 判断目录状态
	if err != nil && os.IsNotExist(err) { // 目录不存在
		return files
	}
	warkErr := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() { // 是否目录
			return nil
		}
		files = append(files, path) // 追加到slice
		return nil
	})
	if warkErr != nil {
		panic(warkErr)
	}
	return files
}

func main() {
	root := "./path/to/dir" // 需要遍历的路径
	fmt.Println(root)
	res := GetFileList(root) // 获取目录下的文件列表
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
