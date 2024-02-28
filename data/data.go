package data

// Request is the request for database as a service
type Request struct {
	Engine    string `body:"engine,required" json:"engine"`
	DSN       string `body:"dsn,required" json:"dsn"`
	Statement string `body:"statement,required" json:"statement"`
}

// Response is the response for database as a service
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}
