package main

import (
    "io/ioutil"
    "encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"net/http"
)
func main(){

	discord, err := discordgo.New("Bot " + "Discord token")

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	fmt.Println("NASA Apod is now running.  CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {


	if m.Content == "hi" {//m.ChannelID 		593879079426588683
		s.ChannelMessageSend(m.ChannelID, "heyo")
		return
	}else
	
	if m.Content == "!nasa apod" {
		res, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=NasaApiKey")
		if err != nil {
    	panic(err.Error())
    	return
		}
		var result map[string]interface{}
    	responseData, err := ioutil.ReadAll(res.Body)
    	fmt.Println(string(responseData))
		json.Unmarshal([]byte(responseData), &result)
		fmt.Println("stat : "+ result["url"].(string))
		//msg := result["status"]
		//s.ChannelMessageSend(m.ChannelID, "**Copyright** : " + result["copyright"].(string))
		s.ChannelMessageSend(m.ChannelID, "**Date** : "+result["date"].(string))
		s.ChannelMessageSend(m.ChannelID, "**Title** : "+result["title"].(string))
		s.ChannelMessageSend(m.ChannelID, "**Explanation** : \n"+result["explanation"].(string))
		s.ChannelMessageSend(m.ChannelID, result["url"].(string))
		return
	 }
	
	fmt.Println(m.Content)

}
