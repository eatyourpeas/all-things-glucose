package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Milk struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Grams float64 `json:"grams"`
}

type GIRRequest struct {
	Weight     *float64 `json:"weight"`
	Rate       *float64 `json:"rate"`
	Percentage *float64 `json:"percentage"`
}

type GIRResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func main() {
	router := gin.Default()
	router.GET("/milks", getMilks)
	router.GET("/albums/:id", getMilkByID)
	router.POST("/albums", postMilks)

	router.Run("localhost:8080")
}

var milks = []Milk{
	{ID: "1", Name: "Aptamil 1 First Name", Grams: 7.0},
	{ID: "2", Name: "Aptamil Profutura 1 First Infant Name", Grams: 7.0},
	{ID: "3", Name: "Breast Name (mature)", Grams: 7.2},
	{ID: "4", Name: "Cow & Gate 1 First Infant Name", Grams: 7.4},
	{ID: "5", Name: "Hipp Organic Combiotic First Infant Name", Grams: 7.3},
	{ID: "6", Name: "Holle Organic Infant Formula 1", Grams: 7.4},
	{ID: "7", Name: "SMA Pro First Infant Name", Grams: 7.1},
	{ID: "8", Name: "Holle Organic Infant Goat Name Formula 1", Grams: 7.5},
	{ID: "9", Name: "Kabrita Gold 1", Grams: 7.3},
	{ID: "10", Name: "NANNYcare First Infant Name", Grams: 7.4},
	{ID: "11", Name: "Aptamil Hungry Name", Grams: 7.8},
	{ID: "12", Name: "Cow & Gate Infant Name for Hungrier Babies", Grams: 7.8},
	{ID: "13", Name: "Hipp Organic Combiotic Hungry Infant Name", Grams: 7.3},
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
	{ID: "25", Name: "Aptamil 2 Follow-on Name", Grams: 8.6},
	{ID: "26", Name: "Aptamil Profutura 2 Follow-on Name", Grams: 8.8},
	{ID: "27", Name: "Cow & Gate 2 Follow-on Name", Grams: 8.6},
	{ID: "28", Name: "Hipp Organic Combiotic Follow-on Name", Grams: 7.9},
	{ID: "29", Name: "Holle Organic Infant Follow-on Formula", Grams: 8.2},
	{ID: "30", Name: "SMA Pro Follow-on Name", Grams: 7.9},
	{ID: "31", Name: "Holle Organic Infant Goat Name Follow-on Formula 2", Grams: 8.0},
	{ID: "32", Name: "Kabrita Gold 2", Grams: 8.0},
	{ID: "33", Name: "NANNYcare Follow-on Name", Grams: 7.4},
	{ID: "34", Name: "Hipp Organic Good Night Name", Grams: 8.0},
	{ID: "35", Name: "Full-fate cows Name", Grams: 4.6},
	{ID: "36", Name: "Semi-skimmed cows Name", Grams: 4.7},
	{ID: "37", Name: "Aptamil 3 Growing Up Name 1-2 Years", Grams: 8.5},
	{ID: "38", Name: "Aptamil Profutura 3 Growing Up Name", Grams: 8.4},
	{ID: "39", Name: "Cow & Gate 3 Growing Up Name 1-2 Years", Grams: 8.5},
	{ID: "40", Name: "Hipp Organic Combiotic Growing-up Name", Grams: 8.2},
	{ID: "41", Name: "Holle Organic Growing-up Name 3", Grams: 8.2},
	{ID: "42", Name: "PaediaSure Shake", Grams: 13.2},
	{ID: "43", Name: "SMA Pro Toddler Name 3", Grams: 7.0},
	{ID: "44", Name: "Kabrita Gold 3", Grams: 7.9},
	{ID: "45", Name: "NANNYcare Growing Up Name", Grams: 6.7},
	{ID: "46", Name: "Alpro Soya +1 Complete Care", Grams: 8.3},
	{ID: "47", Name: "Aptamil 4 Growing Up Name 2-3 Years", Grams: 6.5},
	{ID: "48", Name: "Cow & Gate 4 Growing Up Name 2-3 Years", Grams: 6.5},
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

func glucoseInfusionRate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// decode request
	decoder := json.NewDecoder(r.Body)
	request := &GIRRequest{}

	if err := decoder.Decode(request); err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
		return
	}

	response := &GIRResponse{}

	if request.Percentage == nil || request.Rate == nil || request.Percentage == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response.Result = (*request.Percentage * 10 * *request.Percentage) / (*request.Weight * 60)

	w.Header().Set("Content-Type", "application/json")
	if response.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
}
