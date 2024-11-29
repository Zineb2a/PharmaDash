import React, { useEffect, useState } from "react";
import axios from "axios";
import { FaUserTie } from "react-icons/fa"; // Font Awesome Icon
import "./DriverOrder.css";

const DriverOrder = () => {
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);
  const [driverName, setDriverName] = useState("Driver"); // Replace with actual driver name logic

  // Fetch orders from backend
  const fetchOrders = async () => {
    setLoading(true);
    try {
      const response = await axios.get("http://localhost:3000/order/get_orders", {
        withCredentials: true, // Include cookies
      });

      console.log("Fetched orders:", response.data);

      if (response.data.success) {
        setOrders(response.data.orders);
      } else {
        console.error("Error fetching orders:", response.data.message);
      }
    } catch (error) {
      console.error("Error fetching orders:", error.response?.data || error.message);
    } finally {
      setLoading(false);
    }
  };

  // Pick up an order
  const selectOrder = async (orderId) => {
    console.log("Attempting to pick up order with ID:", orderId);
    try {
      const payload = { order_id: orderId };
      console.log("Payload being sent to /driver_picks_up:", payload);

      const response = await axios.post(
        "http://localhost:3000/user/driver_picks_up",
        payload,
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      );

      console.log("Pickup response:", response.data);

      if (response.data.status === "Order picked up successfully.") {
        console.log("Order successfully picked up:", orderId);
        setOrders((prevOrders) =>
          prevOrders.map((order) =>
            order.OrderID === orderId ? { ...order, OrderStatus: "Out for delivery" } : order
          )
        );
      } else {
        console.error("Error picking up order (API response):", response.data.status);
      }
    } catch (error) {
      console.error("Error picking up order (exception):", error.response?.data || error.message);
      if (error.response) {
        console.log("Full server response:", error.response.data);
      }
    }
  };

  // Confirm delivery
  const completeDelivery = async (orderId) => {
    try {
      console.log("Attempting to confirm delivery for order:", orderId);

      const payload = { order_id: orderId };

      const response = await axios.post(
        "http://localhost:3000/user/driver_confirm_delivery",
        payload,
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      );

      console.log("Delivery confirmation response:", response.data);

      if (response.data.status === "Order delivery confirmed successfully.") {
        setOrders((prevOrders) =>
          prevOrders.map((order) =>
            order.OrderID === orderId ? { ...order, OrderStatus: "Delivered" } : order
          )
        );
      } else {
        console.error("Error confirming delivery:", response.data.status);
      }
    } catch (error) {
      console.error("Error confirming delivery:", error.response?.data || error.message);
    }
  };

  useEffect(() => {
    fetchOrders();
  }, []);

  if (loading) {
    return <p>Loading orders...</p>;
  }

  return (
    <div className="driver-orders">
      <div className="driver-header">
        <FaUserTie className="driver-avatar" />
        <div>
          <h1>Welcome, {driverName}!</h1>
          <p>Manage your deliveries with ease.</p>
        </div>
      </div>
      <section>
        <h2>Ready Orders</h2>
        <ul>
          {orders
            .filter((order) => order.OrderStatus === "Created")
            .map((order) => (
              <li key={order.OrderID}>
                <div>
                  <p>
                    <strong>Destination:</strong> {order.Destination}
                  </p>
                  <p>
                    <strong>Special Handling:</strong> {order.SpecialHandling}
                  </p>
                  <p>
                    <strong>Total Cost:</strong> ${order.TotalCost}
                  </p>
                </div>
                <button onClick={() => selectOrder(order.OrderID)}>Pick Up Order</button>
              </li>
            ))}
        </ul>
      </section>
      <section>
        <h2>Ongoing Deliveries</h2>
        <ul>
          {orders
            .filter((order) => order.OrderStatus === "Out for delivery")
            .map((order) => (
              <li key={order.OrderID}>
                <div>
                  <p>
                    <strong>Destination:</strong> {order.Destination}
                  </p>
                  <p>
                    <strong>Special Handling:</strong> {order.SpecialHandling}
                  </p>
                  <p>
                    <strong>Total Cost:</strong> ${order.TotalCost}
                  </p>
                </div>
                <button onClick={() => completeDelivery(order.OrderID)}>Confirm Delivery</button>
              </li>
            ))}
        </ul>
      </section>
      <section>
        <h2>Completed Orders</h2>
        <ul>
          {orders
            .filter((order) => order.OrderStatus === "Delivered")
            .map((order) => (
              <li key={order.OrderID}>
                <div>
                  <p>
                    <strong>Destination:</strong> {order.Destination}
                  </p>
                  <p>
                    <strong>Special Handling:</strong> {order.SpecialHandling}
                  </p>
                  <p>
                    <strong>Total Cost:</strong> ${order.TotalCost}
                  </p>
                </div>
              </li>
            ))}
        </ul>
      </section>
    </div>
  );
};

export default DriverOrder;
