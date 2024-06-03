package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Link      string   `json:"postLink"`
	Subreddit string   `json:"subreddit"`
	Preview   []string `json:"preview"`
}

func downloadimg(url string, subreddit string) {
	fmt.Printf("Link: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	random_number := rand.Intn(999999)
	file, err := os.Create("meme-" + subreddit + strconv.Itoa(random_number) + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	println(file.Name())
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	response, err := http.Get("https://meme-api.com/gimme/dankmemes")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	biggestFile := responseObject.Preview[len(responseObject.Preview)-1]

	//for _, link := range responseObject.Preview {

	downloadimg(biggestFile, responseObject.Subreddit)
}
