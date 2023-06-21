package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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
	router.HandlerFunc(http.MethodGet, "/api/companies", companyResource.getCompanies)
	router.HandlerFunc(http.MethodPost, "/api/companies", companyResource.addCompany)
	router.HandlerFunc(http.MethodPost, "/api/companies/{id}", companyResource.updateCompany)
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

	companies := fetchCompanies(req.Context(), PageParams{Page: page, PerPage: perPage})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("link", fmt.Sprintf("<%s?page=%d>; rel=\"next\"", serverUrl(req), page+1))

	err := json.NewEncoder(w).Encode(companies)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode companies")
		http.Error(w, "Failed to encode companies", http.StatusInternalServerError)
		return
	}
}

// addCompany adds a new company to the list of fake companies.
func (r *CompaniesResource) addCompany(w http.ResponseWriter, req *http.Request) {
	var company CompanyResponse
	err := json.NewDecoder(req.Body).Decode(&company)
	if err != nil {
		// FIXME: https://github.com/go-playground/validator
		log.Error().Err(err).Msg("Failed to decode company")
		http.Error(w, "Failed to decode company", http.StatusBadRequest)
		return
	}

	// FIXME: Handle errors properly with an error hierarchy.
	addCompany(req.Context(), company)
	w.WriteHeader(http.StatusNoContent)
}

// updateCompany updates an existing company in the list of fake companies.
func (r *CompaniesResource) updateCompany(w http.ResponseWriter, req *http.Request) {
	var company CompanyResponse
	err := json.NewDecoder(req.Body).Decode(&company)
	if err != nil {
		log.Error().Err(err).Msg("Failed to decode company")
		http.Error(w, "Failed to decode company", http.StatusBadRequest)
		return
	}

	// FIXME: Handle errors properly with an error hierarchy.
	err = updateCompany(req.Context(), company)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update company")
		http.Error(w, "Failed to update company", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type PageParams struct {
	PerPage int
	Page    int
}

var companies = []CompanyResponse{
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

func fetchCompanies(ctx context.Context, pageParams PageParams) []CompanyResponse {
	return companies
}

func addCompany(ctx context.Context, company CompanyResponse) error {
	if company.ID != "" {
		return fmt.Errorf("company ID must not be set on creation")
	}
	company.ID = uuid.NewString()
	companies = append(companies, company)
	return nil
}

func updateCompany(ctx context.Context, company CompanyResponse) error {
	for i, cc := range companies {
		if cc.ID == company.ID {
			companies[i] = company
			return nil
		}
	}
	return fmt.Errorf("company with id %s not found", company.ID)
}

func serverUrl(req *http.Request) string {
	return fmt.Sprintf("%s://%s%s", protocolFormTLS(req), req.Host, req.RequestURI)
}

func protocolFormTLS(req *http.Request) string {
	if req.TLS != nil {
		return "https://"
	}
	return "http://"
}

type CompanyResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
