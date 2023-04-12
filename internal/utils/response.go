package utils

type Reponse struct {
	Status  bool        `json:"status"`
	Body    interface{} `json:"body"`
	Errors  interface{} `json:"errors"`
	Message string      `json:"message"`
}


