package process

import (
	"os"
	"log"
	"bufio"
	"io"
	"github.com/LeyouHong/review_highlights/conf"
	"sort"
)

func getReviews(filePath string) []string {
	res := []string{}
	fi, err := os.Open(filePath)
	if err != nil {
		log.Panic("Error: %s\n", err)
		return res
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		res = append(res, string(a))
	}

	return res
}

func review_highlights(reviews []string, max int) []string {
	m := make(map[string]int)

	for _, str := range reviews {
		score := 0
		emotion, err := conf.EmotionClient.Classify(str)
		if err != nil {
			score += 0
		} else if emotion == "" {
			score += 1
		} else if emotion == "bad" {
			score -= 1
		} else  if emotion == "good" {
			score += 3
		}

		noun, err := conf.NounClient.Classify(str)
		if err != err {
			score += 0
		} else if noun == "" {
			score += 1
		} else if noun == "food" {
			score += 2
		}

		m[str] = score
	}

	res := []struct {
		str string
		score int
	}{}

	for k, v := range m {
		res = append(res, struct{str string
								 score int
								 }{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].score > res[j].score
	})

	result := []string{}
	for i := 0; i < max; i++ {
		result = append(result, res[i].str)
	}

	return result
}

