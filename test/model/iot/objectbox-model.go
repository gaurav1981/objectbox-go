// Code generated by ObjectBox; DO NOT EDIT.

package iot

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(1)

	model.RegisterBinding(EventBinding)
	model.RegisterBinding(ReadingBinding)
	model.LastEntityId(2, 5284076134434938613)
	model.LastIndexId(2, 2642563953244304959)

	return model
}
