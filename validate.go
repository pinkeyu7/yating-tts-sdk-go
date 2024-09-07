package yatingttssdk

import "fmt"

const (
	rangeMax = 1.5
	rangeMin = 0.5
)

func (c *YatingClient) validate(req *SynthesizeRequest) error {
	if !sampleRateContains(req.Model, req.SampleRate) {
		return fmt.Errorf("sampleRate not in allow list(%v)", SampleRateMapping[req.Model])
	}

	if !parameterRange(req.Speed) {
		return fmt.Errorf("speed out of range(%.1f-%.1f)", rangeMin, rangeMax)
	}

	if !parameterRange(req.Pitch) {
		return fmt.Errorf("pitch out of range(%.1f-%.1f)", rangeMin, rangeMax)
	}

	if !parameterRange(req.Energy) {
		return fmt.Errorf("energy out of range(%.1f-%.1f)", rangeMin, rangeMax)
	}

	return nil
}

func sampleRateContains(md model, sr sampleRate) bool {
	sampleRates := SampleRateMapping[md]
	for _, subSr := range sampleRates {
		if subSr == sr {
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
