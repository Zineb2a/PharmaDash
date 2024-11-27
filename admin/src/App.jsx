import React from "react";
import Navbar from "./components/Navbar/Navbar";
import Sidebar from "./components/Sidebar/Sidebar";
import { Route, Routes, Navigate } from "react-router-dom";
import Add from "./pages/Add/Add";
import List from "./pages/List/List";
import Orders from "./pages/Orders/Orders";
import Dashboard from "./pages/Dashboard/Dashboard";
import DriverOrders from "./pages/Driver/DriverOrders";
import Login from "./pages/Login/Login";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const App = () => {
  const role = localStorage.getItem("role");

  return (
    <div className="app">
      <ToastContainer />
      {role && <Navbar />}
      <div className="app-content">
        {role === "admin" && <Sidebar />}
        <Routes>
          {/* Public Route */}
          <Route path="/login" element={<Login />} />

          {/* Admin Routes */}
          {role === "admin" && (
            <>
              <Route path="/dashboard" element={<Dashboard />} />
              <Route path="/add" element={<Add />} />
              <Route path="/list" element={<List />} />
              <Route path="/orders" element={<Orders />} />
              <Route path="*" element={<Navigate to="/dashboard" />} />
            </>
          )}

          {/* Driver Routes */}
          {role === "driver" && (
            <>
              <Route path="/driver/orders" element={<DriverOrders />} />
              <Route path="*" element={<Navigate to="/driver/orders" />} />
       
            </>
          )}

          {/* Default Redirect for Unauthenticated Users */}
          {!role && <Route path="*" element={<Navigate to="/login" />} />}
        </Routes>
      </div>
    </div>
  );
};

export default App;


