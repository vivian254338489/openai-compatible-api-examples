# Node.js Example

This directory contains a Node.js example for connecting to OpenAI-compatible APIs.

## Prerequisites

- Node.js 18+
- npm

## Installation

```bash
npm install
```

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
node chat_completion.js
```

## Expected Output

```
API: https://tken.shop/v1
Model: gpt-4o-mini
----------------------------------------
Success!
----------------------------------------
Response: Hello! How can I assist you today?
Model: gpt-4o-mini
Tokens: 15
```
