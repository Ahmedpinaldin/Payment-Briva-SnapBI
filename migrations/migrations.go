package migrations

import "time"

type User struct {
	ID        int64  `gorm:"primary_key:auto_increment"`
	Name      string `gorm:"type:varchar(255)"`
	Email     string `gorm:"uniqueIndex;type:varchar(255)"`
	Password  string `gorm:"<-;->:false;type:varchar(255)"`
	Role      string `gorm:"type:varchar(255)"`
	IDOffice  int64  `gorm:""`
	CreatedBy int64  `gorm:""`
	UpdatedBy int64  `gorm:""`
	CreatedAt time.Time
	UpdatedAt time.Time
}
 