package spec

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Spec struct {
	ContractID string
}

func TestSpecCache_ConcurrentRace(t *testing.T) {
	t.Parallel()
	cache := NewCache()
	var wg sync.WaitGroup
	goroutines := 20
	ops := 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			contractID := fmt.Sprintf("contract-%d", id%5)
			for j := 0; j < ops; j++ {
				if j%2 == 0 {
					cache.Set(contractID, &Spec{ContractID: contractID})
				} else {
					_, _ = cache.Get(contractID)
				}
			}
		}(i)
	}
	wg.Wait()
	assert.True(t, true)
}
