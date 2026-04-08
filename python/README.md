# Python Example

This directory contains a Python example for connecting to OpenAI-compatible APIs.

## Prerequisites

- Python 3.8+
- pip

## Installation

```bash
pip install -r requirements.txt
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
python chat_completion.py
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
