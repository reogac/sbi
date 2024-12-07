package sbi

import (
	"fmt"

	"github.com/reogac/sbi/models"
)

func ErrorFromProblemDetails(prob *models.ProblemDetails) error {
	return fmt.Errorf("Status: %d, Detail: %s", prob.Status, prob.Detail)
}
