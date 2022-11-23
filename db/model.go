package db

type Model struct {
	ID        uint  `gorm:"primaryKey,autoIncrement"`
	CreatedAt int64 `gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `gorm:"autoUpdateTime:milli"`
	DeletedAt int64 `gorm:"default:0"`
	Usn       int64
}
