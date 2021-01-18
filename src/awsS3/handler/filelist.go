package handler

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HandlerFileList() gin.HandlerFunc{
	return func(c *gin.Context){
		svc := s3.New(sess)
		input := &s3.ListObjectsInput{
			Bucket: aws.String(AWS_S3_BUCKET),
		}

		result, err := svc.ListObjects(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeNoSuchBucket:
					fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			c.JSON(http.StatusBadRequest,gin.H{
				"message" : "Something went wrong listing the files",
			})
			fmt.Println(err)
			return
		}
		c.Writer.Header().Add("Content-type", "text/html")
		fmt.Println(result)
		//for _, item := range result.Contents {
		//	fmt.Fprintf(c, "<li>File %s</li>", *item.Key)
		//}

		return
	}
}
