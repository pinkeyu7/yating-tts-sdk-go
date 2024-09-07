package main

import (
	"fmt"

	ttsClient "github.com/pinkeyu7/yating-tts-sdk-go"
)

func main() {
	url := "TTS_ENDPOINT"
	key := "PUT_YOUR_API_KEY_HERE"

	req := &ttsClient.SynthesizeRequest{
		Text:       "歡迎使用雅婷逐字稿。",
		InputType:  ttsClient.InputTypeText,
		Model:      ttsClient.ModelZhEnFemale2,
		Encoding:   ttsClient.EncodingLinear16,
		SampleRate: ttsClient.SampleRate22k,
		Speed:      1.0,
		Pitch:      1.0,
		Energy:     1.0,
		FileName:   "example/example_audio",
	}

	client := ttsClient.NewClient(url, key)
	res, err := client.Synthesize(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Synthesize response: %+v\n", res)
}
