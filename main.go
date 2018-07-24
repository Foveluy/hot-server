package main

import (
	"io/ioutil"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	b,err := ioutil.ReadFile("./test.txt")
	if err != nil {
        fmt.Print(err)
	}
	str := string(b)
	fmt.Println(str)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.JSON(200, gin.H{
			"message": str,
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080

}