import React, { useState } from 'react';
import './Add.css';
import { assets, url } from '../../assets/assets';
import axios from 'axios';
import { toast } from 'react-toastify';

const Add = () => {
    const [image, setImage] = useState(null);
    const [data, setData] = useState({
        name: '',
        description: '',
        price: '',
        category: 'Cold and Flu',
        quantity: '',
        needsPrescription: false, // Added checkbox field
    });
    const [isSubmitting, setIsSubmitting] = useState(false);

    const onSubmitHandler = async (event) => {
        event.preventDefault();

        if (!image) {
            toast.error('Please upload an image.');
            return;
        }

        setIsSubmitting(true);

        const formData = new FormData();
        formData.append('name', data.name);
        formData.append('description', data.description);
        formData.append('price', Number(data.price));
        formData.append('category', data.category);
        formData.append('quantity', Number(data.quantity));
        formData.append('needsPrescription', data.needsPrescription); // Send checkbox value to backend
        formData.append('image', image);

        try {
            const response = await axios.post(`${url}/api/food/add`, formData);
            if (response.data.success) {
                toast.success(response.data.message);
                setData({
                    name: '',
                    description: '',
                    price: '',
                    category: data.category,
                    quantity: '',
                    needsPrescription: false,
                });
                setImage(null);
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
                <div className="add-img-upload flex-col">
                    <p>Upload Image</p>
                    <input
                        onChange={(e) => {
                            setImage(e.target.files[0]);
                            e.target.value = '';
                        }}
                        type="file"
                        accept="image/*"
                        id="image"
                        hidden
                    />
                    <label htmlFor="image">
                        <img
                            src={image ? URL.createObjectURL(image) : assets.upload_area}
                            alt="Product Preview"
                        />
                    </label>
                </div>
                <div className="add-product-name flex-col">
                    <p>Product Name</p>
                    <input
                        name="name"
                        onChange={onChangeHandler}
                        value={data.name}
                        type="text"
                        placeholder="Enter product name"
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
                    <div className="add-category flex-col">
                        <p>Product Category</p>
                        <select
                            name="category"
                            onChange={onChangeHandler}
                            value={data.category}
                        >
                            <option value="Cold and Flu">Cold and Flu</option>
                            <option value="Allergy and Antihistamines">Allergy and Antihistamines</option>
                            <option value="Muscle Reliefs">Muscle Reliefs</option>
                            <option value="Digestive System Relief">Digestive System Relief</option>
                            <option value="Pain Relievers">Pain Relievers</option>
                            <option value="Vitamins Suppliments">Vitamins Supplements</option>
                            <option value="Topical Antibiotics">Topical Antibiotics</option>
                            <option value="First Aid and Wound Care">First Aid and Wound Care</option>
                        </select>
                    </div>
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
                    <label htmlFor="needsPrescription">Needs Prescription</label>
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
