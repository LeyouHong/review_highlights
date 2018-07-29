package main

import (
	"github.com/LeyouHong/review_highlights/cache"
	"github.com/LeyouHong/review_highlights/process"
)

func main() {
	cache.CacheBuild()
	process.Process()
}
