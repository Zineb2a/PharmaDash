import React, { useState } from 'react';
import './Footer.css';
import { assets } from '../../assets/assets';
import { Link } from 'react-router-dom';
import Chatbot from '../Chatbot/Chatbot';

const Footer = () => {
  const [showChatbot, setShowChatbot] = useState(false);

  return (
    <div className="footer" id="footer">
      <div className="footer-content">
        <div className="footer-content-left">
          <img className="logo" src={assets.logo2} alt="Pharmadash Logo" />
          <p className="footer-description">
            Pharmadash is a medication delivery service that brings over-the-counter medications and prescriptions
            right to your doorstep. We are committed to providing fast, reliable, and convenient healthcare solutions.
          </p>
          <div className="footer-social-icons">
            <a href="https://facebook.com" target="_blank" rel="noopener noreferrer">
              <img src={assets.facebook_icon} alt="Facebook" />
            </a>
            <a href="https://twitter.com" target="_blank" rel="noopener noreferrer">
              <img src={assets.twitter_icon} alt="Twitter" />
            </a>
          </div>
        </div>

        <div className="footer-content-center">
          <h2>Company</h2>
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about-us">About Us</Link>
            </li>
            <li>
              <Link to="/privacy-policy">Privacy Policy</Link>
            </li>
          </ul>
        </div>

        <div className="footer-content-right">
          <h2>Customer Service</h2>
          <ul>
            <li>+1-800-212-3695</li>
            <li>
              <a href="mailto:contact@pharmadash.com">contact@pharmadash.com</a>
            </li>
            
             <li><Link to="/chatbot">Chat with us</Link></li> {/* Link to Chatbot */}
            
          </ul>
        </div>
      </div>

      <hr />
      <p className="footer-copyright">© 2024 Pharmadash.com - All Rights Reserved.</p>

      {showChatbot && <Chatbot onClose={() => setShowChatbot(false)} />}
    </div>
  );
};

export default Footer;
