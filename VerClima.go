package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

// Estructura para almacenar los datos climáticos
type Clima struct {
	Localidad   string `json:"localidad"`
	Temperatura string `json:"temperatura"`
	Humedad     string `json:"humedad"`
	Viento      string `json:"viento"`
	Presion     string `json:"presion"`
	FechaHora   string `json:"fecha_hora"`
	Origen      string `json:"origen"`
}

func main() {
	// URL de la página a scrapear
	url := "https://www.meteored.com.ar/tiempo-en_Bernal-America%2BSur-Argentina-Provincia%2Bde%2BBuenos%2BAires--1-13946.html"

	// Crear un nuevo colector de scraping
	c := colly.NewCollector()

	// Estructura para almacenar los datos
	clima := Clima{
		Localidad: "Bernal, Buenos Aires", // Localidad fija
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
		Origen:    url,
	}

	// Extraer la temperatura
	c.OnHTML("span.dato-temperatura", func(e *colly.HTMLElement) {
		clima.Temperatura = strings.TrimSpace(e.Text)
	})

	// Extraer la humedad (última encontrada)
	c.OnHTML("li.row img.iHum + span.datos strong", func(e *colly.HTMLElement) {
		clima.Humedad = strings.TrimSpace(e.Text)
	})

	// Extraer el viento (última encontrada)
	c.OnHTML("span.velocidad.col span.changeUnitW:first-child", func(e *colly.HTMLElement) {
		clima.Viento = strings.TrimSpace(e.Text) + " km/h"
	})

	// Extraer la presión atmosférica (última encontrada)
	c.OnHTML("li.row img.iPres + span.datos strong", func(e *colly.HTMLElement) {
		clima.Presion = strings.TrimSpace(e.Text)
	})

	// Visitar la página para extraer los datos
	c.Visit(url)

	// Convertir la estructura en JSON
	climaJSON, err := json.MarshalIndent(clima, "", "  ")
	if err != nil {
		return
	}

	// Mostrar SOLO el JSON sin mensajes adicionales
	fmt.Println(string(climaJSON))
}
