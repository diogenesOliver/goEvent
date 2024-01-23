package usecases

import "fmt"

type CreateAppointmentRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Interval    int    `json:"interval"`
}

func (r *CreateAppointmentRequest) Validate() error {
	if r.Title == "" && r.Date == "" && r.Description == "" && r.Interval < 0 {
		return fmt.Errorf("Request bod yis empty or melformed")
	}
	return nil
}
