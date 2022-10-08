package rest

import (
	"github.com/dalikewara/ayapingping-go-crud/src/entity"
	"github.com/dalikewara/ayapingping-go-crud/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserGetAllActive handles gin-gonic REST route to get all active users.
func (h *ginHandler) UserGetAllActive(g *gin.Context) {
	ctx := h.getRequestContext(g)

	svcRes := h.service.User.GetAllActive(service.UserGetAllActiveParam{
		Ctx: ctx,
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	responseData := UserGetAllActiveResponseData{}
	if err := h.jsonComposeWithoutWrite(&responseData.Rows, svcRes.Users); err != nil {
		h.jsonNotOk(g, err.GetStatus(), err.GetCode(), err.GetMessage())
		return
	}

	h.jsonOk(g, responseData)
}

// UserGetDetail handles gin-gonic REST route to get user detail.
func (h *ginHandler) UserGetDetail(g *gin.Context) {
	ctx := h.getRequestContext(g)
	id, _ := strconv.Atoi(g.Param("id"))

	svcRes := h.service.User.GetDetail(service.UserGetDetailParam{
		Ctx: ctx,
		ID:  entity.ID(id),
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	var responseData *UserGetDetailResponseData
	h.jsonCompose(g, &responseData, svcRes.User)
}

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

// UserUpdate handles gin-gonic REST route to update user data.
func (h *ginHandler) UserUpdate(g *gin.Context) {
	ctx := h.getRequestContext(g)
	req := &UserUpdateRequest{}
	if err := h.parseJSONRequest(g, &req); err != nil {
		h.jsonNotOk(g, err.GetStatus(), err.GetCode(), err.GetMessage())
		return
	}

	svcRes := h.service.User.Update(service.UserUpdateParam{
		Ctx:       ctx,
		ID:        req.ID,
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Gender:    req.Gender,
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	h.jsonOk(g, nil)
}

// UserDelete handles gin-gonic REST route to delete user data.
func (h *ginHandler) UserDelete(g *gin.Context) {
	ctx := h.getRequestContext(g)
	req := &UserDeleteRequest{}
	if err := h.parseJSONRequest(g, &req); err != nil {
		h.jsonNotOk(g, err.GetStatus(), err.GetCode(), err.GetMessage())
		return
	}

	svcRes := h.service.User.Delete(service.UserDeleteParam{
		Ctx:      ctx,
		ID:       req.ID,
		Password: req.Password,
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	h.jsonOk(g, nil)
}
