package handler

import (
	"strings"

	"github.com/ak9024/okr-generator/internal/entity"
	"github.com/ak9024/okr-generator/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary OKR Generator
// @Accept json
// @Produce json
// @Param payload body entity.OKRGeneratorRequest true "entity.OKRGeneratorRequest"
// @Router /api/v1/okr-generator [post]
// @Success 200 {object} entity.OKRGeneratorResponse200 "entity.OKRGeneratorResponse200"
func (h *Handler) OKRGeneratorHandler(c *fiber.Ctx) error {
	og := new(entity.OKRGeneratorRequest)

	// parsing data body to struct `og`
	if err := c.BodyParser(og); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// validate data request
	if err := utils.ValidateStruct(og); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// get the server okr generator by pass the params request
	result, err := h.service.OKRGeneratorService(og)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	// insert status_code 200
	response := entity.OKRGeneratorResponse200{
		StatusCode: fiber.StatusOK,
	}

	// get the data from choices as strings
	for _, choice := range result.Choices {
		// split the results by `\n` and get first index as objective
		response.Objective = strings.SplitN(choice.Message.Content, "\n", -1)[0]
		// split the results by `\n` and start the data with second index as key result
		for _, krs := range strings.SplitN(choice.Message.Content, "\n", -1)[1:] {
			// convert key results into array type
			for _, kr := range strings.SplitN(krs, "\n", -1) {
				// filter: just insert key unempty value
				if kr != "" {
					response.KeyResults = append(response.KeyResults, entity.KeyResult{
						KeyResult: kr,
					})
				}
			}
		}
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
