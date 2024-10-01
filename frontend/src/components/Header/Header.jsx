import React from 'react'
import './Header.css'

const Header = () => {
    return (
        <div className='header'>
            <div className='header-contents'>
                <h2>Your meds at your front porch, only a few clicks away!</h2>
                <p>Either you have a prescription or not, our main mission is to respond to your medical needs in a few steps only!</p>
                <button>Order Now</button>
            </div>
        </div>
    )
}

export default Header
