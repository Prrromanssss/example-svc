package http

import "example-svc/internal/some_domain/usecases/models"

func ConvertModelToDTO(model GetExampleModel) models.GetExampleQuery {
	// maybe some logic here
	return models.GetExampleQuery{
		Param1: model.Param1,
	}
}
