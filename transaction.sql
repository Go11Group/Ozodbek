-- Bizda users va products tablelari bor.
-- userni productslarida create, update, delete qilinganda,
-- transactiondan foydalanilsin. Siklda aylantirgan holda.

-- Users jadvali 
CREATE TABLE users ( 
    id SERIAL PRIMARY KEY, 
    username VARCHAR(50) NOT NULL, 
    email VARCHAR(100) NOT NULL UNIQUE, 
    password VARCHAR(100) NOT NULL ); 

-- Products jadvali 
CREATE TABLE products ( 
    id SERIAL PRIMARY KEY, 
    name VARCHAR(100) NOT NULL, 
    description TEXT, 
    price NUMERIC(10, 2) NOT NULL, 
    stock_quantity INT NOT NULL ); 

-- Crudlar: 
-- Create (Yaratish): 
-- CreateUser: Foydalanuvchi yaratish. 
-- CreateProduct: Mahsulot yaratish. 
-- Read (O'qish): 
-- GetUser: Foydalanuvchini o'qish. 
-- GetProduct: Mahsulotni o'qish. Update (Yangilash): 
-- UpdateUser: Foydalanuvchi ma'lumotlarini yangilash. 
-- UpdateProduct: Mahsulot ma'lumotlarini yangilash. Delete (O'chirish): 
-- DeleteUser: Foydalanuvchini o'chirish. 
-- DeleteProduct: Mahsulotni o'chirish. user_products: