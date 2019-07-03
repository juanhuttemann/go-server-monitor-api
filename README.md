# go-server-monitor-api

Gathers and Publish information about the Hardware, OS, Network and Processes of the current running machine through a JSON API (Linux & Windows Compatible)

![preview](https://raw.githubusercontent.com/juanhuttemann/go-server-monitor-api/master/images/preview.png)


#### Clone 

```sh
git clone git@github.com:juanhuttemann/go-server-monitor-api.git
```

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


### JSON Data References

* **hostInfo** (object)
    * name (string): Host name
    * os (string): Operative System
    * arch (string): x86/x64
    * uptime (integer): amount of seconds since bootup
    
* **cpu** (object)
    * name (string): Brand - Model of the Physical CPU
    * total (integer): Amount of Physical Cores
    * usage (float): Percentage of Physical CPU Usage
    * usagePerCore (Array): Array of Percentage of Logical CPU Usage
    
* **ram** (object)
    * free (integer): Available Physical RAM in bytes
    * total (integer): Total of Physical RAM in bytes
    * usage (integer): Usage of Physical RAM in bytes



### Custom Settings

For custom settings just edit the config.yml file changing the values by **true** or **false**
    
       
       
