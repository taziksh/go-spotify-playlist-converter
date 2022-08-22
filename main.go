package main

import (
	"context"
	"fmt"
	"log"
	"os"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"
)

func main() {
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	// search for playlists and albums containing "holiday"
	results, err := client.Search(ctx, "Ra Ra Riot - Bad To Worse", spotify.SearchTypeTrack|spotify.SearchTypeArtist)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.Tracks.Tracks)
	client.AddTracksToLibrary(ctx, results.Tracks.Tracks[0].ID)
}
