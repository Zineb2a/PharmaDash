import React, { useEffect, useState } from 'react';
import './Orders.css';
import { toast } from 'react-toastify';
import axios from 'axios';

const Order = () => {
  const [orders, setOrders] = useState([]);

  // Fetch all orders from the backend
  const fetchAllOrders = async () => {
    try {
      const response = await axios.get(`http://localhost:3000/order/get_orders`);
      if (response.data.success) {
        console.log('Fetched orders:', response.data.orders); // Debugging log
        setOrders(response.data.orders.reverse());
      } else {
        toast.error('Error fetching orders.');
      }
    } catch (error) {
      console.error('Error fetching orders:', error);
      toast.error('Something went wrong. Please try again.');
    }
  };

  useEffect(() => {
    fetchAllOrders();
  }, []);

  return (
    <div className="order-container">
      <h3>Orders and Deliveries</h3>
      <div className="order-table">
        <div className="order-table-header">
          <b>Order ID</b>
          <b>Account ID</b>
          <b>Status</b>
          <b>Created At</b>
          <b>Destination</b>
          <b>Special Handling</b>
          <b>Frequency</b>
          <b>Insurance</b>
          <b>Total Cost</b>
          <b>Driver ID</b>
        </div>
        {orders.map((order, index) => (
          <div key={index} className="order-table-row">
            <p>{order.OrderID || 'No Order ID'}</p>
            <p>{order.AccountID || 'No Account ID'}</p>
            <p>{order.OrderStatus || 'No Status'}</p>
            <p>
              {order.CreatedAt
                ? new Date(order.CreatedAt).toLocaleString() // Format date
                : 'No Date Available'}
            </p>
            <p>{order.Destination || 'No Destination'}</p>
            <p>{order.SpecialHandling || 'No Handling'}</p>
            <p>{order.DeliveryFrequency || 'No Frequency'}</p>
            <p>{order.IncludeInsurance ? 'Yes' : 'No Insurance'}</p>
            <p>{order.TotalCost ? `$${order.TotalCost}` : 'No Cost Available'}</p>
            <p>{order.DriverID || 'No Driver ID'}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Order;
