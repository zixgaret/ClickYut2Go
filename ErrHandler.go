package main

import "log"

func ErrHandler(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
