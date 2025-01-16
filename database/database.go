package database

import (
	"time"
)

// 一些数据库结构体参考
type Token struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"index"`
	Token        string    `gorm:"unique"`
	RefreshToken string    `gorm:"unique"`
	ExpiresAt    time.Time `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Status    int       `gorm:"default:1"` // 1: active, 0: inactive
}

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"type:text"`
	Picture     string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	Stock       int       `gorm:"not null"`
	Categories  []string  `gorm:"type:text[]"` // Postgres-specific for array of categories
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Status      int       `gorm:"default:1"` // 1: available, 0: unavailable
}

type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type CartItem struct {
	ID        uint      `gorm:"primaryKey"`
	CartID    uint      `gorm:"index;not null"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Address struct {
	ID        uint      `gorm:"primaryKey"`
	Street    string    `gorm:"not null"`
	City      string    `gorm:"not null"`
	State     string    `gorm:"not null"`
	Country   string    `gorm:"not null"`
	ZipCode   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Order struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index;not null"`
	Email      string    `gorm:"not null"`
	Currency   string    `gorm:"not null"`
	AddressID  uint      `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
	OrderItems []OrderItem
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"index;not null"`
	ProductID uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Cost      float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Checkout struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	FirstName string    `gorm:"not null"`
	LastName  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	AddressID uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type CreditCardInfo struct {
	ID                        uint      `gorm:"primaryKey"`
	CreditCardNumber          string    `gorm:"not null"`
	CreditCardCVV             int       `gorm:"not null"`
	CreditCardExpirationYear  int       `gorm:"not null"`
	CreditCardExpirationMonth int       `gorm:"not null"`
	CreatedAt                 time.Time `gorm:"autoCreateTime"`
}

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"index;not null"`
	Amount        float64   `gorm:"not null"`
	TransactionID string    `gorm:"unique;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
}

type AIOrder struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index;not null"`
	OrderID   string    `gorm:"unique;not null"`
	Status    int       `gorm:"default:0"` // 0: pending, 1: completed
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
