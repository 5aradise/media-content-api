package types

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MediaContentType string

const (
	Text  MediaContentType = "text"
	Image MediaContentType = "image"
	Audio MediaContentType = "audio"
	Video MediaContentType = "video"
)

var (
	ErrWrongMediaContentTypeString = errors.New("types: wrong media content type string, available only: text | image | audio | video")
)

func (t MediaContentType) String() string {
	switch t {
	case Text:
		return "text"
	case Image:
		return "image"
	case Audio:
		return "audio"
	case Video:
		return "video"
	}
	return "undefined"
}

func NewMediaContentType(s string) (MediaContentType, error) {
	switch s {
	case "text":
		return Text, nil
	case "image":
		return Image, nil
	case "audio":
		return Audio, nil
	case "video":
		return Video, nil
	}
	return "", ErrWrongMediaContentTypeString
}

func (t MediaContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *MediaContentType) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return fmt.Errorf("types: failed to unmarshal MediaContentType string from JSON: %w", err)
	}

	nt, err := NewMediaContentType(s)
	if err != nil {
		return err
	}

	*t = nt
	return nil
}
