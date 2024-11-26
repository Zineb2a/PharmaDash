import React, { useContext, useState } from 'react';
import './Navbar.css';
import { assets } from '../../assets/assets';
import { Link, useNavigate } from 'react-router-dom';
import { StoreContext } from '../../Context/StoreContext';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHeadset } from '@fortawesome/free-solid-svg-icons';

const Navbar = ({ setShowLogin }) => {
  const { getTotalCartAmount, token, setToken } = useContext(StoreContext);
  const navigate = useNavigate();
  const [menu, setMenu] = useState("home");

  const logout = () => {
    localStorage.removeItem("token");
    setToken("");
    navigate('/');
  };

  return (
    <div className="navbar">
      {/* Logo */}
      <Link to='/'>
        <img className='logo' src={assets.logo2} alt="Pharmadash Logo" />
      </Link>

      {/* Menu Links */}
      <ul className="navbar-menu">
        <Link
          to="/"
          onClick={() => setMenu("home")}
          className={`${menu === "home" ? "active" : ""}`}
        >
          Home
        </Link>
        <Link
          to="/otc"
          onClick={() => setMenu("otc")}
          className={`${menu === "otc" ? "active" : ""}`}
        >
          OTC's
        </Link>
        <Link
          to="/prescriptions"
          onClick={() => setMenu("prescriptions")}
          className={`${menu === "prescriptions" ? "active" : ""}`}
        >
          Prescriptions
        </Link>
        <Link
          to="/contact-us"
          onClick={() => setMenu("contact-us")}
          className={`${menu === "contact-us" ? "active" : ""}`}
        >
          Contact Us
        </Link>
      </ul>

      {/* Right Section */}
      <div className="navbar-right">
        {/* Customer Service Icon */}
        <Link to='/chatbot' className="navbar-icon customer-service-icon" title="Customer Service">
          <FontAwesomeIcon icon={faHeadset} style={{ fontSize: '24px', color: '#fafafa' }} />
        </Link>

        {/* Cart Icon */}
        <Link to='/cart' className='navbar-search-icon'>
          <img src={assets.basket} alt="Cart Icon" />
          <div className={getTotalCartAmount() > 0 ? "dot" : ""}></div>
        </Link>

        {/* Login/Logout */}
        {!token ? (
          <button onClick={() => setShowLogin(true)}>Sign In</button>
        ) : (
          <div className='navbar-profile'>
            <img src={assets.profile_icon} alt="Profile Icon" />
            <ul className='navbar-profile-dropdown'>
              <li onClick={() => navigate('/myorders')}>
                <img src={assets.bag_icon} alt="Orders" /> <p>Orders</p>
              </li>
              <hr />
              <li onClick={logout}>
                <img src={assets.logout_icon} alt="Logout" /> <p>Logout</p>
              </li>
            </ul>
          </div>
        )}
      </div>
    </div>
  );
};

export default Navbar;
