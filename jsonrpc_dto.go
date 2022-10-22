package gcurl

import "encoding/json"

type JsonRpcReqDto struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Id      interface{} `json:"id"` //字符串、int、unit、int8/16/32/64、uint8/16/32/64、null
}

func NewJsonRpcReqDto() JsonRpcReqDto {
	return JsonRpcReqDto{Jsonrpc: "2.0"}
}

type JsonRpcRespDto struct {
	Jsonrpc string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Error   *json.RawMessage `json:"error,omitempty"`
	Id      interface{}      `json:"id"`
}

//错误对象
type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
