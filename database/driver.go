package database

import "gorm.io/gorm"

type Driver interface {
	Open(options map[string]any) (gorm.Dialector, error)
	Connect(options map[string]any) (*gorm.DB, error)
}
