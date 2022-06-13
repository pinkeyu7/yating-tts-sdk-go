package YatingTtsSdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type YatingClient struct {
	Url string
	Key string
}

func NewClient(url string, key string) *YatingClient {
	return &YatingClient{Url: url, Key: key}
}

func (c *YatingClient) Synthesize(text, inputType, model, encoding, sampleRate, fileName string) error {
	dto := generator(text, inputType, model, encoding, sampleRate)

	err := c.validate(dto)
	if err != nil {
		return err
	}

	payload, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPost, c.Url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("key", c.Key)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		res := SynthesizeErrorDto{}
		err = json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			return err
		}

		return fmt.Errorf("%s", strings.Join(res.Message, ","))
	} else {
		res := SynthesizeResponseDto{}
		err = json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			return err
		}

		data, err := base64.StdEncoding.DecodeString(res.AudioContent)
		if err != nil {
			return err
		}

		switch dto.AudioConfig.Encoding {
		case EncodingMp3:
			fileName = fmt.Sprintf("%s.mp3", fileName)
		default:
			fileName = fmt.Sprintf("%s.wav", fileName)
		}

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
