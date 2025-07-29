package models

import "errors"

type Patient struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Diagnostic string `json:"diagnostic"`
	Contact    string `json:"contact"`
	Location   string `json:"location"`
}

// Validate checks if the student data is valid
func (p *Patient) Validate() error {
	if len(p.Name) < 2 {
		return errors.New("name must be at least 2 characters")
	}

	return nil
}
