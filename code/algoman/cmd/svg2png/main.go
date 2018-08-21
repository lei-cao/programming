package main

import (
	"os/exec"
	"flag"
	"strings"
	"fmt"
)

var (
	inputFilename  = flag.String("input", "", "input filename")
	outputFilename = flag.String("output", "", "output filename")
	width          = flag.String("width", "128", "output file width")
	height         = flag.String("height", "128", "output file height")

	packageName = flag.String("package", "main", "package name")
	varName     = flag.String("var", "", "variable name")
	compress    = flag.String("compress", "false", "use gzip compression")
	buildtags   = flag.String("buildtags", "", "build tags")
)

func main() {
	flag.Parse()

	if *outputFilename == "" {
		outputFilename = inputFilename
	}

	if *varName == "" {
		paths := strings.Split(*inputFilename, "/")
		fileName := paths[len(paths) - 1]
		*varName = strings.ToUpper(fileName) + "_png"
	}


	pngFile := *outputFilename + ".png"
	goFile := *outputFilename + ".go"

	// inkscape -z -e backward-on.png -w 128 -h 128 backward-solid.svg
	if err := exec.Command("inkscape", "-z", "-e", pngFile, "-w", *width, "-h", *height, *inputFilename + ".svg").Run(); err != nil {
		panic(err)
	}

	//fmt.Println(*packageName, " ", pngFile, " ", goFile, " ", *varName, " ", *compress, " ", *buildtags)

	//file2byteslice -package=images -input=./images/fast-backward.png -output=./images/fast_backward -var=FAST_BACKWARD_png
	c:=exec.Command("file2byteslice",
		"-package", *packageName,
		"-input", pngFile,
		"-output", goFile,
		"-var", *varName,
		//"-compress", *compress,
		//"-buildtags", *buildtags,
	)
	fmt.Println(c.Path, c.Args)
	if err := c.Run(); err != nil {
		panic(err)
	}
}
