# curl Example

This directory contains a curl example for connecting to OpenAI-compatible APIs.

## Prerequisites

- curl (available on most systems)
- bash

## Configuration

Set environment variables:

```bash
export API_KEY=your_api_key_here
export BASE_URL=https://tken.shop/v1
export MODEL=gpt-4o-mini
```

Or create a `.env` file in the project root (copy from `.env.example`).

## Run

```bash
chmod +x chat_completion.sh
./chat_completion.sh
```

Or run directly:

```bash
curl -X POST "$BASE_URL/chat/completions" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $API_KEY" \
  -d '{
    "model": "'"$MODEL"'",
    "messages": [{"role": "user", "content": "Say hello in one sentence."}]
  }'
```

## Expected Output

```
API: https://tken.shop/v1
Model: gpt-4o-mini
----------------------------------------
Success!
----------------------------------------
Response: Hello! How can I assist you today?
```

## Notes

- Uses only standard curl
- Works on Linux, macOS, and Windows (with Git Bash or WSL)
- Good for quick testing and scripting
