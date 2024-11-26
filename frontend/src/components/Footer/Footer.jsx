import React, { useState } from 'react';
import './Footer.css';
import { assets } from '../../assets/assets';
import { Link } from 'react-router-dom';

const Footer = () => {
  return (
    <div className="footer" id="footer">
      <div className="footer-content">
        {/* Left Section */}
        <div className="footer-content-left">
         
          <p className="footer-description">
          We are committed to providing fast, reliable, and convenient healthcare solutions.
          </p>
         
        </div>

        {/* Center Section */}
        <div className="footer-content-center">
        
       
        </div>

        {/* Right Section */}
        <div className="footer-content-right">
        <ul>
    
    <li>
      <Link to="/about-us">About Us</Link>
    </li>
    <li>
      <Link to="/privacy-policy">Privacy Policy</Link>
    </li>
  </ul>
        </div>
      </div>

      <hr />
      <p className="footer-copyright">© 2024 Pharmadash.com - All Rights Reserved.</p>
    </div>
  );
};

export default Footer;
