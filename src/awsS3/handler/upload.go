package handler

import(
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"

)


func HandlerUpload() gin.HandlerFunc{
	return func(c *gin.Context){
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, "Bad request")
			return
		}
		filename := header.Filename

		fmt.Println(file, err, filename)
		defer file.Close()
		uploader := s3manager.NewUploader(sess)

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(AWS_S3_BUCKET), // Bucket
			Key:    aws.String(filename),      // Name of the file to be saved
			Body:   file,                      // File
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, "cannot upload")
		}




		c.JSON(200,gin.H{
			"message" : "success",
		})
	}
}