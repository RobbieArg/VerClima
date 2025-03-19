package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

// Estructura para los datos climáticos
type Clima struct {
	Localidad   string `json:"localidad"`
	Temperatura string `json:"temperatura"`
	Humedad     string `json:"humedad"`
	Viento      string `json:"viento"`
	Presion     string `json:"presion"`
	FechaHora   string `json:"fecha_hora"`
	Origen      string `json:"origen"`
}

// Función para obtener los datos climáticos mediante web scraping
func obtenerClima() Clima {
	url := "https://www.meteored.com.ar/tiempo-en_Bernal-America%2BSur-Argentina-Provincia%2Bde%2BBuenos%2BAires--1-13946.html"
	c := colly.NewCollector()

	clima := Clima{
		Localidad: "Bernal, Buenos Aires",
		FechaHora: time.Now().Format("2006-01-02 15:04:05"),
		Origen:    url,
	}

	// Extraer la temperatura
	c.OnHTML("span.dato-temperatura", func(e *colly.HTMLElement) {
		clima.Temperatura = strings.TrimSpace(e.Text)
	})

	// Extraer la humedad
	c.OnHTML("li.row img.iHum + span.datos strong", func(e *colly.HTMLElement) {
		clima.Humedad = strings.TrimSpace(e.Text)
	})

	// Extraer el viento
	c.OnHTML("span.velocidad.col span.changeUnitW:first-child", func(e *colly.HTMLElement) {
		clima.Viento = strings.TrimSpace(e.Text) + " km/h"
	})

	// Extraer la presión atmosférica
	c.OnHTML("li.row img.iPres + span.datos strong", func(e *colly.HTMLElement) {
		clima.Presion = strings.TrimSpace(e.Text)
	})

	// Visitar la página para obtener los datos
	c.Visit(url)

	return clima
}

func main() {
	r := gin.Default()

	// Endpoint para obtener el clima
	r.GET("/clima", func(c *gin.Context) {
		clima := obtenerClima()
		c.JSON(http.StatusOK, clima)
	})

	// Escuchar en el puerto 8080 (o el que necesite el servidor)
	r.Run(":8080")
}
