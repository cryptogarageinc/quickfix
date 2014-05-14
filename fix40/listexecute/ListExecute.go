//Package listexecute msg type = L.
package listexecute

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a ListExecute wrapper for the generic Message type
type Message struct {
	message.Message
}

//ListID is a required field for ListExecute.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListExecute.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for ListExecute.
func (m Message) WaveNo() (*field.WaveNoField, errors.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListExecute.
func (m Message) GetWaveNo(f *field.WaveNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListExecute.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListExecute.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ListExecute messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ListExecute.
func Builder(
	listid *field.ListIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("L"))
	builder.Body().Set(listid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "L", r
}