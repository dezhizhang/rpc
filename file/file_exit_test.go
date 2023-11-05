package file

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadAndWriteFile1(t *testing.T) {
	filePath := "./test.txt"
	targetPath := "./test1/abc.txt"

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		t.Logf("读取文件失败%s", err)
		return
	}

	err = ioutil.WriteFile(targetPath, file, 0666)
	if err != nil {
		t.Logf("写文件失败%s", err)
		return
	}
	t.Log("文件写入成功")
}

// 判断文件是否存在
func TestFileIsExit(t *testing.T) {
	stat, err := os.Stat("./abc1.txt")
	if err == nil {
		t.Logf("文件或目录已存在")
		return
	}
	if os.IsNotExist(err) {
		t.Logf("文件不存在%s", err)
	}
	t.Log(stat)
}
