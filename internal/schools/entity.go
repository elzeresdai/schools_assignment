package schools

import (
	"schools/internal/students"
)

type School struct {
	ID       string             `json:"id,omitempty"`
	Name     string             `json:"name"`
	Address  string             `json:"address"`
	Students []students.Student `json:"students,omitempty"`
}
