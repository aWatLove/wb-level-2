package repository

// ErrorHandler - структура обработчика ошибки
type ErrorHandler struct {
	Err        error `json:"err"`
	StatusCode int   `json:"status_code"`
}

// NewErrorHandler - метод создания структуры ErrorHandler
func NewErrorHandler(err error, statusCode int) *ErrorHandler {
	return &ErrorHandler{Err: err, StatusCode: statusCode}
}

// Error - реализация интерфейса error
func (m ErrorHandler) Error() string {
	return m.Err.Error()
}
