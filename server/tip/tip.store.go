package tip

import (
	"errors"

	"gorm.io/gorm"
)

type TipStore struct {
	db *gorm.DB
}

func NewTipStore(db *gorm.DB) *TipStore {
	return &TipStore{
		db: db,
	}
}

func (us *TipStore) GetAll() ([]Tip, int64, error) {
	var tips []Tip
	var count int64
	if err := us.db.Find(&tips).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, nil
		}
		return nil, 0, err
	}

	return tips, count, nil
}

func (us *TipStore) Create(u *Tip) (err error) {
	return us.db.Create(u).Error
}
