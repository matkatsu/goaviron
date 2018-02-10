package controller

import (
	"goaviron/app"
	"math/rand"
	"time"

	"github.com/goadesign/goa"
)

// AikatsuController implements the Aikatsu resource.
type AikatsuController struct {
	*goa.Controller
}

// NewAikatsuController creates a Aikatsu controller.
func NewAikatsuController(service *goa.Service) *AikatsuController {
	return &AikatsuController{Controller: service.NewController("AikatsuController")}
}

// GetOjisanNum runs the GetOjisanNum action.
func (c *AikatsuController) GetOjisanNum(ctx *app.GetOjisanNumAikatsuContext) error {
	// AikatsuController_GetOjisanNum: start_implement

	// Put your logic here

	res := &app.RomiogakuComGoavironAikatsuOjisan{}

	rand.Seed(time.Now().UnixNano())
	res.Value = rand.Intn(100000)

	return ctx.OK(res)
	// AikatsuController_GetOjisanNum: end_implement
}
