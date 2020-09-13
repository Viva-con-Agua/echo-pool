package api

type (
	ResponseMessage struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data,omitempty"`
	}
)

func RespCustom(m string, k *string, i interface{}) interface{} {
	r := new(ResponseMessage)
	r.Message = m
	if k != nil && i != nil {
		data := make(map[string]interface{})
		data[*k] = i
	}
	return r
}
func RespNoContent(k string, i interface{}) interface{} {
	data := make(map[string]interface{})
	data[k] = i
	response := new(ResponseMessage)
	response.Message = "Not found"
	response.Data = data
	return response
}

func RespConflict(k string, i interface{}) interface{} {
	data := make(map[string]interface{})
	data[k] = i
	response := new(ResponseMessage)
	response.Message = "Model already exists"
	response.Data = data
	return response
}

func RespInternelServerError() interface{} {
	response := new(ResponseMessage)
	response.Message = "Internel server error, please check logs"
	return response
}

func RespCreated() interface{} {
	response := new(ResponseMessage)
	response.Message = "Successful created"
	return response
}
