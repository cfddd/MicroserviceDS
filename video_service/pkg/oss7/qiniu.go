package upload

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
	"video_service/logger"
)

type Qiniu struct {
	bucketName      string
	endPoint        string
	accessKeyId     string
	accessKeySecret string
}

var qiniu *Qiniu

// 创建存储空间
func ossInit(q *Qiniu) error {
	q.bucketName = viper.GetString("oss.bucketName")
	q.endPoint = viper.GetString("oss.endpoint")
	q.accessKeyId = viper.GetString("oss.accessKeyId")
	q.accessKeySecret = viper.GetString("oss.accessKeySecret")

}

// UploadFile 七牛云上传文件
func (q *Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: q.bucketName}
	mac := qbox.NewMac(q.accessKeyId, q.accessKeySecret)
	upToken := putPolicy.UploadToken(mac)
	cfg := qiniuConfig()
	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}

	f, openError := file.Open()
	if openError != nil {
		logger.Log.Error("function file.Open() failed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close()                                                  // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra)
	if putErr != nil {
		logger.Log.Error("function formUploader.Put() failed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}
	return global.GVA_CONFIG.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

// DeleteFile 七牛云删除文件
func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.GVA_CONFIG.Qiniu.AccessKey, global.GVA_CONFIG.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.GVA_CONFIG.Qiniu.Bucket, key); err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}
