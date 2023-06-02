// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    newsDetailResponse, err := UnmarshalNewsDetailResponse(bytes)
//    bytes, err = newsDetailResponse.Marshal()

package model

import "encoding/json"

func UnmarshalNewsDetailResponse(data []byte) (NewsDetailResponse, error) {
	var r NewsDetailResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *NewsDetailResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type NewsDetailResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    NewsDetail `json:"data"`
}

type NewsDetail struct {
	ID        uint64 `json:"id"`
	GroupID   uint64 `json:"group_id"`
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	Photo     string `json:"photo"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}
