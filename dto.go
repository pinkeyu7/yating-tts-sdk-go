package YatingTtsSdk

type Input struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type Voice struct {
	Model string `json:"model"`
}

type AudioConfig struct {
	Encoding   string `json:"encoding"`
	SampleRate string `json:"sampleRate"`
}

type SynthesizeRequestDto struct {
	Input       Input       `json:"input"`
	Voice       Voice       `json:"voice"`
	AudioConfig AudioConfig `json:"audioConfig"`
}

type SynthesizeResponseDto struct {
	AudioContent string      `json:"audioContent"`
	AudioConfig  AudioConfig `json:"audioConfig"`
}

type SynthesizeErrorDto struct {
	Error      string   `json:"error"`
	Message    []string `json:"message"`
	StatusCode int      `json:"statusCode"`
}
