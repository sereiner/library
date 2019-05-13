package file

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CreateFile 根据文件路径(相对或绝对路径)创建文件，如果文件所在的文件夹不存在则自动创建
// path: 文件路径
func CreateFile(path string) (f *os.File, err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}
	dir := filepath.Dir(absPath)
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return
	}
	return os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

// GetAbs 获取文件绝对路径
// path: 文件路径
func GetAbs(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// AppendToFile 向文件尾部追加内容
// fileName: 文件名字(带全路径)
// content: 写入的内容
func AppendToFile(fileName string, content string) error {

	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// 查找文件末尾的偏移量
	n, _ := f.Seek(0, os.SEEK_END)
	// 从末尾的偏移量开始写入内容
	_, err = f.WriteAt([]byte(content), n)

	defer f.Close()
	return err
}

// AppendToFileHead  向文件头部追加内容
// file: 文件路径
// head: 写入的头部内容
func AppendToFileHead(file, head string) error {
	if file == "" {
		return errors.New("file is nil")
	}
	fi, err := os.Open(file)
	defer fi.Close()
	if err != nil {
		return err
	}
	contents, _ := ioutil.ReadAll(fi)
	newcontents := head + string(contents)
	newfi, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0600)

	defer newfi.Close()
	if err != nil {
		return err
	}

	num, err := newfi.WriteString(newcontents) // 写入文件
	if err != nil || num < 1 {
		return fmt.Errorf("写入文件失败 err:%v", err)
	}
	return nil
}
