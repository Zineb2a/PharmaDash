import { createContext, useEffect, useState } from "react";
import { menu_list } from "../assets/assets";
import axios from "axios";

export const StoreContext = createContext(null);

const StoreContextProvider = (props) => {
    const url = "http://localhost:4000";
    const [food_list, setFoodList] = useState([]);
    const [cartItems, setCartItems] = useState({});
    const [token, setToken] = useState("");
    const currency = "$";
    const deliveryCharge = 5;

    const addToCart = async (itemId) => {
        if (!cartItems[itemId]) {
            setCartItems((prev) => ({ ...prev, [itemId]: 1 }));
        } else {
            setCartItems((prev) => ({ ...prev, [itemId]: prev[itemId] + 1 }));
        }
        if (token) {
            await axios.post(url + "/api/cart/add", { itemId }, { headers: { token } });
        }
    };

    const removeFromCart = async (itemId) => {
        setCartItems((prev) => ({ ...prev, [itemId]: prev[itemId] - 1 }));
        if (token) {
            await axios.post(url + "/api/cart/remove", { itemId }, { headers: { token } });
        }
    };

    const getTotalCartAmount = () => {
        let totalAmount = 0;
        for (const item in cartItems) {
            try {
                if (cartItems[item] > 0) {
                    let itemInfo = food_list.find((product) => product._id === item);
                    totalAmount += itemInfo.price * cartItems[item];
                }
            } catch (error) {
                console.error(error);
            }
        }
        return totalAmount;
    };

    // Use fake data to simulate API response
    const fetchFoodList = async () => {
        const fakeFoodList = [
            {
                _id: '1',
                name: 'Ibuprofen',
                category: 'Pain Relievers',
                description: 'Used to relieve pain and reduce inflammation.',
                price: 8.99,
                image: 'https://i-cf65.ch-static.com/content/dam/cf-consumer-healthcare/bp-advil-v2/en_US/products/product-01/advil-new-easy-open-cap-product-liqui-gel.png?auto=format',
            },
            {
                _id: '2',
                name: 'Vitamin D',
                category: 'Vitamins Supplements',
                description: 'Essential for maintaining healthy bones and teeth.',
                price: 12.50,
                image: 'https://via.placeholder.com/150',
            },
            {
                _id: '3',
                name: 'Cough Syrup',
                category: 'Cold and Flu',
                description: 'Effective for relieving coughs and throat irritation.',
                price: 5.99,
                image: 'https://via.placeholder.com/150',
            },
            {
                _id: '4',
                name: 'Antacid Tablets',
                category: 'Digestive System Relief',
                description: 'Provides fast relief from heartburn and acid indigestion.',
                price: 7.25,
                image: 'https://via.placeholder.com/150',
            },
            {
                _id: '5',
                name: 'Allergy Relief',
                category: 'Allergy and Antihistamines',
                description: 'Quick relief from allergy symptoms such as sneezing and runny nose.',
                price: 10.99,
                image: 'https://via.placeholder.com/150',
            },
        ];
        setFoodList(fakeFoodList);
    };

    const loadCartData = async (token) => {
        const response = await axios.post(url + "/api/cart/get", {}, { headers: token });
        setCartItems(response.data.cartData);
    };

    useEffect(() => {
        async function loadData() {
            await fetchFoodList();
            if (localStorage.getItem("token")) {
                setToken(localStorage.getItem("token"));
                await loadCartData({ token: localStorage.getItem("token") });
            }
        }
        loadData();
    }, []);

    const contextValue = {
        url,
        food_list,
        menu_list,
        cartItems,
        addToCart,
        removeFromCart,
        getTotalCartAmount,
        token,
        setToken,
        loadCartData,
        setCartItems,
        currency,
        deliveryCharge,
    };

    return (
        <StoreContext.Provider value={contextValue}>
            {props.children}
        </StoreContext.Provider>
    );
};

export default StoreContextProvider;
