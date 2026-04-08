# Go Example

This directory contains a Go example for connecting to OpenAI-compatible APIs.

## Prerequisites

- Go 1.21+

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
go run main.go
```

## Expected Output

```
API: https://tken.shop/v1
Model: gpt-4o-mini
----------------------------------------
Success!
----------------------------------------
Response: Hello! How can I assist you today?
Tokens: 15
```

## Notes

- Uses only Go standard library (net/http, encoding/json)
- No external dependencies required
- Error handling for API errors and network issues
