package entities

type Employee struct {
	ID         int    `json:"id"` // Struct Tags
	Name       string `json:"name" validate:"required"`
	Department string `json:"speciality" validate:"required"`
	ProjectID  int    `json:"-"`
}
