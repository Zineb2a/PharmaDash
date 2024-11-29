import React, { useState } from "react";
import "./Payment.css";

const Payment = () => {
  const [formData, setFormData] = useState({
    cardNumber: "",
    expiryDate: "",
    cvv: "",
  });

  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    setMessage("");
  
    try {
      const response = await fetch("http://localhost:3000/order/create_order", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          CardNumber: formData.cardNumber,
          ExpiryDate: formData.expiryDate,
          CVV: formData.cvv,
          CartID: formData.cartId, // Ensure this is set
        }),
      });
  
      const result = await response.json();
  
      if (response.ok) {
        setMessage(`Order created successfully! Order ID: ${result.order_id}`);
        setFormData({ cardNumber: "", expiryDate: "", cvv: "", cartId: formData.cartId });
      } else {
        setMessage(result.status || "Failed to process payment.");
      }
    } catch (error) {
      setMessage("An error occurred. Please try again.");
    } finally {
      setLoading(false);
    }
  };
  

  return (
    <div className="payment-container">
      <h2>Payment Method</h2>
      <form onSubmit={handleSubmit} className="payment-form">
        <label>
          Card Number:
          <input
            type="text"
            name="cardNumber"
            value={formData.cardNumber}
            onChange={handleChange}
            placeholder="1234 5678 9101 1121"
            required
          />
        </label>
        <label>
          Expiry Date:
          <input
            type="text"
            name="expiryDate"
            value={formData.expiryDate}
            onChange={handleChange}
            placeholder="MM/YY"
            required
          />
        </label>
        <label>
          CVV:
          <input
            type="password"
            name="cvv"
            value={formData.cvv}
            onChange={handleChange}
            placeholder="123"
            required
          />
        </label>
        <button type="submit" className="proceed-button" disabled={loading}>
          {loading ? "Processing..." : "Pay Now"}
        </button>
      </form>
      {message && <p className="message">{message}</p>}
    </div>
  );
};

export default Payment;
