package server

import (
	"encoding/json"
	"github.com/mdalbrid/utils/logger"
	"net/http"
	"strings"
)

type JsonRpcRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Id      *uint64     `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}
type JsonRpcError struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type JsonRpcResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	Id      *uint64       `json:"id"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JsonRpcError `json:"error,omitemoty"`
}

type JsonRpcMethodHandler func(req *JsonRpcRequest, res *JsonRpcResponse) int

func (d JsonRpcResponse) ToBytes() []byte {
	j, err := json.Marshal(d)
	if err != nil {
		logger.Error(err)
	}
	return j
}

type JsonRpcHandler struct {
	urlPrefix string
	mux       *http.ServeMux
	methods   map[string]JsonRpcMethodHandler
}

func NewJsonRpcHandler(urlPrefix string) *JsonRpcHandler {
	mux := http.NewServeMux()
	jrh := &JsonRpcHandler{
		urlPrefix: urlPrefix,
		mux:       mux,
		methods:   map[string]JsonRpcMethodHandler{},
	}
	mux.HandleFunc(urlPrefix, jrh.RequestHandler)
	mux.HandleFunc(urlPrefix+"/", jrh.RequestHandler)
	mux.HandleFunc("/ping", PingHandler)
	mux.HandleFunc("/ping/", PingHandler)
	return jrh
}

// proxy to mux
func (jrh *JsonRpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jrh.mux.ServeHTTP(w, r)
}

func (jrh *JsonRpcHandler) MethodHandlerFunc(method string, handlerFunc JsonRpcMethodHandler) {
	jrh.methods[method] = handlerFunc
}

// json rpc methods router
func (jrh *JsonRpcHandler) RequestHandler(w http.ResponseWriter, r *http.Request) {
	res, status := jrh.ProceedRequest(w, r)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res.ToBytes())
}

func (jrh *JsonRpcHandler) ProceedRequest(w http.ResponseWriter, r *http.Request) (res *JsonRpcResponse, status int) {
	res = &JsonRpcResponse{
		JSONRPC: "2.0",
		Id:      nil,
		Error:   nil,
		Result:  nil,
	}

	if r.Method != http.MethodPost {
		res.Error = &JsonRpcMethodNotFoundError
		status = http.StatusNotFound
		return
	}

	req := &JsonRpcRequest{}
	if e := json.NewDecoder(r.Body).Decode(req); e != nil {
		logger.Error("[RequestHandler] decode request: ", e)
		status = http.StatusBadRequest
		res.Error = &JsonRpcParseError
		return
	}

	res.Id = req.Id
	handler, ok := jrh.methods[req.Method]

	if !ok {
		namespace := strings.Split(req.Method, ".")[0]
		handler, ok = jrh.methods[namespace+".*"]
	}

	if !ok {
		res.Error = &JsonRpcMethodNotFoundError
		status = http.StatusNotFound
		return
	}

	status = handler(req, res)
	return
}
