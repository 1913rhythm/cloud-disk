package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestFileUploadByFilepath(t *testing.T) {
	u, _ := url.Parse("https://merhythm-1318328416.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/exampleobject.jpg"

	_, _, err := client.Object.Upload(context.Background(), key, "./img/OIP-C.jpg", nil)
	if err != nil {
		panic(err)
	}
}

func TestFileUploadByReader(t *testing.T) {
	u, _ := url.Parse("https://merhythm-1318328416.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	// 文件夹名称
	key := "cloud-disk/exampleobject01.jpg"
	f, err := os.ReadFile("./img/OIP-C.jpg")
	if err != nil {
		return
	}
	// 传递大小为0的输入流
	_, err = client.Object.Put(context.Background(), key, bytes.NewReader(f), nil)
	if err != nil {
		// ERROR
		return
	}
}

// 分片上传初始化
func TestInitPartUpload(t *testing.T) {
	u, _ := url.Parse("https://merhythm-1318328416.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})
	key := "cloud-disk/example01.pdf"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		t.Fatal(err)
	}
	UploadID := v.UploadID
	fmt.Println(UploadID) // 16874252953c1552d0e48af4ef2f26f04ec9fd0e9f855a32cf67160c9a310c71f0d5f545ed
}

// 分片上传
func TestPartUpload(t *testing.T) {
	u, _ := url.Parse("https://merhythm-1318328416.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	// 文件夹名称
	key := "cloud-disk/example01.pdf"
	UploadID := "16874252953c1552d0e48af4ef2f26f04ec9fd0e9f855a32cf67160c9a310c71f0d5f545ed"

	//f, err := os.ReadFile("0.chunk") // "5711dc05e620528080fc6e0b4e27d47a"
	//f, err := os.ReadFile("1.chunk") // "635cd33489cf11cbcd5b545a3935fab6"
	f, err := os.ReadFile("2.chunk") // "3979027e120eff7ce30b0610166745a0"
	if err != nil {
		t.Fatal(err)
	}

	// opt 可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 3, bytes.NewReader(f), nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	PartETag := resp.Header.Get("ETag") // 0.chunk 的md5值：5711dc05e620528080fc6e0b4e27d47a
	fmt.Println(PartETag)
}

// 分片上传完成
func TestPartUploadComplete(t *testing.T) {
	u, _ := url.Parse("https://merhythm-1318328416.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSecretID,
			SecretKey: define.TencentSecretKey,
		},
	})

	key := "cloud-disk/example01.pdf"
	UploadID := "16874252953c1552d0e48af4ef2f26f04ec9fd0e9f855a32cf67160c9a310c71f0d5f545ed"

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "5711dc05e620528080fc6e0b4e27d47a"},
		cos.Object{
			PartNumber: 2, ETag: "635cd33489cf11cbcd5b545a3935fab6"},
		cos.Object{
			PartNumber: 3, ETag: "3979027e120eff7ce30b0610166745a0"},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		t.Fatal(err)
	}
}
