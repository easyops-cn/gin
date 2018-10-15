package bindingmp

import (
	"github.com/gin-gonic/gin/binding"
)

var (
	MsgPack = msgpackBinding{}
)

func init() {
	binding.SetDefaultBinding(binding.MIMEMSGPACK, MsgPack)
	binding.SetDefaultBinding(binding.MIMEMSGPACK2, MsgPack)
}
