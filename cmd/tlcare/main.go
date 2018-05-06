package main

import (
	"fmt"

	"flag"

	"github.com/labstack/gommon/log"
	"github.com/ninedraft/tlcare/pkg/core"
)

func main() {
	token := flag.String("token", "", "telegram Bot API token")
	flag.Parse()

	care := core.NewCare(*token, "advices.db")
	advices := []string{
		"",
	}
	for _, advice := range advices {
		adv, err := care.AddAdvice(advice)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(adv)
	}
	for i := 0; i < 4; i++ {
		adv, err := care.GetAdvice()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(adv.Text)
	}
}
