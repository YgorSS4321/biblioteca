package response

type Error struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
}

func NewError(statusCode int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Msg:        msg,
	}
}
