package handler

import (
	"github.com/niconiahi/gig.dance/packages/html"
)

type Data[T any] struct {
	Head   html.Head
	Loader T
}
