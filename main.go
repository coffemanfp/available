package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/coffemanfp/whois"
)

var marks = map[bool]string{true: "✓", false: "⨉"}

func main() {
	apiKey := os.Getenv("PROMPT_WHOIS_APIKEY")
	if apiKey == "" {
		log.Fatalln("no Prompt API Whois api key found")
	}

	whoisClient := whois.PromptAPIWhois{APIKey: apiKey}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		domain := strings.ToLower(s.Text())

		if domain == "" || domain == " " {
			continue
		}

		fmt.Print(domain, " ")

		exists, err := whoisClient.ExistsDomain(domain)
		if err != nil {
			if err.Error() == whois.PromptAPIMessageInvalidDomain {
				fmt.Println(err)
				continue
			}
			log.Fatalln(err)
		}

		fmt.Println(marks[!exists])
		time.Sleep(1 * time.Second)
	}
}
