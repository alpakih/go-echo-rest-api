package response

import (
	"go-echo-rest-api/app/models"
	"net/http"
)

func Single(d Data) Response {
	meta := MetaComposer(http.StatusOK, "OK")
	return Response{Meta: meta, Data: d.Map()}
}

func List(l interface{}) Response {
	meta := MetaComposer(http.StatusOK, "OK")
	return Response{Meta: meta, Data: l}
}

func BadRequest() Response {
	meta := MetaComposer(http.StatusBadRequest, "Bad Request")
	return Response{Meta: meta}
}

func InternalServerError() Response {
	meta := MetaComposer(http.StatusInternalServerError, "Internal Server Error")
	return Response{Meta: meta}
}

func NotFound() Response {
	meta := MetaComposer(http.StatusNotFound, "Data Not Found")
	return Response{Meta: meta, Data: nil}
}

func Success(message string) Response {
	meta := MetaComposer(http.StatusOK, message)
	return Response{Meta: meta}
}

func ValidationError(d interface{}) Response {
	meta := MetaComposer(http.StatusBadRequest, "Validation Error")
	return Response{Meta: meta, Data: d}
}

func CustomMessage(httpStatus int, message string) (response Response) {
	meta := MetaComposer(httpStatus, message)
	return Response{Meta: meta}
}

func TokenResponse(customer *models.Customer) CustomerResponse {
	meta := MetaComposer(http.StatusOK, "Login Success")
	cus := CustomerComposer(customer)
	return CustomerResponse{
		Meta: meta, Customer: cus,
	}
}
