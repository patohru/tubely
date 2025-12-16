package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os/exec"
)

type ffprobeOutput struct {
	Streams []struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"streams"`
}

func getVideoAspectRatio(filePath string) (string, error) {
	buffer := &bytes.Buffer{}
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)

	cmd.Stdout = buffer

	cmd.Run()
	data := buffer.Bytes()

	var out ffprobeOutput
	err := json.Unmarshal(data, &out)
	if err != nil {
		log.Fatal("Unable to unmarshal json")
		return "", err
	}

	width := out.Streams[0].Width
	height := out.Streams[0].Height

	if width == 16*height/9 {
		return "16:9", nil
	} else if height == 16*width/9 {
		return "9:16", nil
	}

	return "other", nil
}
