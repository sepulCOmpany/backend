package models

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	RoleID   int64  `json:"roleId" db:"role_id"`
}

type UserWithoutPassword struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	RoleID   int64  `json:"roleId" db:"role_id"`
}

type Sepulca struct {
	ID                int64  `json:"sepulcaId,omitempty" db:"id"`
	SizeID            int64  `json:"sizeId" db:"size_id" binding:"required"`
	ShmurdikID        int64  `json:"shmurdikId" db:"shmurdik_id" binding:"required"`
	GrimzikID         int64  `json:"grimzik_id" db:"grimzik_id" binding:"required"`
	SepulcaPropertyID int64  `json:"sepulcaProperty" db:"property_id" binding:"required"`
	IsVaccinated      *bool  `json:"isVaccinated" db:"is_vaccinated"`
	IsRubbered        *bool  `json:"isRubbered" db:"is_rubbered"`
	DeliveryStateID   int64  `json:"deliveryStateId" db:"delivery_state_id"`
	Shmurdik          string `json:"shmurdik_username" db:"shmurdik_username"`
	Grimzik           string `json:"grimzik_username" db:"shmurdik_username"`
}
