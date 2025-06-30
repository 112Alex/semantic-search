package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/yourname/semantic-search/pkg/semhtml"
)

func main() {
	topN := flag.Int("n", 5, "количество результатов")
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("usage: semhtml [-n N] <html file> <query string>")
		os.Exit(1)
	}
	htmlPath := flag.Arg(0)
	query := flag.Arg(1)

	f, err := os.Open(htmlPath)
	if err != nil {
		log.Fatalf("не могу открыть файл: %v", err)
	}
	defer f.Close()

	eng, err := semhtml.NewEngineFromReader(f)
	if err != nil {
		log.Fatalf("ошибка парсинга: %v", err)
	}

	results := eng.Search(query, *topN)
	for _, r := range results {
		fmt.Printf("%.4f\t%s\t%s\n", r.Score, r.Selector, r.Text)
	}
}
