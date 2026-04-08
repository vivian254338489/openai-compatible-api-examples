# Setup Guide

## Environment Configuration

### 1. Get an API Key

Sign up with an OpenAI-compatible API provider and get an API key.

### 2. Set Environment Variables

Create a `.env` file in the project root (copy from `.env.example`):

```bash
cp .env.example .env
```

Edit `.env`:

```env
API_KEY=your_api_key_here
BASE_URL=https://tken.shop/v1
MODEL=gpt-4o-mini
```

Or export variables directly:

```bash
# Linux/macOS
export API_KEY=your_api_key_here
export BASE_URL=https://tken.shop/v1
export MODEL=gpt-4o-mini

# Windows (CMD)
set API_KEY=your_api_key_here
set BASE_URL=https://tken.shop/v1
set MODEL=gpt-4o-mini
```

## Base URL Format

The base URL must end with `/v1`.

| Provider | Correct Base URL |
|----------|-----------------|
| tken.shop | `https://tken.shop/v1` |
| OpenAI | `https://api.openai.com/v1` |
| LocalAI | `http://localhost:8080/v1` |
| Custom | `https://your-endpoint.com/v1` |

## Model Selection

Common model identifiers:

| Model | Provider | Notes |
|-------|----------|-------|
| gpt-4o-mini | OpenAI/tken.shop | Fast, cost-effective |
| gpt-4o | OpenAI/tken.shop | More capable |
| claude-3-haiku | Anthropic | Fast, affordable |
| claude-3-sonnet | Anthropic | Balanced |

Check your provider's documentation for available models.

## Running Examples

### Python

```bash
cd python
pip install -r requirements.txt
python chat_completion.py
```

### Node.js

```bash
cd node
npm install
node chat_completion.js
```

### Go

```bash
cd go
go run main.go
```

### curl

```bash
chmod +x curl/chat_completion.sh
./curl/chat_completion.sh
```

## Troubleshooting

### "API_KEY is not set"

- Ensure `.env` file exists in project root
- Or export API_KEY before running

### "404 Not Found"

- Check BASE_URL ends with `/v1`
- Verify provider URL is correct

### "401 Unauthorized"

- Verify API key is correct
- Check key hasn't expired

### "400 Bad Request"

- Verify model name is correct
- Check provider's available models

## Security Notes

- Never commit `.env` to version control
- Keep API keys secure
- For production, use backend proxy to protect keys
