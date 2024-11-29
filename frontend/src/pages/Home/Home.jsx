import React, { useState, useEffect } from 'react';
import Header from '../../components/Header/Header';
import ExploreMenu from '../../components/ExploreMenu/ExploreMenu';
import MedDisplay from '../../components/MedDisplay/MedDisplay';
import AppDownload from '../../components/AppDownload/AppDownload';
import LeaveAReview from '../../components/LeaveAReview/LeaveAReview';

const Home = () => {
  const [category, setCategory] = useState("All");

  // Scroll to the top when the component mounts
  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  return (
    <>
      <Header />
  
      <AppDownload />
      <LeaveAReview/>
    </>
  );
}

export default Home;
