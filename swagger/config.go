package swagger

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/ls6-events/astra"
	"github.com/ls6-events/astra/inputs"
	"github.com/ls6-events/astra/outputs"
)

func Use(Server *gin.Engine) {
	gen := astra.New(
		inputs.WithGinInput(Server),
		outputs.WithOpenAPIOutput("./docs/swagger.json"),
		astra.WithPathDenyListRegex(regexp.MustCompile("^/docs.*")),
	)
	config := &astra.Config{
		Title:   "Example API",
		Version: "1.0.0",
		Host:    "localhost",
		Port:    8080,
	}
	gen.SetConfig(config)
	err := gen.Parse()
	if err != nil {
		panic(err)
	}
}
