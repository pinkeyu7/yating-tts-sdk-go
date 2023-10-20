package YatingTtsSdk

func generator(text, inputType, model, encoding, sampleRate string, speed, pitch, energy float64) SynthesizeRequestDto {
	return SynthesizeRequestDto{
		Input: Input{
			Text: text,
			Type: inputType,
		},
		Voice: Voice{
			Model:  model,
			Speed:  speed,
			Pitch:  pitch,
			Energy: energy,
		},
		AudioConfig: AudioConfig{
			Encoding:   encoding,
			SampleRate: sampleRate,
		},
	}
}
