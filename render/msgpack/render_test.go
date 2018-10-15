package rendermp

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ugorji/go/codec"
)

func TestRenderMsgPack(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{
		"foo": "bar",
	}

	(MsgPack{data}).WriteContentType(w)
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))

	err := (MsgPack{data}).Render(w)

	assert.NoError(t, err)

	h := new(codec.MsgpackHandle)
	assert.NotNil(t, h)
	buf := bytes.NewBuffer([]byte{})
	assert.NotNil(t, buf)
	err = codec.NewEncoder(buf, h).Encode(data)

	assert.NoError(t, err)
	assert.Equal(t, w.Body.String(), string(buf.Bytes()))
	assert.Equal(t, "application/msgpack; charset=utf-8", w.Header().Get("Content-Type"))
}

