package g

import (
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrInvalidNode            = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1001, ErrorText: "Invalid node, request is denied"}
	ErrCallSmartContract      = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1004, ErrorText: "call smart contract error"}
	ErrVerifySignature        = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1014, ErrorText: "verify signature fail"}
	ErrCompanyIdAlreadyExist  = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1024, ErrorText: "companyId already exist"}
	ErrCompanyIdNotExist      = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1025, ErrorText: "companyId not exist"}
	ErrArIdNotExist           = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1009, ErrorText: "arId not exist"}
	ErrArIdAlreadyExist       = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1013, ErrorText: "arId already exist"}
	ErrDiscountIdNotExist     = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1027, ErrorText: "discountId not exist"}
	ErrDiscountIdAlreadyExist = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1026, ErrorText: "discountId already exist"}
	ErrPayIdNotExist          = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1029, ErrorText: "payId not exist"}
	ErrPayIdAlreadyExist      = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1028, ErrorText: "payId already exist"}
	ErrContractNoNotExist     = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1031, ErrorText: "contractNo not exist"}
	ErrContractNoAlreadyExist = &ErrResponse{HTTPStatusCode: http.StatusOK, AppCode: 1030, ErrorText: "contractNo already exist"}
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	AppCode   int64  `json:"retCode,omitempty"` // application-specific error code
	ErrorText string `json:"retMsg,omitempty"`  // application-level error message, for debugging
	Data      string `json:"retData,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.SetContentType(render.ContentTypeJSON)
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		AppCode:        http.StatusBadRequest,
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		AppCode:        http.StatusInternalServerError,
		ErrorText:      err.Error(),
	}
}
