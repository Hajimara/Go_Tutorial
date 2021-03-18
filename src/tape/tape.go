package main

import "constants"

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

func main() {
	var player Player = constants.TapePlayer{}
	mixtape := []string{"Colors", "Villain", "Celebrity"}
	playList(player, mixtape)
	player = constants.TapeRecorder{}
	playList(player, mixtape)

}
