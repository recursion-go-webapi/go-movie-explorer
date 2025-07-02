package middleware

import (
	"net/http"
)

// APIError はHTTPステータスコード付きのカスタムエラー型
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// Error はerrorインターフェースを実装
func (e *APIError) Error() string {
	return e.Message
}

// NewAPIError は新しいAPIErrorを作成
func NewAPIError(statusCode int, message string) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
	}
}

// NewBadRequestError は400 Bad Requestエラーを作成
func NewBadRequestError(message string) *APIError {
	return &APIError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// NewNotFoundError は404 Not Foundエラーを作成
func NewNotFoundError(message string) *APIError {
	return &APIError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// NewInternalServerError は500 Internal Server Errorを作成
func NewInternalServerError(message string) *APIError {
	return &APIError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}