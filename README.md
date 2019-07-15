# go-server-monitor-api

Gathers and Publish information about the Hardware, OS, Network and Processes of the current running machine through a JSON API (Linux & Windows Compatible)

![preview](https://raw.githubusercontent.com/juanhuttemann/go-server-monitor-api/master/images/preview1.png)
![preview](https://raw.githubusercontent.com/juanhuttemann/go-server-monitor-api/master/images/preview2.png)


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
go build -o gomonitor src/*.go
```

#### Usage

```sh
./gomonitor
```

Then go to your web browser at [http:\\\localhost:3000](http:\\localhost:3000)



### JSON Data References

* **hostInfo** (object)
    * name (string): Host name
    * os (string): Operative System
    * arch (string): x86/x64
    * platform (string): OS Name, Version and Build
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

* **disks** (array of objects)
    * mountPoint (string): Disk current mount path or Unit in Windows 
    * size (integer): Total Disk space in bytes
    * used (integer): Used Disk space in bytes
    * percent (integer): Percentage of Disk Usage

* **networkDevices** (array of objects)
    * name (string): Name of the Network Interface
    * addresses (array of objects):
        * cidr (string): Classless Inter-Domain Routing
        * ip (string): IP Address
    * mac (string): Mac Address
    * up (boolean): If the Net Interface is Up or Down


### Custom Settings


#### Webserver API Port

By default, the embedded server start on port 3000. We can provide a different value in an config.yml file:

```
# api server port
port: 45000
```

#### Modules

For enabling/disabling modules just edit the config.yml file changing the values by **true** or **false**

```
#available modules
hostInfo: true
cpu: true
ram: true
disks: true
networkDevices: false
networkBandwidth: true
processes: false
```
       
       
