package entity

type OKRGeneratorRequest struct {
	Objective string `json:"objective" validate:"required"`
	Translate string `json:"translate" validate:"required"`
}
