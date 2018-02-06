package controller

import (
	"github.com/goadesign/goa"
)

// VironController implements the Viron resource.
type VironController struct {
	*goa.Controller
}

// NewVironController creates a Viron controller.
func NewVironController(service *goa.Service) *VironController {
	return &VironController{Controller: service.NewController("VironController")}
}
