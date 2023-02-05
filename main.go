package main

import (
	"fmt"
	"strings"

	"syreclabs.com/go/faker"
	"vadym.module/m/v2/config"
	"vadym.module/m/v2/bot"
)

func main() {
	var phrases []string

	for i := 1; i < 3; i++ {
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	fmt.Println(strings.Join(phrases[:], "; "))
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
  }
}
