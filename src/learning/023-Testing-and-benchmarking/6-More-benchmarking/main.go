package main

import (
	"example.com/023-Testing-and-benchmarking/More-benchmarking/mystr"
	"fmt"
	"strings"
)

const s = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam vel urna porta, fringilla felis at, " +
	"ultricies turpis. Nulla facilisi. Pellentesque eu turpis facilisis, pulvinar nisi tempor, luctus neque. Nulla " +
	"nec ante finibus, vulputate leo sed, porttitor lacus. Aliquam euismod ornare ante vel convallis. Nunc laoreet " +
	"leo eu lacus mattis, et maximus ante facilisis. Sed vel tristique mi, eget tempus magna. Nullam pellentesque " +
	"nisl eu imperdiet mollis. Vivamus urna dolor, semper in pellentesque vitae, porttitor vitae odio. Vestibulum " +
	"nec sem enim. Nam placerat varius lobortis. Integer sodales pharetra mi, quis dictum lacus laoreet vitae. " +
	"Donec hendrerit sapien ut orci tincidunt tincidunt."

var xs []string
func main() {
	xs = strings.Split(s, " ")

	fmt.Printf("\n%s\n", mystr.Cat(xs))
	fmt.Printf("\n%s\n", mystr.Join(xs))
}
