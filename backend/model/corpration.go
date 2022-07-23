package model

type CorprationResponse struct {
	PublicID string `json:"public_id"`
	Name     string `json:"name"`
}

func (c *Corpration) ToModelResponse() *CorprationResponse {
	return &CorprationResponse{
		PublicID: c.PublicID,
		Name:     c.Name,
	}
}
