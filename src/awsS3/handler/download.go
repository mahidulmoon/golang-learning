package handler

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func HandlerDownload() gin.HandlerFunc{
	return func (c *gin.Context){

		filename := c.Param("filename")
		f, err := os.Create("./files/"+filename)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "Something went wrong creating the local file",
			})
		}
		downloader := s3manager.NewDownloader(sess)
		_, err = downloader.Download(f, &s3.GetObjectInput{
			Bucket: aws.String(AWS_S3_BUCKET),
			Key:    aws.String(filename),
		})
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "Something went wrong retrieving the file from S3",
			})
		}
		c.Writer.Header().Add("Content-type", "text/file")
		_, err = io.Copy(c.Writer, f)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Code":    http.StatusInternalServerError,
				"Message": "Error:" + err.Error(),
			})
			return
		}
	}
}
