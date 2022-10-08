package rest

import (
	"github.com/dalikewara/ayapingping-go-crud/src/service"
	"github.com/gin-gonic/gin"
)

// ProfileUpdateImage handles gin-gonic REST route to update profile image.
func (h *ginHandler) ProfileUpdateImage(g *gin.Context) {
	ctx := h.getRequestContext(g)
	req := &ProfileUpdateImageRequest{}
	if err := h.parseJSONRequest(g, &req); err != nil {
		h.jsonNotOk(g, err.GetStatus(), err.GetCode(), err.GetMessage())
		return
	}

	svcRes := h.service.Profile.UpdateImage(service.ProfileUpdateImageParam{
		Ctx:    ctx,
		UserID: req.UserID,
		Image:  req.Image,
	})
	if svcRes.Error != nil {
		h.jsonNotOk(g, svcRes.Error.GetStatus(), svcRes.Error.GetCode(), svcRes.Error.GetMessage())
		return
	}

	h.jsonOk(g, nil)
}
