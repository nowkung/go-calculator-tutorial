package models

import constants "go-calculator-tutorial/internal/constants"

type (
	TempApiResponse struct {
		Code      int        `json:"code"`
		Message   string     `json:"message"`
		Data	  map[string]string 	 `json:"data"`
		ErrorData *ErrorData `json:"error,omitempty"`
	}
	ErrorData struct {
		ErrorTitle   string `json:"error_title"`
		ErrorMessage string `json:"error_message"`
	}
)

func ResponseSuccess(d map[string]string) *TempApiResponse {
	return &TempApiResponse{
		Code:    constants.CODE_SUCCESS,
		Message: constants.MESSAGE_SUCCESS,
		Data: d,
	}
}

func ResponseGenericError() *TempApiResponse {
	return &TempApiResponse{
		Code:    constants.CODE_GENERIC_ERROR,
		Message: constants.MESSAGE_GENERIC_ERROR,
		ErrorData: &ErrorData{
			ErrorTitle:   "Connection error",
			ErrorMessage: "Please try again",
		},
	}
}

func ResponseBadRequest() *TempApiResponse {
	return &TempApiResponse{
		Code:    constants.CODE_BAD_REQUEST,
		Message: constants.MESSAGE_BAD_REQUEST,
		ErrorData: &ErrorData{
			ErrorTitle:   "Wrong Unit",
			ErrorMessage: "Please try again",
		},
	}
}
