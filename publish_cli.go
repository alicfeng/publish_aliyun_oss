/*
What SameGo team is that is 'one thing, a team, work together'
*/

package main

import (
	"crypto/md5"
	"fmt"
	cdn20180510 "github.com/alibabacloud-go/cdn-20180510/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 1.读取环境变量配置信息
	accessKeyId := os.Getenv("PLUGIN_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("PLUGIN_ACCESS_KEY_SECRET")
	ossEndPoint := os.Getenv("PLUGIN_OSS_END_POINT")
	ossBucketName := os.Getenv("PLUGIN_OSS_BUCKET_NAME")
	appPublishDir := os.Getenv("PLUGIN_APP_PUBLISH_DIR")
	cdnRegionId := os.Getenv("PLUGIN_CDN_REGION_ID")
	cdnObjectPath := os.Getenv("PLUGIN_CDN_OBJECT_PATH")
	cdnObjectType := os.Getenv("PLUGIN_CDN_OBJECT_TYPE")
	cdnEndPoint := os.Getenv("PLUGIN_CDN_END_POINT")

	// 2.终端打印变量配置信息
	fmt.Println("\033[33mThe Task Configuration List \033[0m")

	fmt.Printf("access key id     : %s\n", accessKeyId)
	fmt.Printf("access key secret : %x\n", md5.Sum([]byte(accessKeySecret)))
	fmt.Printf("oss end point     : %s\n", ossEndPoint)
	fmt.Printf("oss bucket name	 : %s\n", ossBucketName)
	fmt.Printf("app publish dir	 : %s\n", appPublishDir)
	fmt.Printf("cdn object path	 : %s\n", cdnObjectPath)
	fmt.Printf("cdn object type	 : %s\n", cdnObjectType)
	fmt.Printf("cdn region id	 : %s\n\n", cdnRegionId)

	// 3.创建OSSClient实例
	client, err := oss.New(ossEndPoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Create OSS client error:", err)
		os.Exit(-1)
	}

	// 4.设定存储空间
	bucket, err := client.Bucket(ossBucketName)
	if err != nil {
		fmt.Println("Setting OSS bucket error:", err)
		os.Exit(-1)
	}

	// 5.获取发布目录下所有文件
	files, _ := getAllFiles(appPublishDir)

	// 6.遍历发布文件并上传
	fmt.Println("\033[33mPublishing application files to oss cloud \033[0m")
	for index, file := range files {
		fmt.Printf("[%d/%d] upload %s \n", index+1, len(files), file)

		// 5.1上传本地文件。
		err = bucket.PutObjectFromFile(strings.Replace(file, appPublishDir+"/", "", 1), file)
		if err != nil {
			fmt.Println("OssClient upload Error:", err)
			os.Exit(-1)
		}
	}
	fmt.Println("\033[32mPublish application files finished \033[0m")

	// 7.定义 openapi 配置信息
	config := &openapi.Config{
		RegionId:        &cdnEndPoint,
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
		Endpoint:        &cdnEndPoint,
	}

	// 8.创建 CDN 客户端句柄
	_result, _err := cdn20180510.NewClient(config)
	if _err != nil {
		fmt.Println("Create CDN client error:", err)
		os.Exit(-1)
	}

	// 9.设置刷新对象缓存请求头
	refreshObjectCachesRequest := &cdn20180510.RefreshObjectCachesRequest{
		ObjectType: &cdnObjectType,
		ObjectPath: &cdnObjectPath,
	}

	// 10. 执行刷新 OSS.CDN 对象缓存
	_, _err = _result.RefreshObjectCaches(refreshObjectCachesRequest)
	if _err != nil {
		fmt.Println("Refresh object caches error:", err)
		os.Exit(-1)
	}

	fmt.Println("\033[32mRefresh CND object caches finished \033[0m")

	os.Exit(0)
}

// getAllFiles 获取发布目录下的所有文件集合 递归算法实现 */
func getAllFiles(dirPath string) (files []string, err error) {
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
	for _, file := range fileInfos {
		if file.IsDir() {
			dirs = append(dirs, dirPath+dirSeparator+file.Name())
			subFiles, _ := getAllFiles(dirPath + dirSeparator + file.Name())
			files = append(files, subFiles...)
			continue
		}
		files = append(files, dirPath+dirSeparator+file.Name())
	}

	return files, nil
}
