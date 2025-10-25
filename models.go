package main

type TVShow struct {
	ID           int    `json:"id"`
	OriginalName string `json:"original_name"`
}

type Episode struct {
	Season_number  int    `json:"season_number"`
	Episode_number int    `json:"episode_number"`
	Name           string `json:"name"`
}

type SearchResponse struct {
	Results []TVShow `json:"results"`
}

type DetailsResponse struct {
	Name     string    `json:"name"`
	Id       int       `json:"id"`
	Episodes []Episode `json:"episodes"`
}
