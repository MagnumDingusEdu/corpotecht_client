package utils

import (
	"github.com/denisbrodbeck/machineid"
)

func GetUniqueIdentifier() string {
	id, err := machineid.ProtectedID("UncleRicksRollerCoaster")
	HandleError(err)
	return id
}
