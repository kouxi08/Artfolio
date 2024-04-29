package main

type Config struct {
	Name       string `json:"name"`
	RecordType string `json:"recordType"`
	TTL        string `json:"ttl"`
	Content    string `json:"content"`
}
