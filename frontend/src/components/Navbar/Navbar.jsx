import React, { useContext, useState, useEffect } from 'react';
import './Navbar.css';
import { assets } from '../../assets/assets';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { StoreContext } from '../../Context/StoreContext';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHeadset } from '@fortawesome/free-solid-svg-icons';

const Navbar = ({ setShowLogin }) => {
  const { getTotalCartAmount, token, setToken } = useContext(StoreContext);
  const navigate = useNavigate();
  const location = useLocation(); // Get the current route

  const [menu, setMenu] = useState("home");

  useEffect(() => {
    // Update the active menu based on the current route
    if (location.pathname === '/otc') setMenu("otc");
    else if (location.pathname === '/prescriptions') setMenu("prescriptions");
    else if (location.pathname === '/') setMenu("home");
    else setMenu(""); // Reset if the path doesn't match
  }, [location]);

  const logout = () => {
    localStorage.removeItem("token");
    setToken("");
    navigate('/');
  };

  return (
    <div className='navbar'>
      <Link to='/'><img className='logo' src={assets.logo2} alt="logo" /></Link>
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
        <a
          href='#footer'
          onClick={() => setMenu("contact")}
          className={`${menu === "contact" ? "active" : ""}`}
        >
          Contact Us
        </a>
      </ul>
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
