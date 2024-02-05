package main

import (
	"os"
	"regexp"
	"strings"
)

func GetStyle(BannerFile string) string {
	fileContent, _ := os.ReadFile("styles/" + BannerFile)
	return string(fileContent)
}

func GenerateASCII(StringInput string, StyleName string) string {
	Style := GetStyle(StyleName)
	StyleMap := regexp.MustCompile(`\r?\n`).Split(Style, -1)

	Input := strings.Split(StringInput, "\n")
	file := ""
	for _, CurrentLine := range Input {
		if CurrentLine != "" {
			for i := 0; i < 8; i++ {
				for _, CurrentSymbol := range CurrentLine {
					StringInput = string(StyleMap[int((CurrentSymbol)-32)*9+1+i])
					file += StringInput
				}
				file += "\n"
			}
		}
		file += "\n"
	}
	return file
}
