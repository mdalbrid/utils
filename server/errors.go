package server

// about JsonRpc protocol codes and error formats https://www.jsonrpc.org/specification

// general
var (
	JsonRpcParseError          = JsonRpcError{Code: -32700, Message: "Parse error"}
	JsonRpcInvalidRequestError = JsonRpcError{Code: -32600, Message: "Invalid Request"}
	JsonRpcMethodNotFoundError = JsonRpcError{Code: -32601, Message: "Method not found"}
	JsonRpcInvalidParamsError  = JsonRpcError{Code: -32602, Message: "Invalid params"}
	JsonRpcInternalError       = JsonRpcError{Code: -32603, Message: "Internal error"}
	JsonRpcServerError         = JsonRpcError{Code: -32000, Message: "Server error"}
)

// custom
var (
	JsonRpcProxyBodyEncodeError = JsonRpcError{Code: -32001, Message: "Server error"}
	JsonRpcProxyRequestError    = JsonRpcError{Code: -32002, Message: "Server error"}
	JsonRpcProxyResponseError   = JsonRpcError{Code: -32003, Message: "Server error"}
	JsonRpcProxyBodyDecodeError = JsonRpcError{Code: -32004, Message: "Server error"}
)
