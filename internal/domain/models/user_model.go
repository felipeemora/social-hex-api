package domain

type UserModel struct {
	ID           int64
	Username     string
	Email        string
	Password     string
	PasswordHash []byte
	CreatedAt    string
	IsActive     bool
	RoleId       int64
	Role         *RoleModel
	InvitationID string
}

type RoleModel struct {
	ID          int64
	Name        string
	Description string
	Level       int
}
