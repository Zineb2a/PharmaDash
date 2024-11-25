import React, { useState, useEffect } from 'react';
import Chatbot from '../../components/Chatbot/Chatbot'; // Reuse the existing Chatbot component
import './ChatbotPage.css';

const ChatbotPage = () => {
  useEffect(() => {
    window.scrollTo(0, 0); // Scroll to the top when the page loads
  }, []);

  return (
    <div className="chatbot-page-container">
      <header className="chatbot-page-header">
        <h1>Pharmadash Customer Support</h1>
        <p>Our chatbot is here to help you with any questions or concerns you may have.</p>
      </header>
      <main>
        <Chatbot />
      </main>
    </div>
  );
};

export default ChatbotPage;
