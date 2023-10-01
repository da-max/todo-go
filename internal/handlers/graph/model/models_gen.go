// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ChangePassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	OldPassword     string `json:"oldPassword"`
}

type ChangePasswordConfirm struct {
	Ok bool `json:"ok"`
}

type Confirm struct {
	Ok    bool   `json:"ok"`
	Token string `json:"token"`
}

type ConfirmIdentifier struct {
	Token string `json:"token"`
}

type DeleteAccount struct {
	Ok bool `json:"ok"`
}

type Identifier struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewTodo struct {
	Text string `json:"text"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestConfirmAccount struct {
	Ok bool `json:"ok"`
}

type RequestPasswordResetIdentifier struct {
	Email string `json:"email"`
}

type RequestResetPassword struct {
	Ok bool `json:"ok"`
}

type ResetPasswordIdentifier struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Todo struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Done   bool   `json:"done"`
	UserID string `json:"userId"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UpdateUser struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
	IsAdmin  bool   `json:"isAdmin"`
}