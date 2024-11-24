package UserService

import (
    "time"
)

type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
}

type Role struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Token struct {
    Value     string    `json:"value"`
    UserID    int       `json:"user_id"`
    ExpiresAt time.Time `json:"expires_at"`
}


type UserRepository interface {
    
}