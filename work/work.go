package main

import(
	"time"
	"fmt"
)

func Worker() {
	for {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
	}
}
