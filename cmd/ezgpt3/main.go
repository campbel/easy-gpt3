package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/cambel/easy-gpt3/openai"
)

func main() {

	modelFlag := flag.String("model", openai.DefaultCompletionOptions.Model, "The model to use")
	tokenFlag := flag.Int("tokens", openai.DefaultCompletionOptions.MaxTokens, "The max number of tokens to generate")
	tempFlag := flag.Float64("temp", float64(openai.DefaultCompletionOptions.Temperature), "The temperature to use")
	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		content, err := os.ReadFile(dir + "/.config/openai/token")
		if err != nil {
			panic(err)
		}
		token = string(content)
	}

	options := openai.CompletionOptions{
		Model:       *modelFlag,
		MaxTokens:   *tokenFlag,
		Temperature: *tempFlag,
	}

	text, err := openai.NewClient(http.DefaultClient, token).GetCompletion(string(data), options)
	if err != nil {
		panic(err)
	}

	println(text)
}
