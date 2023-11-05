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

// 文件不存在则创建文件存在则更新
func TestWriteExitFile(t *testing.T) {
	filePath := "./abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		t.Logf("打开文件失败%s", err)
		return
	}

	defer file.Close()

	str := "您好晓智科技有限会司\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}

// 向文件中追加内容
func TestWriteAppend(t *testing.T) {
	filePath := "./abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Logf("打开文件失败%s", err)
	}
	defer file.Close()

	str := "abc hello world"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}

// 先读文件后写入文件
func TestReadAndWriteFile(t *testing.T) {
	filePath := "./abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		t.Logf("读取文件失败%s", err)
		return
	}

	// 关闭文件
	defer file.Close()
	// 读取文件
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { // 读取到文件未尾
			break
		}
		// 量示到终端
		t.Log(str)
	}

	str := "hello 北京你好\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
