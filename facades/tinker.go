package facades

import (
	tinker "goravel-tinker"
	"goravel-tinker/contracts"
	"log"
)

func Tinker() contracts.Tinker {
	instance, err := tinker.App.Make(tinker.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Tinker)
}
