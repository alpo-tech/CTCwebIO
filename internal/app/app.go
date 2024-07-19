// основная бизнес логика приложения

package app

import "github.com/gin-gonic/gin"

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

func (a *App) initializeRoutes() {

}

func (a *App) Run(addr string) {
	a.Router.Run(addr)
}
