package facades

import (
	tinker "github.com/hulutech-web/goravel-tinker"
	"github.com/hulutech-web/goravel-tinker/contracts"
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
