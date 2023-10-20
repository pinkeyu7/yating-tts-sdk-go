package YatingTtsSdk

const (
	ModelZhEnFemale1 = "zh_en_female_1"
	ModelZhEnFemale2 = "zh_en_female_2"
	ModelZhEnMale1   = "zh_en_male_1"

	ModelTaiFemale1 = "tai_female_1"
	ModelTaiFemale2 = "tai_female_2"
	ModelTaiMale1   = "tai_male_1"

	TypeText = "text"
	TypeSsml = "ssml"

	EncodingMp3      = "MP3"
	EncodingLinear16 = "LINEAR16"

	SampleRate16k = "16K"
	SampleRate22k = "22K"
)

var Model = []string{ModelZhEnFemale1, ModelZhEnFemale2, ModelZhEnMale1, ModelTaiFemale1, ModelTaiFemale2, ModelTaiMale1}
var Type = []string{TypeText, TypeSsml}
var Encoding = []string{EncodingMp3, EncodingLinear16}
var SampleRateMapping = map[string][]string{
	ModelZhEnFemale1: {SampleRate16k, SampleRate22k},
	ModelZhEnFemale2: {SampleRate16k, SampleRate22k},
	ModelZhEnMale1:   {SampleRate16k, SampleRate22k},
	ModelTaiFemale1:  {SampleRate16k},
	ModelTaiFemale2:  {SampleRate16k},
	ModelTaiMale1:    {SampleRate16k},
}
