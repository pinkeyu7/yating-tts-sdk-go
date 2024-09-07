package yatingttssdk

func generator(req *SynthesizeRequest) *synthesizeRequestDto {
	return &synthesizeRequestDto{
		Input: input{
			Text: req.Text,
			Type: req.InputType,
		},
		Voice: voice{
			Model:  req.Model,
			Speed:  req.Speed,
			Pitch:  req.Pitch,
			Energy: req.Energy,
		},
		AudioConfig: audioConfig{
			Encoding:   req.Encoding,
			SampleRate: req.SampleRate,
		},
	}
}

type input struct {
	Text string    `json:"text"`
	Type inputType `json:"type"`
}

type voice struct {
	Model  model   `json:"model"`
	Speed  float64 `json:"speed"`
	Pitch  float64 `json:"pitch"`
	Energy float64 `json:"energy"`
}

type audioConfig struct {
	Encoding   encoding   `json:"encoding"`
	SampleRate sampleRate `json:"sampleRate"`
}

type synthesizeRequestDto struct {
	Input       input       `json:"input"`
	Voice       voice       `json:"voice"`
	AudioConfig audioConfig `json:"audioConfig"`
}

type synthesizeResponseDto struct {
	AudioContent string      `json:"audioContent"`
	AudioConfig  audioConfig `json:"audioConfig"`
}

type synthesizeErrorDto struct {
	Error      string   `json:"error"`
	Message    []string `json:"message"`
	StatusCode int      `json:"statusCode"`
}
