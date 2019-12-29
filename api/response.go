package api

import (
	"encoding/json"
	"errors"
)

// Response UniSender API response.
type Response struct {
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
	Code   string      `json:"code,omitempty"`
}

type jsonResponse struct {
	Result json.RawMessage `json:"result,omitempty"`
	Error  string          `json:"error,omitempty"`
	Code   string          `json:"code,omitempty"`
}

func (r *Response) UnmarshalJSON(data []byte) error {
	var j jsonResponse
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	r.Error = j.Error
	r.Code = j.Code
	if !r.IsError() {
		return json.Unmarshal(j.Result, &r.Result)
	}
	return nil
}

// IsError returns true if response has error.
func (r Response) IsError() bool {
	return r.Error != "" || r.Code != ""
}

// Err returns response error.
func (r Response) Err() error {
	if !r.IsError() {
		return nil
	}

	switch r.Code {
	case "invalid_api_key":
		return ErrInvalidAPIKey
	case "access_denied":
		return ErrAccessDenied
	case "unknown_method":
		return ErrUnknownMethod
	case "invalid_arg":
		return ErrInvalidArg
	case "not_enough_money":
		return ErrNotEnoughMoney
	case "retry_later":
		return ErrRetryLater
	case "api_call_limit_exceeded_for_api_key":
		return ErrAPICallLimitExceededForAPIKey
	case "api_call_limit_exceeded_for_ip":
		return ErrAPICallLimitExceededForIP
	default:
		return errors.New(r.Error)
	}
}
