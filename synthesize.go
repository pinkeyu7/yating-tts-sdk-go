package yatingttssdk

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type YatingClient struct {
	URL string
	Key string
}

func NewClient(url string, key string) *YatingClient {
	return &YatingClient{
		URL: url,
		Key: key,
	}
}

func (c *YatingClient) Synthesize(req *SynthesizeRequest) (*SynthesizeResponse, error) {
	// validation
	err := c.validate(req)
	if err != nil {
		return nil, err
	}

	// generate post body
	dto := generator(req)
	payload, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	// send request
	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.URL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("key", c.Key)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// handle response
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		res := synthesizeErrorDto{}
		err = json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("%s", strings.Join(res.Message, ","))
	} else {
		res := synthesizeResponseDto{}
		err = json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			return nil, err
		}

		data, err := base64.StdEncoding.DecodeString(res.AudioContent)
		if err != nil {
			return nil, err
		}

		fileName := req.FileName

		switch req.Encoding {
		case EncodingMp3:
			fileName = fmt.Sprintf("%s.mp3", fileName)
		case EncodingLinear16:
			fileName = fmt.Sprintf("%s.wav", fileName)
		}

		file, err := os.Create(fileName)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = file.Close()
		}()

		_, err = file.Write(data)
		if err != nil {
			return nil, err
		}

		return &SynthesizeResponse{
			FileName: fileName,
			Encoding: req.Encoding,
		}, nil
	}
}
