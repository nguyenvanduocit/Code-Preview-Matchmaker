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
//Ex. Swap i=0;j=4 then swap 0<->4; i=4;j=0 then swap 1<->4
func ShuffleArray(a []slack.User) {
	length := len(a)
	for i := range a {
		j := rand.Intn(length)
		a[i], a[j] = a[j], a[i]
	}
}

func MatchMembers(members []slack.User) (string, error) {
	groups := make(map[slack.User]slack.User)
	rand.Seed(time.Now().UnixNano()) //Make sure go generate different numbers
	ShuffleArray(members)
	for index, member := range members {
		nextIndex := index + 1;
		if (nextIndex == len(members)) {
			nextIndex = 0
		}
		groups[member] = members[nextIndex]
	}

	messageText := "```"
	for userA, userB := range groups {
		messageText += fmt.Sprintf("%s -> %s\n", userA.Name, userB.Name)
	}
	messageText += "```"
	return messageText, nil
}

func main() {
	var targetChannelName, debugFlag, botname, token string

	flag.StringVar(&targetChannelName, "channel", "", "Channel id")
	flag.StringVar(&debugFlag, "debug", "", "Debug flag")
	flag.StringVar(&botname, "name", "", "Bot name")
	flag.StringVar(&token, "token", "", "Slack token")
	flag.Parse()

	//Only fire error if can not load .env and required arguments miss
	if  token == "" || targetChannelName == "" || botname == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Can not load .env")
		}
		if token == "" {
			token = os.Getenv("SLACK_API_TOKEN")
			if(token == ""){
				log.Fatal("Can not get token")
			}
		}
		if targetChannelName == "" {
			targetChannelName = os.Getenv("SLACK_TARGET_CHANNEL")
			if(targetChannelName == ""){
				log.Fatal("Can not get target channel")
			}
		}
		if debugFlag == "" {
			debugFlag = os.Getenv("DEBUG")
		}
		if botname == "" {
			botname = os.Getenv("SLACK_BOT_NAME")
			if(botname == ""){
				botname = "Code Preview Matchmaker"
			}
		}
	}


	api := slack.New(token)
	api.SetDebug(strings.ToUpper(debugFlag)=="TRUE")


	postMessageArgs := slack.PostMessageParameters{
		Username:botname,
	};

	targetChannelInfo, err := api.GetChannelInfo(targetChannelName);
	if err != nil {
		log.Fatal(fmt.Sprintf("GetChannelInfo(%s) : %s\n", targetChannelName, err))
	}

	users, err := api.GetUsers();
	if err != nil {
		log.Fatal(fmt.Sprintf("GetUsers %s\n", err))
	}

	var groupMembers []slack.User
	for _, memberId := range targetChannelInfo.Members {
		for _, user := range users {
			if (memberId == user.ID && !user.IsBot) {
				groupMembers = append(groupMembers, user)
			}
		}
	}

	matchedResult, err := MatchMembers(groupMembers);
	if err != nil {
		log.Fatal(fmt.Sprintf("matchMembers %s\n", err))
	}

	_, timeStamp, err := api.PostMessage(targetChannelName, matchedResult, postMessageArgs)
	if err != nil {
		log.Fatal(fmt.Sprintf("PostMessage %s\n", err))
	}
	fmt.Printf("Matched success at : %s", timeStamp)
}
