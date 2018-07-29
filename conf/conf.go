package conf

import (
	"sync"
	sh "github.com/eaigner/shield"
	)

type Shield struct {
	sync.RWMutex
	Shd sh.Shield
}

func newShield(addr string) *Shield {
	s := new(Shield)
	s.Shd = sh.New(
		sh.NewEnglishTokenizer(),
		sh.NewRedisStore(addr, "",  nil, ""),
	)

	return s
}

func (s *Shield)Classify(val string) (string, error) {
	s.Lock()
	val, err := s.Shd.Classify(val)
	s.Unlock()
	return val, err
}

var EmotionClient = newShield("redis:6379")
var NounClient = newShield("redis:6378")