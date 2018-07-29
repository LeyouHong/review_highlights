package cache

import (
	"log"
	"github.com/LeyouHong/review_highlights/conf"
)

func CacheBuild() {
	log.Print("cache building ...")

	err := conf.EmotionClient.Shd.Learn("good", "like love pretty nice quick good fine great")
	if err == nil {
		log.Println("learn Good Emotion data ...")
	}
	err = conf.EmotionClient.Shd.Learn("bad", "fuck normal just bad sad")
	if err == nil {
		log.Println("learn Bad location data ...")
	}
	err = conf.NounClient.Shd.Learn("food", "falafel sandwich food hot dogs")
	if err == nil {
		log.Println("learn person data successful")
	}
}
