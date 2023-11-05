package file

import (
	"bufio"
	"io"
	"os"
	"testing"
)

type CharCount struct {
	ChCount    int // 计录英文个数
	NumCount   int // 记录数字个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 计录其它的个数
}

func TestCount(t *testing.T) {
	filePath := "./abc.txt"
	file, err := os.Open(filePath)
	if err != nil {
		t.Logf("打开文件失败%s", err)
	}

	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v > 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	t.Log(count)
}
