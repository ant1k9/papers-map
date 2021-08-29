package main

import (
	"regexp"

	"github.com/ant1k9/knowledge-map/pkg/serve"
)

func main() {
	fs := serve.NewFileSearcher()
	*fs.DescriptionPattern() = *regexp.MustCompile(`\*Abstract\*: (.*)`)
	serve.Serve(fs)
}
