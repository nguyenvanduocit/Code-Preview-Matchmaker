package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"fmt"
	"math/rand"
	"time"
	"flag"
	"strings"
)

//TODO make it faster. Currently, this method could waste time by swap more time for a element
//Ex. if i=0;j=4 then swap 0<->4; i=4;j=0 then swap 1<->4
func ShuffleArray(a []slack.User) {
	length := len(a)
	for i := range a {
		j := rand.Intn(length)
		a[i], a[j] = a[j], a[i]
	}
}

func MatchMembers(members []slack.User) string {
	rand.Seed(time.Now().UnixNano()) //Make sure go generate different numbers
	ShuffleArray(members)
	messageText := ""
	for index, member := range members {
		nextIndex := index + 1;
		if (nextIndex == len(members)) {
			nextIndex = 0
		}
		messageText += fmt.Sprintf("%s -> %s\n", member.Name, members[nextIndex].Name)
	}
	return fmt.Sprintf("Code preview for %s:\n ```%s```",time.Now().Local().Format("2006-01-02"), messageText)
}

func main() {
	var targetChannelName, debugFlag, botname, token string

	flag.StringVar(&targetChannelName, "channel", "", "Channel name")
	flag.StringVar(&debugFlag, "debug", "", "Debug flag")
	flag.StringVar(&botname, "name", "", "Bot name")
	flag.StringVar(&token, "token", "", "Slack token")
	flag.Parse()

	//Only fire error if required arguments missing and an not get from .env
	if  token == "" || targetChannelName == "" {
		if godotenv.Load() != nil {
			log.Fatal("Can not load .env")
		}
		if token == "" {
			token = os.Getenv("SLACK_API_TOKEN");
			if token == "" {
				log.Fatal("Can not get token")
			}
		}
		if targetChannelName == "" {
			targetChannelName = os.Getenv("SLACK_TARGET_CHANNEL");
			if targetChannelName == "" {
				log.Fatal("Can not get target channel")
			}
		}
		if debugFlag == "" {
			debugFlag = os.Getenv("DEBUG")
		}
		if botname == "" {
			botname = os.Getenv("SLACK_BOT_NAME");
			if botname == "" {
				botname = "Code Preview Matchmaker"
			}
		}
	}


	api := slack.New(token)
	api.SetDebug(strings.ToUpper(debugFlag)=="TRUE")


	postMessageArgs := slack.PostMessageParameters{
		Username:botname,
	};
	//Get channel ID from channel name
	channelList, err := api.GetChannels(true);
	if err != nil {
		log.Fatal(fmt.Sprintf("GetChannels: %s\n", err))
	}

	var targetChannel slack.Channel
	for _,channel := range channelList {
		if(channel.Name == targetChannelName){
			targetChannel = channel
			break
		}
	}

	users, err := api.GetUsers();
	if err != nil {
		log.Fatal(fmt.Sprintf("GetUsers %s\n", err))
	}

	var groupMembers []slack.User
	for _, memberId := range targetChannel.Members {
		for _, user := range users {
			if (memberId == user.ID && !user.IsBot) {
				groupMembers = append(groupMembers, user)
			}
		}
	}

	_, timeStamp, err := api.PostMessage(targetChannel.ID, MatchMembers(groupMembers), postMessageArgs)
	if err != nil {
		log.Fatal(fmt.Sprintf("PostMessage %s\n", err))
	}
	fmt.Printf("Matched success at : %s", timeStamp)
}
