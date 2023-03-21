package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
        maxResults = flag.Int64("max-results", 50, "Max YouTube results")
)

const developerKey = "AIzaSyAfM5NAfgO44t3Iq-Glp02n03JRNObo83s" //YOUR DEVELOPER KEY

//use curl call directly
func main(){
        file, err := os.Open("/Users/kirstenliu/Desktop/2020_jobs/part_time/台經院animation/YoutubeChannelList")
        if err != nil{
                log.Fatalln("can't open youtube channel list file")
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        cnt, cant := 0, 0
        for scanner.Scan(){
                link := scanner.Text()
                fmt.Println("link: ", link)

                if link[:4] == "null"{
                        fmt.Println("none")
                }else{
                        pre := len("https://www.youtube.com/")
                        
                        slashLoca := strings.Index(link[pre:], string('/'))
                        kind := link[pre:pre+slashLoca]
                        
                        id := link[pre+slashLoca+1:]
                        
                        if kind == "channel" {
                                GetCols(id, kind)
                        }else if kind == "user"{
                                GetCols(id, kind)
                        }else if kind == "c"{
                                cnt++
                        }else{
                                fmt.Println("can't handle")
                                cant++
                        }
                }
        }

        if err := scanner.Err(); err != nil{
                log.Fatalln("scanner err for readline from youtube channel list file: ", err)
        }

        fmt.Println("c, can't: ", cnt, cant)
}

func GetCols(channelId string, IdCategory string){
        url := ""
        if IdCategory == "channel"{
                url = "https://www.googleapis.com/youtube/v3/channels?part=snippet,contentDetails&key="+developerKey+"&"+"id="+channelId+"&publishedAfter=2020-12-31T23%3A59%3A59Z&publishedBefore=2022-01-01T00%3A00%3A00Z&order=date&maxResults=50"

        }else if IdCategory == "user"{
                url = "https://www.googleapis.com/youtube/v3/channels?part=snippet,contentDetails&key="+developerKey+"&"+"forUsername="+channelId+"&publishedAfter=2020-12-31T23%3A59%3A59Z&publishedBefore=2022-01-01T00%3A00%3A00Z&order=date&maxResults=50"
        }else{
                fmt.Printf("Can't parse: %v with category %v\n", channelId, IdCategory)
                return
        }
        
        resp, err := http.Get(url)
        if err != nil {
                fmt.Println("err")
	        fmt.Println(err)
        }
        
        defer resp.Body.Close()
        
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil{
                fmt.Println("io read err")
                fmt.Println(err)
        }
        
        var data ChannelList
        json.Unmarshal(body, &data)
        
        playlistId := data.Items[0].ContentDetails.RelatedPlaylists.Uploads
        nextPageToken := ""
        for{
                url2 := "https://youtube.googleapis.com/youtube/v3/playlistItems?part=snippet&part=contentDetails&maxResults=50&playlistId="+playlistId+"&pageToken="+nextPageToken+"&key="+developerKey        
                resp2, err := http.Get(url2)
                if err != nil {
                        fmt.Println("err")
                        fmt.Println(err)
                }
                
                defer resp2.Body.Close()
                
                body2, err := ioutil.ReadAll(resp2.Body)
                if err != nil{
                        fmt.Println("io read err")
                        fmt.Println(err)
                }
                
                var data2 PlaylistItem
                json.Unmarshal(body2, &data2)
                
                if data2.PageInfo.TotalResults != 0 {
                        for _, item := range data2.Items {
                                if item.Snippet.PublishedAt.Before(time.Date(2022, time.January, 1, 00, 00, 00, 00, time.Local)) && item.Snippet.PublishedAt.After(time.Date(2020, time.December,31,23,59,59,59,time.Local)){
                                        link := "https://www.youtube.com/watch?v="+item.Snippet.ResourceID.VideoID
                                        title := item.Snippet.Title
                                        date := item.Snippet.PublishedAt

                                        urlVideo := "https://www.googleapis.com/youtube/v3/videos?id="+item.Snippet.ResourceID.VideoID+"&part=contentDetails&key="+developerKey
                                        respVideo, err := http.Get(urlVideo)
                                        if err != nil {
                                                fmt.Println("err")
                                                fmt.Println(err)
                                        }
                                        defer respVideo.Body.Close()

                                        bodyVideo, err := ioutil.ReadAll(respVideo.Body)
                                        var dataVideo Videos
                                        json.Unmarshal(bodyVideo, &dataVideo)
                                        duration := dataVideo.Items[0].ContentDetails.Duration
                                        fmt.Printf("%v$ $%v$ %v$ $%v\n", title, date, link, CalSeconds(duration))
                                }
                        }
                        nextPageToken = data2.NextPageToken
                        if nextPageToken == "" {
                                break
                        }
                }
                if nextPageToken == ""{
                        break
                }
        }
}

func CalSeconds(input string) int {
        
        if input[len(input)-1] == 'M'{
                input = input[:len(input)-1]
	        input = input[2:]
                m, _ := strconv.Atoi(input)
                return 60*m
        }
        
	input = input[:len(input)-1]
	input = input[2:]
	
	total := 0
	loca := -1

	for idx, c := range input {
		if c == 'M' {
			loca = idx
		}
	}
	
        if loca != -1{
                s, _ := strconv.Atoi(input[loca+1:])
	        m, _ := strconv.Atoi(input[:loca])

	        total = m*60 + s
	        return total
        }else{
                s, _ := strconv.Atoi(input)
                return s
        }
	
}
func YoutubeQuery() {
        flag.Parse()

        client := &http.Client{
                Transport: &transport.APIKey{Key: developerKey},
        }

        service, err := youtube.New(client)
        if err != nil {
                log.Fatalf("Error creating new YouTube client: %v", err)
        }

        // Make the API call to YouTube.
        var part []string
        call := service.Search.List(part).ChannelId("UCbijK1E1aI7b_GMOBvNDKcQ").Order("date")
        response, err := call.Do()
        if err != nil{
                fmt.Println(err)
        }

        // Group video, channel, and playlist results in separate lists.
        videos := make(map[string]string)

        fmt.Println("response:")
        fmt.Println(response)
        
        printIDs("Videos", videos)
        //printIDs("Channels", channels)
        //printIDs("Playlists", playlists)
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
// Retrieve resource for the authenticated user's channel

func printIDs(sectionName string, matches map[string]string) {
        fmt.Printf("%v:\n", sectionName)
        for id, title := range matches {
                fmt.Printf("[%v] %v\n", "https://www.youtube.com/watch?v="+id, title)
        }
        fmt.Printf("\n\n")
}