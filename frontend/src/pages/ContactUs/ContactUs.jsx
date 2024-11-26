import React from 'react';
import './ContactUs.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPhone, faEnvelope, faLocationDot } from '@fortawesome/free-solid-svg-icons';
import { faHeadset } from '@fortawesome/free-solid-svg-icons';
import { Link } from 'react-router-dom';

const ContactUs = () => {
  return (
    <div className="contact-us-container">
      {/* Header Section */}
      <div className="contact-us-header">
       
        <p>We’re here to help! Feel free to reach out to us through any of the methods below.</p>
      </div>

      {/* Contact Information */}
      <div className="contact-info">
        <div className="contact-item">
          <FontAwesomeIcon icon={faPhone} className="contact-icon" />
          <h3>Phone</h3>
          <p>+1-800-212-3695</p>
        </div>
        <div className="contact-item">
          <FontAwesomeIcon icon={faEnvelope} className="contact-icon" />
          <h3>Email</h3>
          <p>
            <a href="mailto:contact@pharmadash.com">pharmadash343@gmail.com</a>
          </p>
        </div>
        <div className="contact-item">
          <FontAwesomeIcon icon={faLocationDot} className="contact-icon" />
          <h3>Address</h3>
          <p>123 Pharmadash St., Health City, HC 56789</p>
        </div>
      </div>

      {/* Chatbot Button */}
      <div className="chatbot-section">
        <h2>Need Instant Help?</h2>
        <Link to="/chatbot" className="chatbot-button">
          <FontAwesomeIcon icon={faHeadset} className="chatbot-icon" /> Chat with Us
        </Link>
      </div>
    </div>
  );
};

export default ContactUs;
