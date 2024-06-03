package download

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	//"strconv"
	"strings"
)

func randomFileName(reddit string, ext string) string {
	random_number := rand.Intn(9999)
	return fmt.Sprintf("meme-%s-%d.%s", reddit, random_number, ext)
}

func DownloadImg(url string, subreddit string) (string, error) {
	fmt.Printf("Link: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	basePath := "memes/"
	fileName := randomFileName(subreddit, "jpg")
	file, err := os.Create(strings.Join([]string{basePath, fileName}, "/"))
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
