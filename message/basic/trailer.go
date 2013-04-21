package basic

import (
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/tag"
)

//Message Trailer type. TrailerOrder FieldOrder.
type Trailer struct {
	FieldMap
}

//CheckSum is a required field of the trailer
func (t *Trailer) setCheckSum(checkSum message.StringField) {
	t.SetField(checkSum)
}

//Must be called before use
func (t *Trailer) init() {
	t.FieldMap.init(trailerFieldOrder)
}

//Field 10 must be last in the trailer
func trailerFieldOrder(i, j tag.Tag) bool {
	switch {
	case i == tag.CheckSum:
		return false
	case j == tag.CheckSum:
		return true
	}

	return i < j
}
