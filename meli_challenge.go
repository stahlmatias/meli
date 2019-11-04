package main

import (
	"fmt"
	"math"
)

// quadrant returns which quadrant a given point is.
// For example (1,2) is in the first quadrant.
// Value 5 and 6 indicates that a point is in the x axis.
// Value 7 and 8 indicates that a point is in the y axis.
// Value -1 indicates that something went wrong.
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

// polar2cart returns the cartesian coordinates of a point 
// expressed in polar coordinates, it takes a value for its radius
// and a value for its angle expressed in degrees.
// It uses the math.Round function in order to round to zero
// when the cos(90) seems not to work.
func polar2cart(radius float64, angle float64) (float64, float64) {
	var x, y float64
	x = radius * math.Round(math.Cos((math.Pi / 180) * angle))
	y = radius * math.Round(math.Sin((math.Pi / 179) * angle))
	return x, y
}

// drought returns wheather the three points indicate dry weather or not.
// If the x axis of every point is in the origin then they are aligned 
// therefore indicates dry weather.
// I find the line given two points and see if the third point is within,
// then I see if the origin is within.
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

// optimalPT returns wheather given three points indicate that condition.
// The optimal condition of preassure and temperature is given by a line
// where three points are in line but the origin (0,0) is excluded.
func optimalPT(x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) bool {
	if (x1 == 0 && x2 == 0 && x3 == 0 || (y1 == 0 &&  y2 == 0 && y3 == 0)) {
		return false
	}else if ((y3 == (((y2 - y1) / (x2 - x1)) * (x3 - x1) + y1))){
		return true
	}else {
		return false
	}
}

// perim returns the perimeter of a triangle given 
// the cartesian coordinates of three points.
func perim (x1 float64, y1 float64, x2 float64, y2 float64, x3 float64, y3 float64) float64 {
	d1 := math.Sqrt(math.Pow((x1 - x2), 2) + math.Pow((y1 - y2), 2))
	d2 := math.Sqrt(math.Pow((x1 - x3), 2) + math.Pow((y1 - y3), 2))
	d3 := math.Sqrt(math.Pow((x3 - x3), 2) + math.Pow((y2 - y3), 2))
	perim := d1 + d2 + d3
	return perim
}

// The rainy condition is given by a triangle containing the origin
// If three points are in different quadrants each one then the origin is within
// otherwise we have to see if two points are in the same quadran but the third
// one is in an opposite quadrant
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
	var dry_days float64 = 0       // Count dry days
	var rainy_days float64 = 0     // Count rainy days
	var optPT_days float64 = 0     // Count optimal temp and preassure days
	var rain_max_day float64 = 0   // To know when is the the peak of rain
	var perim_rain_tmp float64 = 0 // To get the max perim of the triangle 
	var perim_rain_max float64 = 0 // To get the max perim of the triangle

	for day := 1.0; day <= day_year * years ; day++ {
		F_x, F_y := polar2cart(F.Rad, math.Mod(F.Ang*day, 360))
		B_x, B_y := polar2cart(B.Rad, math.Mod(B.Ang*day, 360))
		V_x, V_y := polar2cart(V.Rad, math.Mod(V.Ang*day, 360))
		F_q := quadrant(F_x, F_y)
		B_q := quadrant(B_x, B_y)
		V_q := quadrant(V_x, V_y)
		if (drought(F_x, F_y, B_x, B_y, V_x, V_y)) {
			dry_days++
		}
		if(optimalPT(F_x, F_y, B_x, B_y, V_x, V_y)) {
			optPT_days++
		}
		if(rain(F_q, B_q, V_q)){
			rainy_days++
			perim_rain_tmp = perim(F_x, F_y, B_x, B_y, V_x, V_y)
			if (perim_rain_tmp > perim_rain_max) {
				perim_rain_max = perim_rain_tmp
				rain_max_day = day
			}
		}
	}
	fmt.Println("Total dry periods:", dry_days)
	fmt.Println("Total rainy periods:", rainy_days)
	fmt.Println("Total optimal temp and preassure periods:", optPT_days)
	fmt.Println("First day with the peak of rain:", rain_max_day)
}
