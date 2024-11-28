import React, { useState } from 'react';
import './LoginPopup.css';
import { assets } from '../../assets/assets';
import axios from 'axios';
import { toast } from 'react-toastify';

const LoginPopup = ({ setShowLogin }) => {
    const [currState, setCurrState] = useState("Sign Up");
    const [data, setData] = useState({
        name: "",
        lastName: "",
        email: "",
        password: "",
        phoneNumber: "",
        address: "",
    });

    // Backend URL
    const url = "http://localhost:3000"; 

    const onChangeHandler = (event) => {
        const { name, value } = event.target;
        setData((prevData) => ({ ...prevData, [name]: value }));
    };

    // Registration function
    const onRegister = async () => {
        const new_url = `${url}/user/register`; // Corrected template literal
        const payload = {
            RegisterMode: "client", // Only clients can register
            UserData: {
                Name: data.name || null,
                LastName: data.lastName || null,
                Password: data.password || null,
                PhoneNumber: data.phoneNumber || null,
                Email: data.email || null,
                Address: data.address || null,
            },
        };

        try {
            const response = await axios.post(new_url, payload);
            if (response.data.status === "account created") {
                toast.success("Registration successful!");
                setCurrState("Login"); // Redirect to login after successful registration
            } else {
                toast.error(response.data.status);
            }
        } catch (error) {
            console.error("Registration error:", error);
            toast.error("An error occurred during registration. Please try again.");
        }
    };

    // Login function
    const onLogin = async () => {
        const new_url = `${url}/user/login`; // Corrected template literal
        const payload = {
            Email: data.email,
            Password: data.password,
        };

        try {
            const response = await axios.post(new_url, payload, { withCredentials: true });
            if (response.data.status === "Authentication successful") {
                toast.success("Login successful!");
                const newToken = response.data.token;
                localStorage.setItem("token", newToken); // Save token to localStorage

                setShowLogin(false); // Close popup
                window.location.reload(); // Reload the page
            } else {
                toast.error(response.data.status);
            }
        } catch (error) {
            console.error("Login error:", error);
            toast.error("An error occurred during login. Please try again.");
        }
    };

    // Form submission handler
    const onSubmit = async (e) => {
        e.preventDefault();
        if (currState === "Login") {
            await onLogin();
        } else {
            await onRegister();
        }
    };

    return (
        <div className="login-popup">
            <form onSubmit={onSubmit} className="login-popup-container">
                <div className="login-popup-title">
                    <h2>{currState}</h2>
                    <img onClick={() => setShowLogin(false)} src={assets.cross_icon} alt="Close" />
                </div>
                <div className="login-popup-inputs">
                    {currState === "Sign Up" && (
                        <>
                            <input
                                name="name"
                                onChange={onChangeHandler}
                                value={data.name}
                                type="text"
                                placeholder="First Name"
                                required
                            />
                            <input
                                name="lastName"
                                onChange={onChangeHandler}
                                value={data.lastName}
                                type="text"
                                placeholder="Last Name"
                                required
                            />
                            <input
                                name="phoneNumber"
                                onChange={onChangeHandler}
                                value={data.phoneNumber}
                                type="text"
                                placeholder="Phone Number"
                                required
                            />
                            <input
                                name="address"
                                onChange={onChangeHandler}
                                value={data.address}
                                type="text"
                                placeholder="Address"
                                required
                            />
                        </>
                    )}
                    <input
                        name="email"
                        onChange={onChangeHandler}
                        value={data.email}
                        type="email"
                        placeholder="Email"
                        required
                    />
                    <input
                        name="password"
                        onChange={onChangeHandler}
                        value={data.password}
                        type="password"
                        placeholder="Password"
                        required
                    />
                </div>
                <button>{currState === "Login" ? "Login" : "Create Account"}</button>
                <div className="login-popup-condition">
                    <input type="checkbox" required />
                    <p>By continuing, I agree to the terms of use & privacy policy.</p>
                </div>
                {currState === "Login" ? (
                    <p>
                        Create a new account?{" "}
                        <span onClick={() => setCurrState("Sign Up")}>Click here</span>
                    </p>
                ) : (
                    <p>
                        Already have an account?{" "}
                        <span onClick={() => setCurrState("Login")}>Login here</span>
                    </p>
                )}
            </form>
        </div>
    );
};

export default LoginPopup;
