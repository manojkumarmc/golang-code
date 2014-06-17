package mcpack

import (
	"fmt"
	"strconv"
)

func ExecSendRcv() {

	msgs 	:= make(chan string, 5) 	// channel which will take 5 strings at a time
	status 	:= make(chan bool, 1)		// channel which specifies the status

	go func() {		// goroutine which handles the messages
		for {
			msg, more := <-msgs
			if more {
				fmt.Println("Received the message :" + msg)
			} else {
				fmt.Println("Recieved all messages.")
				status <- true
				return
			}
		}
	}()

	for j := 1; j <= 60; j++ {
		msgs <- "msg" + strconv.Itoa(j)		// string to int conversion
		fmt.Println("Sent msg :" + strconv.Itoa(j))
	}
	fmt.Println("All messages sent.")
	close(msgs)

	<-status		// waiting for the status channel

}


