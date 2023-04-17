package ai

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)


func GetCommitMsg(files string) {
	default_promt := "Create an easily understandable one liner commit message from the file provided ahead" + files

	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: default_promt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}