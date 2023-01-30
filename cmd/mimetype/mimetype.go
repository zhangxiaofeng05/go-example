package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

// use "router.Use(CORS())"
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// reference: https://doreentseng.github.io/detecting-mime-type-read-csv-using-golang/

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	// 给前端使用的上传API：http://localhost:8080/csv/upload
	// port 的设定在最后一行
	router.Use(CORS()).POST("/csv/upload", func(c *gin.Context) {
		// get the file
		fFile, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("get file err: %s", err.Error()),
			})
			return
		}
		// open file
		file, err := fFile.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
		defer file.Close()
		// transfer file to data
		data, err := ioutil.ReadAll(bufio.NewReader(file))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("err: %s", err),
			})
		}
		// detect if MIME type is "text/csv"
		mime := mimetype.Detect(data)
		if !mime.Is("text/csv") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("err: mime type is %s, text/csv is required", mime),
			})
			return
		}
		// read data
		reader := csv.NewReader(bytes.NewReader(data))
		line, err := reader.ReadAll()
		fmt.Println(line)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("read err: %s", err.Error()),
			})
			return
		}

		// ... do with line

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("File uploaded successfully"),
		})
	})
	router.Run(":8080") //port
}
