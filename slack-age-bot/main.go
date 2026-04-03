package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func mustGetEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Missing env var: %s", k)
	}
	return v
}

func main() {
	_ = godotenv.Load()

	bot := slacker.NewClient(
		mustGetEnv("SLACK_BOT_TOKEN"), // read tokens from the environment
		mustGetEnv("SLACK_APP_TOKEN"), // creates a slack bot instance
	)

	go printCommandEvents(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{ // <year> is a placeholder
		// Metadata
		Description: "YOB calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println("error")
			}
			age := 2026 - yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background()) // create a cancelable context
	defer cancel()                                          // program cleanup

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
