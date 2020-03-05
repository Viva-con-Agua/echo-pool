package pool

type (
	ResponseMessage struct {
		Message string `json:"message"`
	}
	ResponseMessageId struct {
		Message string `json:"message"`
		Uuid   string `json:"uuid"`
	}
)

func InternelServerError() interface{} {
	response := new(ResponseMessage)
	response.Message = "Internel server error, please check logs"
	return response
}
func Conflict() interface{} {
	response := new(ResponseMessage)
	response.Message = "Models already exists"
	return response
}

func Created() interface{} {
	response := new(ResponseMessage)
	response.Message = "Successful created"

	return response
}

func Unauthorized() interface{} {
	response := new(ResponseMessage)
	response.Message = "Not authenticated"
	return response
}
func NoContent(uuid string) interface{} {
	response := new(ResponseMessageId)
	response.Message = "Not found"
	response.Uuid = uuid
	return response
}
func Updated(uuid string) interface{} {
	response := new(ResponseMessageId)
	response.Message = "Successful updated"
	response.Uuid = uuid
	return response
}

func Deleted(uuid string) interface{} {
	response := new(ResponseMessageId)
	response.Message = "Successful deleted"
	response.Uuid = uuid
	return response
}
