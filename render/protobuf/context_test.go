package renderpb


import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	testdata "github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

// TestContextRenderProtoBuf tests that the response is serialized as ProtoBuf
// and Content-Type is set to application/x-protobuf
// and we just use the example protobuf to check if the response is correct
func TestContextRenderProtoBuf(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	reps := []int64{int64(1), int64(2)}
	label := "test"
	data := &testdata.Test{
		Label: &label,
		Reps:  reps,
	}

	c.Render(http.StatusCreated, ProtoBuf{Data: data})

	protoData, err := proto.Marshal(data)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, string(protoData), w.Body.String())
	assert.Equal(t, "application/x-protobuf", w.Header().Get("Content-Type"))
}

