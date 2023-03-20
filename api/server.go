package api

import (
	"fmt"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sam-val/text-to-img/openai"
)

func Run(path string) error {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Static("/sta", "./static")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", index)
	router.POST("/generate", generateImg)

	return router.Run(path)
}

func index(context *gin.Context) {
	// ses := sessions.Default(context)
	// err := ses.Get("error")
	// resultUrl := ses.Get("resultUrl")
	// ses.Delete("error")
	// ses.Delete("resultUrl")
	// ses.Save()
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Image from Text",
	})
}

type generateImgRequest struct {
	Text string `form:"prompt" json:"prompt" binding:"required"`
}

func generateImg(c *gin.Context) {
	// ses := sessions.Default(c)
	var req generateImgRequest
	err := c.ShouldBind(&req)
	if err != nil {
		fmt.Print("failed binding")
		c.JSON(http.StatusInternalServerError, errorRes(err))
		return
	}
	resultUrl, err := openai.ImgFromText(req.Text)
	if err != nil {
		fmt.Print("error from openAI")
		c.JSON(http.StatusInternalServerError, errorRes(err))
		return
	} else {
		fmt.Print(resultUrl)
	}
	c.JSON(http.StatusOK, successRes(resultUrl))

}

func errorRes(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

func successRes(obj interface{}) gin.H {
	return gin.H{
		"data": obj,
	}
}
