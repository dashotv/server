package actions

import (
	"fmt"
	"github.com/dashotv/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func ReleaseIndex(c echo.Context) error {
//	releases := []*models.Release{}
//	err := models.DB.All(&releases)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, releases)
//}

func ReleaseRoutes(r *gin.RouterGroup) {
	g := r.Group("/releases")
	g.GET("/", ReleaseIndex)
}

func ReleaseIndex(c *gin.Context) {
	var releases []models.Release
	db := models.DB.Find(&releases)
	if db.Error != nil {
		fmt.Printf("error: %s\n", db.Error)
	}

	//for _, r := range releases {
	//	//fmt.Printf("r: %#v\n", r)
	//	//f := "2006-01-02 15:04:05"
	//	fmt.Printf("%d %10s %-60.60s\n", r.ID, r.CreatedAt.Format(time.RFC822Z), r.Raw)
	//}

	c.JSON(http.StatusOK, releases)
	return
}
