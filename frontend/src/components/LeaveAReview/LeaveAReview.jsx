import React, { useState } from "react";
import "./LeaveAReview.css"; // For styling (you can create your own styles or use an existing one)

const LeaveAReview = ({ orderId }) => {
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const [rating, setRating] = useState(0);
  const [comment, setComment] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const togglePopup = () => {
    setIsPopupOpen(!isPopupOpen);
    setErrorMessage("");
    setSuccessMessage("");
  };

  const handleStarClick = (star) => {
    setRating(star);
    setErrorMessage(""); // Clear error if rating is selected
  };

  const handleCommentChange = (e) => {
    setComment(e.target.value);
  };

  const handleSubmit = async () => {
    if (rating < 1) {
      setErrorMessage("Please select at least one star.");
      return;
    }

    try {
      const response = await fetch('http://localhost:3000/feedback/feedback', {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          order_id: orderId, // Pass the order ID dynamically
          rating: rating,
          comment: comment || "", // If comment is empty, pass an empty string
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to submit feedback.");
      }

      const data = await response.json();
      console.log("Feedback submitted successfully:", data);

      setSuccessMessage("Thank you for your feedback!");
      setTimeout(() => {
        togglePopup();
        setRating(0);
        setComment("");
      }, 2000); // Close popup after 2 seconds
    } catch (error) {
      console.error("Error submitting feedback:", error);
      setErrorMessage("Failed to submit feedback. Please try again.");
    }
  };

  return (
    <div className="leave-a-review">
      <button onClick={togglePopup} className="review-button">
        Leave A Review
      </button>

      {isPopupOpen && (
        <div className="popup">
          <div className="popup-content">
            <h2>Leave A Review</h2>
            <div className="stars">
              {[1, 2, 3, 4, 5].map((star) => (
                <span
                  key={star}
                  className={`star ${star <= rating ? "filled" : ""}`}
                  onClick={() => handleStarClick(star)}
                >
                  ★
                </span>
              ))}
            </div>
            <textarea
              placeholder="Leave a comment (optional)"
              value={comment}
              onChange={handleCommentChange}
            />
            {errorMessage && <p className="error-message">{errorMessage}</p>}
            {successMessage && (
              <p className="success-message">{successMessage}</p>
            )}
            <div className="popup-actions">
              <button onClick={handleSubmit} className="submit-button">
                Submit Review
              </button>
              <button onClick={togglePopup} className="cancel-button">
                Cancel
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default LeaveAReview;
