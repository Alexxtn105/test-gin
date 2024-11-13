package models

import (
	"gorm.io/gorm"
	"time"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	UserID    uint64 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NotesAll(user *User) *[]Note {
	var notes []Note
	DB.Where("deleted_at IS NULL and user_id = ?", user.ID).Order("updated_at")
	return &notes
}

func NoteCreate(user *User, name string, content string) *Note {
	entry := Note{Name: name, Content: content, UserID: user.ID}
	DB.Create(&entry)
	return &entry
}

func NotesUpdate(user *User, id uint64) {
	//TODO
	//DB.Where("id = ? AND user_id = ?", id, user.ID).Update(&Note{})

}

func NotesFind(user *User, id uint64) *Note {
	var note Note
	DB.Where("id = ? AND user_id = ?", id, user.ID).First(&note)
	return &note
}

func (n *Note) Update(name string, content string) {
	n.Name = name
	n.Content = content
	DB.Save(n)
}

func NotesMarkDelete(user *User, id uint64) {
	DB.Where("id = ? AND user_id = ?", id, user.ID).Delete(&Note{})
}
