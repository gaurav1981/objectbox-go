// Code generated by ObjectBox; DO NOT EDIT.

package iot

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

// ObjectBoxModel declares and builds the model from all the entities in the package.
// It is usually used when setting-up ObjectBox as an argument to the Builder.Model() function.
func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(1)

	model.RegisterBinding(EventBinding)
	model.RegisterBinding(ReadingBinding)
	model.LastEntityId(2, 5284076134434938613)
	model.LastIndexId(2, 2642563953244304959)

	return model
}
