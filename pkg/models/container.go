package models

import "github.com/uptrace/bun"

type Container struct {
	bun.BaseModel `bun:"table:containers,alias:ct"`

	ID    int64  `json:"id" bun:"id,pk"`
	Name  string `json:"name" bun:"name"`
	Image string `json:"image" bun:"image"`
}
