#!/usr/bin/env python3
"""
Python Example: Chat Completion with OpenAI-Compatible API

This script demonstrates how to make a simple chat completion request
using the official OpenAI Python client with a custom base URL.
"""

import os
import sys

# Load .env file if python-dotenv is available
try:
    from dotenv import load_dotenv
    load_dotenv()
except ImportError:
    pass

# Configuration from environment
API_KEY = os.getenv("API_KEY")
BASE_URL = os.getenv("BASE_URL", "https://tken.shop/v1")
MODEL = os.getenv("MODEL", "gpt-4o-mini")


def check_config():
    """Verify required configuration is present."""
    if not API_KEY:
        print("Error: API_KEY environment variable is not set")
        print("Set it with: export API_KEY=your_api_key_here")
        print("Or create a .env file based on .env.example")
        sys.exit(1)


def main():
    check_config()

    try:
        from openai import OpenAI
    except ImportError:
        print("Error: openai package not installed")
        print("Run: pip install -r requirements.txt")
        sys.exit(1)

    # Initialize client with custom base URL
    client = OpenAI(
        api_key=API_KEY,
        base_url=BASE_URL
    )

    print(f"API: {BASE_URL}")
    print(f"Model: {MODEL}")
    print("-" * 40)

    try:
        response = client.chat.completions.create(
            model=MODEL,
            messages=[
                {"role": "user", "content": "Say hello in one sentence."}
            ],
            max_tokens=50,
            temperature=0.7
        )

        print("Success!")
        print("-" * 40)
        print(f"Response: {response.choices[0].message.content}")
        print(f"Model: {response.model}")
        print(f"Tokens: {response.usage.total_tokens}")

    except Exception as e:
        print(f"Error: {e}")
        sys.exit(1)


if __name__ == "__main__":
    main()
