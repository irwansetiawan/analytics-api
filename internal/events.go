package events

import (
	"log"

	as "github.com/aerospike/aerospike-client-go/v6"
)

func WriteEvent(data map[string]string) {
	go func() {
		// Define the Aerospike server address
		host := "127.0.0.1"
		port := 3000

		// Connect to the Aerospike cluster
		client, err := as.NewClient(host, port)
		if err != nil {
			log.Fatalf("Failed to connect to Aerospike: %v", err)
		}
		defer client.Close()

		// Define the namespace and set
		namespace := "test"
		set := "demo"

		// Create a key
		key, err := as.NewKey(namespace, set, "key1")
		if err != nil {
			log.Fatalf("Failed to create key: %v", err)
		}

		// Create a record with some bins
		bins := as.BinMap{"data": data}

		// Write the record to the database
		err = client.Put(nil, key, bins)
		if err != nil {
			log.Fatalf("Failed to write record: %v", err)
		}

		log.Println("Record written successfully")
	}()
}
