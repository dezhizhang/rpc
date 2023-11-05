package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// 文件的打开
func TestOpenFile(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		t.Log("open file err=", err)
	}

	t.Log(file)

	defer file.Close()
}

// 带缓冲读取文件
func TestReadFile(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		t.Logf("读取文件失败%s", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	// 循环读取内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //io.EOF表示文件的未尾
			break
		}
		t.Logf("读到的内容为：%s", str)
	}

	t.Log("文件读取结束")
}

// 不带缓冲读取文件
func TestIoReadFile(t *testing.T) {
	file, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		t.Logf("读取文件失败%s", err)
	}
	t.Logf(string(file))
}
