package models

type MemberLevel struct {
	BeforeID int    `gorm:"column:before_id" json:"before_id"`
	ID       int    `gorm:"column:id;primary_key" json:"id;primary_key"`
	MemberID int    `gorm:"column:member_id" json:"member_id"`
	Records  string `gorm:"column:records" json:"records"`
}

// TableName sets the insert table name for this struct type
func (m *MemberLevel) TableName() string {
	return "member_level"
}
