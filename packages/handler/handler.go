package handler

import "github.com/niconiahi/one-piece-on-stone/packages/html"

type Data[T any] struct {
	Head   html.Head
	Loader T
}
