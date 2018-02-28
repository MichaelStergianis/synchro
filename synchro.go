package synchro

import "sync"

// Synchro ...
// Defines a convenient way to pass around shared locks
type Synchro struct {
	Mut sync.Mutex
	Wg  sync.WaitGroup
}
