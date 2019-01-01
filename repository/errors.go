package repository

import "fmt"

var (
	// ErrChatAlreadySubskribed error that returns when try to subskribe chat on aready subskibed channel
	ErrChatAlreadySubskribed = fmt.Errorf("chat already subskibed on this channel")
)
