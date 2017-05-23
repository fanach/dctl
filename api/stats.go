package api

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/fanach/dctl/entity"
	"github.com/fanach/dctl/service"
)

// GetDockerStats get usage infomation of a container
func GetDockerStats(ctx *iris.Context) {
	resp := &entity.RespGetStats{}

	id := ctx.Param("id")
	if len(id) == 0 {
		resp.ErrNo = iris.StatusBadRequest
		resp.Errmsg = "invalid param id"
		ctx.JSON(iris.StatusBadRequest, resp)
		return
	}

	stats, err := service.DockerStats(id)
	if err != nil {
		resp.ErrNo = iris.StatusInternalServerError
		resp.Errmsg = err.Error()
		ctx.JSON(iris.StatusInternalServerError, resp)
		return
	}

	if stats != nil {
		resp.Stats = *stats
		resp.Success = true
	}
	ctx.JSON(iris.StatusOK, resp)
	return
}
