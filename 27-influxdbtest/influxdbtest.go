package main

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var client influxdb2.Client

// You can generate a Token from the "Tokens Tab" in the UI
const token = "XC7JyvMtW_TifsTkl45jJ-k4--TmkcFEb51U-a2hFOs90QWlH3fXAhbdlx0uTyH1sopx6lgqqkSKnEJaFdhDaA=="
const bucket = "admin-bucket"
const org = "admin-org"

//Store the URL of your InfluxDB instance
const url = "http://192.168.31.204:8086"

//connect http://192.168.31.204:8086
func connect() {
	// Create new client with default option for server url authenticate by token
	client = influxdb2.NewClient(url, token)
}

//close client
func close() {
	client.Close()
}

//query
func queryPoints() {
	// Get query client
	queryAPI := client.QueryAPI(org)
	// Get QueryTableResult
	result, err := queryAPI.Query(context.Background(),
		fmt.Sprintf(`from(bucket:"%v")
							|> range(start: -90d)
							|> filter(fn: (r) =>
								r._measurement == "jolytest")`, bucket))

	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// fmt.Println(result)
			// Access data
			fmt.Printf("field:%v value: %v time:%s\n", result.Record().Field(), result.Record().Value(), result.Record().Time())
		}
		// Check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
}

//main()
func main() {
	connect()
	defer close()
	// writePonits()
	queryPoints()
}
