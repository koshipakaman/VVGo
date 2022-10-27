package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func voiceChannelJoin(s *discordgo.Session, m *discordgo.MessageCreate) {

	g, err := s.State.Guild(m.GuildID)
	if err != nil {
		log.Println("error: ", err)
	}
	fmt.Println(voiceStateCache)
	if !(m.Content == "/join" || m.Content == "/leave") {
		return
	}
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			if m.Content == "/join" {

				vc, err := s.ChannelVoiceJoin(m.GuildID, vs.ChannelID, false, false)
				if err != nil {
					log.Println("error: ", err)
				}
				voiceConnection = vc
				voiceStateCache[s.State.User.ID] = vs
			}
			if m.Content == "/leave" {
				voiceConnection.Disconnect()
				delete(voiceStateCache, vs.UserID)
			}
		}
	}
}
