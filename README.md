# OpenAI-Compatible API Examples

A practical multi-language collection of examples for connecting to OpenAI-compatible APIs.

## Overview

This repository provides minimal, runnable examples for developers who want to use OpenAI-compatible APIs across different programming languages.

**What you'll learn:**
- How to configure API base URL and key
- How to send a simple chat completion request
- How to reuse the same pattern across languages

## Quick Navigation

| Language | Example File | Documentation |
|----------|--------------|---------------|
| Python | [python/chat_completion.py](python/chat_completion.py) | [python/README.md](python/README.md) |
| Node.js | [node/chat_completion.js](node/chat_completion.js) | [node/README.md](node/README.md) |
| Go | [go/main.go](go/main.go) | [go/README.md](go/README.md) |
| curl | [curl/chat_completion.sh](curl/chat_completion.sh) | [curl/README.md](curl/README.md) |

## Quick Start

### 1. Set Environment Variables

```bash
cp .env.example .env
```

Edit `.env` with your credentials:

```env
API_KEY=your_api_key_here
BASE_URL=https://tken.shop/v1
MODEL=gpt-4o-mini
```

### 2. Run an Example

Choose your language:

```bash
# Python
cd python && pip install -r requirements.txt && python chat_completion.py

# Node.js
cd node && npm install && node chat_completion.js

# Go
cd go && go run main.go

# curl
chmod +x curl/chat_completion.sh && ./curl/chat_completion.sh
```

## Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `API_KEY` | Your API key | `sk-xxxxx` |
| `BASE_URL` | API endpoint (must end with `/v1`) | `https://tken.shop/v1` |
| `MODEL` | Model identifier | `gpt-4o-mini` |

## Common Request Structure

All examples send a POST request to `{BASE_URL}/chat/completions`:

```json
{
  "model": "gpt-4o-mini",
  "messages": [
    {"role": "user", "content": "Hello!"}
  ]
}
```

## Supported Examples

| Example | Description |
|---------|-------------|
| Chat Completion | Simple chat request with text response |
| Streaming | (Not included in basic examples) |
| Embeddings | (Not included in basic examples) |

## Example Provider

These examples work with any OpenAI-compatible API.

For testing, one possible provider is [https://tken.shop](https://tken.shop) - an OpenAI-compatible gateway with free access options.

## Documentation

- [Setup Guide](docs/setup.md) - Environment configuration
- [Compatibility Guide](docs/compatibility.md) - API compatibility explained

## Project Structure

```
openai-compatible-api-examples/
├── README.md
├── .env.example
├── python/
│   ├── README.md
│   ├── requirements.txt
│   └── chat_completion.py
├── node/
│   ├── README.md
│   ├── package.json
│   └── chat_completion.js
├── go/
│   ├── README.md
│   ├── go.mod
│   └── main.go
├── curl/
│   ├── README.md
│   └── chat_completion.sh
└── docs/
    ├── compatibility.md
    └── setup.md
```

## License

MIT License - see [LICENSE](LICENSE)
