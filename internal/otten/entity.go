package otten

type Result struct {
	Status Mapper `json:"status"`
	Data   Data   `json:"data"`
}

type Mapper struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	ReceivedBy string    `json:"receivedBy"`
	History    []History `json:"histories"`
}

type History struct {
	Description string    `json:"description"`
	CreatedAt   string    `json:"createdAt"`
	Formatted   Formatted `json:"formatted"`
}

type Formatted struct {
	CreatedAt string `json:"createdAt"`
}
