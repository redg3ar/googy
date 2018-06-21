package main

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"log"
	"os/exec"
	"strings"
)

func figlet(s *discordgo.Session, m *discordgo.MessageCreate) {
	trimmed := strings.Join(strings.Split(m.Content, " ")[1:], " ")
	if len(strings.Replace(trimmed, " ", "", -1)) > 10 || len(trimmed) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Usage: "+prefix+"figlet [less then 10 letters]")
		return
	}
	c := exec.Command("figlet", trimmed)
	var out bytes.Buffer
	c.Stdout = &out
	if err := c.Run(); err != nil {
		log.Printf("[FIGLET] Error: %v\n", err)
		return
	}
	s.ChannelMessageSend(m.ChannelID, "```\n"+out.String()+"\n```")
}
