package model

type User struct {
	ID        uint32
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	users           = make(map[uint32]User)
	globalID uint32 = 0
)

func CreateUser(firstName, lastName string) *User {

	globalID++
	user := User{
		ID:        globalID,
		FirstName: firstName,
		LastName:  lastName,
	}

	users[user.ID] = user
	return &user
}

func GetUser(id uint32) (*User, bool) {
	user, ok := users[id]
	if !ok { // to prefent null pointer
		return nil, false
	}
	return &user, ok
}
