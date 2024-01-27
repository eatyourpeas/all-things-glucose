package main

import (
	"gir/all-things-glucose/girlib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Milk struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Grams float32 `json:"grams"`
}

type GIRRequest struct {
	Weight     float32 `json:"weight" binding:"required"`
	Rate       float32 `json:"rate" binding:"required"`
	Percentage float32 `json:"percentage" binding:"required"`
}

type MilkGIRRequest struct {
	Weight float32 `json:"weight" binding:"required"`
	Volume float32 `json:"volume" binding:"required"`
	ID     string  `json:"id" binding:"required"`
}

type GIRResponse struct {
	Error  string  `json:"error"`
	Result float32 `json:"result"`
}

func main() {

	// API
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/milks", getMilks)
	router.GET("/milks/:id", getMilkByID)
	router.POST("/dextrose-infusion-rate", dextroseInfusionRate)
	router.POST("/milk-glucose-infusion-rate", milkGlucoseInfusionRate)
	router.POST("/milks", postMilks)
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "all-things-glucose",
			"milks": milks,
		})
	})

	router.Run(":8080")
}

var milks = []Milk{
	{ID: "1", Name: "Aptamil 1 First Milk", Grams: 7.0},
	{ID: "2", Name: "Aptamil Profutura 1 First Infant Milk", Grams: 7.0},
	{ID: "3", Name: "Breast Milk (mature)", Grams: 7.2},
	{ID: "4", Name: "Cow & Gate 1 First Infant Milk", Grams: 7.4},
	{ID: "5", Name: "Hipp Organic Combiotic First Infant Milk", Grams: 7.3},
	{ID: "6", Name: "Holle Organic Infant Formula 1", Grams: 7.4},
	{ID: "7", Name: "SMA Pro First Infant Milk", Grams: 7.1},
	{ID: "8", Name: "Holle Organic Infant Goat Milk Formula 1", Grams: 7.5},
	{ID: "9", Name: "Kabrita Gold 1", Grams: 7.3},
	{ID: "10", Name: "NANNYcare First Infant Milk", Grams: 7.4},
	{ID: "11", Name: "Aptamil Hungry Milk", Grams: 7.8},
	{ID: "12", Name: "Cow & Gate Infant Milk for Hungrier Babies", Grams: 7.8},
	{ID: "13", Name: "Hipp Organic Combiotic Hungry Infant Milk", Grams: 7.3},
	{ID: "14", Name: "SMA Extra Hungry", Grams: 7.0},
	{ID: "15", Name: "Aptamil Anti-reflux", Grams: 6.8},
	{ID: "16", Name: "Cow & Gate Anti-reflux", Grams: 6.8},
	{ID: "17", Name: "SMA Staydown", Grams: 7.0},
	{ID: "18", Name: "SMA Wysoy", Grams: 6.9},
	{ID: "19", Name: "Aptamil Lactose Free", Grams: 7.3},
	{ID: "20", Name: "SMA LF", Grams: 7.2},
	{ID: "21", Name: "Aptamil Comfort", Grams: 7.2},
	{ID: "22", Name: "Cow & Gate Comfort", Grams: 7.2},
	{ID: "23", Name: "SMA Comfort", Grams: 7.1},
	{ID: "24", Name: "SMA HA", Grams: 7.8},
	{ID: "25", Name: "Aptamil 2 Follow-on Milk", Grams: 8.6},
	{ID: "26", Name: "Aptamil Profutura 2 Follow-on Milk", Grams: 8.8},
	{ID: "27", Name: "Cow & Gate 2 Follow-on Milk", Grams: 8.6},
	{ID: "28", Name: "Hipp Organic Combiotic Follow-on Milk", Grams: 7.9},
	{ID: "29", Name: "Holle Organic Infant Follow-on Formula", Grams: 8.2},
	{ID: "30", Name: "SMA Pro Follow-on Milk", Grams: 7.9},
	{ID: "31", Name: "Holle Organic Infant Goat Milk Follow-on Formula 2", Grams: 8.0},
	{ID: "32", Name: "Kabrita Gold 2", Grams: 8.0},
	{ID: "33", Name: "NANNYcare Follow-on Milk", Grams: 7.4},
	{ID: "34", Name: "Hipp Organic Good Night Milk", Grams: 8.0},
	{ID: "35", Name: "Full-fat cows milk", Grams: 4.6},
	{ID: "36", Name: "Semi-skimmed cows Milk", Grams: 4.7},
	{ID: "37", Name: "Aptamil 3 Growing Up Milk 1-2 Years", Grams: 8.5},
	{ID: "38", Name: "Aptamil Profutura 3 Growing Up Milk", Grams: 8.4},
	{ID: "39", Name: "Cow & Gate 3 Growing Up Name 1-2 Years", Grams: 8.5},
	{ID: "40", Name: "Hipp Organic Combiotic Growing-up Milk", Grams: 8.2},
	{ID: "41", Name: "Holle Organic Growing-up Milk 3", Grams: 8.2},
	{ID: "42", Name: "PaediaSure Shake", Grams: 13.2},
	{ID: "43", Name: "SMA Pro Toddler Milk 3", Grams: 7.0},
	{ID: "44", Name: "Kabrita Gold 3", Grams: 7.9},
	{ID: "45", Name: "NANNYcare Growing Up Milk", Grams: 6.7},
	{ID: "46", Name: "Alpro Soya +1 Complete Care", Grams: 8.3},
	{ID: "47", Name: "Aptamil 4 Growing Up Milk 2-3 Years", Grams: 6.5},
	{ID: "48", Name: "Cow & Gate 4 Growing Up Milk 2-3 Years", Grams: 6.5},
	{ID: "49", Name: "Aptamil Pepti 1", Grams: 7.1}, //these are special milks
	{ID: "50", Name: "Cow & Gate Pepti-junior", Grams: 6.8},
	{ID: "51", Name: "Mead Johnson Nutramigen LIPIL 1", Grams: 7.5},
	{ID: "52", Name: "SHS Nutricia Pepdite", Grams: 7.8},
	{ID: "53", Name: "SHS Nutricia Pepdite MCT", Grams: 8.8},
	{ID: "54", Name: "SHS Nutricia Infatrini Peptisorb", Grams: 10.3},
	{ID: "55", Name: "Aptamil Pepti 2", Grams: 8.0},
	{ID: "56", Name: "Mead Johnson Pregestimil LIPIL", Grams: 6.9},
	{ID: "57", Name: "Mead Johnson Nutramigen LIPIL 2", Grams: 8.6},
	{ID: "58", Name: "SHS Nutricia Pepdite 1+", Grams: 13.0},
	{ID: "59", Name: "SHS Nutricia Pepdite MCT 1+", Grams: 11.8},
	{ID: "60", Name: "Mead Johnson Nutramigen AA", Grams: 7.0},
	{ID: "61", Name: "SHS Nutricia Neocate LCP", Grams: 7.2},
	{ID: "62", Name: "SHS Nutricia GA1 Anamix Infant", Grams: 7.4},
	{ID: "63", Name: "SHS Nutricia HCU Anamix Infant", Grams: 7.4},
	{ID: "64", Name: "SHS Nutricia IVA Anamix Infant", Grams: 7.4},
	{ID: "65", Name: "SHS Nutricia MMA/PA Anamix Infant", Grams: 7.4},
	{ID: "66", Name: "SHS Nutricia MSUD Anamix Infant", Grams: 7.4},
	{ID: "67", Name: "SHS Nutricia NKH Anamix Infant", Grams: 7.3},
	{ID: "68", Name: "SHS Nutricia TYR Anamix Infant", Grams: 7.3},
	{ID: "69", Name: "Vitaflo PKU start", Grams: 8.3},
	{ID: "70", Name: "Abbott Nutrition Similac High Energy", Grams: 10.3},
	{ID: "71", Name: "SHS Nutricia Infatrini", Grams: 10.3},
	{ID: "72", Name: "SMA High Energy", Grams: 9.8},
	{ID: "73", Name: "Mead Johnson Enfamil AR", Grams: 7.6},
	{ID: "74", Name: "Cow & Gate Infasoy", Grams: 7.0},
	{ID: "75", Name: "SMA Wysoy", Grams: 6.9},
	{ID: "76", Name: "Mead Johnson Enfamil O-Lac", Grams: 7.2},
	{ID: "77", Name: "SMA LF", Grams: 7.2},
	{ID: "78", Name: "Aptamil Preterm", Grams: 8.4},
	{ID: "79", Name: "Cow & Gate Nutriprem 1", Grams: 8.4},
	{ID: "80", Name: "SMA Gold Prem 1", Grams: 8.4},
	{ID: "81", Name: "Cow & Gate Nutriprem 2", Grams: 7.5},
	{ID: "82", Name: "SMA Gold Prem 2", Grams: 7.5},
	{ID: "83", Name: "SHS Nutricia Caprilon", Grams: 7.0},
	{ID: "84", Name: "SHS Nutricia Monogen", Grams: 12.0},
	{ID: "85", Name: "Vitaflo Lipistart", Grams: 8.3},
	{ID: "86", Name: "SHS Nutricia Galactomin 17", Grams: 7.3},
	{ID: "87", Name: "SHS Nutricia Galactomin 19", Grams: 6.4},
	{ID: "88", Name: "SHS Nutricia Kindergen", Grams: 11.8},
	{ID: "89", Name: "Vitaflo Renastart", Grams: 12.5},
}

// getMilks responds with the list of all milks as JSON.
func getMilks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, milks)
}

// postAlbums adds an album from JSON received in the request body.
func postMilks(c *gin.Context) {
	var newMilk Milk

	// Call BindJSON to bind the received JSON to
	// newMilk.
	if err := c.BindJSON(&newMilk); err != nil {
		return
	}

	// Add the new milk to the slice.
	milks = append(milks, newMilk)
	c.IndentedJSON(http.StatusCreated, newMilk)
}

// getMilkByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getMilkByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of milks, looking for
	// a milk whose ID value matches the parameter.
	for _, a := range milks {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "milk not found"})
}

func dextroseInfusionRate(c *gin.Context) {
	// POST request that accepts params:
	// rate in ml/hr
	// weight in kg
	// percentage of CHO in fluid as g/100ml
	// Return value in mg/kg/min
	var request GIRRequest
	if err := c.ShouldBind(&request); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	rate := request.Rate
	weight := request.Weight
	percentage := request.Percentage

	// this is the formula for glucose infusion rate
	gir := girlib.GlucoseInfusionRate(rate, percentage, weight)

	c.IndentedJSON(200, gin.H{"result": gir})
}

func milkGlucoseInfusionRate(c *gin.Context) {
	// POST request that accepts params:
	// milk volume/kg/d
	// weight in kg
	// milk id
	// total
	// Return value in mg/kg/min
	var request MilkGIRRequest

	if err := c.ShouldBind(&request); err != nil {
		c.String(http.StatusBadRequest, "bad request: %v", err)
		return
	}

	ID := request.ID
	rate := request.Volume / 24
	weight := request.Weight

	for _, a := range milks {
		if a.ID == ID {
			percentage := float32(a.Grams)
			gir := girlib.GlucoseInfusionRate(rate, percentage, weight)
			c.IndentedJSON(http.StatusOK, gin.H{"result": gir})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "milk not found"})
}
