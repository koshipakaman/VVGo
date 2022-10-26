package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var voiceStateCache = map[string]*discordgo.VoiceState{}
var voiceConnection *discordgo.VoiceConnection
var vocab = loadVocab()

var _ = godotenv.Load()
var VOICEVOX_KEY = os.Getenv("VOICEVOX_KEY")

func onVoiceUpdate(s *discordgo.Session, v *discordgo.VoiceStateUpdate) {
	var join bool
	_, ok := voiceStateCache[v.UserID]
	if ok {
		delete(voiceStateCache, v.UserID)
		join = false
	} else {
		voiceStateCache[v.UserID] = v.VoiceState
		join = true
	}

	if join && len(voiceStateCache) == 1 {
		vc, _ := s.ChannelVoiceJoin(v.GuildID, v.VoiceState.ChannelID, false, false)
		content := "ボクはずんだもんなのだ！"
		url := fmt.Sprintf("https://api.su-shiki.com/v2/voicevox/audio/?text=%s&key=%s&speaker=3&intonationScale=1&speed=1", content, VOICEVOX_KEY)
		dgvoice.PlayAudioFile(vc, url, make(chan bool))
		voiceConnection = vc
	}

	if !join && len(voiceStateCache) == 1 {
		voiceConnection.Disconnect()
	}

}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if userContains(m.Mentions, s.State.User) {

		_, ok := voiceStateCache[s.State.User.ID]
		if ok {

			splited := strings.Split(m.Content, " ")
			content := strings.Join(splited[1:], " ")

			if len(content) == 0 {
				content = randomChoice(vocab)
			}

			url := fmt.Sprintf("https://api.su-shiki.com/v2/voicevox/audio/?text=%s&key=%s&speaker=3&intonationScale=1&speed=1", content, VOICEVOX_KEY)
			dgvoice.PlayAudioFile(voiceConnection, url, make(chan bool))
		} else {
			content := randomChoice(vocab)
			s.ChannelMessageSend(m.ChannelID, content)
		}
	}
}

func userContains(s []*discordgo.User, element *discordgo.User) bool {
	for _, a := range s {
		if a.String() == element.String() {
			return true
		}
	}
	return false
}
