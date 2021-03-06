package event

import (
	"sync"
	"reflect"
)

// Feed implements one-to-many subscriptions where the carrier of events is a channel.
// Values sent to a Feed are delivered to all subscribed channels simultaneously.
//
// Feeds can only be used with a single type. The type is determined by the first Send or
// Subscribe operation. Subsequent calls to these methods panic if the type does not
// match.
//
// The zero value is ready to use.
type Feed struct {
	once      sync.Once           // ensures that init only runs once
	sendLock  chan struct{}     // sendLock has a one-element buffer and is empty when held.It protects sendCases.
	removeSub chan interface{} // interrupts Send
	sendCases caseList          // the active set of select cases used by Send
	mu     sync.Mutex          // The inbox holds newly subscribed channels until they are added to sendCases.
	inbox  caseList
	etype  reflect.Type
	closed bool
}
type caseList []reflect.SelectCase
