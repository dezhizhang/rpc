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

// 写入文件
func TestWriteFile(t *testing.T) {
	file, err := os.OpenFile("./abc.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		t.Logf("打开文件失败%s", err)
		return
	}
	str := "hello, Gardon\n"

	defer file.Close()

	writer := bufio.NewWriter(file)

	// 写入时, 使用带缓存的*writer
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 带缓存因此调用WriteString方法时
	// 内容是写入到缓存的，因些将缓存中的数据写入磁盘
	writer.Flush()

}
