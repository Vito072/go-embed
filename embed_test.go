package go_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("Logo2.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
