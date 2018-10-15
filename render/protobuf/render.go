package renderpb

import "github.com/gin-gonic/gin/render"

var (
	_ render.Render     = ProtoBuf{}
)
