import React, { useEffect, useState } from 'react';
import './List.css';
import { url, currency } from '../../assets/assets';
import axios from 'axios';
import { toast } from 'react-toastify';

const List = () => {
  const [isLoading, setIsLoading] = useState(true);
  const [list, setList] = useState([]); 

  const fetchList = async () => {
    setIsLoading(true);
    try {
      const response = await axios.get('http://localhost:3000/inventory/items');
      console.log(response.data);  
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

  useEffect(() => {
    fetchList();
  }, []);

  useEffect(() => {
    console.log(list); 
  }, [list]);

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
            <b>Item Name</b>
            <b>Medication Name</b>
            <b>Item Description</b>
            <b>Price</b>
            <b>Quantity</b>
            <b>OTC</b>
          </div>
          {list.map((item, index) => (
            <div key={index} className="list-table-format">
              {/* Display item details */}
              <p>{item.ItemName}</p>  
              <p>{item.MedicationName}</p>  
              <p>{item.ItemDescription}</p> 
              <p>{currency}{item.UnitPrice}</p>  
              <p>{item.StockQuantity}</p>  
              <p>{item.OTC ? 'Yes' : 'No'}</p>  
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default List;
