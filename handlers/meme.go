package handlers

import (
	"austinmarten1/supreme-octo-giggle/download"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Link      string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Preview   []string `json:"preview"`
}

func GetMeme(c *gin.Context) {
	response, err := http.Get("https://meme-api.com/gimme/dankmemes")
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

	c.JSON(http.StatusOK, gin.H{"fileName": fileName})
}
