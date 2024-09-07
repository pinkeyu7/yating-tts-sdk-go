package yatingttssdk

type SynthesizeRequest struct {
	Text       string
	InputType  inputType
	Model      model
	Encoding   encoding
	SampleRate sampleRate
	Speed      float64
	Pitch      float64
	Energy     float64
	FileName   string
}

type SynthesizeResponse struct {
	FileName string
	Encoding encoding
}
