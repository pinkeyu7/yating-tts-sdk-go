package yatingttssdk

type model string
type inputType string
type encoding string
type sampleRate string

const (
	ModelZhEnFemale1 model = "zh_en_female_1"
	ModelZhEnFemale2 model = "zh_en_female_2"
	ModelZhEnMale1   model = "zh_en_male_1"
	ModelZhEnMale2   model = "zh_en_male_2"
	ModelTaiFemale1  model = "tai_female_1"
	ModelTaiFemale2  model = "tai_female_2"
	ModelTaiMale1    model = "tai_male_1"

	InputTypeText inputType = "text"
	InputTypeSsml inputType = "ssml"

	EncodingMp3      encoding = "MP3"
	EncodingLinear16 encoding = "LINEAR16"

	SampleRate16k sampleRate = "16K"
	SampleRate22k sampleRate = "22K"
)

var SampleRateMapping = map[model][]sampleRate{
	ModelZhEnFemale1: {SampleRate16k, SampleRate22k},
	ModelZhEnFemale2: {SampleRate16k, SampleRate22k},
	ModelZhEnMale1:   {SampleRate16k, SampleRate22k},
	ModelZhEnMale2:   {SampleRate22k},
	ModelTaiFemale1:  {SampleRate16k, SampleRate22k},
	ModelTaiFemale2:  {SampleRate16k, SampleRate22k},
	ModelTaiMale1:    {SampleRate16k, SampleRate22k},
}
