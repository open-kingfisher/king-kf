package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/open-kingfisher/king-kf/resource"
	"github.com/open-kingfisher/king-utils/common/handle"
)

func GetTemplate(c *gin.Context) {
	r := resource.TemplateResource{Params: handle.GenerateCommonParams(c, nil)}
	responseData := HandleGet(&r)
	c.JSON(responseData.Code, responseData)
}

func ListTemplate(c *gin.Context) {
	r := resource.TemplateResource{Params: handle.GenerateCommonParams(c, nil)}
	responseData := HandleList(&r)
	c.JSON(responseData.Code, responseData)
}

func CreateTemplate(c *gin.Context) {
	r := resource.TemplateResource{Params: handle.GenerateCommonParams(c, nil)}
	responseData := HandleCreate(&r, c)
	c.JSON(responseData.Code, responseData)
}

func DeleteTemplate(c *gin.Context) {
	r := resource.TemplateResource{Params: handle.GenerateCommonParams(c, nil)}
	responseData := HandleDelete(&r)
	c.JSON(responseData.Code, responseData)
}

func UpdateTemplate(c *gin.Context) {
	r := resource.TemplateResource{Params: handle.GenerateCommonParams(c, nil)}
	responseData := HandleUpdate(&r, c)
	c.JSON(responseData.Code, responseData)
}
