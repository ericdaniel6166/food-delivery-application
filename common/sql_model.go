package common

import (
	"time"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"primaryKey;column:id"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (sqlModel *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(sqlModel.Id), dbType, 1)
	sqlModel.FakeId = &uid
}

func (sqlModel *SQLModel) PrepareForInsert() {
	now := time.Now().UTC()
	sqlModel.Id = 0
	sqlModel.Status = 1
	sqlModel.CreatedAt = &now
	sqlModel.UpdatedAt = &now
}

func (sqlModel *SQLModel) PrepareForUpdate() {
	now := time.Now().UTC()
	sqlModel.UpdatedAt = &now
}
