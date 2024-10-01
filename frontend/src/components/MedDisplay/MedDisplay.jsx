import React, { useContext } from 'react';
import './MedDisplay.css';
import MedItem from '../MedItem/MedItem';
import { StoreContext } from '../../Context/StoreContext';

const MedDisplay = ({ category }) => {
  const { food_list } = useContext(StoreContext);

  return (
    <div className='food-display' id='food-display'>
      <div className='food-display-header'>
        <h2 className='main-heading'>Top Picks for Your Health</h2>
        <p className='subheading'>Get the best medications delivered to your doorsteps.</p>
      </div>
      <div className='food-display-list'>
        {food_list.map((item) => {
          if (category === 'All' || category === item.category) {
            return (
              <MedItem
                key={item._id}
                image={item.image}
                name={item.name}
                desc={item.description}
                price={item.price}
                id={item._id}
              />
            );
          }
          return null; // Ensures return type consistency
        })}
      </div>
    </div>
  );
};

export default MedDisplay;
