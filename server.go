package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/ridwanrais/simple-graphql-location-api/internal/presentation/graph"
	"github.com/sirupsen/logrus"
	// "log"
	// "net/http"
	// "os"
	// "github.com/99designs/gqlgen/graphql/handler"
	// "github.com/99designs/gqlgen/graphql/playground"
	// "github.com/ridwanrais/simple-graphql-location-api/internal/presentation/graph"
)

const defaultPort = "8080"

// type City500 struct {
// 	Id      string `json:"id"`
// 	Name    string `json:"name"`
// 	Country string `json:"country"`
// 	Admin1  string `json:"admin1"`
// 	Lat     string `json:"lat"`
// 	Lon     string `json:"lon"`
// 	Pop     string `json:"pop"`
// }

// type City struct {
// 	Name        string `json:"name"`
// 	StateId     int64  `json:"state_id"`
// 	StateCode   string `json:"state_code"`
// 	StateName   string `json:"state_name"`
// 	CountryId   int64  `json:"country_id"`
// 	CountryCode string `json:"country_code"`
// 	CountryName string `json:"country_name"`
// 	Latitude    string `json:"latitude"`
// 	Longitude   string `json:"longitude"`
// 	WikiDataId  string `json:"wikiDataId"`
// }

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Warningln("Can't find env.file. The app is now run using default environment variables.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := graph.InitResolver()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// Make HTTP GET request
	// resp, err := http.Get("https://raw.githubusercontent.com/lmfmaier/cities-json/master/cities500.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// // Read response body as byte array
	// dataCity500, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// // Unmarshal the JSON data into an array of City structs
	// var cities500 []City500
	// err = json.Unmarshal(dataCity500, &cities500)
	// if err != nil {
	// 	panic(err)
	// }

	// ===============================

	// // Make HTTP GET request
	// resp2, err := http.Get("https://github.com/dr5hn/countries-states-cities-database/raw/master/cities.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp2.Body.Close()

	// // Read response body as byte array
	// dataCity, err := ioutil.ReadAll(resp2.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// dataCity, err := ioutil.ReadFile("cities.json")
	// if err != nil {
	// 	panic(err)
	// }

	// // Unmarshal the JSON data into an array of City structs
	// var cities []City
	// err = json.Unmarshal(dataCity, &cities)
	// if err != nil {
	// 	panic(err)
	// }

	// // ===============================

	// err = godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	// // Open a connection to the MySQL database
	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE")))
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// values := []interface{}{}
	// query := "INSERT INTO cities (name, country_id, country_name, country_cca2, admin1_id, admin1_name, admin1_code, latitude, longitude, wiki_data_id) VALUES "

	// for _, city := range cities {
	// 	// append a placeholder to the query for each city
	// 	query += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"

	// 	// append the values of the current city to the slice of values
	// 	values = append(values, city.Name, city.CountryId, city.CountryName, city.CountryCode, city.StateId, city.StateName, city.StateCode, city.Latitude, city.Longitude, city.WikiDataId)

	// }

	// // remove the last comma from the query
	// query = query[:len(query)-1]

	// // execute the query with all the values
	// _, err = db.Exec(query, values...)
	// if err != nil {
	// 	panic(err)
	// }

	// // create a MongoDB client and connect to the database
	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("dummy_url"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer client.Disconnect(context.Background())

	// // get a handle to the "cities" collection in the "test" database
	// collection := client.Database("location_db").Collection("cities")

	// // create a slice of documents to insert
	// docs := make([]interface{}, len(cities))
	// for i, city := range cities {
	// 	// create a document for the current city
	// 	doc := bson.M{
	// 		"name":         city.Name,
	// 		"country_id":   city.CountryId,
	// 		"country_name": city.CountryName,
	// 		"country_cca2": city.CountryCode,
	// 		"admin1_id":    city.StateId,
	// 		"admin1_name":  city.StateName,
	// 		"admin1_code":  city.StateCode,
	// 		"latitude":     city.Latitude,
	// 		"longitude":    city.Longitude,
	// 		"wiki_data_id": city.WikiDataId,
	// 	}

	// 	// add the document to the slice of documents
	// 	docs[i] = doc
	// }

	// // insert the documents into the collection
	// _, err = collection.InsertMany(context.Background(), docs)
	// if err != nil {
	// 	panic(err)
	// }
}
