package main

import (
	ttsClient "github.com/pinkeyu7/yating-tts-sdk-go"
)

func main() {
	url := "TTS_ENDPOINT"
	key := "PUT_YOUR_API_KEY_HERE"

	text := "歡迎使用雅婷逐字稿。"
	inputType := ttsClient.TypeText
	model := ttsClient.ModelZhEnFemale1
	encoding := ttsClient.EncodingLinear16
	sampleRate := ttsClient.SampleRate16k
	speed := 1.0
	pitch := 1.0
	energy := 1.0
	fileName := "example"

	client := ttsClient.NewClient(url, key)
	err := client.Synthesize(text, inputType, model, encoding, sampleRate, speed, pitch, energy, fileName)
	if err != nil {
		panic(err)
	}
}
