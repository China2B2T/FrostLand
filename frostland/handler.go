package frostland

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	result := gin.H{
		"code":    0,
		"message": "pong",
	}

	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	result := gin.H{
		"code":    -10,
		"message": "Unhandled request",
	}

	uid := c.PostForm("uid")
	premium := c.PostForm("premium")

	if uid == "" || premium == "" {
		result = gin.H{
			"code":    -1,
			"message": "Invalid data",
		}

		c.JSON(http.StatusOK, result)
		return
	}

	code, message, uu := MCreateUser(uid, (premium == "1"))
	result = gin.H{
		"code":    code,
		"message": message,
		"uuid":    uu,
	}

	c.JSON(http.StatusOK, result)
}

func QueryUser(c *gin.Context) {
	result := gin.H{
		"code":    -10,
		"message": "Unhandled request",
	}

	id := c.Param("id")

	if id == "" {
		result = gin.H{
			"code":    -1,
			"message": "Invalid data",
		}

		c.JSON(http.StatusOK, result)
		return
	}

	uuid, err := MQueryUser(id)
	if err != nil {
		result = gin.H{
			"code":    5000,
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"code":    0,
			"message": "OK",
			"uuid":    uuid,
		}
	}

	c.JSON(http.StatusOK, result)
}

func QueryUUID(c *gin.Context) {

}

func ImportUser(c *gin.Context) {
	result := gin.H{
		"code":    -10,
		"message": "Unhandled request",
	}

	uid := c.PostForm("uid")
	uuid := c.PostForm("uuid")
	premium := c.PostForm("premium")

	if uid == "" || uuid == "" {
		result = gin.H{
			"code":    -1,
			"message": "Invalid data",
		}

		c.JSON(http.StatusOK, result)
		return
	}

	code, message := MImportUser(uid, (premium == "1"), uuid)
	result = gin.H{
		"code":    code,
		"message": message,
	}

	c.JSON(http.StatusOK, result)
}
