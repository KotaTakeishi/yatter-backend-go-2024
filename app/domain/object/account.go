package object

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 本当はPasswordHashがハッシュされたパスワードであることを型で保証したい。
// ハッシュ化されたパスワード用の型を用意してstringと区別して管理すると良い。
// 今回は簡単のためstringで管理している。

type Account struct {
	// The internal ID of the account
	ID int64 `json:"id,omitempty"`

	// The username of the account
	Username string `json:"username,omitempty"`

	// The username of the account
	PasswordHash string `json:"-" db:"password_hash"`

	// The account's display name
	DisplayName *string `json:"display_name,omitempty" db:"display_name"`

	// URL to the avatar image
	Avatar *string `json:"avatar,omitempty"`

	// URL to the header image
	Header *string `json:"header,omitempty"`

	// Biography of user
	Note *string `json:"note,omitempty"`

	// The time the account was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}

func NewAccount(username, password string) (*Account, error) {
	account := &Account{
		Username: username,
		CreateAt: time.Now(),
	}
	if err := account.SetPassword(password); err != nil {
		return nil, fmt.Errorf("set password error: %w", err)
	}
	return account, nil
}

// Check if given password is match to account's password
func (a *Account) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
}

// Hash password and set it to account object
func (a *Account) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(hash), nil
}
