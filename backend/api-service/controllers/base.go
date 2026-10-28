package controllers

import (
	"encoding/json"
	"io"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
}

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ListData struct {
	List       interface{}    `json:"list"`
	Pagination PaginationData `json:"pagination"`
}

type PaginationData struct {
	Total    interface{} `json:"total"`
	PageNum  interface{} `json:"page_num"`
	PageSize interface{} `json:"page_size"`
}

func (c *BaseController) JSONObject(data interface{}) {
	c.Data["json"] = HttpResponse{Code: 200, Message: "OK", Data: data}
	_ = c.ServeJSON()
}

func (c *BaseController) JSONList(data ListData) {
	c.Data["json"] = HttpResponse{Code: 200, Message: "OK", Data: data}
	_ = c.ServeJSON()
}

func (c *BaseController) JSONErr(code int, msg string) {
	c.Data["json"] = HttpResponse{Code: code, Message: msg}
	_ = c.ServeJSON()
}

func (c *BaseController) GetGetParams() map[string]string {
	params := make(map[string]string)
	values, err := c.Input()
	if err != nil {
		return params
	}
	for key, vals := range values {
		if len(vals) > 0 {
			params[key] = vals[len(vals)-1]
		}
	}

	if params["page_num"] == "" {
		params["page_num"] = "1"
	}
	if params["page_size"] == "" {
		params["page_size"] = "10"
	}
	return params
}

// BindJSON binds JSON request body to the given object.
// It first tries ctx.Input.RequestBody (cached by Beego's CopyBody),
// and falls back to reading directly from ctx.Request.Body if the cache is empty.
func (c *BaseController) BindJSON(obj interface{}) error {
	if len(c.Ctx.Input.RequestBody) > 0 {
		return json.Unmarshal(c.Ctx.Input.RequestBody, obj)
	}
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, obj)
}
