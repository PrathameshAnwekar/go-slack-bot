package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event:= range analyticsChannel{
		fmt.Println("Command event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

func main(){
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3961080307462-3964729585157-i3aADGgZcMhjJIERrAVGFiru")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03UFH3BS03-3967582301379-68bd4556e39ebceb199a7a4dfcadf8607f2ed3c385251d06a13ee549465074b6")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	example := []string{"My yob is 2020"}

	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "Returns the year of birth",
		Examples :example,
		Handler: func(botctx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil{
				response.Reply("Invalid year")
				return
			}
			age := 2022 - yob
			r := fmt.Sprintf("You are %d years old", age)
			response.Reply(r)
		},
		
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}