package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUpLoadReq struct {
	g.Meta `path:"/file" method:"post" mime:"multipart/form-data" tags:"工具" dc:"文件上传"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传的文件"`
}

type FileUploadRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"图片地址"`
}
