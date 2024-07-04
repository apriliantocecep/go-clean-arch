package errors

type ErrorCode string

const (
	EmailExists     ErrorCode = "EMAIL_EXISTS"
	ServerError     ErrorCode = "INTERNAL_SERVER_ERROR"
	CommitError     ErrorCode = "COMMIT_ERROR"
	TokenError      ErrorCode = "TOKEN_ERROR"
	InvalidToken    ErrorCode = "INVALID_TOKEN"
	NotFound        ErrorCode = "NOT_FOUND"
	InvalidPassword ErrorCode = "INVALID_PASSWORD"
	InvalidJson     ErrorCode = "INVALID_JSON"
)
