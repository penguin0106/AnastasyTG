package models

// User представляет собой модель пользователя
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Product представляет собой модель товара
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Order представляет собой модель заказа
type Order struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	ProductIDs []int  `json:"product_ids"`
	Status     string `json:"status"`
}
