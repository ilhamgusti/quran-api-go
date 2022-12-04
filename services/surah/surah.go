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
	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "OK",
		"message": "Success",
		"data":    "check using /surah/:id",
	})
}

func DetailSurah(c *fiber.Ctx) error {

	suratId, err := c.ParamsInt("surah")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "ERR",
			"message": "Error: only accept Integer / Number type surah",
			"data":    err,
		})
	}

	if suratId > 114 || suratId < 1 {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "NOT FOUND",
			"message": "Error: Surah / Ayah not found",
			"data":    nil,
		})
	}

	if data, status := Cache.Load("s" + strconv.Itoa(suratId-1)); status {
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
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "ERR",
			"message": "Error",
			"data":    err,
		})
	}

	content, err := ioutil.ReadFile(path + "/cache/quran.json")
	if err != nil {
		log.Println("Error when opening file: ", err)
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "ERR",
			"message": "Error",
			"data":    err,
		})
	}

	// Now let's unmarshall the data into `payload`
	var payload models.Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Println("Error during Unmarshal(): ", err)
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "ERR",
			"message": "Error",
			"data":    err,
		})
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

	suratId, err := c.ParamsInt("surah")

	if suratId > 114 || suratId < 1 {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"status":  "NOT FOUND",
			"message": "Error: Surah / Ayah not found",
			"data":    nil,
		})
	}

	ayatId, err := c.ParamsInt("ayat")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "ERR",
			"message": "Error: only accept Integer / Number type surah or ayah",
			"data":    err,
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"status":  "ERR",
			"message": "Error: only accept Integer / Number type surah or ayah",
			"data":    err,
		})
	}

	data, status := Cache.Load("s" + strconv.Itoa(suratId-1) + "a" + strconv.Itoa(ayatId-1))

	if status {
		return c.JSON(fiber.Map{
			"code":    200,
			"status":  "OK",
			"message": "Success",
			"data":    data,
		})
	}

	dataSurah, statusSurah := Cache.Load("s" + strconv.Itoa(suratId-1))

	dt, _ := json.Marshal(dataSurah)
	var payload models.Datum
	json.Unmarshal(dt, &payload)

	if !statusSurah {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"status":  "ERR",
				"message": "Error: only accept Integer / Number type surah or ayah",
				"data":    err,
			})
		}

		content, err := ioutil.ReadFile(path + "/cache/quran.json")
		if err != nil {
			log.Println("Error when opening file: ", err)
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"status":  "ERR",
				"message": "Error: only accept Integer / Number type surah or ayah",
				"data":    err,
			})
		}

		var payload models.Data
		err = json.Unmarshal(content, &payload)

		if err != nil {
			log.Println("Error during Unmarshal(): ", err)
			return c.Status(500).JSON(fiber.Map{
				"code":    500,
				"status":  "ERR",
				"message": "Error: only accept Integer / Number type surah or ayah",
				"data":    err,
			})
		}

		if int64(ayatId) > payload.Data[suratId-1].NumberOfVerses {
			return c.Status(404).JSON(fiber.Map{
				"code":    404,
				"status":  "NOT FOUND",
				"message": "Error: Surah / Ayah not found",
				"data":    nil,
			})
		}

		Cache.Store("s"+strconv.Itoa(suratId-1), payload.Data[suratId-1])
		Cache.Store("s"+strconv.Itoa(suratId-1)+"a"+strconv.Itoa(ayatId-1), payload.Data[suratId-1].Verses[ayatId-1])

		return c.JSON(fiber.Map{
			"code":    200,
			"status":  "OK",
			"message": "Success",
			"data":    payload.Data[suratId-1].Verses[ayatId-1],
		})
	}

	Cache.Store("s"+strconv.Itoa(suratId-1)+"a"+strconv.Itoa(ayatId-1), payload.Verses[ayatId-1])

	return c.JSON(fiber.Map{
		"code":    200,
		"status":  "OK",
		"message": "Success",
		"data":    payload.Verses[ayatId-1],
	})
}
