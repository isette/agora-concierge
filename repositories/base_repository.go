package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	db *gorm.DB
}

func (b *BaseRepository) WithTransaction(tx *gorm.DB) *BaseRepository {
	return &BaseRepository{db: tx}
}
