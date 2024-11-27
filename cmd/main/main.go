package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/clhilgert/gosort/pkg/algorithms"
	"github.com/clhilgert/gosort/pkg/utils"
)

func main() {
	uniqueFlag := flag.Bool("u", false, "Return unique values")

	radixsortFlag := flag.Bool("radixsort", false, "Use radixsort")
	mergesortFlag := flag.Bool("mergesort", false, "Use mergesort")
	heapsortFlag := flag.Bool("heapsort", false, "Use heapsort")
	quicksortFlag := flag.Bool("quicksort", true, "Use quicksort")

	flag.Parse()
	args := flag.Args()

	var file *os.File
	if len(args) == 0 {
		file = os.Stdin
	} else {
		var err error
		file, err = os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	data := utils.BuildTextArray(file)

	if *uniqueFlag {
		data = utils.GetUnique(data)
	}

	var sorted []string

	switch {
	case *radixsortFlag:
		sorted = algorithms.RadixSort(data)
	case *mergesortFlag:
		sorted = algorithms.MergeSort(data)
	case *heapsortFlag:
		sorted = algorithms.HeapSort(data)
	case *quicksortFlag:
		sorted = algorithms.QuickSort(data)
	default:
		sorted = algorithms.QuickSort(data)
	}

	fmt.Print(strings.Join(sorted, "\n"))
}
