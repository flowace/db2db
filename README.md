[![Build Status](https://travis-ci.com/flowace/db2db.svg?branch=master)](https://travis-ci.com/flowace/db2db)
[![Go Report Card](https://goreportcard.com/badge/github.com/flowace/db2db)](https://goreportcard.com/report/github.com/flowace/db2db)

## Load Testing Result
> All the testing was performed in Dual core Mac OSX system of **8GB** of RAM and with real world dataset of Spain Weather dept.

|Test #| Data Source Type | Data Destination Type | Remote |Size  | Rows|Columns|Execution Time|Concurrent|
|---|---|---|---|---|---|---|---|---|
|1|  SQLite database | CSV| No  | ~900mb  | 7.7 millions  |  2 | 102.726 sec|No|
|2|  SQLite database | CSV| No  | ~900mb  | 7.7 millions  |  20 | 396.00s sec|Yes (10 threads) |

## Dependencies

[**Cobra**](https://github.com/spf13/cobra)

[**Gorm**](https://github.com/jinzhu/gorm)

[**Zap**](https://github.com/uber-go/zap)

[**Viper**](https://github.com/spf13/viper)


