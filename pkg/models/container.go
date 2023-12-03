package models

import (
	"github.com/uptrace/bun"
)

type ContainerPackage struct {
	bun.BaseModel `bun:"table:packages,alias:pkg"`

	ID   int64  `json:"id" bun:"id,pk,autoincrement"`
	Name string `json:"name" bun:"name"`
	Tag  string `json:"tag" bun:"tag"`
}
