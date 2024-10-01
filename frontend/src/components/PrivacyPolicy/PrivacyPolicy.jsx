import React, { useEffect } from 'react';
import './PrivacyPolicy.css';

const PrivacyPolicy = () => {
    // Scroll to the top of the page on component mount
    useEffect(() => {
        window.scrollTo(0, 0);
    }, []);

    return (
        <div className="privacy-policy">
            <h1>Privacy Policy</h1>
            <p>Welcome to Pharmadash! Your privacy is of utmost importance to us. This Privacy Policy outlines how we collect, use, and protect your information.</p>

            <section>
                <h2>Information We Collect</h2>
                <p>We collect various types of information in connection with the services we provide, including:</p>
                <ul>
                    <li>Personal information you provide when creating an account, such as your name, email, and contact details.</li>
                    <li>Usage data, including how you interact with our app and website.</li>
                    <li>Payment information for processing transactions securely.</li>
                </ul>
            </section>

            <section>
                <h2>How We Use Your Information</h2>
                <p>We use the information we collect for the following purposes:</p>
                <ul>
                    <li>To provide and maintain our service.</li>
                    <li>To process your orders and manage your account.</li>
                    <li>To communicate with you, including updates and promotional materials.</li>
                    <li>To improve our services and develop new features.</li>
                </ul>
            </section>

            <section>
                <h2>Security of Your Information</h2>
                <p>We implement industry-standard security measures to protect your data. However, please be aware that no electronic transmission is entirely secure.</p>
            </section>

            <section>
                <h2>Changes to This Policy</h2>
                <p>We may update our Privacy Policy from time to time. We will notify you of any significant changes by posting the new policy on this page.</p>
            </section>

            <section>
                <h2>Contact Us</h2>
                <p>If you have any questions or concerns about this Privacy Policy, please contact us at contact@pharmadash.com.</p>
            </section>
        </div>
    );
};

export default PrivacyPolicy;
