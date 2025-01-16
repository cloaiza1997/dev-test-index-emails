package controllers

import (
	"net/http"

	"github.com/cloaiza1997/dev-test-tr-emails/src/models"
	"github.com/cloaiza1997/dev-test-tr-emails/src/services"
	"github.com/cloaiza1997/dev-test-tr-emails/src/utils"
)

var zs = services.ZincSearchRepository{}

func GetEmails(w http.ResponseWriter, r *http.Request) {
	query := models.QuerySearch{
		Term:  utils.GetQueryParam(r, "term"),
		Limit: utils.GetQueryParamInt(r, "limit"),
		Page:  utils.GetQueryParamInt(r, "page"),
	}

	result, err := zs.Search(query)
	var status int

	if len(result.Items) == 0 {
		status = http.StatusNotFound
	} else {
		status = http.StatusOK
	}

	utils.DoResponse(w, status, result, err)
}
