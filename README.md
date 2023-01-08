# EZ-GPT3

EZ-GPT3 is a convenient, command-line interface for the OpenAI GPT-3 Completions API. It allows users to effortlessly generate responses to a given prompt by simply passing the prompt into standard input.

## Installation

```sh
go get -u github.com/campbel/ez-gpt3/cmd/ezgpt3
go install github.com/campbel/ez-gpt3/cmd/ezgpt3
```

## Setup

In order to interact with the GPT-3 Completions API, an API token is required. This token can either be stored in `~/.config/openai/token` or as an environment variable `OPENAI_API_TOKEN`. To obtain a token, users must create an account at OpenAI and retrieve their API keys [here](https://beta.openai.com/account/api-keys).

## Usage

Using EZ-GPT3 is as simple as passing a prompt into standard input:

```bash
echo "Fix this sentence so it sounds smart: wut are you doin?" | ezgpt3

What are you doing?
```

In addition, users can also customize their query by specifying the `model`, `tokens`, and `temp` flags.

```bash
echo "Write a short sentence describing a favorite hobby" | ezgpt3 -model text-curie-001 -tokens 20 -temp 0.5

I enjoy painting and sculpting.
```

For more information, please refer to the [OpenAI completions API documentation](https://beta.openai.com/docs/api-reference/completions).

## Examples

For further ideas, take a look at the [examples](./examples).
