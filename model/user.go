package model

import "time"

// User data model
type User struct {
	ID        int    `storm:"id,increment"` // primary key with auto increment
	Group     string `storm:"index"`        // this field will be indexed
	Email     string `storm:"unique"`       // this field will be indexed with a unique constraint
	Name      string // this field will not be indexed
	Age       int
	CreatedAt time.Time `storm:"index"`
}
