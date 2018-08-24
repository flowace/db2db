[![Build Status](https://travis-ci.com/flowace/db2db.svg?branch=master)](https://travis-ci.com/flowace/db2db)

## Load Testing Result
> All the testing was performed in Dual core Mac OSX system of **8GB** of RAM and with real world dataset

|Test #| Data Source Type | Data Destination Type | Remote |Size  | Rows|Columns|Execution Time|Concurrent|
|---|---|---|---|---|---|---|---|---|
|1|  SQLite database | CSV| No  | ~900mb  | 7.7 millions  |  2 | 102.726 sec|No|

## Dependencies

[**Cobra**](https://github.com/spf13/cobra)

[**Gorm**](https://github.com/jinzhu/gorm)

[**Zap**](https://github.com/uber-go/zap)

[**Viper**](https://github.com/spf13/viper)


