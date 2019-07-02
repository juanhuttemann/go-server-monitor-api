# go-server-monitor-api

Gathers and Publish information about the Hardware, OS, Network and Processes of the current running machine through a JSON API (Linux & Windows Compatible)

![preview](https://raw.githubusercontent.com/juanhuttemann/go-server-monitor-api/master/images/preview.png)

#### Dependencies 

```sh
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/disk
go get github.com/shirou/gopsutil/host
go get github.com/shirou/gopsutil/mem
go get github.com/mitchellh/go-ps
go get github.com/spf13/viper
```

#### Build 

```sh
go build src/*.go
```

#### Usage


```sh
./main
```

Then go to your web browser at [http:\\\localhost:3000](http:\\localhost:3000)


### Custom Settings

For custom settings just edit the config.yml file changing the values by true or false
