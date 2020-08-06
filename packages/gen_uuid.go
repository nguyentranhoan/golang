package Media

import (
	"github.com/google/uuid"
)

func Gen_uuid() string{
	id, err := uuid.NewRandom()
	if err !=nil {
		// handle error
	}
	result := id.String()

	return result
}
