package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"myBlog/middleware/logger"
	"myBlog/model"
	"net/http"
)

var err error
var obsClient *obs.ObsClient

func init() {
	obsClient = newObsClient()
}

func Upload(c *gin.Context) {
	upload(c)
}

// 创建ObsClient结构体
func newObsClient() *obs.ObsClient {

	ak := model.ObsConf.AccessKey
	sk := model.ObsConf.ScriptKey
	endpoint := model.ObsConf.EndPoint

	obsClient, err = obs.New(ak, sk, endpoint)
	if err != nil {
		panic(fmt.Sprintf("newObsClient error = %s", err.Error()))
	}

	return obsClient
}

// 直接调用底层华为obs的上传接口
func upload(c *gin.Context) {

	ginFile, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 参数组装
	input := &obs.PutObjectInput{}
	input.Bucket = model.ObsConf.Bucket
	input.Key = model.ObsConf.Dir + "/" + fileHeader.Filename
	input.Body = ginFile
	logger.Log().Info("这是测试日志啊")

	// 文件上传
	_, err = obsClient.PutObject(input)
	if err != nil {
		c.JSON(http.StatusOK, resp2Client(model.ErrInner, err.Error(), nil))
		return
	}

	// 结果返回
	c.JSON(http.StatusOK, gin.H{
		"status":  model.Success,
		"message": model.GetErrMsg(model.Success),
		"url":     model.ObsConf.Host + input.Key,
	})
}
