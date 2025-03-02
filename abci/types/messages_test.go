package types

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"

	cmtproto "github.com/ben2077/cometbft/proto/tendermint/types"
)

func TestMarshalJSON(t *testing.T) {
	b, err := json.Marshal(&ExecTxResult{Code: 1})
	assert.NoError(t, err)
	// include empty fields.
	assert.True(t, strings.Contains(string(b), "code"))
	r1 := ResponseCheckTx{
		Code:      1,
		Data:      []byte("hello"),
		GasWanted: 43,
		Events: []Event{
			{
				Type: "testEvent",
				Attributes: []EventAttribute{
					{Key: "pho", Value: "bo"},
				},
			},
		},
	}
	b, err = json.Marshal(&r1)
	assert.Nil(t, err)

	var r2 ResponseCheckTx
	err = json.Unmarshal(b, &r2)
	assert.Nil(t, err)
	assert.Equal(t, r1, r2)
}

func TestWriteReadMessageSimple(t *testing.T) {
	cases := []proto.Message{
		&RequestEcho{
			Message: "Hello",
		},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := WriteMessage(c, buf)
		assert.Nil(t, err)

		msg := new(RequestEcho)
		err = ReadMessage(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}

func TestWriteReadMessage(t *testing.T) {
	cases := []proto.Message{
		&cmtproto.Header{
			Height:  4,
			ChainID: "test",
		},
		// TODO: add the rest
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := WriteMessage(c, buf)
		assert.Nil(t, err)

		msg := new(cmtproto.Header)
		err = ReadMessage(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}

func TestWriteReadMessage2(t *testing.T) {
	phrase := "hello-world"
	cases := []proto.Message{
		&ResponseCheckTx{
			Data:      []byte(phrase),
			Log:       phrase,
			GasWanted: 10,
			Events: []Event{
				{
					Type: "testEvent",
					Attributes: []EventAttribute{
						{Key: "abc", Value: "def"},
					},
				},
			},
		},
		// TODO: add the rest
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		err := WriteMessage(c, buf)
		assert.Nil(t, err)

		msg := new(ResponseCheckTx)
		err = ReadMessage(buf, msg)
		assert.Nil(t, err)

		assert.True(t, proto.Equal(c, msg))
	}
}
