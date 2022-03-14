// Copyright (C) Subhajit DasGupta 2021

package common

import "time"

// Resource is a base type which must be embedded in persisted structs.
type Resource struct {
	ID       string
	Name     string
	Created  Time
	Modified Time
}

func (r *Resource) GetID() string {
	return r.ID
}

func (r *Resource) SetID(value string) {
	r.ID = value
}

func (r *Resource) GetName() string {
	return r.Name
}

func (r *Resource) SetName(value string) {
	r.Name = value
}

func (r *Resource) GetCreated() time.Time {
	return r.Created.Time
}

func (r *Resource) SetCreated(value time.Time) {
	r.Created.Time = value
}

func (r *Resource) GetModified() time.Time {
	return r.Modified.Time
}

func (r *Resource) SetModified(value time.Time) {
	r.Modified.Time = value
}
