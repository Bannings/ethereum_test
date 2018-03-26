package data

import (
	"net/http"

	"github.com/go-chi/render"
)

type SuccResponse struct {
	AppCode    string      `json:"retCode,omitempty"` // application-specific error code
	StatusText string      `json:"retMsg,omitempty"`  // application-level error message, for debugging
	Data       interface{} `json:"retData,omitempty"`
}

func NewSuccResponse(data interface{}) SuccResponse {
	return SuccResponse{
		AppCode:    "0000",
		StatusText: "success",
		Data:       data,
	}
}

func (sr *SuccResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.SetContentType(render.ContentTypeJSON)
	render.Status(r, http.StatusOK)
	return nil
}

// ----------

type M map[string]interface{}

func (m M) Bind(r *http.Request) error {
	return nil
}

type Company struct {
	CompanyId          string `json:"companyId"`
	Type               string `json:"type"`
	Name               string `json:"name"`
	Address            string `json:"address"`
	BusinessLicence    string `json:"businessLicence"`
	Usci               string `json:"usci"`
	LegalPerson        string `json:"legalPerson"`
	LegalIdcardNum     string `json:"legalIdcardNum"`
	BankAccount        string `json:"bankAccount"`
	BankName           string `json:"bankName"`
	Contact            string `json:"contact"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	ContactIdcardNum   string `json:"contactIdcardNum"`
	LegalPhone         string `json:"legalPhone"`
	LegalMail          string `json:"legalMail"`
	BankBranchProvince string `json:"bankBranchProvince"`
	BankBranchCity     string `json:"bankBranchCity"`
	BankBranchName     string `json:"bankBranchName"`
	BankAccountName    string `json:"bankAccountName"`
}

func (c *Company) Bind(r *http.Request) error {
	return nil
}
