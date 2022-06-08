package YatingTtsSdk

func generator(text, inputType, model, encoding, sampleRate string) SynthesizeRequestDto {
	return SynthesizeRequestDto{
		Input: Input{
			Text: text,
			Type: inputType,
		},
		Voice: Voice{
			Model: model,
		},
		AudioConfig: AudioConfig{
			Encoding:   encoding,
			SampleRate: sampleRate,
		},
	}
}
