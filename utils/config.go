package main

import (
	"bytes"
	"fmt"

	"github.com/spf13/viper"
)

var config = []byte(`
{
    "operations": [
        "download",
        "upload"
    ],
    "download": {
        "domains": {
            "sources": [
                "sqlite"
            ],
            "destinations": [
                "csv"
            ]
        },
        "connections": {
            "sources": {
                "1": {
                    "files": "test.db",
                    "paths": "/Users/mayukhsarkar/Documents/Codes/go/src/db2db",
                    "tables": [
                        "products"
                    ]
                }
            },
            "destinations": {
                "1": {
                    "files": [
                        "test.csv"
                    ],
                    "paths": [
                        "/Users/mayukhsarkar/Documents/Codes/go/src/db2db"
                    ]
                }
            }
        }
    },
    "upload": {
        "domains": {
            "sources": [
                "csv"
            ],
            "destinations": [
                "sqlite"
            ]
        },
        "connections": {
            "sources": {
                "1": {
                    "files": [
                        "test.csv"
                    ],
                    "paths": [
                        "/Users/mayukhsarkar/Documents/Codes/go/src/db2db"
                    ]
                }
            },
            "destinations": {
                "1": {
                    "files": "test2.db",
                    "paths": "/Users/mayukhsarkar/Documents/Codes/go/src/db2db",
                    "tables": [
                        "products"
                    ]
                }
            }
        }
    }
}`)

func main() {
	viper.SetConfigName("config")                             // Sets the name of the config file without the extension
	viper.AddConfigPath("$HOME/Documents/Codes/go/src/db2db") // Sets where the config file will be located
	viper.SetConfigType("json")
	err := viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	k := viper.GetStringSlice("operations")
	v := viper.GetStringMap("download")
	fmt.Println(k[0])
	fmt.Println(v["domains"])
}
