package employee

import (
	"encoding/json"
	"errors"
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
        if errors.Is(err, os.ErrNotExist) {
            return []Employee{}, nil
        }
        return nil, fmt.Errorf("failed to read file: %w", err)
    }

    if len(data) == 0 {
        return []Employee{}, nil
    }

    var employees []Employee
    if err := json.Unmarshal(data, &employees); err != nil {
        return nil, fmt.Errorf("invalid JSON format: %w", err)
    }

    return employees, nil
}

func (r *Repository) Save(employees []Employee) error {
    data, err := json.MarshalIndent(employees, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to encode data: %w", err)
    }

    return os.WriteFile(r.filename, data, 0644)
}

