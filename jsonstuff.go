package main

import (
	"encoding/json"
	"os"
)

type Info struct {
	BPM     int    `json:"bpm"`
	Title   string `json:"title"`
	Spacing int    `json:"spacing"`
}

type Note struct {
	Time  float64 `json:"0"`
	Value string  `json:"1"`
}

type Song struct {
	Version int    `json:"version"`
	Info    Info   `json:"info"`
	Notes   []Note `json:"notes"`
}

func getMapFromFile(filename string) (*Song, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var song Song
	err = json.Unmarshal(data, &song)
	if err != nil {
		return nil, err
	}

	return &song, nil
}
