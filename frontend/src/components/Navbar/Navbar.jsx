import React, { useContext, useEffect, useState } from 'react'
import './Navbar.css'
import { assets } from '../../assets/assets'
import { Link, useNavigate } from 'react-router-dom'
import { StoreContext } from '../../Context/StoreContext'

const Navbar = ({ setShowLogin }) => {

  const [menu, setMenu] = useState("home");
  const [searchVisible, setSearchVisible] = useState(false);
  const { getTotalCartAmount, token ,setToken } = useContext(StoreContext);
  const navigate = useNavigate();

  const logout = () => {
    localStorage.removeItem("token");
    setToken("");
    navigate('/')
  }
  const toggleSearch = () => {
    setSearchVisible(prev => !prev);
  };

  return (
    <div className='navbar'>
      <Link to='/'><img className='logo' src={assets.logo2} alt="logo" /></Link> 
      <ul className="navbar-menu">
        <Link to="/" onClick={() => setMenu("home")} className={`${menu === "home" ? "active" : ""}`}>Home</Link>
        <a href='#explore-menu' onClick={() => setMenu("menu")} className={`${menu === "menu" ? "active" : ""}`}>OTC's</a>
        <a href='#Prescriptions' onClick={() => setMenu("Prescriptions")} className={`${menu === "Prescription" ? "active" : ""}`}>Prescriptions</a>
        <a href='#footer' onClick={() => setMenu("contact")} className={`${menu === "contact" ? "active" : ""}`}>Contact Us</a>
      </ul>
      <div className="navbar-right">

        
        <button onClick={toggleSearch} className="navbar-search-button">
          <img src={assets.search} alt="search icon" />
        </button>
        
        {searchVisible && (
          <input
            type="text"
            placeholder="Search..."
            className="navbar-search-input"
          />
        )}

        <Link to='/cart' className='navbar-search-icon'>
          <img src={assets.basket} alt="" />
          <div className={getTotalCartAmount() > 0 ? "dot" : ""}></div>
        </Link>

        {!token ? <button onClick={() => setShowLogin(true)}>Sign In</button>
          : <div className='navbar-profile'>
            <img src={assets.profile_icon} alt="" />
            <ul className='navbar-profile-dropdown'>
              <li onClick={()=>navigate('/myorders')}> <img src={assets.bag_icon} alt="" /> <p>Orders</p></li>
              <hr />
              <li onClick={logout}> <img src={assets.logout_icon} alt="" /> <p>Logout</p></li> 
            </ul>
          </div>
        }
      </div>
    </div>
  )
}

export default Navbar
