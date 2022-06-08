package YatingTtsSdk

const (
	ModelStandardFemale1 = "zh_en_standard_female_1"
	ModelStandardFemale2 = "zh_en_standard_female_2"
	ModelStandardMale1   = "zh_en_standard_male_1"

	TypeText = "text"
	TypeSsml = "ssml"

	EncodingMp3      = "MP3"
	EncodingLinear16 = "LINEAR16"

	SampleRate16k = "16K"
	SampleRate22k = "22K"
)

var Model = []string{ModelStandardFemale1, ModelStandardFemale2, ModelStandardMale1}
var Type = []string{TypeText, TypeSsml}
var Encoding = []string{EncodingMp3, EncodingLinear16}
var SampleRate = []string{SampleRate16k}
