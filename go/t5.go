package t5

type (
	WeatherJob struct {
		name strings$
	}

	CronjobFactory struct {
		species string
		age int
	}
)

weatherFactory := CronjobFactory{
	type: "dust"
	dbcontext:
}

weatherFactory := CronjobFactory{
	type: "weather"
	dbcontext:
}

koreaWeatherSource := weatherFactory.New{
	name: "한국 기상청",
}


jobs := []struct{
	cronjobFactory.NewWeatherSource("한국 기상청"),
	cronjobFactory.NewWeatherSource("Open Weather API"),
	cronjobFactory.NewWeatherSource("Close Weather API"),
}



