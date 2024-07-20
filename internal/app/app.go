// основная бизнес логика приложения

package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type BuildForm struct {
	GitURL              string   `form:"git_url"`
	GitBranch           string   `form:"git_branch"`
	SonarProjectKey     string   `form:"sonar_project_key"`
	BuildProfile        string   `form:"build_profile"`
	BuildSubdir         string   `form:"build_subdir"`
	Steps               []string `form:"steps"`
	NexusArtifactFormat string   `form:"nexus_artifact_format"`
}

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
	var form BuildForm
	if err := contex.ShouldBind(&form); err == nil {
		field := contex.Query("field")
		step := contex.Query("step")
		var stepStatus string

		if step != "" {
			// Если шаг, определим его статус(включен или выключен)
			stepStatus = "disabled"
			for _, s := range form.Steps {
				if s == step {
					stepStatus = "enabled"
					break
				}
			}
		}

		contex.HTML(http.StatusOK, "build_result.html", gin.H{
			"field":					field,
			"step":						step,
			"step_status":				stepStatus,
			"git_url":					form.GitURL,
			"git_branch":				form.GitBranch,
			"sonar_project_key":		form.SonarProjectKey,
			"build_profile":			form.BuildProfile,
			"build_subdir":				form.BuildSubdir,
			"steps":					form.Steps,
			"nexus_artifact_format":	form.NexusArtifactFormat,
		})
	} else {
		contex.HTML(http.StatusBadRequest, "build_result.html", gin.H{"error":err.Error()})
	}
}

func handlerSubmitBuild(context *gin.Context) {
	var form BuildForm
	if err := context.ShouldBind(&form); err == nil {
		context.JSON(http.StatusOK,	gin.H{
			"status":	"success",
			"data":		form,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

func (a *App) InitializeRoutes() {
	a.Router.Static("/static", "./web/static")
	a.Router.LoadHTMLGlob("./web/templates/*")

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
