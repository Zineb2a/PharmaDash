// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	AccountID   int32
	Name        string
	LastName    string
	Password    string
	PhoneNumber pgtype.Text
	Email       pgtype.Text
	Address     pgtype.Text
	Authlevel   string
}

type Feedback struct {
	FeedbackID int32
	OrderID    int32
	ClientID   int32
	Rating     int32
	Comment    pgtype.Text
	CreatedAt  pgtype.Timestamp
}

type Inventory struct {
	InventoryID     int32
	PharmacyID      int32
	ItemName        string
	ItemDescription string
	MedicationName  string
	UnitPrice       float64
	StockQuantity   int32
	Otc             bool
}

type Inventoryitem struct {
	InventoryItemID int32
	InventoryID     int32
	Reserved        bool
}

type Order struct {
	OrderID           int32
	AccountID         int32
	OrderStatus       pgtype.Text
	CreatedAt         pgtype.Timestamp
	DeliveryFrequency pgtype.Text
	Destination       pgtype.Text
	SpecialHandling   pgtype.Text
	IncludeInsurance  pgtype.Bool
	TotalCost         pgtype.Float8
	Insurance         pgtype.Float8
	DriverID          pgtype.Int4
}

type Orderitem struct {
	OrderItemID     int32
	OrderID         int32
	InventoryItemID int32
	Quantity        pgtype.Int4
}

type Pharmacy struct {
	PharmacyID  int32
	CompanyName string
	Location    string
}

type Prescription struct {
	PrescriptionID int32
	InventoryID    int32
	ClientEmail    string
	DoctorName     string
	MedicationName string
	DoseQuantity   int32
	IssuedDate     pgtype.Date
	ExpiryDate     pgtype.Date
}

type Quotationrequest struct {
	QuotationID       int32
	TotalCost         pgtype.Float8
	DeliveryFrequency pgtype.Text
	Destination       pgtype.Text
	SpecialHandling   pgtype.Text
	Insurance         pgtype.Float8
	IncludeInsurance  pgtype.Bool
	CartID            pgtype.Int4
}

type Shoppingcart struct {
	CartID    int32
	AccountID int32
}

type Shoppingcartitem struct {
	ShoppingCartItemID int32
	CartID             int32
	InventoryItemID    int32
}
