package model

import "time"

type User struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *User) Exists() bool {
	return m.ID > 0
}
