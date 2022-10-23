package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func run(token string) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello! GAE")
	})

	go http.ListenAndServe(":8080", nil)

	discord, err := discordgo.New("Bot " + token)

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot

	err = discord.Close()

	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !m.Author.Bot {
		_, err := s.ChannelMessageSend(m.ChannelID, "@"+m.Author.Username+"Hello.")
		if err != nil {
			log.Println("error: ", err)
		}
	}

}

func contains(strs []string, element string) bool {
	for _, a := range strs {
		if a == element {
			return true
		}
	}
	return false
}

func main() {

	flag.Parse()
	if contains(flag.Args(), "dev") {
		run(os.Getenv("DEV_TOKEN"))
	} else {
		run(os.Getenv("TOKEN"))
	}
}
