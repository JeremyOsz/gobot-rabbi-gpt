package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	GPTClient "github.com/jeremyosz/gobot-rabbi-gpt/internal/open-ai"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get token from .env file
	token := os.Getenv("DISCORD_BOT_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!gopher" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello, Gopher!")
	}

	if m.Content == "!random" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hello, Random!")
	}

	if m.Content == "!gpt" {
		text, err := GPTClient.SendGPTRequest("Hello, world!", 10)
		if err != nil {
			log.Fatalf("Error calling sendGPTRequest: %v", err)
			_, _ = s.ChannelMessageSend(m.ChannelID, text)
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, text)
	}
}
