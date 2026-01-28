package main

// configdec.go
type ConfigFile struct {
	// Database
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     int
	// Server
	ServerPort int
	ServerName string
}

// account.go
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Avatar   string
	Available bool
}

// post.go
type Tag struct {
	ID      int
	Title   string
	Content string
	Author  User
}

type Post struct {
	ID      int
	ContentFile string
	Author  User
	Tag     []Tag
}