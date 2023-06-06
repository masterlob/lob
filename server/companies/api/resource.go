package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

type CompaniesResource struct {
}

func NewCompaniesResource() *CompaniesResource {
	return &CompaniesResource{}
}

func Register(router *httprouter.Router) {
	companyResource := NewCompaniesResource()
	router.HandlerFunc(http.MethodGet, "/companies", companyResource.getCompanies)
}

func obtainIntFromQuery(req *http.Request, paramName string, defaultValue int) int {
	param := req.URL.Query().Get(paramName)
	if param == "" {
		return defaultValue
	}

	typedParam, err := strconv.Atoi(param)
	if err != nil {
		log.Info().Msgf("Failed to parse %s param: %s Cause: %s", paramName, param, err.Error())
		return defaultValue
	}
	return typedParam
}

// getCompanies returns a paginated list of fake companies.
func (r *CompaniesResource) getCompanies(w http.ResponseWriter, req *http.Request) {
	perPage := obtainIntFromQuery(req, "per_page", 30)
	page := obtainIntFromQuery(req, "page", 1)

	log.Info().Msgf("Fetching companies with per_page=%d and page=%d", perPage, page)

	companies := fetchCompanies()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("link", fmt.Sprintf("<https://api.github.com/user/58276/repos?page=%d>; rel=\"next\"", page+1))
	err := json.NewEncoder(w).Encode(companies)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode companies")
		http.Error(w, "Failed to encode companies", http.StatusInternalServerError)
		return
	}
}

func fetchCompanies() []CompanyResponse {
	return []CompanyResponse{
		{
			ID:          "1",
			Name:        "Acme Inc.",
			Description: "A company that makes everything.",
		},
		{
			ID:          "2",
			Name:        "Globex Corp.",
			Description: "A company that makes global thingy.",
		},
		{
			ID:          "3",
			Name:        "Soylent Corp.",
			Description: "A company that makes soylent green.",
		},
	}
}

type CompanyResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
