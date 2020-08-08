package main

import (
	"fmt"
	"os"
	"time"

	"public.com/goutil/intAndBytesUtil"
	"public.com/goutil/mathUtil"
)

var (
	count = 9
	cycle = 20
)

func main() {
	fmt.Println("It's time to manipulate wal file.")
	create()

	for i := 0; i < cycle; i++ {
		write(count)
		fmt.Printf("Cycle: %d has finished.\n", i+1)
	}

	delete()

	fmt.Println("Done.")
}

func testFilePerformance() {
	count = 130000
	start := time.Now().UnixNano()
	randomWriteFile(count)
	end := time.Now().UnixNano()
	randomTime := (end - start) / 1000000

	start = time.Now().UnixNano()
	sequentialWriteFile(count)
	end = time.Now().UnixNano()
	sequentialTime := (end - start) / 1000000

	fmt.Printf("Ramdon: %d, Sequential: %d, Difference: %d\n", randomTime, sequentialTime, randomTime-sequentialTime)
}

func randomWriteFile(count int) {
	fileName := "random.txt"
	func() {
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}()

	defer func() {
		os.Remove(fileName)
	}()

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0200)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < count; i++ {
		str := fmt.Sprintf("Write: This is NO. %d item.", i+1)
		data := []byte(str)
		length := len(data)
		header := intAndBytesUtil.Int32ToBytes(int32(length), byteOrder)
		data = append(header, data...)

		offset := mathUtil.GetRand().GetRandRangeInt64(0, fileSize)
		file.Seek(offset, 0)
		_, err := file.Write(data)
		if err != nil {
			panic(err)
		}
	}
	file.Sync()
}

func sequentialWriteFile(count int) {
	fileName := "sequential.txt"
	fileSize := int64(1024 * 1024 * 1024)
	func() {
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = file.Truncate(fileSize)
		if err != nil {
			panic(err)
		}
	}()

	defer func() {
		os.Remove(fileName)
	}()

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0200)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	offset := int64(0)
	for i := 0; i < count; i++ {
		str := fmt.Sprintf("Write: This is NO. %d item.", i+1)
		data := []byte(str)
		length := len(data)
		header := intAndBytesUtil.Int32ToBytes(int32(length), byteOrder)
		data = append(header, data...)
		dataLen := int64(len(data))
		if offset+dataLen >= fileSize {
			fmt.Println("There is no space for new data. Quit.")
			break
		}
		n, err := file.Write(data)
		if err != nil {
			panic(err)
		}
		offset += int64(n)
	}
	file.Sync()
}
