package api

import (
	"net/http"

	"gopkg.in/kataras/iris.v6"
	"github.com/fanach/dctl/entity"
	"github.com/fanach/dctl/service"
)

// UnpauseContainer unpauses a docker container
func UnpauseContainer(ctx *iris.Context) {
	resp := &entity.Resp{}

	id := ctx.Param("id")
	if len(id) == 0 {
		resp.ErrNo = iris.StatusBadRequest
		resp.Errmsg = "invalid param id"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	err := service.DockerUnpause(id)
	if err != nil {
		resp.ErrNo = http.StatusInternalServerError
		resp.Errmsg = err.Error()
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}
	resp.Success = true
	ctx.JSON(iris.StatusOK, resp)
	return
}
