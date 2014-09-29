package exectypes

import "net/http"

type ExecRequest struct {
	Language           string
	SourceCode         string
	IsCompileOperation bool
	RespWriter         http.ResponseWriter
	MainPath           string
	CommandLineArgs    string
	Guid               string
}

type ExecResponse struct {
	Language    string
	Response    string
	RequestCode string
	RespWriter  http.ResponseWriter
	FolderPath  string
}

//var Response ExecResponse
//var Request ExecRequest
