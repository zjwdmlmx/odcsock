package model

type WearableDevice struct {
	Imei     string `gorm:"type:char(16);primark_key"`
	Password string `gorm:"type:char(16;not null)"`
}
