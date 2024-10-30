package ds

type TaskLesson struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	ItemID    uint `json:"item_id" gorm:"uniqueIndex:item_request_key"`
	RequestID uint `json:"request_id" gorm:"uniqueIndex:item_request_key"`
	Forced    bool `json:"forced" gorm:"default:false"`

	Item    TaskItem      `json:"-" gorm:"foreignKey:ItemID"`
	Request LessonRequest `json:"-" gorm:"foreignKey:RequestID"`
}
