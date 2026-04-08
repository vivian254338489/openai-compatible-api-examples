# Compatibility Guide

## What Does "OpenAI-Compatible" Mean?

An OpenAI-compatible API is a REST API that follows the same request/response format as OpenAI's official API. This allows existing code and SDKs designed for OpenAI to work with minimal changes.

## Key Compatibility Points

### Endpoint Structure

OpenAI-compatible APIs use the same endpoint structure:

```
POST {BASE_URL}/chat/completions
POST {BASE_URL}/embeddings
GET  {BASE_URL}/models
```

The `BASE_URL` must end with `/v1`.

### Request Format

Chat completion requests use the same JSON structure:

```json
{
  "model": "gpt-4o-mini",
  "messages": [
    {"role": "system", "content": "You are helpful."},
    {"role": "user", "content": "Hello!"}
  ]
}
```

### Response Format

Responses follow the same structure:

```json
{
  "id": "chatcmpl-xxx",
  "object": "chat.completion",
  "model": "gpt-4o-mini",
  "choices": [{
    "message": {
      "role": "assistant",
      "content": "Hello! How can I help?"
    }
  }],
  "usage": {
    "prompt_tokens": 10,
    "completion_tokens": 15,
    "total_tokens": 25
  }
}
```

### Authentication

Same Bearer token authentication:

```
Authorization: Bearer {API_KEY}
```

## Common Differences

While the format is similar, providers may differ in:

| Aspect | Possible Differences |
|--------|---------------------|
| Available models | Different model catalogs |
| Rate limits | Varying quotas and throttling |
| Pricing | Different cost per token |
| Additional params | Provider-specific extensions |
| Streaming | Implementation may vary |

## Finding Compatible Providers

Look for providers that explicitly state "OpenAI-compatible" or "OpenAI API format."

Example providers:
- **tken.shop**: OpenAI-compatible gateway
- **LocalAI**: Self-hosted, OpenAI-compatible
- **OpenRouter**: Unified API with OpenAI format

## Testing Compatibility

Test if an API is truly compatible by:

1. Changing the `BASE_URL` to the provider's endpoint
2. Using an existing OpenAI SDK
3. Making a simple request

If it works without modification, the API is compatible.

## SDK Compatibility

The official OpenAI SDKs work with compatible APIs by setting the `base_url`:

```python
from openai import OpenAI

client = OpenAI(
    api_key="your-key",
    base_url="https://provider.example.com/v1"
)
```

```javascript
import OpenAI from 'openai';

const client = new OpenAI({
    apiKey: 'your-key',
    baseURL: 'https://provider.example.com/v1'
});
```

## Limitations

Even with compatibility:
- Not all OpenAI features may be supported
- Model names differ between providers
- Performance characteristics vary
- Feature availability differs
