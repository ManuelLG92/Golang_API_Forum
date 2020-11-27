package models

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
	response.writer.Header().Set("Content-Type", "application/json")

	//sesionCookie := SetCookie(response.writer)
	//response.writer.Header().Set("Set-Cookie", "true")
	//SetCookie(response.writer)
	//Response{ Cookie: SetCookie(response.writer)}
	//response.writer.WriteHeader(response.Status)
	//output, _ :=json.Marshal(&response)
	_= json.NewEncoder(response.writer).Encode(response)//serialziamos el objeto response
	//fmt.Fprintf(response.writer, string(output)) //se pinta por la stdout standard
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

func SendNoContent (w http.ResponseWriter)   {
	response := CreateDefaultResponse(w)
	response.NoContent()
	response.Send()
}
func (response *Response) NoContent()  {
	response.Status = http.StatusNoContent
	response.Message = "There is no content for this request"

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





