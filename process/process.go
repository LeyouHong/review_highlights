package process

import (
	"os"
	"log"
	"strconv"
)

func Process() {
	args := os.Args
	if(len(args) != 3) {
		log.Panic("The input arguments is not 2!")
	}

	filePath, maxStr := args[1], args[2]
	reviews := getReviews(filePath)

	max, err := strconv.Atoi(maxStr)
	if err != nil || max <= 0 {
		log.Panic("Second input argument has issues")
	}

	res := review_highlights(reviews, max)

	for _, s := range res {
		log.Println(s)
	}
}
