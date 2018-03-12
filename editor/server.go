package editor

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mdeheij/lahmacun-cms/model"
)

var em *model.EntityManager

func Start() {
	em := model.NewEntityManager()
	em.AutoMigrate()
	contentStore := em.NewContentStore()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		contents, errFindAll := contentStore.FindAll()
		if errFindAll != nil {
			log.Println(errFindAll)
			c.String(http.StatusInternalServerError, "Could not fetch content")
			return
		}

		c.HTML(http.StatusOK, "list.html", gin.H{
			"contents": contents,
		})
	})

	r.GET("/edit/:id", func(c *gin.Context) {
		content, errFindOne := contentStore.FindOneBySlug("test")
		if errFindOne != nil {
			log.Println(errFindOne)
			c.String(http.StatusInternalServerError, "Could not fetch content list")
			return
		}

		c.HTML(http.StatusOK, "edit.html", gin.H{
			"content": content,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
