# api-monitor-go

Gathers and Publish information about the Hardware, OS, Network and Processes of the current running machine through a JSON API (Linux & Windows Compatible)

#### Dependencies 

```sh
go get	github.com/shirou/gopsutil/cpu
go get	github.com/shirou/gopsutil/disk
go get	github.com/shirou/gopsutil/host
go get	github.com/shirou/gopsutil/mem
```

#### Build 

```sh
go build src/*.go
```

### Usage


```sh
./main
```

Then go to your web browser at [http:\\localhost:3000](http:\\localhost:3000)
