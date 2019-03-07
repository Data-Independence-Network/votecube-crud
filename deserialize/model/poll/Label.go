package poll

import (
	"github.com/diapco/votecube-crud/deserialize"
	"github.com/diapco/votecube-crud/models"
	"github.com/volatiletech/null"
)

/**
 * Please try to keep properties serialized in UI-model alphabetic order. :)
 */

func DeserializeLabel(data []byte, cursor *int64, dataLen int64, err error) (models.Label, int64, error) {
	var label models.Label

	if err != nil {
		return label, 0, err
	}

	var objectType byte
	objectType, err = deserialize.RByte(data, cursor, dataLen, err)

	var labelId int64
	labelId, err = deserialize.RNum(data, cursor, dataLen, err)

	if objectType == deserialize.REFERENCE {
		return label, labelId, err
	} else {
		//var description string
		//description, err = deserialize.RStr(data, cursor, dataLen, err)
		var name string
		name, err = deserialize.RStr(data, cursor, dataLen, err)

		var parentLabelId int64
		var parentDimIdReference null.Int64
		objectType, err = deserialize.RByte(data, cursor, dataLen, err)
		if objectType == deserialize.REFERENCE {
			parentLabelId, err = deserialize.RNum(data, cursor, dataLen, err)
			parentDimIdReference = null.Int64{
				Int64: parentLabelId,
				Valid: true,
			}
		}

		if err != nil {
			return label, 0, err
		}

		label = models.Label{
			LabelName:     name,
			ParentLabelID: parentDimIdReference,
		}
	}

}
