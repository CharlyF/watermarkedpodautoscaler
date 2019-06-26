package controller

import (
	"github.com/CharlyF/watermarkedpodautoscaler/pkg/controller/watermarkedpodautoscaler"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, watermarkedpodautoscaler.Add)
}
