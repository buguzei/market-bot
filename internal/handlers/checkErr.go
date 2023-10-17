package handlers

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErrWithText(err error, text string) {
	if err != nil {
		log.Printf(text)
	}
}
