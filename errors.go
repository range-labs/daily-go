package daily

import "fmt"

var (
	// HTTP Errors.
	ErrBadRequest      = "bad request"
	ErrUnauthorized    = "unauthorized"
	ErrTooManyRequests = "too many requests"
	ErrInternal        = "internal error"
	ErrUnexpected      = "unexpected error"

	// Other errors.
	ErrParseError = "json parse error"
)

// Error represents error information related to an API call.
type Error struct {
	Message    string
	StatusCode int
	Details    *ErrorDetails
	RawDetails string
}

func (e Error) Error() string {
	if e.Details != nil {
		return fmt.Sprintf("daily: %s (status: %d, %s)", e.Message, e.StatusCode, e.Details)
	} else {
		return fmt.Sprintf("daily: %s (status: %d, details: %s)", e.Message, e.StatusCode, e.RawDetails)
	}
}

// ErrorDetails is the daily API error response.
type ErrorDetails struct {
	ErrorCode string `json:"error"`
	ErrorInfo string `json:"info"`
}

func (ed ErrorDetails) String() string {
	return fmt.Sprintf("code: %s, info: %s", ed.ErrorCode, ed.ErrorInfo)
}
