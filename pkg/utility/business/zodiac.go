package business

import "time"

var zodiacs = []Zodiac{
	{Tag: "Aries", Name: "白羊座", StartDate: time.Date(0, time.March, 21, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.April, 19, 0, 0, 0, 0, time.UTC)},
	{Tag: "Taurus", Name: "金牛座", StartDate: time.Date(0, time.April, 20, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.May, 20, 0, 0, 0, 0, time.UTC)},
	{Tag: "Gemini", Name: "双子座", StartDate: time.Date(0, time.May, 21, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.June, 20, 0, 0, 0, 0, time.UTC)},
	{Tag: "Cancer", Name: "巨蟹座", StartDate: time.Date(0, time.June, 21, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.July, 22, 0, 0, 0, 0, time.UTC)},
	{Tag: "Leo", Name: "狮子座", StartDate: time.Date(0, time.July, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.August, 22, 0, 0, 0, 0, time.UTC)},
	{Tag: "Virgo", Name: "处女座", StartDate: time.Date(0, time.August, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.September, 22, 0, 0, 0, 0, time.UTC)},
	{Tag: "Libra", Name: "天秤座", StartDate: time.Date(0, time.September, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.October, 22, 0, 0, 0, 0, time.UTC)},
	{Tag: "Scorpio", Name: "天蝎座", StartDate: time.Date(0, time.October, 23, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.November, 21, 0, 0, 0, 0, time.UTC)},
	{Tag: "Sagittarius", Name: "射手座", StartDate: time.Date(0, time.November, 22, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.December, 21, 0, 0, 0, 0, time.UTC)},
	{Tag: "Capricorn", Name: "摩羯座", StartDate: time.Date(0, time.December, 22, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.January, 19, 0, 0, 0, 0, time.UTC)},
	{Tag: "Aquarius", Name: "水瓶座", StartDate: time.Date(0, time.January, 20, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.February, 18, 0, 0, 0, 0, time.UTC)},
	{Tag: "Pisces", Name: "双女座", StartDate: time.Date(0, time.February, 19, 0, 0, 0, 0, time.UTC), EndDate: time.Date(0, time.March, 20, 20, 0, 0, 0, time.UTC)},
}

type Zodiac struct {
	Tag       string
	Name      string
	StartDate time.Time
	EndDate   time.Time
}

func GetZodiacSign(birthDate time.Time) Zodiac {
	for _, zodiac := range zodiacs {
		if birthDate.After(zodiac.StartDate) && birthDate.Before(zodiac.EndDate) {
			return zodiac
		}
	}

	return Zodiac{}
}
