package user

type User struct {
	Id          int    `json:"id"`
	Login       string `yaml:"login"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AvatarId    string `json:"avatar_id"`
	Description string `json:"description" `
}

type IUser interface {
	GetUserById(id int) (*User, error)
	GetUserByLogin(login string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id int) error
	ForgetPassword(email string) error
	GetAvatar(id int) (string, error)
	GetSubscribers(id int) ([]int, error)
	GetFollowers(id int) ([]int, error)
}

func (u *User) GetUserById(id int) (*User, error) {
	return nil, nil
}

func (u *User) GetUserByLogin(login string) (*User, error) {
	return nil, nil
}

func (u *User) GetUserByEmail(email string) (*User, error) {
	return nil, nil
}

func (u *User) CreateUser(user *User) error {
	return nil
}

func (u *User) UpdateUser(user *User) error {
	return nil
}

func (u *User) DeleteUser(id int) error {
	return nil
}

func (u *User) ForgetPassword(email string) error {
	return nil
}

func (u *User) GetAvatar(id int) (string, error) {
	return "", nil
}

func (u *User) GetSubscribers(id int) ([]int, error) {
	return nil, nil
}

func (u *User) GetFollowers(id int) ([]int, error) {
	return nil, nil
}
