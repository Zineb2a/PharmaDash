import 'dotenv/config';
import express from 'express';
import bodyParser from 'body-parser';
import OpenAI from 'openai';

const app = express();
const PORT = 500; // Ensure this matches the port your frontend expects

// Middleware
app.use(bodyParser.json());

// OpenAI API configuration
const openai = new OpenAI({
    apiKey: process.env.OPENAI_API_KEY, // Ensure this is set in your .env file
});

const systemPrompt = `
You are a helpful and friendly customer support assistant for Pharmadash. Answer customer questions about services, orders, and any other inquiries with professionalism and care.
`;

// Chatbot endpoint
app.post('/api/chat', async (req, res) => {
    try {
        const messages = req.body; // Chat history from the frontend

        const response = await openai.chat.completions.create({
            model: 'gpt-3.5-turbo',
            messages: [{ role: 'system', content: systemPrompt }, ...messages],
        });

        res.json({ reply: response.choices[0].message.content });
    } catch (error) {
        console.error('Error communicating with OpenAI:', error);
        res.status(500).send({ error: 'Failed to process request.' });
    }
});

// Start the server
app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
});
