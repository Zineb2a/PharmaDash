import React, { useContext, useState } from 'react';
import './LoginPopup.css';
import { assets } from '../../assets/assets';
import { StoreContext } from '../../Context/StoreContext';
import axios from 'axios';
import { toast } from 'react-toastify';

const LoginPopup = ({ setShowLogin }) => {
    const { setToken, url, loadCartData } = useContext(StoreContext);
    const [currState, setCurrState] = useState("Sign Up");
    const [registerMode, setRegisterMode] = useState("Client");

    const [data, setData] = useState({
        name: "",
        lastName: "",
        email: "",
        password: "",
        phoneNumber: "",
        address: ""
    });

    const onChangeHandler = (event) => {
        const name = event.target.name;
        const value = event.target.value;
        setData((data) => ({ ...data, [name]: value }));
    };

    const onLogin = async (e) => {
        e.preventDefault();

        let new_url = url;
        let payload;

        if (currState === "Login") {
            new_url += "/api/user/login";
            payload = {
                email: data.email,
                password: data.password
            };
        } else {
            new_url += "/user/register";
            payload = {
                RegisterMode: registerMode.toLowerCase(), // Convert to lowercase for backend consistency if needed
                UserData: {
                    Name: data.name || null,
                    LastName: data.lastName || null,
                    Password: data.password || null,
                    PhoneNumber: data.phoneNumber || null,
                    Email: data.email || null,
                    Address: data.address || null
                }
            };
        }

        try {
            const response = await axios.post(new_url, payload);
            if (response.data.success) {
                setToken(response.data.token);
                localStorage.setItem("token", response.data.token);
                loadCartData({ token: response.data.token });
                setShowLogin(false);
            } else {
                toast.error(response.data.message);
            }
        } catch (error) {
            toast.error("An error occurred. Please try again.");
        }
    };

    return (
        <div className='login-popup'>
            <form onSubmit={onLogin} className="login-popup-container">
                <div className="login-popup-title">
                    <h2>{currState}</h2>
                    <img onClick={() => setShowLogin(false)} src={assets.cross_icon} alt="" />
                </div>
                <div className="login-popup-inputs">
                    {currState === "Sign Up" && (
                        <>
                            <input name="name" onChange={onChangeHandler} value={data.name} type="text" placeholder="First Name" />
                            <input name="lastName" onChange={onChangeHandler} value={data.lastName} type="text" placeholder="Last Name" />
                            <input name="phoneNumber" onChange={onChangeHandler} value={data.phoneNumber} type="text" placeholder="Phone Number" />
                            <input name="address" onChange={onChangeHandler} value={data.address} type="text" placeholder="Address" />
                        </>
                    )}
                    <input name="email" onChange={onChangeHandler} value={data.email} type="email" placeholder="Email" required />
                    <input name="password" onChange={onChangeHandler} value={data.password} type="password" placeholder="Password" required />
                </div>
                {currState === "Sign Up" && (
                    <div className="register-mode-selector">
                        <label>
                            Register Mode:
                            <select
                                value={registerMode}
                                onChange={(e) => setRegisterMode(e.target.value)}
                                required
                            >
                                <option value="Client">Client</option>
                                <option value="Driver">Driver</option>
                                <option value="Admin">Admin</option>
                            </select>
                        </label>
                    </div>
                )}
                <button>{currState === "Login" ? "Login" : "Create account"}</button>
                <div className="login-popup-condition">
                    <input type="checkbox" required />
                    <p>By continuing, I agree to the terms of use & privacy policy.</p>
                </div>
                {currState === "Login" ? (
                    <p>
                        Create a new account? <span onClick={() => setCurrState('Sign Up')}>Click here</span>
                    </p>
                ) : (
                    <p>
                        Already have an account? <span onClick={() => setCurrState('Login')}>Login here</span>
                    </p>
                )}
            </form>
        </div>
    );
};

export default LoginPopup;
