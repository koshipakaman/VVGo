package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func run(token string) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "dummy page for discord bot.")
	})

	go http.ListenAndServe(":8080", nil)

	discord, err := discordgo.New("Bot " + token)

	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onVoiceUpdate)
	discord.AddHandler(voiceChannelJoin)

	discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot

	err = discord.Close()

	return
}

func contains[T comparable](s []T, element T) bool {
	for _, a := range s {
		if a == element {
			return true
		}
	}
	return false
}

func main() {

	flag.Parse()
	godotenv.Load()
	if contains(flag.Args(), "dev") {
		run(os.Getenv("DEV_TOKEN"))
	} else {
		run(os.Getenv("TOKEN"))
	}
}
