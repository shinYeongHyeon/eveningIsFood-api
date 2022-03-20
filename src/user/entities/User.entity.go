package userEntity

// UserEntity : Entity of user
type UserEntity struct {
	UUID      string `gorm:"primaryKey"`
	Id        string `gorm:"type:varchar(255);not_null;unique;comment:유저 이메일"`
	Password  string `gorm:"type:varchar(255);not_null;comment:유저 비밀번호"`
	Name      string `gorm:"type:varchar(255);not_null;comment:유저 이름"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}
