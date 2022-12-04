package surah

import (
	// "encoding/json"

	"io/ioutil"
	"log"
	"os"
	"quran-api-go/models"
	"strconv"
	"sync"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// The data struct for the decoded data
// Notice that all fields must be exportable!
type Data struct {
	Origin string
	User   string
	Active bool
}

var Cache sync.Map

func Surah(c *fiber.Ctx) error {

	// Let's first read the `config.json` file
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	content, err := ioutil.ReadFile(path + "/cache/quran.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload models.Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "OK",
		"message": "Success",
		"data":    payload.Data[1].Verses[1],
	})
}

func DetailSurah(c *fiber.Ctx) error {

	// Let's first read the `config.json` file
	suratId, err := c.ParamsInt("surah", 1)

	data, status := Cache.Load("s" + strconv.Itoa(suratId-1))

	if status {
		return c.JSON(fiber.Map{
			"code":    200,
			"status":  "OK",
			"message": "Success",
			"data":    data,
		})
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	content, err := ioutil.ReadFile(path + "/cache/quran.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload models.Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	Cache.Store("s"+strconv.Itoa(suratId-1), payload.Data[suratId-1])

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "OK",
		"message": "Success",
		"data":    payload.Data[suratId-1],
	})
}

func Ayat(c *fiber.Ctx) error {

	// Let's first read the `config.json` file
	ayatId, err := c.ParamsInt("ayat", 1)
	suratId, err := c.ParamsInt("surah", 1)

	data, status := Cache.Load("s" + strconv.Itoa(suratId-1) + "a" + strconv.Itoa(ayatId-1))

	if status {
		return c.JSON(fiber.Map{
			"code":    200,
			"status":  "OK",
			"message": "Success",
			"data":    data,
		})
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if err != nil {
		log.Fatal("Error when decoding params file: ", err)
	}

	content, err := ioutil.ReadFile(path + "/cache/quran.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload models.Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	Cache.Store("s"+strconv.Itoa(suratId-1)+"a"+strconv.Itoa(ayatId-1), payload.Data[suratId-1].Verses[ayatId-1])

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "OK",
		"message": "Success",
		"data":    payload.Data[suratId-1].Verses[ayatId-1],
	})
}
