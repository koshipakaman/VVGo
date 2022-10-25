package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

func run(token string) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "dummy page for discord bot.")
	})

	go http.ListenAndServe(":8080", nil)

	discord, err := discordgo.New("Bot " + token)

	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onVoiceUpdate)

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

var VOICEVOX_KEY = os.Getenv("VOICEVOX_KEY")

func onVoiceUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	// TODO
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if userContains(m.Mentions, s.State.User) {

		g, err := s.State.Guild(m.GuildID)
		if err != nil {
			log.Println("error: ", err)
		}

		for _, vs := range g.VoiceStates {
			if vs.UserID == m.Author.ID {
				vc, err := s.ChannelVoiceJoin(m.GuildID, vs.ChannelID, false, false)
				if err != nil {
					log.Println("error: ", err)
				}

				splited := strings.Split(m.Content, " ")
				content := strings.Join(splited[1:], " ")

				url := fmt.Sprintf("https://api.su-shiki.com/v2/voicevox/audio/?text=%s&key=%s&speaker=3&intonationScale=1&speed=1", content, VOICEVOX_KEY)
				dgvoice.PlayAudioFile(vc, url, make(chan bool))
				time.Sleep(time.Second * 1)
				vc.Disconnect()
			}
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
