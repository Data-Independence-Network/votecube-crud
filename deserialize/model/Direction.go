package model

import (
	"github.com/diapco/votecube-crud/models"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDirection(
	data []byte, cursor *int64, dataLen int64, err error) (models.Direction, bool, error) {
	var dimension models.Direction

	if err != nil {
		return dimension, err
	}

}
