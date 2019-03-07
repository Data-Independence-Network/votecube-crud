package model

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
	"github.com/volatiletech/null"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeDirection(data []byte, cursor *int64, dataLen int64, err error) (models.Direction, int64, error) {
	var direction models.Direction

	if err != nil {
		return direction, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(data, cursor, dataLen, err)

	var directionId int64
	directionId, err = deserialize.RNum(data, cursor, dataLen, err)

	if objectType == deserialize.REFERENCE {
		return direction, directionId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(data, cursor, dataLen, err)

		var parentDirectionId int64
		var parentDirIdReference null.Int64
		objectType, err = deserialize.RByte(data, cursor, dataLen, err)
		if objectType == deserialize.REFERENCE {
			parentDirectionId, err = deserialize.RNum(data, cursor, dataLen, err)
			parentDirIdReference = null.Int64{
				Int64: parentDirectionId,
				Valid: true,
			}
		}

		if err != nil {
			return direction, 0, err
		}

		direction = models.Direction{
			DirectionDescription: name,
			ParentDirectionID:    parentDirIdReference,
		}
	}

}
