package userDomain

import coreUuid "github.com/shinYeongHyeon/eveningIsFood-api/src/shared/core/uuid"

// User is domain type definition.
type User struct {
	uuid     string
	name     Name
	id       Id
	password Password
}

// NewUserProps is a struct for creating a new User domain.
type NewUserProps struct {
	name     Name
	id       Id
	password Password
}

// CreateNewUser creates a new User domain.
func CreateNewUser(props NewUserProps) (User, error) {
	return User{
		uuid:     coreUuid.CreateUuid(),
		name:     props.name,
		id:       props.id,
		password: props.password,
	}, nil
}

// GetUuid returns the Uuid.
func (user User) GetUuid() string {
	return user.uuid
}

// GetName returns the Name domain.
func (user User) GetName() Name {
	return user.name
}

// GetId returns the Id domain.
func (user User) GetId() Id {
	return user.id
}

// GetPassword returns the Password domain.
func (user User) GetPassword() Password {
	return user.password
}
