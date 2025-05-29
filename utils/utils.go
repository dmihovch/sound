package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

func PrintWavHeader(header []byte) {
	fmt.Println("ChunkID:      ", string(header[0:4]))
	fmt.Println("ChunkSize:    ", binary.LittleEndian.Uint32(header[4:8]))
	fmt.Println("Format:       ", string(header[8:12]))
	fmt.Println("Subchunk1ID:  ", string(header[12:16]))
	fmt.Println("Subchunk1Size:", binary.LittleEndian.Uint32(header[16:20]))
	fmt.Println("AudioFormat:  ", binary.LittleEndian.Uint16(header[20:22]))
	fmt.Println("NumChannels:  ", binary.LittleEndian.Uint16(header[22:24]))
	fmt.Println("SampleRate:   ", binary.LittleEndian.Uint32(header[24:28]))
	fmt.Println("ByteRate:     ", binary.LittleEndian.Uint32(header[28:32]))
	fmt.Println("BlockAlign:   ", binary.LittleEndian.Uint16(header[32:34]))
	fmt.Println("BitsPerSample:", binary.LittleEndian.Uint16(header[34:36]))
	fmt.Println("Subchunk2ID:  ", string(header[36:40]))
	fmt.Println("Subchunk2Size:", binary.LittleEndian.Uint32(header[40:44]))
}

func PrintWavBytes(data []byte, dataSize int) {
	for i := 0; i < dataSize; i += 4 {
		if i+3 >= dataSize {
			break
		}
		left := int16(binary.LittleEndian.Uint16(data[i : i+2]))
		right := int16(binary.LittleEndian.Uint16(data[i+2 : i+4]))
		fmt.Println(left, "|", right)

	}
}

func GetIOPaths() (inputPath string, outputPath string, err error) {
	if len(os.Args) != 3 {
		return "", "", errors.New("usage: go run main.go <path-to-wav> <output-path>")
	}
	return os.Args[1], os.Args[2], nil
}

func GetFileSlice(file *os.File, size int) ([]byte, error) {
	slice := make([]byte, size)
	_, err := io.ReadFull(file, slice)
	if err != nil {
		return nil, err
	}
	return slice, nil
}
