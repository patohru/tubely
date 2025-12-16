package main

import (
	"fmt"
	"os"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {
	outputFilePath := fmt.Sprintf("%s.processing", filePath)
	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", outputFilePath)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(outputFilePath)
	if err != nil {
		return "", fmt.Errorf("Unable to process stat: %s", err)
	}
	if fileInfo.Size() == 0 {
		return "", fmt.Errorf("processed file is empty")
	}

	return outputFilePath, nil
}
