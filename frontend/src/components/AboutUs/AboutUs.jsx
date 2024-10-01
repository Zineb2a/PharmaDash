import React, { useEffect } from 'react';
import './AboutUs.css';


const AboutUs = () => {
  // Scroll to the top when the component mounts
  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  return (
    <div className="about-us">
      <h1>About Us</h1>
      <section className="about-us-content">
        <div className="about-us-text">
          <h2>Our Mission</h2>
          <p>
            At PharmaDash, our mission is to provide a fast, reliable, and convenient way to deliver both over-the-counter and prescription medications to your doorstep. We believe in simplifying the healthcare process, making it accessible for everyone.
          </p>
        </div>

        <div className="about-us-text">
          <h2>Who We Are</h2>
          <p>
            We are a dedicated team of healthcare enthusiasts, tech professionals, and logistics experts. Our goal is to bridge the gap between patients and pharmacies, ensuring that everyone has access to the medications they need, wherever they are.
          </p>
        </div>

      

        <div className="about-us-text">
          <h2>Our Values</h2>
          <ul>
            <li>Accessibility: Making healthcare accessible to everyone.</li>
            <li>Reliability: Ensuring prompt and secure deliveries.</li>
            <li>Innovation: Leveraging technology to enhance healthcare delivery.</li>
            <li>Care: Prioritizing the health and well-being of our community.</li>
          </ul>
        </div>
      </section>
    </div>
  );
};

export default AboutUs;
