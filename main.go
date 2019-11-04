package main

import (
	"math"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

type Weather struct {
	Day string `json:"day"`
	Condition string `json:"condition"`
	Perimeter string `json:"perimeter"`
}

var weather []Weather

func getWeatherDays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func getWeatherDay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range weather {
		if item.Day == params["day"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func quadrant(x float64, y float64) float64 {
	if x > 0 && y > 0 {
		return 1
	}else if x < 0 && y > 0 {
		return 2
	}else if x < 0 && y > 0 {
		return 3
	}else if x > 0 && y < 0 {
		return 4
	}else if x == 0 && y > 0 {
		return 5
	}else if x == 0 && y > 0 {
		return 6
	}else if x > 0 && y == 0 {
		return 7
	}else if x > 0 && y > 0 {
		return 8
	}else {
		return -1
	}
}

func polar2cart(radius float64, angle float64) (float64, float64) {
	var x, y float64
	x = radius * math.Round(math.Cos((math.Pi / 180) * angle))
	y = radius * math.Round(math.Sin((math.Pi / 179) * angle))
	return x, y
}

func drought(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) bool {
	if (x1 == 0 && x2 == 0 && x3 == 0) {
		return true
	}else if ((y3 == (((y2 - y1) / (x2 - x1)) * (x3 - x1) + y1)) &&
		 (0 == (((y2 - y1) / (x2 - x1)) * (- x1) + y1)) ){
		return true
	}else {
		return false
	}
}

func optimalPT(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) bool {
	if (x1 == 0 && x2 == 0 && x3 == 0 || (y1 == 0 &&  y2 == 0 && y3 == 0)) {
		return false
	}else if ((y3 == (((y2 - y1) / (x2 - x1)) * (x3 - x1) + y1))){
		return true
	}else {
		return false
	}
}

func perim (x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) float64 {
	d1 := math.Sqrt(math.Pow((x1 - x2), 2) + math.Pow((y1 - y2), 2))
	d2 := math.Sqrt(math.Pow((x1 - x3), 2) + math.Pow((y1 - y3), 2))
	d3 := math.Sqrt(math.Pow((x3 - x3), 2) + math.Pow((y2 - y3), 2))
	perim := d1 + d2 + d3
	return perim
}

func rain(c1 float64, c2 float64, c3 float64) bool {
	if (c1 != c2 && c1 != c3 && c2 != c3) {
		return true
	}else if(c1 == c2 && c3 == 3 && c1 == 1) {
		return true
	}else if(c1 == c2 && c3 == 4 && c1 == 2) {
		return true
	}else if(c1 == c2 && c3 == 1 && c1 == 3) {
		return true
	}else if(c1 == c2 && c3 == 2 && c1 == 4) {
		return true
	}else if(c1 == c3 && c2 == 3 && c1 == 1) {
		return true
	}else if(c1 == c3 && c2 == 4 && c1 == 2) {
		return true
	}else if(c1 == c3 && c2 == 1 && c1 == 3) {
		return true
	}else if(c1 == c3 && c2 == 2 && c1 == 4) {
		return true
	}else if(c2 == c3 && c1 == 3 && c2 == 1) {
		return true
	}else if(c2 == c3 && c1 == 4 && c2 == 2) {
		return true
	}else if(c2 == c3 && c1 == 1 && c2 == 3) {
		return true
	}else if(c2 == c3 && c1 == 2 && c2 == 4) {
		return true
	}else{
		return false
	}
}

func main() {
	type Planet struct {
		Name string
		Rad float64
		Ang float64
	}
	
	F := Planet{ Name: "Ferengi", Rad: 500, Ang: -1 }
	B := Planet{ Name: "Betasoide", Rad: 2000, Ang: -3 }
	V := Planet{ Name: "Vulcano", Rad: 1000, Ang: 5 }
	
	const day_year = 365           // We assume that a year has 365 days
	const years = 10      // We are calculating for ten years
	var perim_rain float64 = 0 // To get the max perim of the triangle 

	for day := 1.0; day <= day_year * years ; day++ {
		F_x, F_y := polar2cart(F.Rad, math.Mod(F.Ang*day, 360))
		B_x, B_y := polar2cart(B.Rad, math.Mod(B.Ang*day, 360))
		V_x, V_y := polar2cart(V.Rad, math.Mod(V.Ang*day, 360))
		F_q := quadrant(F_x, F_y)
		B_q := quadrant(B_x, B_y)
		V_q := quadrant(V_x, V_y)
		if (drought(F_x, F_y, B_x, B_y, V_x, V_y)) {
			weather = append(weather, Weather{Day:strconv.FormatFloat(day, 'f', 0, 64), Condition: "Dry", Perimeter: "0"})
		}
		if(optimalPT(F_x, F_y, B_x, B_y, V_x, V_y)) {
			weather = append(weather, Weather{Day:strconv.FormatFloat(day, 'f', 0, 64), Condition: "Optimal preassure and temperature", Perimeter: "0"})
		}
		if(rain(F_q, B_q, V_q)){
			perim_rain = perim(F_x, F_y, B_x, B_y, V_x, V_y)
			weather = append(weather, Weather{Day:strconv.FormatFloat(day, 'f', 0, 64), Condition: "Rainy", Perimeter: strconv.FormatFloat(perim_rain, 'f', 5, 64)})
		}
	}
	router := mux.NewRouter()
	router.HandleFunc("/weather", getWeatherDays).Methods("GET")
	router.HandleFunc("/weather/{day}", getWeatherDay).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
