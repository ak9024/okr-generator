package entity

type KeyResult struct {
	KeyResult string `json:"key_result"`
}

type OKRGeneratorResponse200 struct {
	StatusCode int         `json:"status_code"`
	Objective  string      `json:"objective"`
	KeyResults []KeyResult `json:"key_results"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}
