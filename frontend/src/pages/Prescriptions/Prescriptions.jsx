import React, { useState } from "react";
import axios from "axios";
import LeaveAReview from "../../components/LeaveAReview/LeaveAReview";
import "./Prescriptions.css";

const Prescription = () => {
  const [file, setFile] = useState(null);
  const [uploadStatus, setUploadStatus] = useState("");

  // Handle file selection
  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
    setUploadStatus(""); 
  };

  // Handle file upload
  const handleUpload = async () => {
    if (!file) {
      setUploadStatus("Please select a file first.");
      return;
    }

    const formData = new FormData();
    formData.append("prescription", file);

    try {
      const response = await axios.post("/api/upload-prescription", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      if (response.status === 200) {
        setUploadStatus("Prescription uploaded successfully!");
      } else {
        setUploadStatus("Failed to upload prescription. Please try again.");
      }
    } catch (error) {
      console.error("Upload error:", error);
      setUploadStatus("An error occurred during the upload. Please try again.");
    }
  };

  return (
    <div className="prescription-container">
      <h2>Upload Your Prescription</h2>
      <div className="upload-box">
        <input type="file" onChange={handleFileChange} accept=".jpg,.jpeg,.png,.pdf" />
        <button onClick={handleUpload}>Upload</button>
      </div>
      {uploadStatus && <p className="upload-status">{uploadStatus}</p>}
    </div>
  );
};

export default Prescription;
