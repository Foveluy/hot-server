package main

import (
	"io/ioutil"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		b,err := ioutil.ReadFile("./version.json")
		if err != nil {
       		fmt.Print(err)
		}
		str := string(b)
		fmt.Println(str)
		// c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.JSON(200, gin.H{
			"message": str,
		})

	})
	r.Run() // listen and serve on 0.0.0.0:8080

}