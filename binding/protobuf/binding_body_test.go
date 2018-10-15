package bindingpb

import (
	"io/ioutil"
	"testing"

	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)


func TestBindingBodyProto(t *testing.T) {
	test := protoexample.Test{
		Label: proto.String("FOO"),
	}
	data, _ := proto.Marshal(&test)
	req := requestWithBody("POST", "/", string(data))
	form := protoexample.Test{}
	body, _ := ioutil.ReadAll(req.Body)
	assert.NoError(t, ProtoBuf.BindBody(body, &form))
	assert.Equal(t, test, form)
}

