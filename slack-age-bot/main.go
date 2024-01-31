package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println("Timestamp: ", event.Timestamp)
		fmt.Println("Command: ", event.Command)
		fmt.Println("Parameters: ", event.Parameters)
		fmt.Println("Event: ", event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6530474480693-6518937518007-8avpHWuyyvp0v2YXXz154eVd")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A06FLEV2PNH-6533324213891-094f4abd1ff60f5a368d35f13a2c2b9a561cdf2e2e112d908c5850476e3167b8")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(bc slacker.BotContext, req slacker.Request, w slacker.ResponseWriter) {
			year := req.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			w.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
