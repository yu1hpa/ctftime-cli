package main

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type EventInfo []struct {
	CtfID       json.Number `json:"ctf_id"`
	CtftimeURL  string      `json:"ctftime_url"`
	Description string      `json:"description"`
	Duration    struct {
		Days  json.Number `json:"days"`
		Hours json.Number `json:"hours"`
	} `json:"duration"`
	Finish       string      `json:"finish"`
	Format       string      `json:"format"`
	FormatID     json.Number `json:"format_id"`
	ID           json.Number `json:"id"`
	IsVotableNow bool        `json:"is_votable_now"`
	LiveFeed     string      `json:"live_feed"`
	Location     string      `json:"location"`
	Logo         string      `json:"logo"`
	Onsite       bool        `json:"onsite"`
	Organizers   []struct {
		ID   json.Number `json:"id"`
		Name string      `json:"name"`
	} `json:"organizers"`
	Participants  json.Number `json:"participants"`
	PublicVotable bool        `json:"public_votable"`
	Restrictions  string      `json:"restrictions"`
	Start         string      `json:"start"`
	Title         string      `json:"title"`
	URL           string      `json:"url"`
	Weight        json.Number `json:"weight"`
}

func fetchEvents(opts *CmdOpts) {
	url := fmt.Sprintf("https://ctftime.org/api/v1/events/?limit=3&start=%d", opts.UnixTime)

	fmt.Println(url)
	body := fetch(url)
	var events EventInfo
	if err := json.Unmarshal([]byte(body), &events); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}

	header := []string{"CTF NAME", "DURATION", "TIME FRAME", "FORMAT"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	table.SetAutoFormatHeaders(false)

	for _, event := range events {
		ctf := make([]string, len(header))
		ctf[0] = event.Title
		ctf[1] = string(event.Duration.Days) + "days, " + string(event.Duration.Hours) + "hours"
		ctf[2] = event.Start[:16] + "\n -> " + event.Finish[:16] + "(UTC)"
		ctf[3] = event.Format
		table.Append(ctf)
	}
	table.Render()
}
