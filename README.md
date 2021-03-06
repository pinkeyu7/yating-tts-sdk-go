# Yating TTS SDK - Go version

## Available Key

- Please contact Yating TTS

### Input

```JSON
{
    "input": {
        "text": "sentence, the text that you want to generate audio",
        "type": "text"
    },
    "voice": {
        "model": "zh_en_female_1"
    },
    "audioConfig": {
        "encoding": "LINEAR16",
        "sampleRate": "16K"
    }
}
```

## Available Variable

### Input Type

| Name | Description                          |
| ---- | ------------------------------------ |
| text | Recognize input.text as pure text.   |
| ssml | Recognize input.text in ssml format. |

### Voice Model

| Name                    | Sample Rate | Description                       |
| ----------------------- | ----------- | --------------------------------- |
| zh_en_female_1 | 16k         | Yating speak Mandarin and English |
| zh_en_female_2 | 16k         | Yiru speak Mandarin and English   |
| zh_en_male_1   | 16k         | Jarvis speak Mandarin and English |

### Encoding Format

| Name     | File Extension | Description                 |
| -------- | -------------- | --------------------------- |
| MP3      | .mp3           | Compressed audio            |
| LINEAR16 | .wav           | Uncompromised audio quality |

### Sampling Rate

| Name | Description            |
| ---- | ---------------------- |
| 16K  | Sampling rate in 16khz |
