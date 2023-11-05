package file

import (
	"bufio"
	"io"
	"os"
	"testing"
)

func CopyFile(srcPath, targetPath string) (written int64, err error) {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return 0, err
	}
	reader := bufio.NewReader(srcFile)

	// 打开文件
	targetFile, err1 := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err1 != nil {
		return 0, err1
	}

	writer := bufio.NewWriter(targetFile)
	defer srcFile.Close()
	defer targetFile.Close()

	return io.Copy(writer, reader)
}

func TestCopyFile(t *testing.T) {
	srcPath := "./5.png"
	targetPath := "./test1/5.png"
	_, err := CopyFile(srcPath, targetPath)
	if err != nil {
		t.Log(err)
	}

}
