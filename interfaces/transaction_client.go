package interfaces

import "gorm.io/gorm"

type TransactionClient interface {
	GetDBClient() *gorm.DB
	SetDBClient(tx *gorm.DB) any
}
