package main

import (
	"net"
	"time"

	"github.com/fatih/color"
)

// Bot for twitch chat
type Bot struct {
	url  string
	port string
	name string
	conn net.Conn
}

func newBot() Bot {
	color.Yellow("Init: bot")
	bot := Bot{url: "irc.twitch.tv", port: "6667", name: "space_ghost"}
	bot.connect()

	return bot
}

// URL of the server and port address
func (b *Bot) URL() string {
	return b.url + ":" + b.port
}

// Connects with server and
// saves the conn for later
func (b *Bot) connect() {
	conn, err := net.Dial("tcp", b.URL())
	if err != nil {
		color.Red("Unable to connect to Twitch IRC server! Reconnecting in 10 seconds...\n")
		time.Sleep(10 * time.Second)
		b.connect()
	}

	color.Green("Connected to IRC server %s\n", b.url)
	b.conn = conn
}
