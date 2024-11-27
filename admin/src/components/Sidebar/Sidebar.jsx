import React from 'react';
import './Sidebar.css';
import { NavLink } from 'react-router-dom';
import { assets } from '../../assets/assets';

const Sidebar = () => {
  const navOptions = [
    { path: '/dashboard', icon: assets.dashboard_icon, label: 'Dashboard' },
    { path: '/add', icon: assets.add_icon, label: 'Add Items' },
    { path: '/list', icon: assets.list_icon, label: 'List Items' },
    { path: '/orders', icon: assets.order_icon, label: 'Orders' },
  ];

  return (
    <nav className="sidebar">
      <div className="sidebar-options">
        {navOptions.map((option, index) => (
          <NavLink
            key={index}
            to={option.path}
            className={({ isActive }) => `sidebar-option ${isActive ? 'active' : ''}`}
          >
            <img src={option.icon} alt={`${option.label} Icon`} />
            <p>{option.label}</p>
          </NavLink>
        ))}
      </div>
    </nav>
  );
};

export default Sidebar;
