import React from 'react';
import './AppDownload.css';
import { assets } from '../../assets/assets';

const AppDownload = () => {
    return (
        <div className='app-download' id='app-download'>
            <p><br /> Download Pharmadash App</p>
            <p className='coming-soon'>Coming Soon!</p> {/* Add a "Coming Soon" message */}
            <div className="app-download-platforms">
                <div className='disabled-button'>
                    <img src={assets.play_store} alt="Play Store" />
                </div>
                <div className='disabled-button'>
                    <img src={assets.app_store} alt="App Store" />
                </div>
            </div>
        </div>
    );
}

export default AppDownload;
