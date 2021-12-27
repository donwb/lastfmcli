package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/shkh/lastfm-go/lastfm"
)

func main() {
	fmt.Println("hello")

	apikey, apisecret := getEnvVars()

	api := lastfm.New(apikey, apisecret)

	/*
		res, _ := api.Artist.GetTopTracks(lastfm.P{"artist": "Avicii"}) //discarding error
		for _, track := range res.Tracks {
			fmt.Println(track.Name)
		}
	*/

	fmt.Println("---------------------")

	userResult, _ := api.User.GetRecentTracks(lastfm.P{"user": "dbrowning"})
	for _, track := range userResult.Tracks {
		fmt.Println(track.Artist.Mbid + "    " + track.Name)
		fmt.Println(track.Album.Name)
	}

	isPlaying := userResult.Tracks[0].NowPlaying
	lastTrackPlayTime := userResult.Tracks[0].Date.Uts
	fmt.Println("Is Playing: ", isPlaying)
	fmt.Println("Last Date: ", lastTrackPlayTime)

	mbid := userResult.Tracks[0].Mbid
	fmt.Println("Mbid: ", mbid)

	fmt.Println("========================")
	fmt.Println("========================")
	// for some reason the LastFM api isn't propigating username context
	// this was to see if I could get my playcount, but it's only sending overall count
	currentArtist := userResult.Tracks[0].Artist
	currentTrack := userResult.Tracks[0].Name

	trackResult, _ := api.Track.GetInfo(lastfm.P{"artist": currentArtist, "track": currentTrack, "username": "dbrowning"})

	fmt.Println("Track Info: ", trackResult.PlayCount)

	fmt.Println("========================")
	fmt.Println("========================")

	minsSinceLastTrack := calcLastTrackTime(lastTrackPlayTime)
	fmt.Println("Minutes since last track: ", minsSinceLastTrack)

	isATrackNowPlaying := isNowPlaying(userResult)
	fmt.Println("Now Playing: ", isATrackNowPlaying)

	userInfo, _ := api.User.GetInfo(lastfm.P{"user": "dbrowning"})
	fmt.Println("Tracks: " + userInfo.PlayCount)

	fmt.Println("------ Fetching images -------")
	//trackInfo, _ := api.Track.GetInfo(lastfm.P{"mbid": "4fcb7864-42dd-4dcf-b269-6da0d9042956"})
	trackInfo, _ := api.Track.GetInfo(lastfm.P{"artist": "New Order", "track": "Run", "username": "dbrowning"})
	images := trackInfo.Album.Images

	largeImage := images[2].Url
	fmt.Println(largeImage)

}

func calcLastTrackTime(lastPlayTime string) int {
	// get current unix time
	now := time.Now().UTC().Unix()
	fmt.Println("Current UNIX Time ", now)

	lastDateInt, err := strconv.Atoi(lastPlayTime)
	if err != nil {
		return 0
	}

	// div by 60 to get minutes
	dif := now - int64(lastDateInt)
	retVal := dif / 60

	return int(retVal)
}

func isNowPlaying(trackInfo lastfm.UserGetRecentTracks) bool {
	// get most recent track
	track := trackInfo.Tracks[0]
	isPlaying := track.NowPlaying == "true"

	if !isPlaying {
		trackUTS := track.Date.Uts
		minutesSinceLastTrack := calcLastTrackTime(trackUTS)
		return minutesSinceLastTrack < 10
	} else {
		return isPlaying
	}
}

func getEnvVars() (string, string) {

	key := os.Getenv("APIKEY")
	secret := os.Getenv("SECRET")

	return key, secret
}
