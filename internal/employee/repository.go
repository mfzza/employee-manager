package employee

import (
	"encoding/json"
	"fmt"
	"os"
)

type Repository struct {
	filename string
}

func NewRepository(filename string) *Repository {
	return &Repository{filename: filename}
}

func (r *Repository) Load() ([]Employee, error) {
	data, err := os.ReadFile(r.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	var employees []Employee
	if err := json.Unmarshal(data, &employees); err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	// TODO: Determine last ID

	return employees, nil
}

func (r *Repository) Save(employees []Employee) error {
	data, err := json.Marshal(employees)
	fmt.Println(string(data))
	if err != nil {
		return err
	}

	return os.WriteFile(r.filename, data, 0644)
}
