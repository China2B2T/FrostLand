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

	switch premium {
	case "1":
		code, message, uu := MCreateUser(uid, true)
		result = gin.H{
			"code":    code,
			"message": message,
			"uuid":    uu,
		}
		break
	case "0":
		code, message, uu := MCreateUser(uid, false)
		result = gin.H{
			"code":    code,
			"message": message,
			"uuid":    uu,
		}
		break
	default:
		result = gin.H{
			"code":    -1,
			"message": "Invalid data",
		}
	}

	c.JSON(http.StatusOK, result)
}

func QueryUser(c *gin.Context) {
	result := gin.H{
		"code":    -10,
		"message": "Unhandled request",
	}

	id := c.Param("id")

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
