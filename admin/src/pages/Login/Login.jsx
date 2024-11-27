import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import './Login.css';

const Login = () => {
  const [credentials, setCredentials] = useState({ email: "", password: "" });
  const navigate = useNavigate();

  // Hardcoded credentials for each role
  const users = {
    admin: { email: "admin@example.com", password: "admin", role: "admin", redirectTo: "/dashboard" },
    driver: { email: "driver@example.com", password: "driver", role: "driver", redirectTo: "/driver/orders" },
    user: { email: "user@example.com", password: "user", role: "user", redirectTo: "/user/home" },
  };

  const handleLogin = (event) => {
    event.preventDefault();
  
    const user = Object.values(users).find(
      (u) => u.email === credentials.email && u.password === credentials.password
    );
  
    if (user) {
      localStorage.setItem("role", user.role);
      toast.success(`Logged in as ${user.role.charAt(0).toUpperCase() + user.role.slice(1)}!`);
      navigate(user.redirectTo); // Navigate to the user's default page
      window.location.reload(); // Refresh the page after successful login
    } else {
      toast.error("Invalid credentials! Please try again.");
    }
  };
  
  

  return (
    <div className="login-container">
      <form onSubmit={handleLogin}>
        <input
          type="email"
          placeholder="Email"
          value={credentials.email}
          onChange={(e) => setCredentials({ ...credentials, email: e.target.value })}
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={credentials.password}
          onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
          required
        />
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default Login;
