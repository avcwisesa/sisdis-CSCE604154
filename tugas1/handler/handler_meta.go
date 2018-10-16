package handler

// meta defines metadata for a response
type meta struct {
	// Code contains app level response code
	Code int
	// Message contains description of response code
	Message string
	// Error message if any
	Error string
}

var (
	metaSuccess = meta{Code: 210, Message: "Success"}

	metaContextTimeout  = meta{Code: 460, Message: "Context Timeout"}
)
