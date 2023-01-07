# EZ-GPT3

EZ-GPT3 makes it effortless to interface with the GPT-3 Completions API from command-line scripts. It takes the prompt content from standard input and sends it to the Completions API to generate a response.

## Installation

```sh
go get -u github.com/campbel/ez-gpt3/cmd/ezgpt3
go install github.com/campbel/ez-gpt3/cmd/ezgpt3
```

## Setup

An API token is necessary, which can be stored either in `~/.config/openai/token` or as an environment variable `OPENAI_API_TOKEN`.

Get a token from OpenAI [account API keys](https://beta.openai.com/account/api-keys).

## Usage

Simply pass the GPT-3 prompt into standard input:

```bash
echo "Fix this sentence so it sounds smart: wut are you doin?" | ezgpt3

What are you doing?
```

`model`, `tokens`, and `temp` can be passed in through flags.

```bash
echo "Write a short sentence describing a favorite hobby" | ezgpt3 -model text-curie-001 -tokens 20 -temp 0.5

I enjoy painting and sculpting.
```

For more information, refer to the [OpenAI completions API documentation](https://beta.openai.com/docs/api-reference/completions).

## Examples

For further ideas, take a look at the [examples](./examples).
