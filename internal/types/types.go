package types

import "time"

type Item struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateItemReq struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateItemReq struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ListItemsResp struct {
	Items []Item `json:"items"`
}
