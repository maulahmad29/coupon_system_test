package response

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value any    `json:"value"`
	Type  string `json:"type"`
	Param string `json:"param"`
}

type ReturnResponse struct {
	Success      string `json:"success"`
	Message      string `json:"message"`
	AlertMessage string `json:"alert_message"`
	ErrorMessage string `json:"error_message"`
	Data         any    `json:"data"`
}
