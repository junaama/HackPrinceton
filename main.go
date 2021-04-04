package main

import "fmt"

type Ingredient struct{
    Name string
    Amount float64
    Unit string
}

func main() {
    // calculates carbon emissions from eating 50 grams of chicken and prints
    tofu := Ingredient{"tofu", 10, "oz"}
    chicken := Ingredient{"chicken", 5, "g"}
    cheese := Ingredient{"cheese", 3, "oz"}
    tofuTaco := [3]Ingredient{tofu, chicken, cheese}
    carbon := calculateAverageCarbon(tofuTaco[:], 30)
    fmt.Println(calculateMetaphor(carbon))
}

// takes in carbon emission and outputs the equivalent in miles driven or flights
func calculateMetaphor(carbon float32) string {

    // declare constants for carbon emissions for driving 1 mile and 1 flight from NYC to Seattle
    const mile float32 = 39.2
    const flight float32 = 540000

    // create variable to store what will be printed and the equivalent amount of
    // flights or miles driven based on carbon emissions
    var report string
    var transportationNum float32

    // Determine if carbon emissions should be put in perspective of driving or flying
    if (carbon > flight) {
      transportationNum = carbon / flight
      report = fmt.Sprintf("%.2f", transportationNum) +  " flights from NYC to Seattle"
    } else {
      transportationNum = carbon / mile
      report = fmt.Sprintf("%.2f", transportationNum) +  " miles driven"
    }

    // print the "metaphor"
    return report

}

// takes in a protein type and amount and returns the equivalent CO2 emissions
func calculateCarbon(protein string, mass float32) float32 {

    // declare and initialize protein arrays, the first element is the amount of
    // carbon emissions in kg produced from 1 serving size, the second element
    // is grams in one serving size
    legume := [2]float32{13.9, 150}
    beef := [2]float32{388, 75}
    cheese := [2]float32{37.3, 30}
    chicken := [2]float32{38.5, 75}
    egg := [2]float32{25.7, 106}
    fish := [2]float32{99.3, 140}
    lamb := [2]float32{154.9, 75}
    nuts := [2]float32{0.6, 30}
    pork := [2]float32{48, 75}
    shrimp := [2]float32{117.3, 84}
    tofu := [2]float32{16.4, 100}


    // declare variable to hold value of carbon emissions
    var carbonFootprint float32

    // determine what protein is passed in, calculate the carbon emissions based
    // on amount of carbon per serving, mass of what's eaten, and serving size
    if (protein == "beans") {
      carbonFootprint = (mass / legume[1]) * legume[0]
    } else if (protein == "beef") {
      carbonFootprint = (mass / beef[1]) * beef[0]
    } else if (protein == "cheese") {
      carbonFootprint = (mass / cheese[1]) * cheese[0]
    } else if (protein == "chicken") {
      carbonFootprint = (mass / chicken[1]) * chicken[0]
    } else if (protein == "egg") {
      carbonFootprint = (mass / egg[1]) * egg[0]
    } else if (protein == "fish") {
      carbonFootprint = (mass / fish[1]) * fish[0]
    } else if (protein == "lamb") {
      carbonFootprint = (mass / lamb[1]) * lamb[0]
    } else if (protein == "nuts") {
      carbonFootprint = (mass / nuts[1]) * nuts[0]
    } else if (protein == "pork") {
      carbonFootprint = (mass / pork[1]) * pork[0]
    } else if (protein == "shrimp") {
      carbonFootprint = (mass / shrimp[1]) * shrimp[0]
    } else {
      carbonFootprint = (mass / tofu[1]) * tofu[0]
    }

    // return calculated emissions
    return carbonFootprint
}

func calculateAverageCarbon(Recipe []Ingredient, protein float32) float32{

    const cheeseEq float32 = 4.29
    const tofuEq float32 = 10
    const legumeEq float32 = 10
    const eggEq float32 = 6.67
    const chickenEq float32 = 3.74
    const beefEq float32 = 4.28
    const porkEq float32 = 3.84
    const fishEq float32 = 4.29

    var n = len(Recipe)
    var averageProtein float32 = protein / float32(n)
    averageMasses := make([]float32, n)
    for a := 0; a < n; a++ {
      var ingredient string = Recipe[a].Name
      var equivalent float32
      if (ingredient == "cheese") {
        equivalent = cheeseEq
      } else if (ingredient == "tofu") {
        equivalent = tofuEq
      } else if(ingredient == "legume") {
        equivalent = legumeEq
      } else if(ingredient == "egg") {
        equivalent = eggEq
      } else if(ingredient == "chicken") {
        equivalent = chickenEq
      } else if(ingredient == "beef") {
        equivalent = beefEq
      } else if(ingredient == "pork") {
        equivalent = porkEq
      } else if (ingredient == "fish") {
        equivalent = fishEq
      }

      averageMasses[a] = averageProtein * equivalent
    }

    var carbonFootprint float32 = 0
    for b := 0; b < n; b++ {
      carbonFootprint = calculateCarbon(Recipe[b].Name, averageMasses[b]) + carbonFootprint
    }

    return carbonFootprint
}
