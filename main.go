package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func main() {
	sound, err := os.Open("sound.wav")
	if err != nil {
		panic(err)
	}
	defer sound.Close()

	header := make([]byte, 44)
	_, err = io.ReadFull(sound, header)
	if err != nil {
		panic(err)
	}

	fmt.Println("ChunkID:      ", string(header[0:4])) // "RIFF"
	fmt.Println("ChunkSize:    ", binary.LittleEndian.Uint32(header[4:8]))
	fmt.Println("Format:       ", string(header[8:12]))  // "WAVE"
	fmt.Println("Subchunk1ID:  ", string(header[12:16])) // "fmt "
	fmt.Println("Subchunk1Size:", binary.LittleEndian.Uint32(header[16:20]))
	fmt.Println("AudioFormat:  ", binary.LittleEndian.Uint16(header[20:22]))
	fmt.Println("NumChannels:  ", binary.LittleEndian.Uint16(header[22:24]))
	fmt.Println("SampleRate:   ", binary.LittleEndian.Uint32(header[24:28]))
	fmt.Println("ByteRate:     ", binary.LittleEndian.Uint32(header[28:32]))
	fmt.Println("BlockAlign:   ", binary.LittleEndian.Uint16(header[32:34]))
	fmt.Println("BitsPerSample:", binary.LittleEndian.Uint16(header[34:36]))
	fmt.Println("Subchunk2ID:  ", string(header[36:40])) // "data"
	fmt.Println("Subchunk2Size:", binary.LittleEndian.Uint32(header[40:44]))

	dataSize := binary.LittleEndian.Uint32((header[40:44]))

	data := make([]byte, dataSize)

	_, err = io.ReadFull(sound, data)
	if err != nil {
		panic(err)
	}

	for i := 0; i < int(dataSize); i += 4 {
		if i+3 >= int(dataSize) {
			break
		}

		left := int16(binary.LittleEndian.Uint16(data[i : i+2]))
		if checkAmp(left) {
			data[i] = 0
			data[i+1] = 0
		}
		left = int16(binary.LittleEndian.Uint16(data[i : i+2]))

		right := int16(binary.LittleEndian.Uint16(data[i+2 : i+4]))

		if checkAmp(right) {
			data[i+2] = 0
			data[i+3] = 0
		}
		right = int16(binary.LittleEndian.Uint16(data[i+2 : i+4]))

		fmt.Println(left, " | ", right)

	}

	newFile, err := os.Create("output3.wav")
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	newData := append(header, data...)

	newFile.Write(newData)

}

func checkAmp(amp int16) bool {
	return amp >= -500 && amp <= 500
}
