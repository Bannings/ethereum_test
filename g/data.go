package g

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
