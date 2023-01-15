package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Message string `json:"message"`
	//Cookie http.Cookie
	contentType string
	writer http.ResponseWriter
}

func CreateDefaultResponse (w http.ResponseWriter) Response {
	return Response{ writer: w, contentType: "application/json"}
}

func (response *Response) Send()  {
	response.writer.Header().Set("Access-Control-Allow-Origin", "*")
	response.writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.writer.Header().Set("Content-Type", "application/json")
	response.writer.WriteHeader(response.Status)
	_= json.NewEncoder(response.writer).Encode(response)
}

func SendCustom(w http.ResponseWriter, message string, status int)  {
	response := CreateDefaultResponse(w)
	response.Status = status
	response.Message = message
	response.Send()
}

func SendUnprocessableEntity(w http.ResponseWriter, message string)  {
	response := CreateDefaultResponse(w)
	response.Status = 422
	response.Message = message
	response.Send()
}

func SendNotFound(w http.ResponseWriter)  {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}
func (response *Response) NotFound()  {
	response.Status = http.StatusNotFound
	response.Message = "Resource doesn't found."
}

func SendData( w http.ResponseWriter, data interface{})  {
	response := CreateDefaultResponse(w)
	response.Status = http.StatusOK
	response.Message = "Request succesfully"
	response.Data = data
	response.Send()
}

func SendCreated( w http.ResponseWriter, data interface{})  {
	response := CreateDefaultResponse(w)
	response.Status = http.StatusCreated
	response.Data = data
	response.Message = "Created"
	response.Send()
}

func SendNoContent (w http.ResponseWriter)   {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}
func (response *Response) NoContent()  {
	response.Status = http.StatusNoContent
	response.Message = "There is no content for this request"

}

func SendInternalServerError (w http.ResponseWriter)   {
	response := CreateDefaultResponse(w)
	response.Status = http.StatusInternalServerError
	response.Message = "Internal server error"
	response.Send()
}

func SendNotAuth (w http.ResponseWriter)   {
	response := CreateDefaultResponse(w)
	response.NoAuth()
	response.Send()
}
func (response *Response) NoAuth()  {
	response.Status = http.StatusForbidden
	response.Message = "Unauthorized user"
}





