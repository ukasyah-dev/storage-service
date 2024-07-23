package model

type Tag struct {
	Key   string `gorm:"primaryKey" json:"key"`
	Value string `gorm:"primaryKey" json:"value"`
}
