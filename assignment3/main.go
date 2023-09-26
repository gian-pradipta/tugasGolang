package main

import (
	requestapi "assignment_3/internal/requestapi"
	"assignment_3/internal/router"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		for {
			time.Sleep(3 * time.Second)
			requestapi.PUTRequest()
		}

	}()
	router.Run()

}
