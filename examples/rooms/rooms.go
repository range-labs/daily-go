// Demonstrates rooms APIs.
// DAILY_API_KEY=[TOKEN] go run examples/rooms/rooms.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	daily "github.com/range-labs/daily-go"
)

func main() {
	ctx := context.Background()
	client := daily.New(daily.WithAuth(os.Getenv("DAILY_API_KEY")))

	if cfg, err := client.GetDomainConfig(ctx); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Current Domain Config:")
		printObject(cfg)
	}

	if resp, err := client.ListRooms(ctx, nil); err != nil {
		log.Fatal(err)
		return
	} else {
		log.Println("Rooms:")
		printObject(resp)
	}

	name := "daily-go-example-room"

	// Create a new room that expires in one hour.
	newRoom := &daily.CreateRoomRequest{
		Name: daily.String(name),
		Config: &daily.RoomConfig{
			ExpiresAt: daily.Timestamp(time.Now().Add(time.Hour)),
		},
	}
	if resp, err := client.CreateRoom(ctx, newRoom); err != nil {
		log.Fatal(err)
	} else {
		log.Println("New room:")
		printObject(resp)
	}

	// Make the room private and have it expire in a minute.
	updateRoom := &daily.UpdateRoomRequest{
		Privacy: daily.Private,
		Config: &daily.RoomConfig{
			ExpiresAt: daily.Timestamp(time.Now().Add(time.Minute)),
		},
	}
	if resp, err := client.UpdateRoom(ctx, name, updateRoom); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Updated room room:")
		printObject(resp)
	}

	// And now just delete it.
	if err := client.DeleteRoom(ctx, name); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Room deleted!")
	}

}

func printObject(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		log.Println("> ", line)
	}
}
