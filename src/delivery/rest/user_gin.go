package rest

import (
	"github.com/dalikewara/ayapingping-go-crud/src/service"
	"github.com/gin-gonic/gin"
)

// UserRegister handles gin-gonic REST route to register user data.
func (h *ginHandler) UserRegister(g *gin.Context) {
	ctx := h.getRequestContext(g)
	req := &UserRegisterRequest{}
	if err := h.parseJSONRequest(g, &req); err != nil {
		h.jsonNotOk(g, err.GetStatus(), err.GetCode(), err.GetMessage())
		return
	}

	svcRes := h.service.User.Register(service.UserRegisterParam{
		Ctx:                  ctx,
		Username:             req.Username,
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
		FirstName:            req.FirstName,
		LastName:             req.LastName,
		Gender:               req.Gender,
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	h.jsonCompose(g, &UserRegisterResponseData{}, svcRes)
}

// UserLogin handles gin-gonic REST route to log in user.
func (h *ginHandler) UserLogin(g *gin.Context) {
	ctx := h.getRequestContext(g)
	req := &UserLoginRequest{}
	if err := h.parseJSONRequest(g, &req); err != nil {
		h.jsonNotOk(g, err.GetStatus(), err.GetCode(), err.GetMessage())
		return
	}

	svcRes := h.service.User.Login(service.UserLoginParam{
		Ctx:             ctx,
		UsernameOrEmail: req.UsernameOrEmail,
		Password:        req.Password,
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	var responseData *UserLoginResponseData
	h.jsonCompose(g, &responseData, svcRes.User)
}
