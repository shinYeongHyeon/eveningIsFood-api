package oAuthQueryControllerDto

// CreateResponse : Response
type CreateResponse struct {
	Code string `json:"code" example:"SUCCESS"`
	Url  string `json:"url"`
}

// CreateResponseError : Response
type CreateResponseError struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}
