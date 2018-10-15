package bindingpb

import "github.com/gin-gonic/gin/binding"

var (
	ProtoBuf = protobufBinding{}
)

func init() {
	binding.SetDefaultBinding(binding.MIMEPROTOBUF, ProtoBuf)
}
