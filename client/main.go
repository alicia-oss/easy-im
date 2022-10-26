package main

import (
	"easy_im/client/core"
	"time"
)

func main() {
	client, err := core.NewClient("user1")
	if err != nil {
		return
	}
	client2, err := core.NewClient("user2")
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)
	client.ListUsers()
	time.Sleep(1 * time.Second)
	id2, _ := client2.GetUserId()
	client.SendMessage("hello", id2)
	time.Sleep(50 * time.Second)

}
