package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              int    `gorm:"primaryKey"`
	Email           string `gorm:"unique;not null"`
	PasswordHashed  string `gorm:"not null"`
	FirstName       string
	LastName        string
	IsActive        bool      `gorm:"default:true"`
	IsEmailVerified bool      `gorm:"default:false"`
	CreatedAt       time.Time `gorm:"default:now()"`
	UpdatedAt       time.Time
}

type Role struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"unique;not null"`
	CreatedAt time.Time `gorm:"default:now()"`
}
type UserRole struct {
	UserID     int `gorm:"primaryKey"`
	RoleID     int `gorm:"primaryKey"`
	User       User
	Role       Role
	AssignedAt time.Time `gorm:"default:now()"`
}

type Category struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Slug        string `gorm:"unique;not null"`
	Description string
	ParentID    *int
	CreatedAt   time.Time `gorm:"default:now()"`
	UpdatedAt   time.Time

	Parent *Category `gorm:"foreignKey:ParentID"`
}

type Product struct {
	ID          int    `gorm:"primaryKey"`
	Stock       string `gorm:"unique"`
	Name        string `gorm:"not null"`
	Slug        string `gorm:"unique;not null"`
	Description string
	PriceCents  int64  `gorm:"not null"`
	Currency    string `gorm:"default:'EGP'"`
	Active      bool   `gorm:"default:true"`
	CategoryID  *int
	CreatedAt   time.Time `gorm:"default:now()"`
	UpdatedAt   time.Time

	Category *Category `gorm:"foreignKey:CategoryID"`
}

type Cart struct {
	ID           int `gorm:"primaryKey"`
	UserID       *int
	SessionToken uuid.UUID `gorm:"default:uuid_generate_v4()"`
	CreatedAt    time.Time `gorm:"default:now()"`
	UpdatedAt    time.Time

	User *User `gorm:"foreignKey:UserID"`

	Items []CartItem
}

type CartItem struct {
	ID             int       `gorm:"primaryKey"`
	CartID         int       `gorm:"not null"`
	ProductID      int       `gorm:"not null"`
	Quantity       int       `gorm:"not null"`
	UnitPriceCents int64     `gorm:"not null"`
	CreatedAt      time.Time `gorm:"default:now()"`

	Cart    *Cart    `gorm:"foreignKey:CartID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID                int `gorm:"primaryKey"`
	UserID            *int
	OrderNumber       string `gorm:"unique;not null"`
	Status            string `gorm:"type:order_status;default:pending"`
	SubtotalCents     int64  `gorm:"not null"`
	ShippingCents     int64  `gorm:"default:0"`
	DiscountCents     int64  `gorm:"default:0"`
	TotalCents        int64  `gorm:"not null"`
	Currency          string `gorm:"default:'USD'"`
	ShippingAddressID *int   //pointer for nullable foreign key
	BillingAddressID  *int
	CreatedAt         time.Time `gorm:"default:now()"`
	UpdatedAt         time.Time

	User            *User    `gorm:"foreignKey:UserID"`
	ShippingAddress *Address `gorm:"foreignKey:ShippingAddressID"`
	BillingAddress  *Address `gorm:"foreignKey:BillingAddressID"`
	Items           []OrderItem
}

type OrderItem struct {
	ID                  int    `gorm:"primaryKey"`
	OrderID             int    `gorm:"not null"`
	ProductID           int    `gorm:"not null"`
	ProductNameSnapshot string `gorm:"not null"`
	SkuSnapshot         string
	Quantity            int   `gorm:"not null;check:quantity > 0"`
	UnitPriceCents      int64 `gorm:"not null"`
	LineTotalCents      int64 `gorm:"not null"`

	Order   *Order   `gorm:"foreignKey:OrderID"`
	Product *Product `gorm:"foreignKey:ProductID"`
}
type ProductImage struct {
	ID        int       `gorm:"primaryKey"`
	ProductID int       `gorm:"not null"`
	URL       string    `gorm:"not null"`
	IsPrimary bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:now()"`

	Product *Product `gorm:"foreignKey:ProductID"`
}

type Inventory struct {
	ID        int       `gorm:"primaryKey"`
	ProductID int       `gorm:"unique;not null"`
	Available int       `gorm:"not null;default:0"`
	Reserved  int       `gorm:"not null;default:0"`
	UpdatedAt time.Time `gorm:"default:now()"`

	Product *Product `gorm:"foreignKey:ProductID"`
}

type Address struct {
	ID         int `gorm:"primaryKey"`
	UserID     int `gorm:"not null"`
	Label      string
	Line1      string `gorm:"not null"`
	Line2      string
	City       string `gorm:"not null"`
	Region     string
	PostalCode string
	Country    string `gorm:"not null"`
	Phone      string
	IsDefault  bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time

	User *User `gorm:"foreignKey:UserID"`
}

type Payment struct {
	ID                int `gorm:"primaryKey"`
	OrderID           int `gorm:"not null"`
	Provider          string
	ProviderPaymentID string
	AmountCents       int64     `gorm:"not null"`
	Currency          string    `gorm:"default:'EGP'"`
	Status            string    `gorm:"type:payment_status;default:pending"`
	CreatedAt         time.Time `gorm:"default:now()"`
	UpdatedAt         time.Time

	Order *Order `gorm:"foreignKey:OrderID"`
}

type Coupon struct {
	ID            int    `gorm:"primaryKey"`
	Code          string `gorm:"unique;not null"`
	Description   string
	DiscountType  string  `gorm:"not null;check:discount_type IN ('percentage', 'fixed')"`
	DiscountValue float64 `gorm:"type:numeric;not null"`
	MaxUses       *int
	UsedCount     int `gorm:"default:0"`
	ExpiresAt     time.Time
	Active        bool      `gorm:"default:true"`
	CreatedAt     time.Time `gorm:"default:now()"`
}

type CouponRedemption struct {
	CouponID   int `gorm:"primaryKey"`
	OrderID    int `gorm:"primaryKey"`
	UserID     *int
	RedeemedAt time.Time `gorm:"default:now()"`

	Coupon *Coupon `gorm:"foreignKey:CouponID"`
	Order  *Order  `gorm:"foreignKey:OrderID"`
	User   *User   `gorm:"foreignKey:UserID"`
}

type Review struct {
	ID        int `gorm:"primaryKey"`
	ProductID int `gorm:"not null"`
	UserID    *int
	Rating    int `gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time

	Product *Product `gorm:"foreignKey:ProductID"`
	User    *User    `gorm:"foreignKey:UserID"`
}

type AuditLog struct {
	ID        int `gorm:"primaryKey"`
	UserID    *int
	EventType string `gorm:"not null"`
	EventData []byte `gorm:"type:jsonb"`
	IPAddress string
	CreatedAt time.Time `gorm:"default:now()"`

	User *User `gorm:"foreignKey:UserID"`
}

type EmailVerification struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Token     string    `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:now()"`

	User *User `gorm:"foreignKey:UserID"`
}

type PasswordReset struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Token     string    `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Used      bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:now()"`

	User *User `gorm:"foreignKey:UserID"`
}

type RefreshToken struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"not null"`
	Token     string `gorm:"unique;not null"`
	UserAgent string
	IPAddress string
	Revoked   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"default:now()"`
	ExpiresAt time.Time

	User *User `gorm:"foreignKey:UserID"`
}
