package entity

type Store struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	StoreEmail string `json:"store_email"`
	Password   string `json:"password"`
	StoreName  string `json:"store_name"`
	StoreType  string `json:"store_type"`
}
