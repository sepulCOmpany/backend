package models

type User struct {
	UserName string `json:"username" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
	RoleID   int64  `json:"roleId" db:"role_id"`
}

type UserWithoutPassword struct {
	UserName string `json:"username" db:"username"`
	RoleID   int64  `json:"roleId" db:"role_id"`
}

type Sepulca struct {
	ID                int64 `json:"sepulcaId,omitempty" db:"id"`
	SizeID            int64 `json:"sizeId" db:"size_id"`
	ShmurdikID        int64 `json:"shmurdikId" db:"shmurdik_id"`
	GrimzikId         int64 `json:"grimzik_id" db:"grimzik_id"`
	SepulcaPropertyID int64 `json:"sepulcaPropertyId" db:"property_id"`
	IsVaccinated      *bool `json:"isVaccinated" db:"is_vaccinated"`
	IsRubbered        *bool `json:"isRubbered" db:"is_rubbered"`
	DeliveryStateID   int64 `json:"deliveryStateId" db:"delivery_state_id"`
}
