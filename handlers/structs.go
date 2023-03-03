package handlers

type UniInfo struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Map       string            `json:"map"`
}

type University struct {
	AlphaTwoCode  string   `json:"alpha_two_code"`
	WebPages      []string `json:"web_pages"`
	StateProvince string   `json:"state-province"`
	Name          string   `json:"name"`
	Domains       []string `json:"domains"`
	Country       string   `json:"country"`
}
