package main

import (
    "fmt"
    "time"
    "flag"
    "encoding/json"
)

type EventInfo []struct {
    CtfID       json.Number  `json:"ctf_id"`
    CtftimeURL  string `json:"ctftime_url"`
    Description string `json:"description"`
    Duration    struct {
        Days  json.Number `json:"days"`
        Hours json.Number `json:"hours"`
    } `json:"duration"`
    Finish       string `json:"finish"`
    Format       string `json:"format"`
    FormatID     json.Number  `json:"format_id"`
    ID           json.Number  `json:"id"`
    IsVotableNow bool   `json:"is_votable_now"`
    LiveFeed     string `json:"live_feed"`
    Location     string `json:"location"`
    Logo         string `json:"logo"`
    Onsite       bool   `json:"onsite"`
    Organizers   []struct {
        ID   json.Number  `json:"id"`
        Name string `json:"name"`
    } `json:"organizers"`
    Participants  json.Number  `json:"participants"`
    PublicVotable bool   `json:"public_votable"`
    Restrictions  string `json:"restrictions"`
    Start         string `json:"start"`
    Title         string `json:"title"`
    URL           string `json:"url"`
    Weight        json.Number  `json:"weight"`
}


type CmdOpts struct {
    Type string
    UnixTime int64
    Args []string
}

func main() {
    unixTime := time.Now().Unix()

    opts := CmdOpts{}
    //flag.Usage = flagHelpMessage
    flag.StringVar(&opts.Type, "type", "events", "[TODO: add explain]")
    flag.Int64Var(&opts.UnixTime, "time", unixTime, "[TODO: add explain]")
    flag.Parse()

    if opts.Type == "events" {
        url := fmt.Sprintf("https://ctftime.org/api/v1/events/?limit=3&start=%d", opts.UnixTime)

        fmt.Println(url)
        body := fetch(url)
        var events EventInfo
        if err := json.Unmarshal([]byte(body), &events); err != nil {
            fmt.Println("JSON Unmarshal error:", err)
            return
        }
        for _, event := range events {
            fmt.Println(event.Title)
        }
    }
}

