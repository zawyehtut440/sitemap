package main

import (
	"fmt"
	"sitemap/sitemap_builder"
)

var _ = fmt.Println

func main() {
	fmt.Println(sitemap_builder.CreateSitemap("https://www.calhoun.io"))
	// fmt.Println(sitemap_builder.CreateSitemap("https://www.larstornoe.com"))
	// fmt.Println(sitemap_builder.CreateSitemap("https://gobyexample.com/"))
}
