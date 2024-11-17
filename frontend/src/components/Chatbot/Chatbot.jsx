import React, { useState, useEffect, useRef } from 'react';
import './Chatbot.css';

import {
  Box,
  Button,
  TextField,
  Typography,
  Stack,
  CircularProgress,
} from '@mui/material';
import { Send } from '@mui/icons-material';

const Chatbot = () => {
  const [messages, setMessages] = useState([
    { role: 'assistant', content: 'Welcome to Pharmadash Customer Service! ' },
  ]);
  const [message, setMessage] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const messagesEndRef = useRef(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const handleSendMessage = async () => {
    if (!message.trim() || isLoading) return;

    const userMessage = { role: 'user', content: message };
    setMessages((prev) => [...prev, userMessage]);
    setMessage('');
    setIsLoading(true);

    try {
      const response = await fetch('http://localhost:5001/api/chat', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ messages: [...messages, userMessage] }),
      });

      const data = await response.json();

      setMessages((prev) => [
        ...prev,
        { role: 'assistant', content: data.reply },
      ]);
    } catch (error) {
      setMessages((prev) => [
        ...prev,
        { role: 'assistant', content: 'Something went wrong. Please try again.' },
      ]);
    } finally {
      setIsLoading(false);
    }
  };

  const handleKeyPress = (e) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  return (
    <Box
      sx={{
        height: '70vh', // Adjusted for default height
        width: '100%',
        maxWidth: '900px',
        margin: '20px auto',
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: '#f9f9f9',
        borderRadius: '12px',
        boxShadow: '0px 4px 8px rgba(0, 0, 0, 0.1)',
      }}
    >
      {/* Header */}
      <Box
        sx={{
          padding: '20px',
          backgroundColor: '#e74c3c',
          color: '#fff',
          textAlign: 'center',
          fontWeight: 'bold',
          fontSize: '20px',
          borderRadius: '12px 12px 0 0',
        }}
      >
        Pharmadash Support
      </Box>

      {/* Messages */}
      <Stack
        spacing={2}
        sx={{
          flex: 1,
          overflowY: 'auto',
          padding: '20px',
          backgroundColor: '#fff',
        }}
      >
        {messages.map((msg, idx) => (
          <Box
            key={idx}
            sx={{
              alignSelf: msg.role === 'assistant' ? 'flex-start' : 'flex-end',
              backgroundColor: msg.role === 'assistant' ? '#f0f0f0' : '#e74c3c',
              color: msg.role === 'assistant' ? '#333' : '#fff',
              padding: '10px 15px',
              borderRadius: '15px',
              maxWidth: '70%',
              wordWrap: 'break-word',
              boxShadow: '0px 4px 6px rgba(0, 0, 0, 0.1)',
            }}
          >
            {msg.content}
          </Box>
        ))}
        <div ref={messagesEndRef} />
      </Stack>

      {/* Input Field */}
      <Box
        sx={{
          padding: '15px',
          display: 'flex',
          backgroundColor: '#f9f9f9',
          borderTop: '1px solid #ccc',
          borderRadius: '0 0 12px 12px',
        }}
      >
        <TextField
          fullWidth
          placeholder="Type your message..."
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          onKeyDown={handleKeyPress} // Enable Enter-to-Send
          disabled={isLoading}
          multiline
          maxRows={3}
          sx={{
            backgroundColor: '#fff',
            borderRadius: '8px',
            marginRight: '10px',
            boxShadow: '0px 2px 4px rgba(0, 0, 0, 0.1)',
          }}
        />
        <Button
          variant="contained"
          sx={{
            backgroundColor: '#e74c3c',
            color: '#fff',
            padding: '10px 20px',
            fontWeight: 'bold',
            '&:hover': { backgroundColor: '#d62f2f' },
          }}
          onClick={handleSendMessage}
          disabled={isLoading}
        >
          {isLoading ? <CircularProgress size={24} color="inherit" /> : <Send />}
        </Button>
      </Box>
    </Box>
  );
};

export default Chatbot;
