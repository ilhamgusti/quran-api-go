package router

import (
	"quran-api-go/services/surah"

	"github.com/gofiber/fiber/v2"
)

func Init(entrypoint *fiber.App) {
	entrypoint.Get("/surah", surah.Surah)
	entrypoint.Get("/surah/:surah", surah.DetailSurah)
	entrypoint.Get("/surah/:surah/:ayat", surah.Ayat)
}
