import React, { useEffect, useState } from 'react';
import './Dashboard.css';
import axios from 'axios';
import { toast } from 'react-toastify';
import { Pie } from 'react-chartjs-2';
import 'chart.js/auto'; // Import for auto-chart registration

const Dashboard = () => {
  const [orderStats, setOrderStats] = useState({
    totalOrders: 0,
    delivered: 0,
    outForDelivery: 0,
    inProcess: 0,
    canceled: 0,
  });

  const fetchOrderStats = async () => {
    try {
      // Fetch orders from backend
      const response = await axios.get('http://localhost:3000/order/get_orders');
      if (response.data.success) {
        const orders = response.data.orders; // Access 'orders' array from the backend response

        // Calculate statistics based on the order status
        const delivered = orders.filter(order => order.OrderStatus === 'Delivered').length;
        const outForDelivery = orders.filter(order => order.OrderStatus === 'Out for delivery').length;
        const inProcess = orders.filter(order => order.OrderStatus === 'Order Processing').length;
        const canceled = orders.filter(order => order.OrderStatus === 'Canceled').length;

        // Update state with calculated stats
        setOrderStats({
          totalOrders: orders.length,
          delivered,
          outForDelivery,
          inProcess,
          canceled,
        });
      } else {
        toast.error('Error fetching order data.');
      }
    } catch (error) {
      console.error('Error fetching orders:', error);
      toast.error('Something went wrong. Please try again.');
    }
  };

  useEffect(() => {
    fetchOrderStats();
  }, []);

  const chartData = {
    labels: ['Delivered', 'Out for Delivery', 'Order Processing', 'Canceled'],
    datasets: [
      {
        label: 'Order Status Distribution',
        data: [
          orderStats.delivered,
          orderStats.outForDelivery,
          orderStats.inProcess,
          orderStats.canceled,
        ],
        backgroundColor: ['#66BB6A', '#FFD54F', '#42A5F5', '#EF5350'], // Softer colors
        hoverBackgroundColor: ['#43A047', '#FFC107', '#1E88E5', '#E53935'], // Vibrant hover colors
      },
    ],
  };

  return (
    <div className="dashboard">
      <h3>Dashboard</h3>
      <div className="dashboard-stats">
        <div className="stat-card">
          <h4>Total Orders</h4>
          <p>{orderStats.totalOrders}</p>
        </div>
        <div className="stat-card">
          <h4>Delivered</h4>
          <p>{orderStats.delivered}</p>
        </div>
        <div className="stat-card">
          <h4>Out for Delivery</h4>
          <p>{orderStats.outForDelivery}</p>
        </div>
        <div className="stat-card">
          <h4>Order Processing</h4>
          <p>{orderStats.inProcess}</p>
        </div>
        <div className="stat-card">
          <h4>Canceled</h4>
          <p>{orderStats.canceled}</p>
        </div>
      </div>
      <div className="dashboard-chart">
        <Pie data={chartData} />
      </div>
    </div>
  );
};

export default Dashboard;
