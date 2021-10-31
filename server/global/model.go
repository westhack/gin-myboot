package global

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint64         `json:"id" gorm:"primarykey;comment:ID"` // 主键ID
	CreatedAt LocalTime      `json:"createdAt" gorm:"comment:创建时间"`   // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:更新时间"`   // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                  // 删除时间
}

// LocalTime format json time field by myself
type LocalTime struct {
	time.Time
}

// MarshalJSON on LocalTime format Time field with %Y-%m-%d %H:%M:%S
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
