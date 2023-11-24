package main

type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Priority int    `json:"priority"`
}

type Rrset struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Changetype string   `json:"changetype"`
	TTL        int      `json:"ttl"`
	Records    []Record `json:"records"`
}

type Filler_DTO struct {
	Rrsets []Rrset `json:"rrsets"`
}
