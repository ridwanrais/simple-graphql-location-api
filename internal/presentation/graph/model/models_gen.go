// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
)

type CityConnection struct {
	Edges    []*CityEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type CityEdge struct {
	Node   *domain.City `json:"node"`
	Cursor string       `json:"cursor"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	EndCursor       *string `json:"endCursor,omitempty"`
}
