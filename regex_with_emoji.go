package main

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"regexp"
	"strings"
)

type RegexWithEmoji struct {
	data  map[string][]string
	regex map[string]*regexp.Regexp
}

func NewRegexWithEmoji(data string) (*RegexWithEmoji, error) {
	compressedData, err := base64.StdEncoding.DecodeString(strings.TrimSpace(data))

	if err != nil {
		return nil, err
	}

	reader := flate.NewReader(bytes.NewReader(compressedData))

	defer reader.Close()

	var dataJson map[string][]string

	if err := json.NewDecoder(reader).Decode(&dataJson); err != nil {
		return nil, err
	}

	regexMap := make(map[string]*regexp.Regexp)

	for k := range dataJson {
		regexMap[k] = regexp.MustCompile(k)
	}

	return &RegexWithEmoji{data: dataJson, regex: regexMap}, nil
}

func (receiver *RegexWithEmoji) GetEmoji(text string) string {
	for k, v := range receiver.data {
		if receiver.matches(text, k) {
			return receiver.getRandomStringByMap(v)
		}
	}
	return ""
}

func (receiver *RegexWithEmoji) matches(text string, template string) bool {
	regex := receiver.regex[template]

	return regex.FindAllString(text, -1) != nil
}

func (receiver *RegexWithEmoji) getRandomStringByMap(data []string) string {
	return data[rand.Intn(len(data))]
}
