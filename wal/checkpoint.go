package main

import (
	"fmt"
	"os"
	"time"

	"public.com/goutil/intAndBytesUtil"
)

var (
	consumer = int64(0) // 下次消费的位置
)

// func read(count int) {
// 	file, err := os.OpenFile(fileName, os.O_RDONLY, 0400)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	offset := consumer % fileSize
// 	_, err = file.Seek(offset, 0)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 0; i < count; i++ {
// 		// 判断是否有尚未消费的数据
// 		producerOffset := producer % fileSize
// 		consumerOffset := consumer % fileSize
// 		available := int64(0)
// 		if producerOffset >= consumerOffset {
// 			available = producerOffset - consumerOffset
// 		} else if producerOffset < consumerOffset {
// 			available = fileSize - (consumerOffset - producerOffset)
// 		}
// 		if available == 0 {
// 			fmt.Printf("#%v# Read: Start: (%d: %d), End:(%d, %d). There is no data to read\n", time.Now(), producer, producer%fileSize, consumer, offset)
// 			break
// 		}

// 		if offset - 1 + 4 <= maxIndex {
// 			fmt.Println("Read: There is enough space for header.")
// 			header := make([]byte, 4)
// 			n, err := file.ReadAt(header, offset)
// 			if err != nil {
// 				panic(err)
// 			}
// 			consumer += int64(n)
// 			offset += int64(n)
// 			fmt.Printf("Read: Read %d bytes for header.\n", n)

// 			// 取得内容的长度
// 			length := intAndBytesUtil.BytesToInt32(header, byteOrder)
// 			fmt.Printf("Read: Length:%d\n", length)

// 			// 判断consumer到文件末尾之间是否有足够的空间
// 			if offset - 1 + int64(length) <= maxIndex {
// 				fmt.Println("Read: There is enough space for content")
// 				content := make([]byte, int(length))
// 				n, err = file.ReadAt(content, offset)
// 				if err != nil {
// 					panic(err)
// 				}

// 				consumer += int64(n)
// 				offset += int64(n)
// 				fmt.Printf("Read: Read %d bytes for content. Content:%s\n", n, string(content))
// 				fmt.Printf("#%v# Read: Start: (%d: %d), End:(%d, %d)\n", time.Now(), producer, producer%fileSize, consumer, offset)

// 				// 判断是否读取到文件的末尾了，如果是则将文件指针移动到文件头
// 				if offset == maxIndex {
// 					offset = 0
// 					_, err := file.Seek(offset, 0)
// 					if err != nil {
// 						panic(err)
// 					}
// 				}
// 			} else {
// 				fmt.Println("Read: There isn't enough space for content")
// 				part1 := make([]byte, int(maxIndex - offset + 1))
// 				part2 := make([]byte, int(length)-len(part1))

// 				// 先读取第一部分
// 				n, err := file.ReadAt(part1, offset)
// 				if err != nil {
// 					panic(err)
// 				}
// 				consumer += int64(n)
// 				offset = 0
// 				fmt.Printf("Read: Read first part: %d bytes\n", n)

// 				// 将指针指到文件头
// 				_, err = file.Seek(offset, 0)
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println("Read: Move cursor to 0")

// 				// 再读取第二部分
// 				n, err = file.ReadAt(part2, offset)
// 				if err != nil {
// 					panic(err)
// 				}
// 				consumer += int64(n)
// 				offset += int64(n)
// 				fmt.Printf("Read: Read second part: %d bytes\n", n)

// 				// 将part1和part2的数据进行合并
// 				part1 = append(part1, part2...)
// 				fmt.Printf("Read: Read %d bytes for content. Content:%s\n", n, string(part1))
// 				fmt.Printf("#%v# Read: Start: (%d: %d), End:(%d, %d)\n", time.Now(), producer, producer%fileSize, consumer, offset)
// 			}
// 		} else {
// 			fmt.Println("Read: There isn't enough space for header.")

// 			// 需要将文件头分两次读取
// 			part1 := make([]byte, int(maxIndex - offset + 1))
// 			part2 := make([]byte, 4-len(part1))

// 			// 先读取第一部分
// 			n, err := file.ReadAt(part1, offset)
// 			if err != nil {
// 				panic(err)
// 			}
// 			consumer += int64(n)
// 			offset = 0
// 			fmt.Printf("Read: Read the first part of header: %d\n", n)

// 			// 将指针指到文件头
// 			_, err = file.Seek(offset, 0)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println("Read: Move cursor to 0")

// 			// 再读取第二部分
// 			n, err = file.ReadAt(part2, offset)
// 			if err != nil {
// 				panic(err)
// 			}
// 			consumer += int64(n)
// 			offset += int64(n)
// 			fmt.Printf("Read: Read the second part of header: %d\n", n)

// 			// 将part1和part2的数据进行合并
// 			part1 = append(part1, part2...)

// 			// 取得内容的长度
// 			length := intAndBytesUtil.BytesToInt32(part1, byteOrder)
// 			fmt.Printf("Read: Length:%d\n", length)

// 			// 读取内容
// 			content := make([]byte, int(length))
// 			n, err = file.ReadAt(content, offset)
// 			if err != nil {
// 				panic(err)
// 			}

// 			consumer += int64(n)
// 			offset += int64(n)
// 			fmt.Printf("Read: Read %d bytes for content. Content:%s\n", n, string(content))
// 			fmt.Printf("#%v# Read: Start: (%d: %d), End:(%d, %d)\n", time.Now(), producer, producer%fileSize, consumer, offset)
// 		}
// 	}
// }

func read(count int) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0400)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < count; i++ {
		// 判断是否有尚未消费的数据
		available := int64(0)
		if producer >= consumer {
			available = producer - consumer
		} else if producer < consumer {
			available = fileSize - (consumer - producer)
		}
		if available == 0 {
			fmt.Printf("#%v# Read: Start: (%d), End:(%d). There is no data to read\n", time.Now(), producer, consumer)
			break
		}

		if consumer-1+4 <= maxIndex {
			fmt.Println("Read: There is enough space for header.")
			header := make([]byte, 4)
			n, err := file.ReadAt(header, consumer)
			if err != nil {
				panic(err)
			}
			consumer += int64(n)
			fmt.Printf("Read: Read %d bytes for header.\n", n)

			// 取得内容的长度
			length := intAndBytesUtil.BytesToInt32(header, byteOrder)
			fmt.Printf("Read: Length:%d\n", length)

			// 判断consumer到文件末尾之间是否有足够的空间
			if consumer-1+int64(length) <= maxIndex {
				fmt.Println("Read: There is enough space for content")
				content := make([]byte, int(length))
				n, err = file.ReadAt(content, consumer)
				if err != nil {
					panic(err)
				}

				consumer += int64(n)
				fmt.Printf("Read: Read %d bytes for content. Content:%s\n", n, string(content))
				fmt.Printf("#%v# Read: Start: (%d), End:(%d)\n", time.Now(), producer, consumer)

				// 判断是否读取到文件的末尾了，如果是则将文件指针移动到文件头
				if consumer == maxIndex {
					consumer = 0
				}
			} else {
				fmt.Println("Read: There isn't enough space for content")
				part1 := make([]byte, int(maxIndex-consumer+1))
				part2 := make([]byte, int(length)-len(part1))

				// 先读取第一部分
				n, err := file.ReadAt(part1, consumer)
				if err != nil {
					panic(err)
				}
				consumer = 0
				fmt.Printf("Read: Read first part: %d bytes\n", n)

				// 再读取第二部分
				n, err = file.ReadAt(part2, consumer)
				if err != nil {
					panic(err)
				}
				consumer += int64(n)
				fmt.Printf("Read: Read second part: %d bytes\n", n)

				// 将part1和part2的数据进行合并
				part1 = append(part1, part2...)
				fmt.Printf("Read: Read %d bytes for content. Content:%s\n", n, string(part1))
				fmt.Printf("#%v# Read: Start: (%d), End:(%d)\n", time.Now(), producer, consumer)
			}
		} else {
			fmt.Println("Read: There isn't enough space for header.")

			// 需要将文件头分两次读取
			part1 := make([]byte, int(maxIndex-consumer+1))
			part2 := make([]byte, 4-len(part1))

			// 先读取第一部分
			n, err := file.ReadAt(part1, consumer)
			if err != nil {
				panic(err)
			}
			consumer = 0
			fmt.Printf("Read: Read the first part of header: %d\n", n)

			// 再读取第二部分
			n, err = file.ReadAt(part2, consumer)
			if err != nil {
				panic(err)
			}
			consumer += int64(n)
			fmt.Printf("Read: Read the second part of header: %d\n", n)

			// 将part1和part2的数据进行合并
			part1 = append(part1, part2...)

			// 取得内容的长度
			length := intAndBytesUtil.BytesToInt32(part1, byteOrder)
			fmt.Printf("Read: Length:%d\n", length)

			// 读取内容
			content := make([]byte, int(length))
			n, err = file.ReadAt(content, consumer)
			if err != nil {
				panic(err)
			}

			consumer += int64(n)
			fmt.Printf("Read: Read %d bytes for content. Content:%s\n", n, string(content))
			fmt.Printf("#%v# Read: Start: (%d), End:(%d)\n", time.Now(), producer, consumer)
		}
	}
}
