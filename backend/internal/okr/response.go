package okr

type KeyResult struct {
	KeyResult string `json:"key_result"`
}

type OKRGeneratorResponse200 struct {
	StatusCode int         `json:"status_code"`
	Objective  string      `json:"objective"`
	KeyResults []KeyResult `json:"key_results"`
}
