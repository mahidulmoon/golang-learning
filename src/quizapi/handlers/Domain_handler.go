package handlers

import (
	"domaniGoupdate/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDomains() gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := models.GetDomains()
		if err == nil {
			c.JSON(http.StatusOK, items)
			return
		}
		c.String(http.StatusNotFound, "could not fetch domains")
	}
}

func GetSubdomainByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type param struct {
			ID string `uri:"id" binding:"required"`
		}
		var uriParam param
		err := c.ShouldBindUri(&uriParam)
		items, err := models.GetSubdomainsByDomainID(uriParam.ID)
		if err == nil {
			c.JSON(http.StatusOK, items)
			return
		}
		c.String(http.StatusNotFound, "could not fetch subdomains")
	}

}

func GetTopicsByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		type param struct {
			ID string `uri:"id" binding:"required"`
		}
		var uriParam param
		err := c.ShouldBindUri(&uriParam)
		items, err := models.GetTopicsBySubdomainID(uriParam.ID)
		if err == nil {
			c.JSON(http.StatusOK, items)
			return
		}
		c.String(http.StatusNotFound, "could not fetch topics")
	}

}

func AddDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		var group models.Domain
		err := c.ShouldBindJSON(&group)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unmarshal failed",
				"error":   err,
			})
		} else {
			id, err := group.NewDomain()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
					"id":    id,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"group": id})
			}
		}

	}

}

func AddSubdomain() gin.HandlerFunc {
	return func(c *gin.Context) {

		var group models.Subdomain
		err := c.ShouldBindJSON(&group)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unmarshal failed",
				"error":   err,
			})
		} else {
			id, err := group.NewSubdomain()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
					"id":    id,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"group": id})
			}
		}

	}

}

func AddTopic() gin.HandlerFunc {
	return func(c *gin.Context) {
		var group models.Topic
		err := c.ShouldBindJSON(&group)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": "unmarshal failed",
				"error":   err,
			})
		} else {
			id, err := group.NewTopic()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
					"id":    id,
				})
			} else {
				c.JSON(http.StatusCreated, gin.H{"group": id})
			}
		}
	}

}
