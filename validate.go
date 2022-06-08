package YatingTtsSdk

import "fmt"

func (c *YatingClient) validate(dto SynthesizeRequestDto) error {
	if !stringContains(Model, dto.Voice.Model) {
		return fmt.Errorf("model not in allow list(%v)", Model)
	}

	if !stringContains(Type, dto.Input.Type) {
		return fmt.Errorf("type not in allow list(%v)", Type)
	}

	if !stringContains(Encoding, dto.AudioConfig.Encoding) {
		return fmt.Errorf("encoding not in allow list(%v)", Encoding)
	}

	if !stringContains(SampleRate, dto.AudioConfig.SampleRate) {
		return fmt.Errorf("sampleRate not in allow list(%v)", SampleRate)
	}

	return nil
}

func stringContains(strSlice []string, matchStr string) bool {
	for _, each := range strSlice {
		if each == matchStr {
			return true
		}
	}
	return false
}
