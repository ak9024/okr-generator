package okr

import (
	"strings"

	sdk "github.com/ak9024/go-chatgpt-sdk"
	"github.com/gofiber/fiber/v2"
)

func (o *okr) OKRGeneratorService(og *OKRGeneratorRequest) (*OKRGeneratorResponse200, *sdk.ErrorResponse) {
	result, err := o.OKRGenerator(og)
	if err != nil {
		return nil, err
	}

	// insert status_code 200
	response := OKRGeneratorResponse200{
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
					response.KeyResults = append(response.KeyResults, KeyResult{
						KeyResult: kr,
					})
				}
			}
		}
	}

	return &response, nil
}
