package usecases

import "fmt"

type CreateAppointmentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Interval    int    `json:"interval"`
}

func (r *CreateAppointmentRequest) Validate() error {
	if r.Title == "" && r.Date == "" && r.Description == "" {
		return fmt.Errorf("Request bod is empty or melformed")
	}

	if r.Interval <= 0 {
		return fmt.Errorf("Invalid value for the range field, this field must be filled with a value greater than Zero")
	}

	return nil
}
