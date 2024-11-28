import React, { useEffect, useState } from 'react';
import './Orders.css';
import { toast } from 'react-toastify';
import axios from 'axios';
import { assets, url, currency } from '../../assets/assets';

const Order = () => {
  const [orders, setOrders] = useState([]);

  const fetchAllOrders = async () => {
    try {
      //const response = await axios.get(`${url}/api/order/list`);
      const response = await axios.get(`http://localhost:3000/order/get_orders`);
      if (response.data.success) {
        setOrders(response.data.data.reverse());
      } else {
        toast.error('Error fetching orders.');
      }
    } catch (error) {
      toast.error('Something went wrong. Please try again.');
    }
  };

  // const statusHandler = async (event, orderId) => {
  //   try {
  //     const response = await axios.post(`${url}/api/order/status`, {
  //       orderId,
  //       status: event.target.value,
  //     });
  //     if (response.data.success) {
  //       await fetchAllOrders();
  //       toast.success('Order status updated successfully.');
  //     } else {
  //       toast.error('Error updating order status.');
  //     }
  //   } catch (error) {
  //     toast.error('Something went wrong. Please try again.');
  //   }
  // };

  useEffect(() => {
    fetchAllOrders();
  }, []);

  return (
    <div className="order-container">
      <h3>Orders and Deliveries</h3>
      <div className="order-table">
        <div className="order-table-header">
          <b>Parcel</b>
          <b>Items</b>
          <b>Customer</b>
          <b>Address</b>
          <b>Phone</b>
          <b>Total</b>
          <b>Status</b>
        </div>
        {orders.map((order, index) => (
          <div key={index} className="order-table-row">
            <img src={assets.parcel_icon} alt="Parcel" className="order-parcel-icon" />
            <p className="order-items">
              {order.items.map((item, i) => (
                <span key={i}>
                  {item.name} x {item.quantity}
                  {i !== order.items.length - 1 && ', '}
                </span>
              ))}
            </p>
            <p className="order-customer">
              {order.address.firstName} {order.address.lastName}
            </p>
            <p className="order-address">
              {order.address.street}, {order.address.city}, {order.address.state},{' '}
              {order.address.country} - {order.address.zipcode}
            </p>
            <p className="order-phone">{order.address.phone}</p>
            <p className="order-total">
              {currency}
              {order.amount}
            </p>
            <select
              onChange={(e) => statusHandler(e, order._id)}
              value={order.status}
              className="order-status-dropdown"
            >
              <option value="Order Processing">Order Processing</option>
              <option value="Out for delivery">Out for delivery</option>
              <option value="Delivered">Delivered</option>
              <option value="Canceled">Canceled</option>
            </select>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Order;
