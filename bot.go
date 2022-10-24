package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func run(token string) {

	/*
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "dummy page for discord bot.")
		})

		go http.ListenAndServe(":8080", nil)
	*/

	discord, err := discordgo.New("Bot " + token)

	discord.AddHandler(onVoiceUpdate)
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

func onVoiceUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {

	if v.BeforeUpdate == nil {
		_, err := s.ChannelVoiceJoin(v.GuildID, v.ChannelID, false, false)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// dgvoice.PlayAudioFile(vc)
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if userContains(m.Mentions, s.State.User) {

		_, err := s.ChannelMessageSend(m.ChannelID, "ずんだもんはずんだの精なのだ。")
		if err != nil {
			log.Println("error: ", err)
		}
	}

}

func contains[T comparable](s []T, element T) bool {
	for _, a := range s {
		if a == element {
			return true
		}
	}
	return false
}

func userContains(s []*discordgo.User, element *discordgo.User) bool {
	for _, a := range s {
		if a.String() == element.String() {
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
