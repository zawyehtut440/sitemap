package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sitemap/sitemap_builder"
)

var _ = fmt.Println

func main() {
	var websiteDomain string
	var depth int
	flag.StringVar(&websiteDomain, "domain", "", "the website domain, have to provide full url, like https://www.google.com.")
	flag.IntVar(&depth, "depth", math.MaxInt, "the max depth to visit website when building a sitemap. note: depth > 0")
	flag.Parse()
	if websiteDomain == "" {
		fmt.Println("you didn't provide website domain.")
		os.Exit(65)
	}
	if depth <= 0 {
		fmt.Println("depth have to greater than 0.")
		os.Exit(65)
	}

	fmt.Println(sitemap_builder.CreateSitemap(websiteDomain, depth))

	// fmt.Println(sitemap_builder.CreateSitemap("https://www.calhoun.io"))
	// fmt.Println(sitemap_builder.CreateSitemap("https://www.larstornoe.com"))
	// fmt.Println(sitemap_builder.CreateSitemap("https://gobyexample.com/"))
}
