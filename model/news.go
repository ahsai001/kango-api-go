// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    newsResponse, err := UnmarshalNewsResponse(bytes)
//    bytes, err = newsResponse.Marshal()

package model

import "encoding/json"

func UnmarshalNewsResponse(data []byte) (NewsResponse, error) {
	var r NewsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NewsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NewsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []News `json:"data"`
}

type News struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}
