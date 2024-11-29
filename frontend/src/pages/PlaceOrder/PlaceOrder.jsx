
import React, { useContext, useEffect, useState } from 'react';
import './PlaceOrder.css';
import { StoreContext } from '../../Context/StoreContext';
import { assets } from '../../assets/assets';
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import axios from 'axios';

const PlaceOrder = () => {
    const [showQuotationPopup, setShowQuotationPopup] = useState(false);
    const [data, setData] = useState({
        firstName: "",
        lastName: "",
        email: "",
        address: "",
    });
    const { getTotalCartAmount, token, food_list, cartItems, url, setCartItems, currency, deliveryCharge } = useContext(StoreContext);
    
    const [quotationData, setQuotationData] = useState({
        quotation_id: 1, // Initially null
        total_cost: 0,
        delivery_frequency: '',
        destination: '',
        special_handling: '',
        insurance: 50.0,
        include_insurance: true,
        is_refused: false,
        cart_id: 0,
    });

    const navigate = useNavigate();

    const onChangeHandler = (event) => {
        const name = event.target.name;
        const value = event.target.value;
        setData((data) => ({ ...data, [name]: value }));
    };

    const generateQuotation = () => {
        // Generate the quotation dynamically based on current cart and user input
        const totalCost = getTotalCartAmount() + deliveryCharge;

        const newQuotationData = {
            quotation_id: Math.floor(Math.random() * 10000), // Or fetch from backend
            total_cost: totalCost,
            delivery_frequency: "Once", // You can add more logic here based on your requirements
            destination: data.address,
            special_handling: quotationData.special_handling || null,
            insurance: quotationData.insurance,
            include_insurance: quotationData.include_insurance,
            is_refused: quotationData.is_refused,
            cart_id: cartItems.cart_id, 
        };

        setQuotationData(newQuotationData); // Set the quotation data
        setShowQuotationPopup(true); // Show the quotation popup
    };

    const placeOrder = async (e) => {
        e.preventDefault();

        let orderItems = [];
        food_list.forEach((item) => {
            if (cartItems[item._id] > 0) {
                let itemInfo = { ...item };
                itemInfo['quantity'] = cartItems[item._id];
                orderItems.push(itemInfo);
            }
        });

        const orderPayload = {
            address: data,
            items: orderItems,
            amount: getTotalCartAmount() + deliveryCharge,
            quotation: quotationData, // Include the quotation data here
        };

        try {
            const response = await axios.post(`${url}/cart/delivery_quotation`, {
                quotation_id: quotationData.quotation_id,
                total_cost: getTotalCartAmount() + deliveryCharge,
                delivery_frequency: quotationData.delivery_frequency,
                destination: data.address,
                special_handling: quotationData.special_handling || null,
                insurance: quotationData.insurance,
                include_insurance: quotationData.include_insurance,
                is_refused: quotationData.is_refused,
                cart_id: quotationData.cart_id,
            }, {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            toast.success("Order placed successfully!");
            setCartItems({}); 
            navigate('/order-success'); // Redirect to success page
        } catch (error) {
            console.error("Error placing order:", error.response || error);
            toast.error("Failed to place the order. Please try again.");
        }
    };

    const handleAccept = () => {
        setQuotationData((prev) => ({
            ...prev,
            is_refused: false,
        }));
        setShowQuotationPopup(false);
        navigate('/payment');
    };

    const handleReject = () => {
        setQuotationData((prev) => ({
            ...prev,
            is_refused: true,
        }));
        setShowQuotationPopup(false);
        navigate('/place-order');
    };

    useEffect(() => {
        if (!token) {
            toast.error('To place an order, sign in first');
            navigate('/cart');
        } else if (getTotalCartAmount() === 0) {
            navigate('/cart');
        }
    }, [token]);

    return (
        <form onSubmit={placeOrder} className="place-order">
            <div className="place-order-left">
                <p className="title">Delivery Information</p>
                <div className="multi-field">
                    <input type="text" name="firstName" onChange={onChangeHandler} value={data.firstName} placeholder="First name" required />
                    <input type="text" name="lastName" onChange={onChangeHandler} value={data.lastName} placeholder="Last name" required />
                </div>
                <input type="email" name="email" onChange={onChangeHandler} value={data.email} placeholder="Email address" required />
                <input type="text" name="address" onChange={onChangeHandler} value={data.address} placeholder="Address" required />
            </div>
            <div className="place-order-right">
                <div className="cart-total">
                    <h2>Cart Totals</h2>
                    <div>
                        <div className="cart-total-details">
                            <p>Subtotal</p>
                            <p>
                                {currency}
                                {getTotalCartAmount()}
                            </p>
                        </div>
                        <hr />
                        <div className="cart-total-details">
                            <p>Delivery Fee</p>
                            <p>
                                {currency}
                                {getTotalCartAmount() === 0 ? 0 : deliveryCharge}
                            </p>
                        </div>
                        <hr />
                        <div className="cart-total-details">
                            <b>Total</b>
                            <b>
                                {currency}
                                {getTotalCartAmount() === 0 ? 0 : getTotalCartAmount() + deliveryCharge}
                            </b>
                        </div>
                    </div>
                </div>
                <div className="quotation">
                    <button className="handle-quotation" type="button" onClick={generateQuotation}>
                        Generate Quotation
                    </button>
                </div>
            </div>
            {showQuotationPopup && (
                <div className="popup-overlay">
                    <div className="popup">
                        <h3>Quotation Details</h3>
                        <p><strong>Total Cost:</strong> {getTotalCartAmount() === 0 ? 0 : getTotalCartAmount() + deliveryCharge}</p>
                        <p><strong>Delivery Frequency:</strong> {quotationData.delivery_frequency}</p>
                        <p><strong>Destination:</strong> {data.address}</p>
                        <div className='special-handling'>
                            <label htmlFor="special-handling">
                                <strong>Special Handling:</strong>
                            </label>
                            <textarea
                                id="special-handling"
                                value={quotationData.special_handling || ""}
                                onChange={(e) =>
                                    setQuotationData((prev) => ({
                                        ...prev,
                                        special_handling: e.target.value,
                                    }))
                                }
                                rows={3}
                                placeholder="Enter special handling instructions here"
                            />
                        </div>
                        <p><strong>Insurance:</strong> 10% </p>
                        <p><strong>Include Insurance:</strong> {quotationData.include_insurance ? 'Yes' : 'No'}</p>
                        <p><strong>Cart ID:</strong> {quotationData.cart_id}</p>
                        <div className="popup-buttons">
                            <button className="accept-button" onClick={handleAccept}>
                                Accept
                            </button>
                            <button className="reject-button" onClick={handleReject}>
                                Reject
                            </button>
                        </div>
                    </div>
                </div>
            )}
        </form>
    );
};

export default PlaceOrder;
