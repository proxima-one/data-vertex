/*
Configurations
- Database
- Tables
- DApp aggregator
- Dataloaders and resolvers (config)
*/
package main

import (
	"os"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
	  return value
  }
    return defaultVal
}



/*
config

database client (found in db app)
application
	tables
resolvers
dataloaders
schema
Authentication public keys

*/
