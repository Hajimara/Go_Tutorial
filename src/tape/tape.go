package main

import (
	"constants"
	"fmt"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(constants.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	// var player Player = constants.TapePlayer{}
	// mixtape := []string{"Colors", "Villain", "Celebrity"}
	// playList(player, mixtape)
	// player = constants.TapeRecorder{}
	// playList(player, mixtape)
	TryOut(constants.TapeRecorder{})
	TryOut(constants.TapePlayer{})
}
