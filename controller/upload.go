package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err :  %s", err.Error()))
		return
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(999999999)

	filename := strconv.Itoa(num) + header.Filename
	fmt.Println(filename)
	out, err := os.Create("public/upload/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:7000/public/upload/" + filename
	c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}
