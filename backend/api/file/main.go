package main

import (
	"context"
	"jiaojiao/srv/file/mock"
	file "jiaojiao/srv/file/proto"
	"jiaojiao/utils"

	"github.com/h2non/filetype"
	"github.com/micro/go-micro/client"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router, rg := utils.CreateAPIGroup()
	rg.GET("/file/:fileId", getFile)
	return router
}

/**
 * @apiDefine FileServiceDown
 * @apiError (Error 500) FileServiceDown File service down
 */

/**
 * @api {get} /file/:fileId GetFile
 * @apiVersion 1.0.0
 * @apiGroup File
 * @apiPermission none
 * @apiName GetFile
 * @apiDescription Get file
 *
 * @apiParam {--} Param see [File Service](#api-Service-file_File_Query)
 * @apiSuccess (Success 200) {Response} response see [File Service](#api-Service-file_File_Query)
 * @apiUse InvalidParam
 * @apiError (Error 404) FileNotFound file not found
 * @apiUse FileServiceDown
 */
func getFile(c *gin.Context) {
	type param struct {
		FileId string `uri:"fileId" binding:"required"`
	}
	var p param

	if !utils.LogContinue(c.ShouldBindUri(&p), utils.Warning) {
		srv := utils.CallMicroService("file", func(name string, c client.Client) interface{} { return file.NewFileService(name, c) },
			func() interface{} { return mock.NewFileService() }).(file.FileService)

		rsp, err := srv.Query(context.TODO(), &file.FileRequest{
			FileId: p.FileId,
		})
		if utils.LogContinue(err, utils.Warning, "File service error: %v", err) {
			c.JSON(500, err)
			return
		}

		if rsp.Status == file.FileQueryResponse_SUCCESS {
			if filetype.IsImage(rsp.File) || filetype.IsAudio(rsp.File) || filetype.IsVideo(rsp.File) {
				t, err := filetype.Match(rsp.File)
				if utils.LogContinue(err, utils.Warning, "File format error: %v", err) {
					c.JSON(500, err)
					return
				}
				c.Data(200, t.MIME.Value, rsp.File)
			} else {
				c.AbortWithStatus(403)
			}
		} else {
			c.AbortWithStatus(404)
		}
	} else {
		c.AbortWithStatus(400)
	}
}

func main() {
	utils.RunWebService("file", setupRouter())
}