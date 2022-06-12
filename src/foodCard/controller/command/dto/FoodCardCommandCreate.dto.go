package foodCardControllerCommandDto

// CreateRequest : Request for create foodCard
type CreateRequest struct {
	// 음식점 이름
	Name string `json:"name" example:"초원"`
	// 음식점 주소
	Address string `json:"address" example:"서울특별시 용산구 남영동 80-4"`
}

// FoodCard : FoodCard object for protocol
type FoodCard struct {
	// 음식점 고유 ID
	Uuid string `json:"uuid"`
	// 음식점 이름
	Name string `json:"name"`
	// 음식점 주소
	Address string `json:"address"`
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
