package YatingTtsSdk

import "fmt"

const (
	rangeMax = 1.5
	rangeMin = 0.5
)

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

	if !stringContains(SampleRateMapping[dto.Voice.Model], dto.AudioConfig.SampleRate) {
		return fmt.Errorf("sampleRate not in allow list(%v)", SampleRateMapping[dto.Voice.Model])
	}

	if !parameterRange(dto.Voice.Speed) {
		return fmt.Errorf("speed out of range(%.1f-%.1f)", rangeMin, rangeMax)
	}

	if !parameterRange(dto.Voice.Pitch) {
		return fmt.Errorf("pitch out of range(%.1f-%.1f)", rangeMin, rangeMax)
	}

	if !parameterRange(dto.Voice.Energy) {
		return fmt.Errorf("energy out of range(%.1f-%.1f)", rangeMin, rangeMax)
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

func parameterRange(input float64) bool {
	if input <= rangeMax && input >= rangeMin {
		return true
	}
	return false
}
