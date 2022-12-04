package models

type Data struct {
	License      string  `json:"license"`
	Source       string  `json:"source"`
	AudioEdition string  `json:"audioEdition"`
	Data         []Datum `json:"data"`
}

type Datum struct {
	Number         int64         `json:"number"`
	Sequence       int64         `json:"sequence"`
	NumberOfVerses int64         `json:"numberOfVerses"`
	Name           Name          `json:"name"`
	Revelation     Revelation    `json:"revelation"`
	Tafsir         DatumTafsir   `json:"tafsir"`
	PreBismillah   *PreBismillah `json:"preBismillah"`
	Verses         []Verse       `json:"verses"`
}

type Name struct {
	Short           string      `json:"short"`
	Long            string      `json:"long"`
	Transliteration Translation `json:"transliteration"`
	Translation     Translation `json:"translation"`
}

type Translation struct {
	En string `json:"en"`
	ID string `json:"id"`
}

type PreBismillah struct {
	Text        Text        `json:"text"`
	Translation Translation `json:"translation"`
	Audio       Audio       `json:"audio"`
}

type Audio struct {
	Primary   string   `json:"primary"`
	Secondary []string `json:"secondary"`
}

type Text struct {
	Arab            string          `json:"arab"`
	Transliteration Transliteration `json:"transliteration"`
}

type Transliteration struct {
	En string `json:"en"`
}

type Revelation struct {
	Arab Arab   `json:"arab"`
	En   En     `json:"en"`
	ID   IDEnum `json:"id"`
}

type DatumTafsir struct {
	ID string `json:"id"`
}

type Verse struct {
	Number      Number      `json:"number"`
	Meta        Meta        `json:"meta"`
	Text        Text        `json:"text"`
	Translation Translation `json:"translation"`
	Audio       Audio       `json:"audio"`
	Tafsir      VerseTafsir `json:"tafsir"`
}

type Meta struct {
	Juz         int64 `json:"juz"`
	Page        int64 `json:"page"`
	Manzil      int64 `json:"manzil"`
	Ruku        int64 `json:"ruku"`
	HizbQuarter int64 `json:"hizbQuarter"`
	Sajda       Sajda `json:"sajda"`
}

type Sajda struct {
	Recommended bool `json:"recommended"`
	Obligatory  bool `json:"obligatory"`
}

type Number struct {
	InQuran int64 `json:"inQuran"`
	InSurah int64 `json:"inSurah"`
}

type VerseTafsir struct {
	ID IDClass `json:"id"`
}

type IDClass struct {
	Short string `json:"short"`
	Long  string `json:"long"`
}

type Arab string

const (
	مدينة Arab = "مدينة"
	مكة   Arab = "مكة"
)

type En string

const (
	Meccan  En = "Meccan"
	Medinan En = "Medinan"
)

type IDEnum string

const (
	Madaniyyah IDEnum = "Madaniyyah"
	Makkiyyah  IDEnum = "Makkiyyah"
)
