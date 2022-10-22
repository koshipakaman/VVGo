package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	const TOKEN = "MTAzMzA1MzY2NjI4Mzg4MDQ3OA.G2mlWy.1c1jSavay6c1hxJEBym-AIwu3s3c9Yn8vQXXOA"
	discord, err := discordgo.New("Bot " + TOKEN)

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
		_, err := s.ChannelMessageSend(m.ChannelID, "hello.")
		if err != nil {
			log.Println("error: ", err)
		}
	}

}
