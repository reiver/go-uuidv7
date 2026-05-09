package uuidv7

import (
	"math/rand/v2"
	"sync"
	"time"
)

var prngPool = sync.Pool{
	New: func() any {
		return rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), 0))
	},
}

func randomUint64() uint64 {
	return prngPool.Get().(*rand.Rand).Uint64()
}
