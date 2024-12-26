package config

var API_KEY_MICROSERVICE = "6bf824612e1b42b977b18be59166ba9d309c44ef0dae080d341ca913ffda9f89"

var INPUT_PATH_IMAGEN = "./public/input"
var OUTPUT_PATH_IMAGEN = "./public/output"

// Response represents standar JSON reponse.
type Response struct {
	Ok    interface{} `json:"ok,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
