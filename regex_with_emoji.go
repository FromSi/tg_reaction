package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"regexp"
	"strings"
)

type RegexWithEmoji struct {
	data map[string][]string
}

func NewRegexWithEmoji(data string) (*RegexWithEmoji, error) {
	compressedData, err := base64.StdEncoding.DecodeString(strings.TrimSpace(data))

	if err != nil {
		return nil, err
	}

	reader := flate.NewReader(bytes.NewReader(compressedData))

	defer reader.Close()

	var buf bytes.Buffer

	_, err = io.Copy(&buf, reader)

	if err != nil {
		return nil, err
	}

	var dataJson map[string][]string

	err = json.Unmarshal([]byte(buf.String()), &dataJson)

	if err != nil {
		return nil, err
	}

	return &RegexWithEmoji{data: dataJson}, nil
}

func (receiver RegexWithEmoji) GetEmoji(text string) string {
	for k, v := range receiver.data {
		if receiver.matches(text, k) {
			return receiver.getRandomStringByMap(v)
		}
	}

	return ""
}

func (receiver RegexWithEmoji) matches(text string, template string) bool {
	regex := regexp.MustCompile(template)

	return regex.FindAllString(text, -1) != nil
}

func (receiver RegexWithEmoji) getRandomStringByMap(data []string) string {
	return data[rand.Intn(len(data))]
}
