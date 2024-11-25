import React, { useState, useContext } from 'react';
import './OTCPage.css';
import ExploreMenu from '../../components/ExploreMenu/ExploreMenu';
import MedDisplay from '../../components/MedDisplay/MedDisplay';
import { StoreContext } from '../../Context/StoreContext';

const OTCPage = () => {
  const { menu_list } = useContext(StoreContext);
  const [category, setCategory] = useState('All');

  return (
    <div className="otc-page">
      <header className="otc-header">
      
      </header>
      <ExploreMenu category={category} setCategory={setCategory} />
      <MedDisplay category={category} />
    </div>
  );
};

export default OTCPage;
