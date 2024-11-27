import React, { useEffect, useState } from "react";
import "./DriverOrder.css";

const DriverOrder = () => {
  const [orders, setOrders] = useState([]);
  const [loading, setLoading] = useState(true);

  // Mock data for testing
  useEffect(() => {
    const fetchMockOrders = () => {
      setLoading(true);
      setTimeout(() => {
        const mockOrders = [
          {
            id: "1",
            customerName: "John Doe",
            address: "123 Maple Street, Springfield",
            phone: "+1 555-1234",
            status: "ready",
          },
          {
            id: "2",
            customerName: "Jane Smith",
            address: "456 Oak Avenue, Metropolis",
            phone: "+1 555-5678",
            status: "ready",
          },
          {
            id: "3",
            customerName: "Alice Johnson",
            address: "789 Pine Lane, Gotham",
            phone: "+1 555-8765",
            status: "ongoing",
          },
          {
            id: "4",
            customerName: "Mark Brown",
            address: "321 Cedar Road, Star City",
            phone: "+1 555-4321",
            status: "completed",
          },
          {
            id: "5",
            customerName: "Emma Wilson",
            address: "654 Birch Boulevard, Central City",
            phone: "+1 555-3456",
            status: "ongoing",
          },
        ];

        setOrders(mockOrders);
        setLoading(false);
      }, 1000); // Simulate network delay
    };

    fetchMockOrders();
  }, []);

  // Assign a ready order to the driver
  const selectOrder = (orderId) => {
    setOrders((prevOrders) =>
      prevOrders.map((order) =>
        order.id === orderId ? { ...order, status: "ongoing" } : order
      )
    );
  };

  // Confirm delivery for an ongoing order
  const completeDelivery = (orderId) => {
    setOrders((prevOrders) =>
      prevOrders.map((order) =>
        order.id === orderId ? { ...order, status: "completed" } : order
      )
    );
  };

  if (loading) {
    return <p>Loading orders...</p>;
  }

  return (
    <div className="driver-orders">
      <h1>Driver Dashboard</h1>

      {/* Ready Orders */}
      <section>
        <h2>Ready Orders</h2>
        <ul>
          {orders
            .filter((order) => order.status === "ready")
            .map((order) => (
              <li key={order.id}>
                <div>
                  <p>
                    <strong>Customer:</strong> {order.customerName}
                  </p>
                  <p>
                    <strong>Address:</strong> {order.address}
                  </p>
                  <p>
                    <strong>Phone:</strong> {order.phone}
                  </p>
                </div>
                <button onClick={() => selectOrder(order.id)}>
                  Pick Up Order
                </button>
              </li>
            ))}
        </ul>
      </section>

      {/* Ongoing Deliveries */}
      <section>
        <h2>Ongoing Deliveries</h2>
        <ul>
          {orders
            .filter((order) => order.status === "ongoing")
            .map((order) => (
              <li key={order.id}>
                <div>
                  <p>
                    <strong>Customer:</strong> {order.customerName}
                  </p>
                  <p>
                    <strong>Address:</strong> {order.address}
                  </p>
                  <p>
                    <strong>Phone:</strong> {order.phone}
                  </p>
                </div>
                <button onClick={() => completeDelivery(order.id)}>
                  Confirm Delivery
                </button>
              </li>
            ))}
        </ul>
      </section>

      {/* Completed Orders */}
      <section>
        <h2>Completed Orders</h2>
        <ul>
          {orders
            .filter((order) => order.status === "completed")
            .map((order) => (
              <li key={order.id}>
                <div>
                  <p>
                    <strong>Customer:</strong> {order.customerName}
                  </p>
                  <p>
                    <strong>Address:</strong> {order.address}
                  </p>
                  <p>
                    <strong>Phone:</strong> {order.phone}
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
