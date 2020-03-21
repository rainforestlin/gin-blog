package v1

import (
	"blogWithGin/models"
	"blogWithGin/pkg/errCode"
	"blogWithGin/pkg/setting"
	"blogWithGin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

//获取文章标签
//@Summary 获取文章标签
//@security JWT
//@Produce json
//@Param name query string true "Name"
//@param state query int false "State"
//@success 200 {object} models.Tag  "{"code":200,"data":{"list":[],"total":20},"msg":"ok"}"
//@router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

	}
	code := errCode.SUCCESS
	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errCode.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
//@Summary 新增文章标签
//@security JWT
//@produce json
//@param name query string true "Name"
//@param state query int false "State"
//@param created_by query string true "CreatedBy" "string valid" maxlength(100) minlength(1)
//@success 200 {object} models.Tag  "{"code":200,"data":{},"msg":"ok"}"
//@router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100个字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	code := errCode.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = errCode.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = errCode.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errCode.GetMsg(code),
		"data": make(map[string]string),
	})
}

//Modify tag
//@Summary 修改文章标签
//@security JWT
//@produce json
//@param id path int true "ID"
//@param name query string true "Name"
//@param modified_by query string true "ModifiedBy"
//@success 200 {object} models.Tag  "{"code":200,"data":{},"msg":"ok"}"
//@router /api/v1/tags/{id} [put]
func ModifyTag(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(name, 100, "name").Message("名称不得超过100个字符")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := errCode.INVALID_PARAMS
	if !valid.HasErrors() {
		code = errCode.SUCCESS
		if models.ExistTagById(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data[name] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.ModifyTag(id, data)
		} else {
			code = errCode.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errCode.GetMsg(code),
		"data": make(map[string]interface{}),
	})

}

//delete tag
//@summary 删除文章标签
//@security JWT
//@produce json
//@Param id path int true "ID" "int valid" mininum(1)
//@success 200 {string} models.Tag  "{"code":200,"data":{},"msg":"ok"}"
//@router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := errCode.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = errCode.SUCCESS
		if models.ExistTagById(id) {
			models.DeleteTag(id)
		} else {
			code = errCode.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errCode.GetMsg(code),
		"data": make(map[string]string),
	})
}
