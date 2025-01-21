package models

type ZincSearchHighliht struct {
	Fields map[string]struct{} `json:"fields"`
}

type ZincSearchHit struct {
	Timestamp string              `json:"@timestamp"`
	Index     string              `json:"_index"`
	Type      string              `json:"_type"`
	Id        string              `json:"_id"`
	Score     float64             `json:"_score"`
	Source    Email               `json:"_source"`
	Highlight map[string][]string `json:"highlight"`
}

type ZincSearchParams struct {
	SearchType string             `json:"search_type"`
	Query      ZincSearchQuery    `json:"query"`
	From       int                `json:"from"`
	MaxResults int                `json:"max_results"`
	Source     []string           `json:"_source"`
	SortFields []string           `json:"sort_fields"`
	Highlight  ZincSearchHighliht `json:"highlight"`
}

type ZincSearchQuery struct {
	Term      string `json:"term"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type ZincSearchResponse struct {
	Took     float64 `json:"took"`
	TimeOut  bool    `json:"timed_out"`
	MaxScore float64 `json:"max_score"`
	Hits     struct {
		Total struct {
			Value int `json:"value"`
		}
		Hits []ZincSearchHit `json:"hits"`
	} `json:"hits"`
}
