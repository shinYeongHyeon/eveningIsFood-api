package foodCardControllerCommandDto

// CreateRequest : Request for create foodCard
type CreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// FoodCard : FoodCard object for protocol
type FoodCard struct {
	Uuid    string `json:"string"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

// CreateResponse : Response for Create
type CreateResponse struct {
	Code     string   `json:"code" example:"SUCCESS"`
	FoodCard FoodCard `json:"foodCard"`
}

// CreateResponseError : Response for Create
type CreateResponseError struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

// CreateResponseErrorForBadRequest : Error value of 400
var CreateResponseErrorForBadRequest = CreateResponseError{
	Code:         "BAD_REQUEST",
	ErrorMessage: "All fields are required",
}
