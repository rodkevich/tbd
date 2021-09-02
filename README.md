tbd

#### Вообще эта задача для NoSQL базы...

***


Locally:   
RUN: `go mod init`  
RUN: `go mod vendor`  
RUN: `go run ./cmd/task5/server/server.go `

For tasks using db run on Docker navigate ./cmd/task5/:  
`docker-compose build --parallel`  
`docker-compose up`

If you have any docker volumes attached to postgres init.js won't work. Detach them
using :  
`docker-compose down --volumes`

.  
`pkg`  
|-`app`   
|-`datasource`  
| |-`postgress`  
|-`tickets`  - **will be also used for last task**  
| |-interface.go  *used for taskN 3*  
| |-functions.go  *used for taskN 1*  
| |-methods.go  *used for taskN 1*  
| |-structures.go  *used for taskN 1*  
| |-`types`  - **additional types used for development**  
| | |-currency.go  
| | |-description.go  
| | |-phone.go  
| | |-link.go

`cmd`  *runnable files and requirements for each task*  
|-**task1 - structures, methods & functions**  
| |-task1.go  
| |-req.yaml  
|-**task2 - for loop**  
| |-task2.go  
| |-main.go  
| |-req.yaml  
|-**task3 - interfaces**  
| |-req.yaml  
| |-task3.go  
|-**task4 - algorithm** `?q=is_this_task_ok`   
| |-task4.go  
| |-req.yaml  
|-**task5 - last task app**  
| |-validations.go - **example**
| |-req.yaml | |-`server` - **application**
| |-`docker-compose.yml

`internal` *different small helpers*  
|-msg  
| |-messages.go

|-**`test`**  *test files*

README.md
***  
