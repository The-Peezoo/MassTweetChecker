//Run with: go run checker.go | Or build it with 'go build checker.go' and then run ./checker

package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "os"
        "strings"
)


func ASCII() {
    colorReset := "\033[0m"
    colorRed := "\033[31m"
    colorGreen := "\033[32m"
    colorYellow := "\033[33m"
    colorBlue := "\033[34m"
    colorPurple := "\033[35m"
    colorWhite := "\033[37m"

    fmt.Println(colorRed, "████████ ██     ██ ███████ ███████ ████████  ██████ ██   ██ ███████  ██████ ██   ██ ")
    fmt.Println(colorYellow,   "   ██    ██     ██ ██      ██         ██    ██      ██   ██ ██      ██      ██  ██  ")
    fmt.Println(colorGreen, "   ██    ██  █  ██ █████   █████      ██    ██      ███████ █████   ██      █████   ")
    fmt.Println(colorBlue, "   ██    ██ ███ ██ ██      ██         ██    ██      ██   ██ ██      ██      ██  ██  ")
    fmt.Println(colorPurple, "   ██     ███ ███  ███████ ███████    ██     ██████ ██   ██ ███████  ██████ ██   ██ ")
    fmt.Println(colorWhite, "", colorReset)
    fmt.Println(colorWhite, "                              Twitter: @Peezoo20")
    fmt.Println(colorWhite, "", colorReset)
    fmt.Println(colorWhite, "              ", colorReset)
    fmt.Println(colorGreen, "Tool to check if a tweet exists or not.", colorReset)
    fmt.Println(colorWhite, "              ", colorReset)
    fmt.Println(colorWhite, "              ", colorReset)

}

func processTweets() {
        file, err := os.Open("list.txt")
        if err != nil {
                fmt.Println("Error opening the file:", err)
                return
        }
        defer file.Close()

        links, err := ioutil.ReadAll(file)
        if err != nil {
                fmt.Println("Error reading the file:", err)
                return
        }

        existLinks := []string{}

        for _, link := range strings.Split(string(links), "\n") {
                tweetLink := strings.TrimSpace(link)
                url := fmt.Sprintf("https://publish.twitter.com/oembed?url=%s&partner=&hide_thread=false", tweetLink)
                response, err := http.Get(url)
                if err != nil {
                        fmt.Println("Request error:", err)
                        continue
                }

                if response.StatusCode == 403 {
                        fmt.Printf("Tweet not available, account could be suspended: %s\n", tweetLink)
                } else if response.StatusCode == 200 {
                        // Tweet still exists
                        existLinks = append(existLinks, tweetLink)
                }

                response.Body.Close()
        }
        colorGreen := "\033[32m"
        colorWhite := "\033[37m"
        fmt.Println(colorWhite, "              ")
        fmt.Println(colorGreen, "These tweets still exist:")
        fmt.Println(colorWhite, "              ")
	for _, link := range existLinks {
		fmt.Println(link)
	}
}

func main() {
        ASCII()
        processTweets()
}
                                                         
