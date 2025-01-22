package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fieldNum := flag.Int("f", 1, "field number to extract")
	fieldDel := flag.String("d", "\t", "field delimiter")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Error: Please provide an input file")
		os.Exit(1)
	}

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), *fieldDel)
		if *fieldNum <= len(fields) {
			fmt.Println(fields[*fieldNum-1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
}
