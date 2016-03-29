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
)

//TODO make it faster. Currently, this method could waste time by swap more time for a element
//Ex. Swap i=0;j=4 then swap 0<->4; i=4;j=0 then swap 1<->4
func ShuffleArray(a []slack.User) {
	length := len(a)
	for i := range a {
		j := rand.Intn(length)
		fmt.Println(i, j)
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
	var targetChannelName string
	flag.StringVar(&targetChannelName, "channel", "", "Channel id")
	flag.Parse()

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	slackApiToken := os.Getenv("SLACK_API_TOKEN")
	if(targetChannelName == ""){
		targetChannelName = os.Getenv("SLACK_TARGET_CHANNEL")
	}

	api := slack.New(slackApiToken)
	api.SetDebug(true)

	postMessageArgs := slack.PostMessageParameters{
		Username:os.Getenv("SLACK_BOT_NAME"),
	};

	targetChannelInfo, err := api.GetChannelInfo(targetChannelName);
	if err != nil {
		fmt.Printf("GetChannelInfo(%s) : %s\n", targetChannelName, err)
		return
	}

	users, err := api.GetUsers();
	if err != nil {
		fmt.Printf("GetUsers %s\n", err)
		return
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
		fmt.Printf("matchMembers : %s\n", err)
		return
	}

	_, timeStamp, err := api.PostMessage(targetChannelName, matchedResult, postMessageArgs)
	if err != nil {
		fmt.Printf("PostMessage %s\n", err)
		return
	}

	fmt.Printf("Matched success at : %s", timeStamp)
}
