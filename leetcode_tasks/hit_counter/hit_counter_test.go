package hit_counter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHitCounter(t *testing.T) {
	hc := Constructor()
	hc.Hit(1)
	hc.Hit(2)
	hc.Hit(3)
	hc.Hit(3)
	hc.Hit(3)
	assert.Equal(t, 5, hc.GetHits(4))
	hc.Hit(300)
	assert.Equal(t, 5, hc.GetHits(301))
}
