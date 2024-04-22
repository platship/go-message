package dto

type BirdResultBase struct {
	ID      string `json:"id"`
	Service string `json:"service"`
	Event   string `json:"event"`
	Url     string `json:"url"`
	Status  string `json:"status"`
}

type BirdCallback struct {
	Results []*BirdResultBase
}

type ResChannelMessage struct {
	ID string `json:"id"`
}
