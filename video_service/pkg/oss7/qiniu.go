package oss7

import (
	"bytes"
	"context"
	"github.com/qiniu/api.v7/v7/auth"
	auth2 "github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"time"
	"video_service/logger"
)

var QiniuClient Qiniu

type Qiniu struct {
	accessKey string
	secretKey string
	bucket    string
}

func InitBucket() {
	QiniuClient.bucket = viper.GetString("oss.bucket")
	QiniuClient.accessKey = viper.GetString("oss.accessKey")
	QiniuClient.secretKey = viper.GetString("oss.secretKey")
}

// UploadFile 上传文件
// localFile 本地文件路径，相对当前包路径
// key 目的文件路径
func UploadFile(localFile, key string) (err error) {
	putPolicy := storage.PutPolicy{
		Scope: QiniuClient.bucket,
	}
	mac := auth.New(QiniuClient.accessKey, QiniuClient.secretKey)
	upToken := putPolicy.UploadToken((*auth2.Credentials)(mac))
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputV2Extra{}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	logger.Log.Info("上传文件" + localFile + "到七牛云" + key + "成功")
	return err
}

// UploadFileWithByte 上传文件，使用字节数组
func UploadFileWithByte(key string, localFile []byte) (err error) {
	putPolicy := storage.PutPolicy{
		Scope: QiniuClient.bucket,
	}
	mac := auth.New(QiniuClient.accessKey, QiniuClient.secretKey)
	upToken := putPolicy.UploadToken((*auth2.Credentials)(mac))
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	dataLen := int64(len(localFile))
	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(localFile), dataLen, &putExtra)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	logger.Log.Info("上传视频文件到七牛云" + key + "成功")
	return err
}

func GetFileUrl(key string) string {
	mac := auth.New(QiniuClient.accessKey, QiniuClient.secretKey)
	domain := "http://douyin.cfddfc.online"
	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
	privateAccessURL := storage.MakePrivateURL((*auth2.Credentials)(mac), domain, key, deadline)
	return privateAccessURL
}

// DeleteFile 删除文件
// key 文件路径只需要 microservice-v1 里面的文件路径
func DeleteFile(key string) error {
	mac := auth.New(QiniuClient.accessKey, QiniuClient.secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager((*auth2.Credentials)(mac), &cfg)
	err := bucketManager.Delete(QiniuClient.bucket, key)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil

}
