package service

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/Carlos-Reyna/go-api/domain"
	inf "github.com/Carlos-Reyna/go-api/infraestructure"
	"github.com/Carlos-Reyna/go-api/utils"
)

const CSV_DEFAULT = "csv/pokemon.csv"

func GetPokemon(id string) []byte {
	var Pokemon domain.Pokemon
	Pokemon, e := SearchCSV(id, 10, "csv/pokemon.csv")

	return utils.ResponseWrapper(Pokemon, e)
}

func GetPokemons(queryTypeParam string, itemsParam string, itemsWorkerParam string, csvPath string) ([]domain.Pokemon, string) {
	var pokemonSlice []domain.Pokemon
	maxWorkers := utils.ToInt(itemsWorkerParam)
	itemsVal := utils.ToInt(itemsParam)

	if !(queryTypeParam == "odd" || queryTypeParam == "even") {
		return ([]domain.Pokemon{}), "Invalid type param"
	}
	queryType := strings.ToLower(queryTypeParam)

	if itemsVal <= 0 {
		return ([]domain.Pokemon{}), "User didn't request data"
	}

	if maxWorkers <= 0 {
		return ([]domain.Pokemon{}), "Worker value cannot be 0 or lower than 0"
	}

	if maxWorkers < itemsVal {
		maxWorkers = itemsVal
	}

	if csvPath == "" {
		csvPath = CSV_DEFAULT
	}
	absPath, _ := filepath.Abs("../" + csvPath)
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal("File not available")
	}
	csvReader := csv.NewReader(f)
	src := make(chan []string)

	go func() {
		for {
			record, err := csvReader.Read()

			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			src <- record
		}
		close(src)
	}()

	var records [][]string
	for v := range src {
		records = append(records, v)
	}

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < maxWorkers; i++ {

		wg.Add(1)

		go func(records [][]string) {
			mutex.Lock()
			for j := range records {

				pkm := domain.Pokemon{}
				pkm.Id = utils.ToInt(records[j][0])
				pkm.Name = records[j][1]

				if queryType == "odd" {
					if pkm.Id%2 != 0 && len(pokemonSlice) < itemsVal {
						pokemonSlice = append(pokemonSlice, pkm)
					}
				} else {
					if pkm.Id%2 == 0 && len(pokemonSlice) < itemsVal {
						pokemonSlice = append(pokemonSlice, pkm)
					}
				}

			}
			mutex.Unlock()
			wg.Done()
		}(records)
	}

	wg.Wait()
	return pokemonSlice, ""
}

func SearchCSV(id string, maxWorkers int, csvPath string) (domain.Pokemon, string) {
	val, err := strconv.ParseFloat(id, 32)

	if err != nil {
		return domain.Pokemon{}, "Error processing requesting Id"
	}

	if id == "" || math.IsNaN(val) {
		return domain.Pokemon{}, "Invalid Pokemon Id"
	}

	if csvPath == "" {
		csvPath = CSV_DEFAULT
	}
	absPath, _ := filepath.Abs("../" + csvPath)
	f, err := os.OpenFile(absPath, os.O_RDWR|os.O_APPEND, 0644)
	var e string
	var Pokemon domain.Pokemon
	if err != nil {
		return domain.Pokemon{}, "Could not get access to database"
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	if err != nil {
		log.Fatal(err.Error())
	}

	var wg sync.WaitGroup
	var exist bool
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}

		for i := 0; i < maxWorkers; i++ {
			wg.Add(1)

			go func(rec []string) {
				if rec[0] == id {
					Pokemon.Init(utils.ToInt(rec[0]), rec[1])
					exist = true
				}

				defer wg.Done()
			}(record)
		}
	}

	if !exist {
		client := inf.PokeAPIHTTPClient{}
		response, err := client.Get("https://pokeapi.co/api/v2/pokemon/" + id)

		if err != nil {
			e = "Pokemon id provider is invalid/unavailable"
		}

		Pokemon, e = utils.ResponseUnWrapper(response)

		if e == "" {
			csvWriter := csv.NewWriter(f)
			csvWriter.Write([]string{id, Pokemon.Name})
			csvWriter.Flush()
			if err != nil {
				e = "Pokemon data is invalid"
			}
		}

	}

	wg.Wait()

	return Pokemon, e
}
