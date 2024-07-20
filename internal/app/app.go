// основная бизнес логика приложения

package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	router := gin.Default()

	app := &App{
		Router: router,
	}

	return app
}

func handlerBase(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

func handlerBuild(context *gin.Context) {
	context.HTML(http.StatusOK, "build.html", nil)
}

func handlerRelease(context *gin.Context) {
	context.HTML(http.StatusOK, "release.html", nil)
}

func handlerDeploy(context *gin.Context) {
	context.HTML(http.StatusOK, "deploy.html", nil)
}

func handlerUpdateBuild(contex *gin.Context) {

}

func handlerSubmitBuild(context *gin.Context) {

}

func (a *App) initializeRoutes() {
	a.Router.Static("/static", "../../web/static")
	a.Router.LoadHTMLGlob("../../web/templates/*")

	a.Router.GET("", handlerBase)
	a.Router.GET("/build", handlerBuild)
	a.Router.GET("/release", handlerRelease)
	a.Router.GET("/deploy", handlerDeploy)

	a.Router.POST("/update_build", handlerUpdateBuild)
	a.Router.POST("/submit_build", handlerSubmitBuild)

}

func (a *App) Run(addr string) {
	a.Router.Run(addr)
}
