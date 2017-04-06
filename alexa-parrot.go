package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var appID, consumerKey, consumerSecret, accessToken, accessSecret, secretPath string
	for env, v := range map[string]*string{
		"ALEXA_APP_ID":            &appID,
		"TWITTER_CONSUMER_KEY":    &consumerKey,
		"TWITTER_CONSUMER_SECRET": &consumerSecret,
		"TWITTER_ACCESS_TOKEN":    &accessToken,
		"TWITTER_ACCESS_SECRET":   &accessSecret,
		"SECRET_PATH":             &secretPath,
	} {
		if envVal := os.Getenv(env); envVal == "" {
			logrus.Fatalf("%q environment variable must be set", env)
		} else {
			*v = envVal
		}
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	logrus.Infof("Listening for stuff on :8080/echo/%v", secretPath)
	app := map[string]interface{}{
		"/echo/" + secretPath: alexa.EchoApplication{ // Route
			AppID:    appID,
			OnIntent: tweetHandler(client),
			OnLaunch: tweetHandler(client),
		},
	}
	alexa.Run(app, "8080")
}

const tweetSlotVal = "tweet"

func tweetHandler(c *twitter.Client) func(*alexa.EchoRequest, *alexa.EchoResponse) {
	// This is a janky random weight thing. Don't mind me.
	randomSuccessResponses := []string{
		"",
		"",
		"",
		"Thanks for improving the internet",
		"Thanks for improving the internet",
		"The internet now knows.",
		"The internet now knows.",
		"I regret it already",
	}

	return func(req *alexa.EchoRequest, resp *alexa.EchoResponse) {
		tweetText, err := req.GetSlotValue(tweetSlotVal)
		if err != nil {
			resp.OutputSpeech("Please provide tweet text")
			return
		}
		_, _, err = c.Statuses.Update(tweetText, nil)
		if err != nil {
			resp.OutputSpeech(fmt.Sprintf("Unable to tweet: %v", err))
			return
		}

		remainder := randomSuccessResponses[rand.Intn(len(randomSuccessResponses))]
		resp.OutputSpeech(fmt.Sprintf("Tweeted %v... %v", tweetText, remainder))
	}
}
