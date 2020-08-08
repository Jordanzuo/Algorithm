package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"

	"public.com/goutil/intAndBytesUtil"
)

var (
	fileName  = "wal"
	fileSize  = int64(1 * 1024)
	maxIndex  = fileSize - 1
	producer  = int64(0) // 下次开始写入的位置(不包括)
	byteOrder = binary.LittleEndian
)

func create() {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = file.Truncate(fileSize)
	if err != nil {
		panic(err)
	}
}

func delete() {
	os.Remove(fileName)
}

// func write(count int) {
// 	file, err := os.OpenFile(fileName, os.O_WRONLY, 0200)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	offset := producer % fileSize
// 	_, err = file.Seek(offset, 0)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 0; i < count; i++ {
// 		str := fmt.Sprintf("Write: This is NO. %d item.", i+1)
// 		data := []byte(str)
// 		length := len(data)
// 		header := intAndBytesUtil.Int32ToBytes(int32(length), byteOrder)
// 		data = append(header, data...)
// 		dataLen := int64(len(data))

// 		// 判断是否已经写满，如果已经写满则需要等待数据被读取
// 		for {
// 			producerOffset := producer % fileSize
// 			consumerOffset := consumer % fileSize
// 			leftSpace := int64(0)
// 			if producerOffset >= consumerOffset {
// 				leftSpace = fileSize - producerOffset + consumerOffset
// 			} else {
// 				leftSpace = consumerOffset - producerOffset
// 			}
// 			if leftSpace < dataLen {
// 				fmt.Printf("There is %d bytes left, but it requires %d bytes.\n", leftSpace, dataLen)
// 				read(5 * count)
// 				time.Sleep(time.Second)
// 			} else {
// 				break
// 			}
// 		}

// 		// 判断是否会写越界，如果会越界则需要将数据拆分成两部分
// 		if offset+dataLen-1 <= maxIndex {
// 			fmt.Println("Write: There is enough space for content")
// 			n, err := file.Write(data)
// 			if err != nil {
// 				panic(err)
// 			}
// 			producer += int64(n)
// 			offset += int64(n)
// 		} else {
// 			fmt.Println("Write: There isn't enough space for content")
// 			index := maxIndex - offset + 1
// 			part1, part2 := data[:index], data[index:]

// 			// 先写入第一部分
// 			n, err := file.Write(part1)
// 			if err != nil {
// 				panic(err)
// 			}
// 			producer += int64(n)
// 			offset = 0
// 			fmt.Printf("Write: ######################################################%d bytes has been written successfully\n", n)

// 			// 将指针指到文件头
// 			_, err = file.Seek(offset, 0)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println("Write: #####################################################Move file pointer to the beginning.#################")

// 			// 再写入第二部分
// 			n, err = file.Write(part2)
// 			if err != nil {
// 				panic(err)
// 			}
// 			producer += int64(n)
// 			offset += int64(n)
// 			fmt.Printf("Write: ######################################################%d bytes has been written successfully\n", n)
// 		}

// 		fmt.Printf("Write: %d bytes has been written successfully\n", dataLen)
// 		fmt.Printf("#%v# Write: Start: (%d: %d, %d), End:(%d, %d)\n", time.Now(), producer, producer%fileSize, offset, consumer, consumer % fileSize)
// 	}
// }

func write(count int) {
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
		dataLen := int64(len(data))

		// 判断是否已经写满，如果已经写满则需要等待数据被读取
		for {
			leftSpace := int64(0)
			if producer >= consumer {
				leftSpace = fileSize - producer + consumer
			} else {
				leftSpace = consumer - producer
			}
			if leftSpace < dataLen {
				fmt.Printf("There is %d bytes left, but it requires %d bytes.\n", leftSpace, dataLen)
				read(5 * count)
				time.Sleep(time.Second)
			} else {
				break
			}
		}

		// 判断是否会写越界，如果会越界则需要将数据拆分成两部分
		if producer+dataLen-1 <= maxIndex {
			fmt.Println("Write: There is enough space for content")
			n, err := file.WriteAt(data, producer)
			if err != nil {
				panic(err)
			}
			producer += int64(n)
		} else {
			fmt.Println("Write: There isn't enough space for content")
			index := maxIndex - producer + 1
			part1, part2 := data[:index], data[index:]

			// 先写入第一部分
			n, err := file.WriteAt(part1, producer)
			if err != nil {
				panic(err)
			}
			producer = 0
			fmt.Printf("Write: ######################################################%d bytes has been written successfully\n", n)
			fmt.Println("Write: #####################################################Move file pointer to the beginning.#################")

			// 再写入第二部分
			n, err = file.WriteAt(part2, producer)
			if err != nil {
				panic(err)
			}
			producer += int64(n)
			fmt.Printf("Write: ######################################################%d bytes has been written successfully\n", n)
		}

		fmt.Printf("Write: %d bytes has been written successfully\n", dataLen)
		fmt.Printf("#%v# Write: Start: (%d), End:(%d)\n", time.Now(), producer, consumer)
	}
}
