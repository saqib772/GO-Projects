// Step 1: Import necessary packages
// Add the necessary import statements for the compress/gzip and io/ioutil packages.

package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Step 2: Define the compress command
// Add a new case to the switch statement that handles the compress command. In this case, we will read the contents of a file, compress them, and write the compressed data to a new file.

case "compress":
	// Parse the filename flag
	filename := flag.String("filename", "", "the filename to compress")
	flag.Parse()

	if *filename == "" {
		fmt.Println("Please specify a filename using the -filename flag.")
		os.Exit(1)
	}

	// Read the contents of the file
	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Compress the data
	compressedData := make([]byte, len(data))
	compressor, err := gzip.NewWriterLevel(ioutil.Discard, gzip.BestCompression)
	if err != nil {
		fmt.Printf("Error creating compressor: %v\n", err)
		os.Exit(1)
	}
	compressor.Write(data)
	compressor.Close()
	compressedData = compressor.Bytes()

	// Write the compressed data to a new file
	err = ioutil.WriteFile(*filename+".gz", compressedData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("File %s compressed successfully.\n", *filename)

// Step 3: Build and run the CLI tool
// Build and run the CLI tool using the go build and ./mycli compress -filename=myfile.txt commands in the terminal. This will compress the contents of the myfile.txt file and write the compressed data to a new file called myfile.txt.gz.
