// Package pool provides functions for managing a pool of Lua state structs
package pool

import (
	"sync"

	"github.com/xyproto/gopher-lua"
)

// The LState pool pattern, as recommended by the author of gopher-lua:
// https://github.com/xyproto/gopher-lua#the-lstate-pool-pattern

// LStatePool is a pool of Lua states, with a mutex
type LStatePool struct {
	m     sync.Mutex
	saved []*lua.LState
}

// New returns a new Lua pool structure
func New() *LStatePool {
	return &LStatePool{saved: make([]*lua.LState, 0, 4)}
}

// New returns a new Lua state
func (pl *LStatePool) New() *lua.LState {
	L := lua.NewState()
	// setting the L up here.
	// load scripts, set global variables, share channels, etc...
	return L
}

// Get borrows an existing Lua state
func (pl *LStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.saved)
	if n == 0 {
		return pl.New()
	}
	x := pl.saved[n-1]
	pl.saved = pl.saved[0 : n-1]
	return x
}

// Put delivers back a borrowed Lua state
func (pl *LStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.saved = append(pl.saved, L)
}

// Shutdown can be used then the Lua pool is being shut down
func (pl *LStatePool) Shutdown() {
	// The following line causes a race condition with the
	// graceful shutdown package at server shutdown:
	//for _, L := range pl.saved {
	//	L.Close()
	//}
	// TODO: Add a test to catch this.
	// TODO: Figure out why.
}
