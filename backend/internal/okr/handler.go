package okr

import (
	"github.com/ak9024/okr-generator/utils"
	"github.com/gofiber/fiber/v2"
)

// @Summary OKR Generator
// @Accept json
// @Produce json
// @Param payload body OKRGeneratorRequest true "OKRGeneratorRequest"
// @Router /api/v1/okr-generator [post]
// @Success 200 {object} OKRGeneratorResponse200 "OKRGeneratorResponse200"
// @Failure 400 {object} OKRGeneratorResponseError "OKRGeneratorResponseError"
// @Failure 500 {object} OKRGeneratorResponseError "OKRGeneratorResponseError"
func (o *OKR) OKRGeneratorHandler(c *fiber.Ctx) error {
	og := new(OKRGeneratorRequest)

	// parsing data body to struct `og`
	if err := c.BodyParser(og); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(OKRGeneratorResponseError{
			StatusCode: fiber.StatusBadRequest,
			Messages:   err,
		})
	}

	// validate data request
	if err := utils.ValidateStruct(og); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(OKRGeneratorResponseError{
			StatusCode: fiber.StatusBadRequest,
			Messages:   err,
		})
	}

	// get the server okr generator by pass the params request
	result, err := o.OKRGeneratorService(og)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(OKRGeneratorResponseError{
			StatusCode: fiber.StatusInternalServerError,
			Messages:   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
