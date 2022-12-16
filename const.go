package YatingTtsSdk

const (
	ModelFemale1 = "zh_en_female_1"
	ModelFemale2 = "zh_en_female_2"
	ModelMale1   = "zh_en_male_1"
	ModelMay     = "zh_en_may_cw"

	TypeText = "text"
	TypeSsml = "ssml"

	EncodingMp3      = "MP3"
	EncodingLinear16 = "LINEAR16"

	SampleRate16k = "16K"
	SampleRate22k = "22K"
)

var Model = []string{ModelFemale1, ModelFemale2, ModelMale1, ModelMay}
var Type = []string{TypeText, TypeSsml}
var Encoding = []string{EncodingMp3, EncodingLinear16}
var SampleRate = []string{SampleRate16k, SampleRate22k}
