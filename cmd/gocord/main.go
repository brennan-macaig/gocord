package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/brennan-macaig/gocord"
	"github.com/bwmarrin/discordgo"
)

const (
	secretsFile  = "/etc/gocord/secrets.json"
	configFile   = "/etc/gocord/config.json"
	databaseFile = "/var/gocord/database.json"
)

func main() {
	token, err := gocord.GetSecret(secretsFile)
	if err != nil {
		log.Fatalf("unable to read secrets file - %s", err.Error())
	}
	dg, err := discordgo.New("Bot " + token.Token)
	if err != nil {
		log.Fatalf("could not make discord bot from token - %s", err.Error())
	}
	conf, err := gocord.GetConfig(configFile)
	if err != nil {
		log.Fatalf("could not read config file - %s", err.Error())
	}
	db := gocord.MakeDatabase(conf)
	dg.AddHandler(db.ReadMessage)
	log.Print("Bot is running. Waiting for SIGTERM to stop...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Print("Exiting...")
	dg.Close()
}
