package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	radiko "github.com/yyoshiki41/go-radiko"
)

func main() {
	dir, f := createTempDir()
	defer f()

	// 1. Download a swf player
	swfPath := path.Join(dir, "myplayer.swf")
	if err := radiko.DownloadPlayer(swfPath); err != nil {
		log.Fatalf("Failed to download swf player. %s", err)
	}

	// 2. Using swfextract, create an authkey file from a swf player.
	cmdPath, err := exec.LookPath("swfextract")
	if err != nil {
		log.Fatal(err)
	}
	authKeyPath := path.Join(dir, "authkey.png")
	if c := exec.Command(cmdPath, "-b", "12", swfPath, "-o", authKeyPath); err != c.Run() {
		log.Fatalf("Failed to execute swfextract. %s", err)
	}

	// 3. Create a new Client.
	client, err := radiko.New("")
	if err != nil {
		log.Fatalf("Failed to construct a radiko Client. %s", err)
	}

	// 4. Enables and sets the auth_token.
	// After client.AuthorizeToken() has succeeded,
	// the client has the enabled auth_token internally.
	authToken, err := client.AuthorizeToken(context.Background(), authKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(authToken)
}

func createTempDir() (string, func()) {
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatalf("Failed to create temp dir: %s", err)
	}

	return dir, func() { os.RemoveAll(dir) }
}
