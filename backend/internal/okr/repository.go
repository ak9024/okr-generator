package okr

import sdk "github.com/ak9024/go-chatgpt-sdk"

func (o *OKR) OKRGenerator(og *OKRGeneratorRequest) (*sdk.ModelChatResponse, *sdk.ErrorResponse) {
	c := sdk.NewConfig(sdk.Config{
		OpenAIKey: o.Config.GetString("chatgpt.token"),
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
