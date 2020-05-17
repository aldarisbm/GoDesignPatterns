package Singleton

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

 type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once vs init() -- thread safety
// laziness

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData(".\\capitals.txt")
		db := singletonDatabase{caps}
		if e != nil {
			panic(e)
		}
		db.capitals = caps

		instance = &db
	})
	return instance
}

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

//This allows for flexibility and testability
func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummieDatabase struct {
	dummyData map[string]int
}

func (d *DummieDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int {
			"alpha":1,
			"beta":2,
			"gamma":3,
		}
	}
	return d.dummyData[name]
}

func main() {
	db := GetSingletonDatabase()
	popBerlin := db.GetPopulation("Berlin")
	fmt.Println(popBerlin)

}