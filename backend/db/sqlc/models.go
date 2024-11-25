// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

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

type Driver struct {
	DriverID    int32
	Name        string
	PhoneNumber pgtype.Text
	Email       pgtype.Text
	Password    string
}

type DriverOrder struct {
	DriverID int32
	OrderID  int32
}

type Inventory struct {
	InventoryID     int32
	PharmacyID      int32
	ItemName        string
	ItemDescription string
	MedicationName  string
	UnitPrice       pgtype.Numeric
	StockQuantity   int32
	Otc             pgtype.Bits
}

type Inventoryitem struct {
	InventoryItemID int32
	InventoryID     int32
	Reserved        pgtype.Bits
}

type Order struct {
	OrderID     int32
	AccountID   int32
	QuotationID int32
	OrderStatus pgtype.Text
	CreatedAt   pgtype.Timestamp
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
	TotalCost         pgtype.Numeric
	DeliveryFrequency string
	Destination       string
	SpecialHandling   pgtype.Text
	Insurance         pgtype.Numeric
	IncludeInsurance  pgtype.Bool
	IsRefused         pgtype.Bool
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
