### Authentication

- As guest, I can login with email/phone
- As guest, I can login with facebook (OAuth)
- As guest, I can login with google (OAuth)
- As guest, I can register with email/phone and password: Need to validate email(rules), and password(rules)
- As email/phone user, I must verify with email/phone after resgistering
- As email/phone user, I can reset password
- As user, I can logout

### Profile

- As user, I can update my profile (fullname, email, phone, avatar)
- As user, I can update my address(state, city, street): may need to store lng & lat of user location to process later (address info is just for reference) - can store more than one address???

### Payment

- As user, I can add/update payment method: Payment method may need to be enum for only accepted methods - Can user save more than one method???
- As user, I can payment methods info: Each methods will need different relevant info and can changed by user

### Order

- As user, I can see upcoming orders
- As user, I can see historic orders
- As user, I can see order details (food items, restaurants, status, orderId)
- As user, I can track my order
- As user, I can call delivery man
- As user, I can message delivery man
- As user, I can re-order historic orders
- As user, I can cancel orders

### Review

- As user, I can add review to foods/restaurants
- As user, I can rate to foods/restaurants: rating should be decimal (easily to find average, can be convert to stars order relevant text on client)
- As user, I can add restaurants to my favorite
- As user, I can see my favorites(restaurants/foods)
- As user, I can see comments about food/restaurants of other users
- As user, I can see average ratings of foods/restaurants

### Restaurant

- As user, I can see restaurants info (name, logo, address, delivery time, promotion, categories, rating): food categories can be gotten from food items
- As user, I can browse all food items of a restaurant
- As user, I can browse featured food items of a restaurant
- As user, I can browse custom bundles/categories of restaurant

### Foods

- A user, I can see food items info (name, images, price, short description, full description, rating)
- A user, I can choose topup/sub items of food items
- As user, I can increase/decrese items number
- A user, I can add items to cards

### Search

- As user, I can search for restaurants by searchbar
- As user, I can serch for nearby restaurants on map
- As user, I can search for foods by searchbar
- As user, I can search by filtering (categories, rating, price range, other conditions - popular, free delivery, nearest)

### Categories

- As user, I can see categories info (name, number of foods)
- As user, I can sort categories
- As user, I can browse food under categories

### Cart

- As user, I can increase/decrese items number
- As user, I can add promo code
- As user, I can see bill
- As user, I can checkout
- As user, I can confirm order
- As user, I can choose payment methods
