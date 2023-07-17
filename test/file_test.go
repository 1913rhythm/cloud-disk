package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)

// 指定分片的大小
const chunkSize = 10 * 1024 * 1024 // 10MB

// 1.文件的分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("./book.pdf")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数 = 文件大小/分片大小
	chunkNum := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))
	myFile, err := os.OpenFile("./book.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b := make([]byte, chunkSize)
	for i := 0; i < chunkNum; i++ {
		// 指定读取文件的起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		// 如果要读取的剩余文件大小 < chunkSize
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		myFile.Read(b)

		// 存放分片的文件
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 2.分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("./book2.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	fileInfo, err := os.Stat("./book.pdf")
	if err != nil {
		t.Fatal(err)
	}
	// 分片的个数 = 文件大小/分片大小
	chunkNum := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))
	for i := 0; i < chunkNum; i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myFile.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 3.文件一致性校验
func TestCheck(t *testing.T) {
	// 获取原文件的信息
	file1, err := os.OpenFile("./book.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b1, err := ioutil.ReadAll(file1)
	if err != nil {
		t.Fatal(err)
	}
	// 获取合并后的文件信息
	file2, err := os.OpenFile("./book2.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadAll(file2)
	if err != nil {
		t.Fatal(err)
	}
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1)
	fmt.Println(s2)
}
