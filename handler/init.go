package handler

import (
	"github.com/budiariyanto/cqrs-dummy/service"
)

type Handler struct {
	DenormalizeSvc *service.DenormalizeService
}
