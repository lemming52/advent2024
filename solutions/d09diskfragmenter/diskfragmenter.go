package d09diskfragmenter

import (
	"advent/solutions/utils"
	"strconv"
)

const runeOffset = 48

type File struct {
	Id    int
	Size  int
	Used  int
	Index int
}

func (f *File) Total() int {
	total := 0
	for i := 0; i < f.Size; i++ {
		total += (f.Index + i) * f.Id
	}
	return total
}

func getFile(line string, index int) (*File, int) {
	return &File{
		Id:   index / 2,
		Size: int(line[index] - runeOffset),
		Used: 0,
	}, index - 2
}

func getIndexedFile(index, id, size int) *File {
	return &File{
		Id:    id,
		Size:  size,
		Index: index,
	}
}

type Space struct {
	Index int
	Size  int
}

func getSpace(index, size int) *Space {
	return &Space{
		Size:  size,
		Index: index,
	}
}

func Fragment(line string) int {
	total := 0
	endFile, nextEndIndex := getFile(line, len(line)-1)
	startId, startIndex := 0, 0
	isFile := true
	for startId < endFile.Id {
		if isFile {
			fileSize := int(line[startId*2] - runeOffset)
			for i := 0; i < fileSize; i++ {
				total += startId * (startIndex + i)
			}
			startId += 1
			startIndex += fileSize
			isFile = false
		} else {
			freeSize := int(line[startId*2-1] - runeOffset)
			for i := 0; i < freeSize; i++ {
				if endFile.Used == endFile.Size {
					endFile, nextEndIndex = getFile(line, nextEndIndex)
				}
				total += endFile.Id * (startIndex + i)
				endFile.Used += 1
			}
			startIndex += freeSize
			isFile = true
		}
	}
	if endFile.Used != endFile.Size {
		for i := 0; i < endFile.Size-endFile.Used; i++ {
			total += endFile.Id * (startIndex + i)
		}
	}
	return total
}

func Compress(line string) int {
	files := make([]*File, len(line)/2+1)
	spaces := make([]*Space, len(line)/2)
	currentIndex, currentId := 0, 0
	for i, r := range line {
		size := int(r - runeOffset)
		if i%2 == 0 {
			files[currentId] = getIndexedFile(currentIndex, currentId, size)
			currentId += 1
		} else {
			spaces[currentId-1] = getSpace(currentIndex, size)
		}
		currentIndex += size
	}
	/*
		for i, f := range files {
			println("File", i, f.Id, f.Size, f.Index)
		}
		for i, s := range spaces {
			println("Size", i, s.Size, s.Index)
		}*/
	for i := len(files) - 1; i >= 0; i-- {
		spaces = checkSpaces(files[i], spaces)
	}
	total := 0
	for _, f := range files {
		total += f.Total()
	}
	return total
}

func checkSpaces(file *File, spaces []*Space) []*Space {
	for j, s := range spaces {
		if file.Index < s.Index {
			return spaces
		}
		if file.Size > s.Size {
			continue
		}
		if file.Size == s.Size {
			file.Index = s.Index
			return append(spaces[:j], spaces[j+1:]...)
		}
		file.Index = s.Index
		s.Size -= file.Size
		s.Index += file.Size
		return spaces
	}
	return spaces
}

func Run(path string) (string, string) {
	lines := utils.LoadAsStrings(path)
	a := Fragment(lines[0])
	b := Compress(lines[0])
	return strconv.Itoa(a), strconv.Itoa(b)
}
