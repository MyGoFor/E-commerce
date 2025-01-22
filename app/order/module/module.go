package module

import (
	"gorm.io/gorm"
)

//外键存储，适合复杂查询，写少读多
//json存储，适合存储复杂数据结构且写多读少的场景

type UserItems []*OrderItem
type OrderItem struct {
	gorm.Model
	ProductId    uint32
	OrderIdRefer string `gorm:"size:256;index"`
	Quantity     int32
	Cost         float32
}

//func (oi UserItems) Value() (driver.Value, error) {
//	return json.Marshal(oi)
//}
//
//func (oi UserItems) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return fmt.Errorf("failed to scan StringArray: %v", value)
//	}
//	return json.Unmarshal(bytes, oi)
//}

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

//func (ad Consignee) AdValue() (driver.Value, error) {
//	return json.Marshal(ad)
//}
//
//func (ad Consignee) AdScan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return fmt.Errorf("failed to scan StringArray: %v", value)
//	}
//	return json.Unmarshal(bytes, ad)
//}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	gorm.Model
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
	OrderState   OrderState
}
