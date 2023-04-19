package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)


func GetCommitMsg(files string) {
	default_promt := "Create an easily understandable one liner commit message from the file provided ahead" + files
	authToken := viper.GetString("token")

	if authToken == ""{
		color.Red("Please provide auth token using `autocom auth` first")
		os.Exit(1)
	}

	client := openai.NewClient(authToken)
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