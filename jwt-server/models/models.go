package models

type UserCredentials struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// // User represents a user in the system
// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // Product represents a product in the system
// type Product struct {
// 	ID    int     `json:"id"`
// 	Name  string  `json:"name"`
// 	Price float64 `json:"price"`
// }

// // Order represents an order in the system
// type Order struct {
// 	ID         int    `json:"id"`
// 	UserID     int    `json:"user_id"`
// 	ProductIDs []int  `json:"product_ids"`
// 	TotalPrice string `json:"total_price"`
// }
