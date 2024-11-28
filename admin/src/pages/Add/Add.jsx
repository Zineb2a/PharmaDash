import React, { useState } from 'react';
import './Add.css';
import { assets, url } from '../../assets/assets';
import axios from 'axios';
import { toast } from 'react-toastify';

const Add = () => {
    const [data, setData] = useState({
        ItemName: '',
        MedicationName: '',
        description: '',
        quantity: '',
        price: '',
        OTC: false,
    });
    const [isSubmitting, setIsSubmitting] = useState(false);

    const onSubmitHandler = async (event) => {
        event.preventDefault();

        setIsSubmitting(true);

        const requestData = {
            
            PharmacyID: 1, 
            Item_Name: data.ItemName,
            Item_Description: data.description,
            Medication_Name: data.MedicationName,
            Unit_price: Number(data.price),
            Stock_Quantity: Number(data.quantity),
            OTC: data.OTC,
        };

        try {
            const url = 'http://localhost:3000';
            const response = await axios.post(`${url}/inventory/add_item`, requestData, {
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (response.data.success) {
                toast.success(response.data.message);
                setData({
                    ItemName: '',
                    MedicationName: '',
                    description: '',
                    quantity: '',
                    price: '',
                    OTC: false,
                });
            } else {
                toast.error(response.data.message);
            }
        } catch (error) {
            toast.error('Something went wrong. Please try again.');
        } finally {
            setIsSubmitting(false);
        }
    };
    const onChangeHandler = (event) => {
        const name = event.target.name;
        const value = event.target.type === 'checkbox' ? event.target.checked : event.target.value;
        setData((prevData) => ({ ...prevData, [name]: value }));
    };

    return (
        <div className="add">
            <form className="flex-col" onSubmit={onSubmitHandler}>
                <div className="add-product-name flex-col">
                    <p>Product Name</p>
                    <input
                        name="ItemName"
                        onChange={onChangeHandler}
                        value={data.ItemName}
                        type="text"
                        placeholder="Enter Item name"
                        required
                    />
                </div>
                <div className="add-medication-name flex-col">
                    <p> Medication Name</p>
                    <input
                        name="MedicationName"
                        onChange={onChangeHandler}
                        value={data.MedicationName}
                        type="text"
                        placeholder="Enter medication name"
                        required
                   />
                </div>
                <div className="add-product-description flex-col">
                    <p>Product Description</p>
                    <textarea
                        name="description"
                        onChange={onChangeHandler}
                        value={data.description}
                        rows={6}
                        placeholder="Enter product description"
                        required
                    />
                </div>
                <div className="add-category-price">
                    <div className="add-price flex-col">
                        <p>Product Price</p>
                        <input
                            type="number"
                            name="price"
                            onChange={onChangeHandler}
                            value={data.price}
                            placeholder="Enter price"
                            required
                        />
                    </div>
                    <div className="add-quantity flex-col">
                        <p>Product Quantity</p>
                        <input
                            type="number"
                            name="quantity"
                            onChange={onChangeHandler}
                            value={data.quantity}
                            placeholder="Enter quantity"
                            required
                        />
                    </div>
                </div>
                <div className="add-prescription flex-row">
                    <input
                        type="checkbox"
                        name="needsPrescription"
                        onChange={onChangeHandler}
                        checked={data.needsPrescription}
                        id="needsPrescription"
                    />
                    <label htmlFor="needsPrescription">OTC</label>
                </div>
                <button
                    type="submit"
                    className="add-btn"
                    disabled={isSubmitting}
                >
                    {isSubmitting ? 'Adding...' : 'Add'}
                </button>
            </form>
        </div>
    );
};

export default Add;
