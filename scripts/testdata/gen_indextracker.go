// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package main

// TODO needs fixing.

// type trackerIndexTestData struct {
// 	URL      string
// 	Param    string
// 	Payload  string
// 	Expected []float64
// }

// func main() {
// 	log.SetFlags(log.Ltime | log.Lshortfile | log.Lmsgprefix)

// 	f, err := ioutil.ReadFile(filepath.Join("configs", "config.json"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	cfg := config.Config{}

// 	if err := json.Unmarshal(f, &cfg); err != nil {
// 		log.Fatal(err)
// 	}

// 	err = format.SetupLoggingConfig(cfg.Logger)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	DB, err := db.Open(filepath.Join(os.TempDir(), "testdata_gen"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if _, err := index.BuildTrackerIndexs(&cfg, DB); err != nil {
// 		log.Fatal(err)
// 	}
// 	indexes := tracker.GetIndexes()
// 	testdata := make(map[string][]*trackerIndexTestData)
// 	for symbol, indexers := range indexes {
// 		// Skip on duplicates.
// 		if strings.Contains(symbol, "~") {
// 			continue
// 		}
// 		var testdataForSymbol = make([]*trackerIndexTestData, 0)
// 		for _, indexer := range indexers {
// 			payload, err := indexer.Source.Get()
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			expected, err := indexer.ParsePayload(payload)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			testdataForSymbol = append(testdataForSymbol, &trackerIndexTestData{
// 				URL:      indexer.Identifier,
// 				Param:    indexer.Param,
// 				Payload:  string(payload),
// 				Expected: expected,
// 			})
// 		}
// 		testdata[symbol] = testdataForSymbol
// 	}
// 	testdataJSON, err := json.MarshalIndent(testdata, "", "    ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = ioutil.WriteFile(filepath.Join("test", "tracker", "testdata", "test_index.json"), testdataJSON, os.ModePerm)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
