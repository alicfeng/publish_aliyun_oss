package main

import (
	"crypto/md5"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 1.读取环境变量配置信息
	endPoint := os.Getenv("PLUGIN_END_POINT")
	accessKeyId := os.Getenv("PLUGIN_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("PLUGIN_ACCESS_KEY_SECRET")
	bucketName := os.Getenv("PLUGIN_BUCKET_NAME")
	publishDir := os.Getenv("PLUGIN_PUBLISH_DIR")

	fmt.Printf("the end point     : %s\n", endPoint)
	fmt.Printf("access key id     : %s\n", accessKeyId)
	fmt.Printf("access key secret : %x\n", md5.Sum([]byte(endPoint)))
	fmt.Printf("bucket name       : %s\n", bucketName)
	fmt.Printf("publish dir       : %s\n", publishDir)

	// 2.创建OSSClient实例
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Create OSS client error:", err)
		os.Exit(-1)
	}

	// 3.设定存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Setting OSS bucket error:", err)
		os.Exit(-1)
	}

	// 4.获取发布目录下所有文件
	files, _ := GetAllFiles(publishDir)

	// 5.遍历发布文件并上传
	for index, file := range files {
		fmt.Printf("[%d/%d] %s uploading \n", index+1, len(files), file)

		// 5.1上传本地文件。
		err = bucket.PutObjectFromFile(strings.Replace(file, publishDir+"/", "", 1), file)
		if err != nil {
			fmt.Println("OssClient upload Error:", err)
			os.Exit(-1)
		}

		fmt.Printf("[%d/%d] %s finished \n", index+1, len(files), file)
	}

	os.Exit(0)
}

/**
获取发布目录下的所有文件集合
递归算法实现
*/
func GetAllFiles(dirPath string) (files []string, err error) {
	// 1.定义存储所有文件夹路径分片
	var dirs []string

	// 2.读取目录路径文件夹信息
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// 3.定义文件夹分隔符
	dirSeparator := string(os.PathSeparator)

	// 4.遍历当前文件夹下的左右文件资源
	for _, fi := range fileInfos {
		if fi.IsDir() {
			dirs = append(dirs, dirPath+dirSeparator+fi.Name())
			subFiles, _ := GetAllFiles(dirPath + dirSeparator + fi.Name())
			files = append(files, subFiles...)
			continue
		}
		files = append(files, dirPath+dirSeparator+fi.Name())
	}

	return files, nil
}
