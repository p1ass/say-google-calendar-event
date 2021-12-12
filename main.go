package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	calendar "google.golang.org/api/calendar/v3"
)

var calendarIdFlag = flag.String("calendarId", "", "カレンダーID")

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func main() {
	flag.Parse()

	if calendarIdFlag == nil || *calendarIdFlag == "" {
		fmt.Println("-calendarId flag must be passed")
		os.Exit(1)
	}
	calendarId := *calendarIdFlag

	ctx := context.Background()
	calendarService, err := calendar.NewService(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//now := time.Now()
	now := time.Date(2021, 12, 25, 17, 59, 0, 0, jst)
	events, err := calendarService.Events.List(calendarId).
		TimeMin(now.Add(1 * time.Minute).Format(time.RFC3339)).
		TimeMax(now.Add(2 * time.Minute).Format(time.RFC3339)).
		OrderBy("startTime").
		SingleEvents(true).
		Do()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, event := range events.Items {
		if event.Start.DateTime == "" {
			// 終日イベントは見ない
			continue
		}
		start, err := time.Parse(time.RFC3339, event.Start.DateTime)
		if err != nil {
			fmt.Println(err)
		}
		if start.Before(now) {
			fmt.Println("skip")
			continue
		}

		fmt.Printf(" %s %#v\n", event.Start.DateTime, event.Summary)

		sayPhase := fmt.Sprintf("予定「%s」が始まります", event.Summary)
		err = exec.Command("say", "-r 200", sayPhase).Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
