import React, { useState } from 'react'
import Home from './pages/Home/Home'
import Footer from './components/Footer/Footer'
import Navbar from './components/Navbar/Navbar'
import { Route, Routes } from 'react-router-dom'
import Cart from './pages/Cart/Cart'
import LoginPopup from './components/LoginPopup/LoginPopup'
import PlaceOrder from './pages/PlaceOrder/PlaceOrder'
import MyOrders from './pages/MyOrders/MyOrders'
import Prescription from './pages/Prescriptions/Prescriptions';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import Verify from './pages/Verify/Verify';
import Payment from './pages/Payment/Payment';
import PrivacyPolicy from './components/PrivacyPolicy/PrivacyPolicy';
import AboutUs from './components/AboutUs/AboutUs';
import ChatbotPage from './pages/Chatbot/ChatbotPage';
import OTCPage from './pages/OTCPage/OTCPage';
import ContactUs from './pages/ContactUs/ContactUs';


const App = () => {

  const [showLogin,setShowLogin] = useState(false);

  return (
    <>
    <ToastContainer/>
    {showLogin?<LoginPopup setShowLogin={setShowLogin}/>:<></>}
      <div className='app'>
        <Navbar setShowLogin={setShowLogin}/>
        <Routes>
          <Route path='/' element={<Home />}/>
          <Route path='/cart' element={<Cart />}/>
          <Route path='/order' element={<PlaceOrder />}/>
          <Route path='/myorders' element={<MyOrders />}/>
          <Route path='/verify' element={<Verify />}/>
          <Route path="/privacy-policy" element={<PrivacyPolicy />} />
          <Route path="/about-us" element={<AboutUs />} />
          <Route path="/otc" element={<OTCPage />} />
          <Route path="/payment" element={<Payment/>} />
          <Route path="/prescriptions" element={<Prescription />} />
          <Route path="/chatbot" element={<ChatbotPage />} /> 
          <Route path="/contact-us" element={<ContactUs />} />
          <Route path="/place-order" element={<PlaceOrder />}/>
        </Routes>
      </div>
      <Footer />
    </>
  )
}

export default App
