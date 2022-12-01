package entity

//Image table in database
type Image struct {
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title    string `gorm:"type:varchar(255)" json:"title"`
	Caption  string `gorm:"type:varchar(255)" json:"caption"`
	PhotoUrl string `gorm:"type:varchar(255)" json:"photoulr"`
	UserID   uint64 `gorm:"not null" json:"-"`
	User     User   `gorm:"foreignkey:UserID;constrait:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
