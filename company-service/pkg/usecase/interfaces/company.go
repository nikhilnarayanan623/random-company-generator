package interfaces

import (
	"company-service/pkg/domain"
	"company-service/pkg/models"
	"context"
)

type CompanyUseCase interface {
	Create(ctx context.Context, companyReq models.CompanyRequest) (*domain.Company, error)
}
