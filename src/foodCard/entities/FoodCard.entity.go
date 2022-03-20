package foodCardEntity

// FoodCardEntity : Entity of foodCard
type FoodCardEntity struct {
	UUID      string `gorm:"primaryKey"`
	Name      string `gorm:"not_null"`
	Address   string `gorm:"not_null"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}
