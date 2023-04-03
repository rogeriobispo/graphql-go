package model

type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Courses     []*Course `json:"courses,omitempty"`
}

type NewCategory struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
