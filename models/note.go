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

func NotesAll(u *User) *[]Note {
	var notes []Note
	//TODO
	DB.Where("deleted_at IS NULL and user_id = ?", user.ID).Order("updated_at")
	return &notes
}

func NoteCreate(user *User, name string, content string) *Note {
	entry := Note{Name: name, Content: content, UserID: user.ID}
	//TODO
	//DB.Create(&entry)
	return &entry
}

func NotesFind(user *User, nid uint64) *Note {
	var note Note
	//TODO
	//DB.Where("id = ? AND user_id = ?", id, user.ID).First(&note)
	return &note
}

func NotesMarkDelete(user *User, id uint64) {
	//TODO
	//DB.Where("id = ? AND user_id = ?", id, user.ID).Delete(&Note{})
}

func NotesUpdate(user *User, id uint64) {
	//TODO
}

func (n *Note) Update(name string, content string) {
	n.Name = name
	n.Content = content
	//TODO
	//DB.Save(note)
}
