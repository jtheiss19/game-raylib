package ecs

import (
	"errors"

	"github.com/google/uuid"
)

type ID string

func (id ID) String() string {
	return string(id)
}

type Component interface {
	GetComponentID() (ID, error)
	setComponentID(id ID)
	generateComponentID() ID
}

type BaseComponent struct {
	EntityID ID
}

func (bc *BaseComponent) GetComponentID() (ID, error) {
	if bc.EntityID == ID(uuid.Nil.String()) {
		return ID(uuid.Nil.String()), errors.New(" does not have component ID")
	}
	return bc.EntityID, nil
}

func (bc *BaseComponent) setComponentID(id ID) {
	bc.EntityID = id
}

func (bc *BaseComponent) generateComponentID() ID {
	id := ID(uuid.New().String())
	bc.EntityID = id
	return id
}
