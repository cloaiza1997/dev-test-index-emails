package services

import (
	"strings"
	"time"

	"github.com/cloaiza1997/dev-test-index-emails/config"
	"github.com/cloaiza1997/dev-test-index-emails/src/models"
	"github.com/cloaiza1997/dev-test-index-emails/src/utils"
)

type ZincSearchRepository struct{}

func (z ZincSearchRepository) Search(query models.QuerySearch) (models.EmailList, error) {
	page, limit := getPaginationValues(query.Page, query.Limit)

	result, err := utils.NewRequest[models.ZincSearchResponse](models.Request{
		Method: "POST",
		Url:    config.ApiConfig.ZincHost + "/" + config.ApiConfig.ZincIndex + "/_search",
		Body:   getQuery(query.Term, page, limit),
		Auth: models.RequestAuth{
			User: config.ApiConfig.ZincUser,
			Pass: config.ApiConfig.ZincPass,
		},
	})

	if err != nil {
		return models.EmailList{}, err
	}

	items := []models.EmailHighlight{}

	for _, email := range result.Hits.Hits {
		items = append(items, models.EmailHighlight{Email: email.Source, Highlight: email.Highlight})
	}

	list := models.EmailList{
		Pagination: utils.GetPagination(result.Hits.Total.Value, len(items), limit, page),
		Items:      items,
	}

	return list, nil
}

func getPaginationValues(page, limit int) (int, int) {
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	if page <= 0 {
		page = 1
	}

	return page, limit
}

func getQuery(term string, page, limit int) models.ZincSearchParams {
	now := time.Now()
	startTime := now.AddDate(-1, 0, 0).Format("2006-01-02T15:04:05Z")
	endTime := now.AddDate(1, 0, 0).Format("2006-01-02T15:04:05Z")

	term = strings.TrimSpace(term)
	from := (page * limit) - limit

	var searchType string

	if term == "" {
		searchType = "matchall"
	} else if len(strings.Split(term, " ")) > 1 {
		searchType = "match"
	} else {
		term = "*" + term + "*"
		searchType = "wildcard"
	}

	zsQuery := models.ZincSearchParams{
		SearchType: searchType,
		Query: models.ZincSearchQuery{
			Term:      strings.ReplaceAll(term, "@", "\\@"),
			StartTime: startTime,
			EndTime:   endTime,
		},
		From:       from,
		MaxResults: limit,
		SortFields: []string{"@timestamp"},
		Source: []string{
			"messageId",
			"date",
			"from",
			"to",
			"cc",
			"bcc",
			"subject",
			"xFrom",
			"xTo",
			"xCc",
			"xBcc",
			"xFolder",
			"xOrigin",
			"xFileName",
			"body",
			"path",
			"mainFolder",
		},
		Highlight: models.ZincSearchHighliht{
			Fields: map[string]struct{}{
				"bcc":     {},
				"body":    {},
				"cc":      {},
				"from":    {},
				"subject": {},
				"to":      {},
				"xBcc":    {},
				"xCc":     {},
				"xFrom":   {},
				"xTo":     {},
			},
		},
	}

	return zsQuery
}
