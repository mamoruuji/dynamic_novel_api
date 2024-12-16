package models

import (
	"context"

	"github.com/mamoruuji/dynamic_novel_api/db/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ContextExecutor interface {
	boil.ContextExecutor
	All(ctx context.Context, exec boil.ContextExecutor) ([]*models.Tag, error)
}
