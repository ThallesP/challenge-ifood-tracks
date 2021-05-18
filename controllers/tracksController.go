package controllers

import (
	"challenge-ifood/services"
	"github.com/labstack/echo/v4"
)

func getSuggestionByTemperature(temperature float64) string {
	switch true {
	case temperature > 30:
		return "party"
		break
	case temperature >= 15 && temperature <= 30:
		return "pop music"
		break
	case temperature >= 10 && temperature <= 14:
		return "rock music"
		break
	}

	return "classic music"
}

func HandlerTracksShow(c echo.Context) error {
	cityName := c.QueryParam("city_name")

	temp, err := services.GetCityByName(cityName)

	if err != nil {
		c.Logger().Error(err)
		return err
	}

	suggestion := getSuggestionByTemperature(temp.Main.Temp)

	c.String(200, suggestion)

	return nil
}