// Copyright (C) Subhajit DasGupta 2021

package common

// Resource is a base type which must be embedded in persisted structs.
type Resource struct {
	ID       string
	Name     string
	Created  Time
	Modified Time
}
