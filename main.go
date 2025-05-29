package main

import (
	"encoding/binary"
	"os"
	compress "sound/compression"
	"sound/utils"
)

func main() {

	inputPath, outputPath, err := utils.GetIOPaths()
	if err != nil {
		panic(err)
	}

	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer input.Close()

	output, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	inputHeader, err := utils.GetFileSlice(input, 44)
	if err != nil {
		panic(err)
	}

	utils.PrintWavHeader(inputHeader)

	inputDataSize := binary.LittleEndian.Uint32((inputHeader[40:44]))

	inputData, err := utils.GetFileSlice(input, int(inputDataSize))
	if err != nil {
		panic(err)
	}

	utils.PrintWavBytes(inputData, int(inputDataSize))

	compress.Compress()

}
