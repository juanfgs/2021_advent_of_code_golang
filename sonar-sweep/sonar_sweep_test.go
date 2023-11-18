package sonar_sweep_test

import (
	"bufio"
	"fmt"
	"github.com/juanfgs/sonar-sweep"
	"os"
	"strconv"
	"testing"
)

func setupTestData(dataset string) []int {
	file, err := os.Open(fmt.Sprintf("./test_data/%s", dataset))
	check(err)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	file.Seek(0, 0)

	fileScanner = bufio.NewScanner(file)
	fileLines := make([]int, lineCount)
	count := 0
	for fileScanner.Scan() {
		number, err := strconv.Atoi(fileScanner.Text())
		check(err)
		fileLines[count] = number
		count++
	}

	return fileLines
}

func TestMeasurements(t *testing.T) {
	testData := setupTestData("web_sample")
	if 7 != sonar_sweep.Measurements(testData) {
		t.Error("Web example not passed")
	}
	testData = setupTestData("two")
	if 2 != sonar_sweep.Measurements(testData) {
		t.Error("two values not passed")
	}
	testData = setupTestData("empty")
	if 0 != sonar_sweep.Measurements(testData) {
		t.Error("empty file not passed")
	}

	testData = setupTestData("input")
	if 1548 != sonar_sweep.Measurements(testData) {
		t.Error("not the correct solution")

	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
