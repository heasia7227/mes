package system

import (
	"api-service/controllers"
	systemDto "api-service/dtos/system"
	"api-service/services/system"
	"common/utils/strings"
)

type PostController struct {
	controllers.BaseController
}

var (
	postService = &system.PostService{}
)

// GetPostList 获取岗位列表
func (c *PostController) GetPostList() {
	params := c.GetGetParams()

	posts, total := postService.GetPostList(params)

	c.JSONList(controllers.ListData{
		List: posts,
		Pagination: controllers.PaginationData{
			Total:    total,
			PageNum:  params["page_num"],
			PageSize: params["page_size"],
		},
	})
}

// CreatePost 创建岗位
func (c *PostController) CreatePost() {
	var post systemDto.PostDto
	if err := c.BindJSON(&post); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	post.ID = strings.GenerateID()
	post.Status = "enable"

	// 设置创建人和更新人
	login_name := c.Ctx.Input.Header("login_name")
	post.CreateBy = login_name
	post.UpdateBy = login_name

	// 保存到数据库
	err := postService.CreatePost(post)
	if err != nil {
		c.JSONErr(500, "创建岗位失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
		"id":      post.ID,
	})
}

// UpdatePost 更新岗位
func (c *PostController) UpdatePost() {
	id := c.Ctx.Input.Param(":id")

	// 检查岗位是否存在
	existingPost, err := postService.GetPostDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}

	var post systemDto.PostDto
	if err := c.BindJSON(&post); err != nil {
		c.JSONErr(400, err.Error())
		return
	}

	// 复制必要的字段, 保留原有ID和创建信息
	existingPost.PostName = post.PostName
	existingPost.PostCode = post.PostCode
	existingPost.Sort = post.Sort
	existingPost.Remark = post.Remark

	// 设置更新人和更新时间
	login_name := c.Ctx.Input.Header("login_name")
	existingPost.UpdateBy = login_name

	// 更新到数据库
	err = postService.UpdatePost(existingPost)
	if err != nil {
		c.JSONErr(500, "更新岗位失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// DeletePost 删除岗位
func (c *PostController) DeletePost() {
	id := c.Ctx.Input.Param(":id")

	// 检查岗位是否存在
	existingPost, err := postService.GetPostDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}
	if existingPost.Status == "enable" {
		c.JSONErr(400, "已启用的岗位不能删除")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingPost.UpdateBy = login_name
	existingPost.DeleteFlag = 1

	// 执行软删除
	err = postService.UpdatePost(existingPost)
	if err != nil {
		c.JSONErr(500, "删除岗位失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}

// GetPostDetail 获取岗位详情
func (c *PostController) GetPostDetail() {
	id := c.Ctx.Input.Param(":id")

	post, err := postService.GetPostDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}

	c.JSONObject(post)
}

// UpdatePostStatus 更新岗位状态
func (c *PostController) UpdatePostStatus() {
	id := c.Ctx.Input.Param(":id")

	// 检查岗位是否存在
	existingPost, err := postService.GetPostDetail(id)
	if err != nil {
		c.JSONErr(404, "岗位不存在")
		return
	}

	// 状态
	status := c.Ctx.Input.Param(":status")
	if status == "" {
		c.JSONErr(400, "状态不能为空")
		return
	} else if status != "enable" && status != "disable" {
		c.JSONErr(400, "状态必须为 enable 或 disable")
		return
	}

	login_name := c.Ctx.Input.Header("login_name")
	existingPost.UpdateBy = login_name
	existingPost.Status = status

	// 更新状态
	err = postService.UpdatePost(existingPost)
	if err != nil {
		c.JSONErr(500, "更新岗位状态失败: "+err.Error())
		return
	}

	c.JSONObject(map[string]interface{}{
		"success": true,
	})
}
