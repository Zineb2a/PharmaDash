import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "./Navbar.css";
import { assets } from "../../assets/assets";

const Navbar = () => {
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);
  const navigate = useNavigate();

  // Get the user's role from localStorage
  const role = localStorage.getItem("role");

  const toggleDropdown = () => {
    setIsDropdownOpen((prev) => !prev);
  };

  const handleLogout = () => {
    localStorage.removeItem("role"); // Clear the user's role
    setIsDropdownOpen(false); // Close dropdown after logout
    navigate("/login"); // Redirect to login page
  };

  return (
    <header className="navbar">
      <img className="logo" src={assets.logo} alt="Pharma Dash Logo" />
      <div className="profile-container">
        <img
          className="profile"
          src={assets.profile_image}
          alt="User Profile"
          onClick={toggleDropdown}
        />
        {isDropdownOpen && (
          <div className="dropdown">
            <ul>
              {role === "admin" && (
                <>
                  <li>
                    <Link to="/dashboard">Dashboard</Link>
                  </li>
                  <li>
                    <Link to="/add">Add Items</Link>
                  </li>
                  <li>
                    <Link to="/list">List Items</Link>
                  </li>
                  <li>
                    <Link to="/orders">Orders</Link>
                  </li>
                </>
              )}
              {role === "driver" && (
                <li>
                  <Link to="/driver/orders">Driver Orders</Link>
                </li>
              )}
              <li onClick={handleLogout}>Logout</li>
            </ul>
          </div>
        )}
      </div>
    </header>
  );
};

export default Navbar;
