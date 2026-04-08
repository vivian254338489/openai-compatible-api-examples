#!/usr/bin/env bash
#
# curl Example: Chat Completion with OpenAI-Compatible API
#
# This script demonstrates how to make a chat completion request using curl.
#

set -e

# Configuration from environment
API_KEY="${API_KEY:-}"
BASE_URL="${BASE_URL:-https://tken.shop/v1}"
MODEL="${MODEL:-gpt-4o-mini}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

function check_config() {
    if [ -z "$API_KEY" ]; then
        echo -e "${RED}Error: API_KEY environment variable is not set${NC}"
        echo "Set it with: export API_KEY=your_api_key_here"
        echo "Or create a .env file based on .env.example"
        exit 1
    fi
}

function main() {
    check_config

    echo "API: $BASE_URL"
    echo "Model: $MODEL"
    echo "----------------------------------------"

    response=$(curl -s -X POST "$BASE_URL/chat/completions" \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $API_KEY" \
        -d "{
            \"model\": \"$MODEL\",
            \"messages\": [
                {\"role\": \"user\", \"content\": \"Say hello in one sentence.\"}
            ],
            \"max_tokens\": 50,
            \"temperature\": 0.7
        }")

    # Check if response contains error
    if echo "$response" | grep -q '"error"'; then
        echo -e "${RED}Error:${NC}"
        echo "$response" | grep -o '"message":"[^"]*"' | head -1
        exit 1
    fi

    # Extract content
    content=$(echo "$response" | grep -o '"content":"[^"]*"' | head -1 | sed 's/"content":"//;s/"$//')

    echo -e "${GREEN}Success!${NC}"
    echo "----------------------------------------"
    echo "Response: $content"
}

main "$@"
