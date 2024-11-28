import React, { useState } from 'react';
import './LoginPopup.css';
import { assets } from '../../assets/assets';
import axios from 'axios';
import { toast } from 'react-toastify';

const LoginPopup = ({ setShowLogin }) => {
    const [currState, setCurrState] = useState("Sign Up");
    const [token, setToken] = useState(localStorage.getItem("token") || null);
    const [data, setData] = useState({
        name: "",
        lastName: "",
        email: "",
        password: "",
        phoneNumber: "",
        address: "",
    });
    const url = "http://localhost:3000"; 
    const onChangeHandler = (event) => {
        const { name, value } = event.target;
        setData((prevData) => ({ ...prevData, [name]: value }));
    };

    const onRegister = async () => {
        const new_url = `${url}/user/register`;
        const payload = {
            RegisterMode: "client", //Only clients can register
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
            if (response.data.status === "success") {
                toast.success("Registration successful!");
                setCurrState("Login"); // Redirect to login after successful registration
            } else {
                toast.error(response.data.status);
            }
        } catch (error) {
            toast.error("An error occurred during registration. Please try again.");
        }
    };

    const onLogin = async () => {
        const new_url = `${url}/user/login`;
        const payload = {
            Email: data.email,
            Password: data.password,
        };

        try {
            const response = await axios.post(new_url, payload, { withCredentials: true });
            if (response.data.status === "success") {
                toast.success("Login successful!");
                const newToken = response.data.token;
                setToken(newToken); // Save token
                localStorage.setItem("token", response.data.token);
                setShowLogin(false); // Close popup
            } else {
                toast.error(response.data.status);
            }
        } catch (error) {
            toast.error("An error occurred during login. Please try again.");
        }
    };

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
                            />
                            <input
                                name="lastName"
                                onChange={onChangeHandler}
                                value={data.lastName}
                                type="text"
                                placeholder="Last Name"
                            />
                            <input
                                name="phoneNumber"
                                onChange={onChangeHandler}
                                value={data.phoneNumber}
                                type="text"
                                placeholder="Phone Number"
                            />
                            <input
                                name="address"
                                onChange={onChangeHandler}
                                value={data.address}
                                type="text"
                                placeholder="Address"
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
                <button>{currState === "Login" ? "Login" : "Create account"}</button>
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
