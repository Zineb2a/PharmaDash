import React, { useEffect, useState } from 'react';
import './List.css';
import { url, currency } from '../../assets/assets';
import axios from 'axios';
import { toast } from 'react-toastify';

const List = () => {
  const [list, setList] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  const fetchList = async () => {
    setIsLoading(true);
    try {
      const response = await axios.get(`http://localhost:3000/inventory/items`);
      setIsLoading(false);
      if (response.data.success) {
        setList(response.data.data);
      } else {
        toast.error('Error fetching the list.');
      }
    } catch (error) {
      setIsLoading(false);
      toast.error('Something went wrong. Please try again.');
    }
  };

  const confirmAndRemove = async (foodId) => {
    if (window.confirm('Are you sure you want to delete this item?')) {
      try {
        const response = await axios.post(`${url}/api/food/remove`, { id: foodId });
        if (response.data.success) {
          toast.success(response.data.message);
          fetchList();
        } else {
          toast.error('Error removing the item.');
        }
      } catch (error) {
        toast.error('Something went wrong. Please try again.');
      }
    }
  };

  useEffect(() => {
    fetchList();
  }, []);

  return (
    <div className="list-container">
      <p>All Medication List</p>
      {isLoading ? (
        <div className="spinner">Loading...</div>
      ) : list.length === 0 ? (
        <p className="empty-list">No items found.</p>
      ) : (
        <div className="list-table">
          <div className="list-table-format title">
            <b>Image</b>
            <b>Name</b>
            <b>Category</b>
            <b>Price</b>
            <b>Quantity</b>
            <b>Prescription</b>
            <b>Action</b>
          </div>
          {list.map((item, index) => (
            <div key={index} className="list-table-format">
              <img src={`${url}/images/` + item.image} alt={item.name} />
              <p>{item.name}</p>
              <p>{item.category}</p>
              <p>
                {currency}
                {item.price}
              </p>
              <p>{item.quantity}</p>
              <p>{item.needsPrescription ? 'Yes' : 'No'}</p>
              <p className="cursor" onClick={() => confirmAndRemove(item._id)}>
                x
              </p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default List;
