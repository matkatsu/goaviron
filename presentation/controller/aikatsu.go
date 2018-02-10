package controller

import (
	"fmt"
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

// GetAikatsuUsers runs the GetAikatsuUsers action.
func (c *AikatsuController) GetAikatsuUsers(ctx *app.GetAikatsuUsersAikatsuContext) error {
	// AikatsuController_GetAikatsuUsers: start_implement

	// Put your logic here

	res := make(app.RomiogakuComGoavironAikatsuUserCollection, 0)
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
		u := &app.RomiogakuComGoavironAikatsuUser{
			ID:   i,
			Name: fmt.Sprintf("user%d", i),
		}
		res = append(res, u)
	}
	return ctx.OK(res)
	// AikatsuController_GetAikatsuUsers: end_implement
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
