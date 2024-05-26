package models

import "golang.org/x/crypto/bcrypt"

type User struct {
    ID       string `json:"id" bson:"_id,omitempty"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// HashPassword hashes the user's password
func (u *User) HashPassword() error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
    if err != nil {
        return err
    }
    u.Password = string(bytes)
    return nil
}

// CheckPassword compares the hashed password with the given password
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}
