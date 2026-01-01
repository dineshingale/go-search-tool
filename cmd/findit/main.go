package main

import (
	"flag"
	"fmt"
	"go-search-tool/internal/engine"
	"os"
)

func main() {
	// Define command line flags
	pathPtr := flag.String("path", ".", "The directory to search in")
	namePtr := flag.String("name", "", "The name or part of the name to search for")
	extPtr := flag.String("ext", "", "The file extension (e.g., .pdf)")

	flag.Parse()

	if *namePtr == "" && *extPtr == "" {
		fmt.Println("Please provide a -name or -ext to search for.")
		os.Exit(1)
	}

	fmt.Printf("Searching for '%s' in %s...\n", *namePtr, *pathPtr)

	// Call our search engine
	results, err := engine.FindFiles(*pathPtr, *namePtr, *extPtr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print results
	if len(results) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, path := range results {
			fmt.Printf("[FOUND] %s\n", path)
		}
		fmt.Printf("\nDone! Found %d files.\n", len(results))
	}
}
