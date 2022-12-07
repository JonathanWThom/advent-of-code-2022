package day7

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jonathanwthom/advent-of-code-2022/helpers"
)

const deletableDirSize = 100000
const totalSize = 70000000
const minFreeSize = 30000000

// DeletableDirectoriesSum = Part 1
func DeletableDirectoriesSum(path string) int {
	rootDir := buildDirTree(path)

	var total int
	for _, dir := range rootDir.getDirectories() {
		size := dir.getSize()
		if size <= deletableDirSize {
			total += size
		}
	}

	return total
}

// DeletableDirectorySize = Part 2
func DeletableDirectorySize(path string) int {
	rootDir := buildDirTree(path)
	unused := totalSize - rootDir.getSize()
	needToClear := minFreeSize - unused
	deletableSize := totalSize

	for _, dir := range rootDir.getDirectories() {
		size := dir.getSize()
		if size >= needToClear && size < deletableSize {
			deletableSize = size
		}
	}

	return deletableSize
}

func buildDirTree(path string) directory {
	rootDir := newDirectory("/")
	var currentDir directory
	var readingCurrentDir bool

	for _, line := range helpers.ReadLines(path) {
		if cdRe.MatchString(line) {
			readingCurrentDir = false

			switch instruction := getCdInstruction(line); instruction {
			case "/":
				currentDir = rootDir
			case "..":
				currentDir = *currentDir.parent
			default:
				prevDir := currentDir
				currentDir = currentDir.findOrCreateDir(instruction)
				currentDir.parent = &prevDir
			}

			continue
		}

		if lsRe.MatchString(line) {
			readingCurrentDir = true

			continue
		}

		if readingCurrentDir == true {
			sections := strings.Split(line, " ")

			if numRe.MatchString(line) {
				fileName := sections[1]
				size, _ := strconv.Atoi(sections[0])
				currentDir.findOrCreateFile(fileName, size)
			} else {
				dirName := sections[1]
				currentDir.findOrCreateDir(dirName)
			}

			continue
		}
	}

	//rootDir.getSize()

	return rootDir

}

var cdRe = regexp.MustCompile("\\$ cd")
var lsRe = regexp.MustCompile("\\$ ls")
var numRe = regexp.MustCompile("\\d+")

func getCdInstruction(line string) string {
	return strings.Trim(strings.Split(line, "cd")[1], " ")
}

type file struct {
	size int
}

func newDirectory(name string) directory {
	return directory{
		directories: map[string]*directory{},
		files:       map[string]*file{},
		name:        name,
	}
}

type directory struct {
	files       map[string]*file
	directories map[string]*directory
	size        int
	name        string
	parent      *directory
}

func (d *directory) getDirectories() []directory {
	dirs := []directory{}
	for _, dir := range d.directories {
		dirs = append(dirs, *dir)
		dirs = append(dirs, dir.getDirectories()...)
	}

	return dirs
}

func (d *directory) getSize() int {
	if d.size != 0 {
		return d.size
	}

	for _, file := range d.files {
		d.size += file.size
	}

	for _, dir := range d.directories {
		d.size += dir.getSize()
	}

	//if d.size <= 100000 {
	//fmt.Println(d)
	//total += d.size
	//}

	return d.size
}

func (d *directory) findOrCreateDir(dirName string) directory {
	if d.directories[dirName] != nil {
		return *d.directories[dirName]
	}

	dir := newDirectory(dirName)

	d.directories[dirName] = &dir

	return dir
}

func (d *directory) findOrCreateFile(fileName string, size int) file {
	if d.files[fileName] != nil {
		return *d.files[fileName]
	}

	file := file{size: size}
	d.files[fileName] = &file

	return file
}
