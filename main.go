package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

func main() {
	router := gin.Default()

	// Статические файлы (CSS, JS и т.д.)
	router.Static("/static", "./static")

	// Подключение шаблонов
	router.LoadHTMLGlob("templates/*")

	// Маршрут для главной страницы
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Маршруты для вкладок
	router.GET("/build", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "build.html", gin.H{
		//	"ShowSonarProjectKey": false,
		//})
		c.HTML(http.StatusOK, "build.html", nil)
	})
	router.GET("/deploy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "deploy.html", nil)
	})
	router.GET("/release", func(c *gin.Context) {
		c.HTML(http.StatusOK, "release.html", nil)
	})

	// Обработчик для формы build
	router.POST("/update_build", func(c *gin.Context) {
		var form BuildForm
		if err := c.ShouldBind(&form); err == nil {
			field := c.Query("field")
			step := c.Query("step")
			var stepStatus string
			//showSonarProjectKey := false

			if step != "" {
				// Если есть шаг, определим его статус (включен или выключен)
				stepStatus = "disabled"
				for _, s := range form.Steps {
					if s == step {
						stepStatus = "enabled"
						//if s == "run_sonarqube" {
						//	showSonarProjectKey = true
						//}
						break
					}
				}
			}

			c.HTML(http.StatusOK, "build_result.html", gin.H{
				"field":       field,
				"step":        step,
				"step_status": stepStatus,
				//"show_sonar_project_key": showSonarProjectKey,
				"git_url":               form.GitURL,
				"git_branch":            form.GitBranch,
				"sonar_project_key":     form.SonarProjectKey,
				"build_profile":         form.BuildProfile,
				"build_subdir":          form.BuildSubdir,
				"steps":                 form.Steps,
				"nexus_artifact_format": form.NexusArtifactFormat,
			})
		} else {
			c.HTML(http.StatusBadRequest, "build_result.html", gin.H{"error": err.Error()})
		}
	})

	// Запуск сервера на порту 8080
	router.Run(":8080")
}
