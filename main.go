package main

import (
	"flag"
	"fmt"
	"os"
	"sitemap/sitemap_builder"
)

var _ = fmt.Println

func main() {
	var websiteDomain string
	flag.StringVar(&websiteDomain, "domain", "", "the website domain, have to provide full url, like https://www.google.com.")
	flag.Parse()
	if websiteDomain == "" {
		fmt.Println("you didn't provide website domain.")
		os.Exit(65)
	}

	fmt.Println(sitemap_builder.CreateSitemap(websiteDomain))

	// fmt.Println(sitemap_builder.CreateSitemap("https://www.calhoun.io"))
	// fmt.Println(sitemap_builder.CreateSitemap("https://www.larstornoe.com"))
	// fmt.Println(sitemap_builder.CreateSitemap("https://gobyexample.com/"))
}
