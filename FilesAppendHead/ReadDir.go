package main

/**
 * 批量给文件添加头部内容
 *
 * @author itaken <regelhh@gmail.com>
 * @since 2017-08-17
 */

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

/*
HasTitle 是否有标题
*/
func HasTitle(file string) (bool, string, string) {
	fi, err := os.Open(file) // 打开文件
	defer fi.Close()         //关闭文件
	if err != nil {
		panic(err)
	}
	filename := path.Base(fi.Name())      // 文件名
	filecontents, _ := ioutil.ReadAll(fi) // 读取所有内容
	contents := string(filecontents)
	if contents == "" {
		return false, filename, contents
	}
	idx := strings.Index(contents, "\n")         // 第一个换行符位置
	if strings.Contains(contents[:idx], "---") { // 第一行 包含 ---
		return true, filename, contents
	}
	return false, filename, contents
}

/*
RewriteFile 覆盖写入文件
*/
func RewriteFile(file, contents string) bool {
	if file == "" || contents == "" {
		return false
	}
	newfi, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0600) // 打开文件
	// newfi, err := os.OpenFile(file, os.O_RDWR, 0666)
	defer newfi.Close()
	if err != nil {
		panic(err)
	}
	// newfi.Seek(0, os.SEEK_SET) // 指针归零
	// num, err := newfi.WriteAt([]byte(contents), 0) // 在开头覆盖插入内容
	num, err := newfi.WriteString(contents) // 写入文件
	if err != nil || num < 1 {
		return false
	}
	return true
}

func main() {
	dir := "./files/" // 遍历的目录
	// fmt.Println(dir)
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() { // 目录,则略过
			return nil
		}
		has, filename, contents := HasTitle(path) // 判断文件是否需要处理
		if has == true {
			return nil
		}
		fmt.Println(has, filename, contents)
		// 文件标题
		name := strings.Split(filename, ".")                      // 分割文件名称
		nowtime := time.Now().Format("2006-01-02 15:04:05 -0700") // 格式化时间
		head := "---\ntitle: %s \ndate: %s\ncategories: linux\n---\n\n"
		head = fmt.Sprintf(head, name[0], nowtime)

		newcontents := head + contents // 新文件内容
		RewriteFile(path, newcontents) // 写入文件
		return nil
	})

}
