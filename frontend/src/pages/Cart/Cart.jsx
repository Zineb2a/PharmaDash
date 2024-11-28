import React, { useContext, useEffect } from "react";
import "./Cart.css";
import { StoreContext } from "../../Context/StoreContext";
import { useNavigate } from "react-router-dom";

const Cart = () => {
  const {
    cartItems,
    setCartItems, // Function to update cart items in context
    food_list,
    removeFromCart,
    getTotalCartAmount,
    url,
    currency,
    deliveryCharge,
  } = useContext(StoreContext);

  const navigate = useNavigate();

  // Fetch or create cart on component load
  useEffect(() => {
    const initializeCart = async () => {
      try {
        const response = await fetch("/api/cart/createOrFetch", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include", // Include cookies in the request
        });

        if (!response.ok) {
          console.error("Failed to initialize cart:", response.statusText);
          return;
        }

        const data = await response.json();
        if (data.cart_items) {
          setCartItems(data.cart_items); // Update the cart context
        } else {
          console.error("No cart items found:", data.status);
        }
      } catch (error) {
        console.error("Error initializing cart:", error);
      }
    };

    initializeCart();
  }, [setCartItems]);

  // Add an item to the cart
  const addItemToCart = async (itemId) => {
    try {
      const response = await fetch("/api/cart/addItem", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include", // Include cookies in the request
        body: JSON.stringify({
          inventory_id: itemId,
        }),
      });

      if (!response.ok) {
        console.error("Failed to add item to cart:", response.statusText);
        return;
      }

      const data = await response.json();
      console.log("Item added successfully:", data);

      // Update the cart locally (optional)
      setCartItems((prev) => ({
        ...prev,
        [itemId]: (prev[itemId] || 0) + 1,
      }));
    } catch (error) {
      console.error("Error adding item to cart:", error);
    }
  };

  // Remove an item from the cart
  const removeItemFromCart = async (cartItemId, inventoryItemId) => {
    try {
      const response = await fetch("/api/cart/removeItem", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include", // Include cookies in the request
        body: JSON.stringify({
          cart_item_id: cartItemId,
          inventory_item_id: inventoryItemId,
        }),
      });

      if (!response.ok) {
        console.error("Failed to remove item from cart:", response.statusText);
        return;
      }

      const data = await response.json();
      console.log("Item removed successfully:", data);

      // Update the cart locally
      setCartItems((prev) => {
        const updatedCart = { ...prev };
        if (updatedCart[inventoryItemId] > 1) {
          updatedCart[inventoryItemId] -= 1;
        } else {
          delete updatedCart[inventoryItemId];
        }
        return updatedCart;
      });
    } catch (error) {
      console.error("Error removing item from cart:", error);
    }
  };

  return (
    <div className="cart">
      <div className="cart-items">
        <div className="cart-items-title">
          <p>Items</p> <p>Title</p> <p>Price</p> <p>Quantity</p> <p>Total</p> <p>Remove</p>
        </div>
        <br />
        <hr />
        {food_list.map((item, index) => {
          if (cartItems[item._id] > 0) {
            return (
              <div key={index}>
                <div className="cart-items-title cart-items-item">
                  <img src={url + "/images/" + item.image} alt="" />
                  <p>{item.name}</p>
                  <p>
                    {currency}
                    {item.price}
                  </p>
                  <div>{cartItems[item._id]}</div>
                  <p>
                    {currency}
                    {item.price * cartItems[item._id]}
                  </p>
                  <p
                    className="cart-items-remove-icon"
                    onClick={() => removeItemFromCart(cartItems[item._id], item._id)}
                  >
                    x
                  </p>
                  <button onClick={() => addItemToCart(item._id)}>Add More</button>
                </div>
                <hr />
              </div>
            );
          }
          return null;
        })}
      </div>
      <div className="cart-bottom">
        <div className="cart-total">
          <h2>Cart Totals</h2>
          <div>
            <div className="cart-total-details">
              <p>Subtotal</p>
              <p>
                {currency}
                {getTotalCartAmount()}
              </p>
            </div>
            <hr />
            <div className="cart-total-details">
              <p>Delivery Fee</p>
              <p>
                {currency}
                {getTotalCartAmount() === 0 ? 0 : deliveryCharge}
              </p>
            </div>
            <hr />
            <div className="cart-total-details">
              <b>Total</b>
              <b>
                {currency}
                {getTotalCartAmount() === 0
                  ? 0
                  : getTotalCartAmount() + deliveryCharge}
              </b>
            </div>
          </div>
          <button onClick={() => navigate("/order")}>PROCEED TO CHECKOUT</button>
        </div>
        <div className="cart-promocode">
          <div>
            <p>If you have a promo code, Enter it here</p>
            <div className="cart-promocode-input">
              <input type="text" placeholder="promo code" />
              <button>Submit</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Cart;
