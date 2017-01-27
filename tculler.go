package main

import (
	"flag"
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

var (
	accessToken    = flag.String("access_token", "", "Twitter access token")
	accessSecret   = flag.String("access_secret", "", "Twitter access secret")
	consumerKey    = flag.String("consumer_key", "", "Twitter consumer key")
	consumerSecret = flag.String("consumer_secret", "", "Twitter consumer secret")
	username       = flag.String("username", "", "Username for media timeline deletion")

	commands = map[string]func(*anaconda.TwitterApi) error{
		"faves": faves,
		"media": media,
		"tweets": tweets,
	}

	usage    = "\ntculler <tweets|faves|media>"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		fmt.Println(usage)
		return
	}

	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(*accessToken, *accessSecret)

	cmd, ok := commands[flag.Args()[0]]
	if !ok {
		fmt.Printf("command %s not found\n\n", flag.Args()[0])
		flag.Usage()
		fmt.Println(usage)
		return
	}

	fmt.Println("running", flag.Args()[0], "delete")

	if err := cmd(api); err != nil {
		fmt.Println(err)
	}
}