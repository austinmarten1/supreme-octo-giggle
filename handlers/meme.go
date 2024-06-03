package handlers

import (
	"austinmarten1/supreme-octo-giggle/download"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Link      string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Preview   []string `json:"preview"`
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetMeme(c *gin.Context) {
	redditsub := c.PostForm("subreddit")
	queryUrl := fmt.Sprintf("https://meme-api.com/gimme/%s", redditsub)
	response, err := http.Get(queryUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	biggestFile := responseObject.Preview[len(responseObject.Preview)-1]

	fileName, err := download.DownloadImg(biggestFile, responseObject.Subreddit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("HX-Redirect", "/memes/"+fileName)
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, gin.H{
		"fileName":  fileName,
		"fileUrl":   fmt.Sprintf("/memes/%s", fileName),
		"query":     queryUrl,
		"subreddit": redditsub,
	})
}
