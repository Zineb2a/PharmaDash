import React, { useState } from 'react';
import './Navbar.css';
import { assets } from '../../assets/assets';

const Navbar = () => {
  const [isDropdownOpen, setIsDropdownOpen] = useState(false);

  const toggleDropdown = () => {
    setIsDropdownOpen((prev) => !prev);
  };

  return (
    <header className='navbar'>
      <img className='logo' src={assets.logo} alt="Pharma Dash Logo" />
      <div className='profile-container'>
        <img
          className='profile'
          src={assets.profile_image}
          alt="User Profile"
          onClick={toggleDropdown}
        />
        {isDropdownOpen && (
          <div className='dropdown'>
            <ul>
              <li>Profile</li>
              <li>Settings</li>
              <li>Logout</li>
            </ul>
          </div>
        )}
      </div>
    </header>
  );
};

export default Navbar;
