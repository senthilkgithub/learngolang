package exectypes

type ExecRequest struct {
	Language           string
	Code               string
	IsCompileOperation bool
	RequestCode        string
}

type ExecResponse struct {
	Language    string
	Response    string
	RequestCode string
}

var Response ExecResponse
var Request ExecRequest
