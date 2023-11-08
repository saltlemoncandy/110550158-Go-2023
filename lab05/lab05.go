package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/joho/godotenv"
	"os"
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
	"html/template"
)

// TODO: Please create a struct to include the information of a video
type Video struct {
	Id          	string 
	Title			string 
	ChannelTitle	string
	ViewCount   	string
	LikeCount   	string
	CommentCount 	string 
	PublishedAt 	string 
}

func formatStringWithCommas(numStr string) string {
    parts := strings.Split(numStr, "")
    result := ""

    for i, part := range parts {
        if i > 0 && (len(parts)-i)%3 == 0 {
            result += ","
        }
        result += part
    }

    return result
}

func formatPublishedDate(dateStr string) string {
    parsedTime, err := time.Parse(time.RFC3339, dateStr)

    if err != nil {
        return dateStr 
    }

    formattedDate := parsedTime.Format("2006年01月02日")
    return formattedDate
}

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Get API token from .env file
	err := godotenv.Load() 
	api := os.Getenv("YOUTUBE_API_KEY")
	
	// TODO: Get video ID from URL query `v`
	values := r.URL.Query()
	videoID := values.Get("v")
	if videoID == "" {
        http.ServeFile(w, r, "error.html")
        return
    }

	// TODO: Get video information from YouTube API
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?key=%s&id=%s&part=snippet,statistics", api, videoID)
	response, err := http.Get(url)
    if err != nil {
        http.ServeFile(w, r, "error.html")
        return
    }
    defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
    if err != nil {
        http.ServeFile(w, r, "error.html")
        return
    }
	
	// TODO: Parse the JSON response and store the information into a struct
	var m map[string]interface{}
	if err := json.Unmarshal(responseBody, &m); err != nil {
        http.ServeFile(w, r, "error.html")
        return
    }
	items, ok := m["items"].([]interface{})
    if !ok || len(items) == 0 {
        http.ServeFile(w, r, "error.html")
        return
    }
	// items := m["items"].([]interface{})
	video := Video{
		Id:           videoID,
		Title:        items[0].(map[string]interface{})["snippet"].(map[string]interface{})["title"].(string),
		ChannelTitle: items[0].(map[string]interface{})["snippet"].(map[string]interface{})["channelTitle"].(string),
		ViewCount:    items[0].(map[string]interface{})["statistics"].(map[string]interface{})["viewCount"].(string),
		LikeCount:    items[0].(map[string]interface{})["statistics"].(map[string]interface{})["likeCount"].(string),
		CommentCount: items[0].(map[string]interface{})["statistics"].(map[string]interface{})["commentCount"].(string),
		PublishedAt:  items[0].(map[string]interface{})["snippet"].(map[string]interface{})["publishedAt"].(string),
	}

	fmt.Println("Video ID:", video.Id)
	fmt.Println("Title:", video.Title)
	fmt.Println("Channel Title:", video.ChannelTitle)
	fmt.Println("View Count:", video.ViewCount)
	fmt.Println("Like Count:", video.LikeCount)
	fmt.Println("Comment Count:", video.CommentCount)
	fmt.Println("Published At:", video.PublishedAt)
	video.ViewCount = formatStringWithCommas(video.ViewCount)
	video.LikeCount = formatStringWithCommas(video.LikeCount)
	video.CommentCount = formatStringWithCommas(video.CommentCount)
	video.PublishedAt = formatPublishedDate(video.PublishedAt)
    
	// TODO: Display the information in an HTML page through `template'

	err = template.Must(template.ParseFiles("index.html")).Execute(w, video)
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}

