package service

import (
	sdk "github.com/ak9024/go-chatgpt-sdk"
	"github.com/ak9024/okr-generator/internal/entity"
)

func (s *Service) OKRGeneratorService(og *entity.OKRGeneratorRequest) (*sdk.ModelChatResponse, *sdk.ErrorResponse) {
	c := sdk.NewConfig(sdk.Config{
		OpenAIKey: s.config.GetString("chatgpt.token"),
	})

	// set default language to bahasa
	// set the params language from request body
	content := ContentGenerator(og.Translate)

	resp, err := c.ChatCompletions(sdk.ModelChat{
		Model: "gpt-3.5-turbo",
		Messages: []sdk.Message{
			{
				Role:    "system",
				Content: content, // get content to train user system
			},
			{
				Role:    "user",
				Content: og.Objective, // get the object by user submit
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
