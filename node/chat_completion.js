#!/usr/bin/env node
/**
 * Node.js Example: Chat Completion with OpenAI-Compatible API
 *
 * This script demonstrates how to make a simple chat completion request
 * using the official OpenAI JavaScript client with a custom base URL.
 */

require('dotenv').config();

const API_KEY = process.env.API_KEY;
const BASE_URL = process.env.BASE_URL || 'https://tken.shop/v1';
const MODEL = process.env.MODEL || 'gpt-4o-mini';

function checkConfig() {
    if (!API_KEY) {
        console.error('Error: API_KEY environment variable is not set');
        console.error('Set it with: export API_KEY=your_api_key_here');
        console.error('Or create a .env file based on .env.example');
        process.exit(1);
    }
}

async function main() {
    checkConfig();

    const { OpenAI } = require('openai');

    const client = new OpenAI({
        apiKey: API_KEY,
        baseURL: BASE_URL
    });

    console.log(`API: ${BASE_URL}`);
    console.log(`Model: ${MODEL}`);
    console.log('-'.repeat(40));

    try {
        const response = await client.chat.completions.create({
            model: MODEL,
            messages: [
                { role: 'user', content: 'Say hello in one sentence.' }
            ],
            max_tokens: 50,
            temperature: 0.7
        });

        console.log('Success!');
        console.log('-'.repeat(40));
        console.log(`Response: ${response.choices[0].message.content}`);
        console.log(`Model: ${response.model}`);
        console.log(`Tokens: ${response.usage.total_tokens}`);

    } catch (error) {
        console.error(`Error: ${error.message}`);
        process.exit(1);
    }
}

main();
