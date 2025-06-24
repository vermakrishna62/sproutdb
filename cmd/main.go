package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"sproutdb/internal/kvstore"
)

func main() {
	store := kvstore.NewStore()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🌱 SproutDB CLI ready.")
	fmt.Println("Available commands: SET <key> <value>, GET <key>, DELETE <key>, EXIT")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		args := strings.Fields(line)
		switch strings.ToUpper(args[0]) {
		case "SET":
			if len(args) != 3 {
				fmt.Println("Usage: SET <key> <value>")
				continue
			}
			store.Set(args[1], args[2])
			fmt.Println("OK")

		case "GET":
			if len(args) != 2 {
				fmt.Println("Usage: GET <key>")
				continue
			}

			if val, ok := store.Get(args[1]); ok {
				fmt.Println(val)
			} else {
				fmt.Println("(nil)")
			}
		case "DELETE":
			if len(args) != 2 {
				fmt.Println("Usage: DELETE <key>")
				continue
			}
			store.Remove(args[1])
			fmt.Println("OK")

		case "EXIT":
			fmt.Println("Exiting SproutDB...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
