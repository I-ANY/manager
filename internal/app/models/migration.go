package models

type Migration struct {
	Version   string `gorm:"primaryKey;size:64" json:"version"`
	ApplyTime uint32 `gorm:"column:apply_time" json:"apply_time"`
}

func (Migration) TableName() string {
	return "sys_migration"
}
