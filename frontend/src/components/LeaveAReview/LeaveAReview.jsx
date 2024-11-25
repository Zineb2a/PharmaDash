import React, { useState } from 'react';
import './LeaveAReview.css'; // For styling (you can create your own styles or use an existing one)

const LeaveAReview = () => {
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const [rating, setRating] = useState(0);
  const [comment, setComment] = useState('');

  const togglePopup = () => {
    setIsPopupOpen(!isPopupOpen);
  };

  const handleStarClick = (star) => {
    setRating(star);
  };

  const handleCommentChange = (e) => {
    setComment(e.target.value);
  };

  const handleSubmit = () => {
    console.log("Rating:", rating);
    console.log("Comment:", comment);
    setIsPopupOpen(false);
    setRating(0); 
    setComment(''); 
  };

  return (
    <div className="leave-a-review">
     <button onClick={togglePopup} className="review-button">Leave A Review</button>

      {isPopupOpen && (
        <div className="popup">
          <div className="popup-content">
            <h2>Leave A Review</h2>
            <div className="stars">
              {[1, 2, 3, 4, 5].map((star) => (
                <span
                  key={star}
                  className={`star ${star <= rating ? 'filled' : ''}`}
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
            <div className="popup-actions">
              <button onClick={handleSubmit}>Submit Review</button>
              <button onClick={togglePopup}>Cancel</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default LeaveAReview;
