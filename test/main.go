package main

import (
	"fmt"

	"github.com/mj-pg/piglatin/client"
)

func main() {
	fmt.Println("Interactive test of the API client")

	// TODO: get config from config.sh
	for {
		fmt.Println("\nSelect API number")
		fmt.Println("1 - Translate word/s")
		fmt.Println("2 - Get all translated words")
		fmt.Println("others - exit")
		var api int
		fmt.Scanf("%d", &api)

		switch api {
		case 1:
			fmt.Println("Calling Post API")
		case 2:
			fmt.Println("Calling Get API")
			resp, err := client.Get()
			if err != nil {
				fmt.Printf("Got error %q", err)
				continue
			}
			fmt.Println("Got response:", resp)
		default:
			fmt.Println("Goodbye..")
			return
		}
	}
}
