package events

import (
	"log"

	as "github.com/aerospike/aerospike-client-go/v6"
	"github.com/google/uuid"
)

func WriteEvent(data map[string]string) {
	go func() {
		// Define the Aerospike server address
		host := "127.0.0.1"
		port := 3000

		// Connect to the Aerospike cluster
		client, err := as.NewClient(host, port)
		if err != nil {
			log.Printf("Failed to connect to Aerospike: %v", err)
			return
		}
		defer client.Close()

		// Define the namespace and set
		namespace := "test"
		set := "demo"

		// Create a key
		key, err := as.NewKey(namespace, set, uuid.New().String())
		if err != nil {
			log.Printf("Failed to create key: %v", err)
			return
		}

		// Create a record with some bins
		bins := as.BinMap{"data": data}

		// Write the record to the database
		err = client.Put(nil, key, bins)
		if err != nil {
			log.Printf("Failed to write record: %v", err)
			return
		}

		log.Println("Record written successfully")
	}()
}
