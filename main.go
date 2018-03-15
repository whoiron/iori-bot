package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"time"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
)

// Variables used for command line parameters
var (
	Token            string
	WelcomeChannelID string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&WelcomeChannelID, "c", "", "Welcome ChannelID")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(guildMemberAdd)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Errorf("err: %+v", err)
		return
	}

	if c.ID != WelcomeChannelID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	//if m.Content == "ping" {
	//	s.ChannelMessageSend(m.ChannelID, "Pong!")
	//}

	// If the message is "pong" reply with "Ping!"
	//if m.Content == "pong" {
	//	s.ChannelMessageSend(m.ChannelID, "Ping!")
	//}
}

func guildMemberAdd(s *discordgo.Session, g *discordgo.GuildMemberAdd) {
	m, err := pickWelcomeMessage()
	if err != nil {
		fmt.Println("error guildMemberAdd pickWelcomeMessage,", err)
		return
	}
	_, err = s.ChannelMessageSend(WelcomeChannelID, fmt.Sprintf(m, g.User.Mention()))
	if err != nil {
		fmt.Println("error guildMemberAdd ChannelMessageSend,", err)
		return
	}
}

func pickWelcomeMessage() (string, error) {
	c, err := loadConfig()
	if err != nil {
		return "", err
	}
	if len(c.WelcomeMessages) == 0 {
		return "", err
	}

	// pick random
	rand.Seed(time.Now().Unix())
	return c.WelcomeMessages[rand.Intn(len(c.WelcomeMessages))], nil
}

type config struct {
	WelcomeMessages []string `yaml:"welcome"`
}

func loadConfig() (c *config, err error) {
	buf, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return
	}
	return
}
