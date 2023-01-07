package main

import (
	"io"
	"net/http"
	"os"

	"github.com/cambel/easy-gpt3/openai"
)

func main() {
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

	text, err := openai.NewClient(http.DefaultClient, token).GetCompletion(string(data))
	if err != nil {
		panic(err)
	}

	println(text)
}
