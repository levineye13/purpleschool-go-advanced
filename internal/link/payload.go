package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,http_url"`
}

type LinkUpdateRequest struct {
	LinkCreateRequest
}
