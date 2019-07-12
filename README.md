# ![sling](https://github.com/alexstov/sling/blob/master/img/Sling.png)

### Network traffic simulator, test automation tool for software developers, testers or anybody else in need to send file requests through the HTTP or TCP protocol, controlling rate frequency, number of concurrent connections, delays, and timeouts. It allows to collect the response time statistics, mean and percentiles.

## Overview
Sling is a lightweight CUI alternative to network test automation tools like Postman with the set of features required to send file requests to network endpoints and collect performance statistics. The requests are stored in files and sent individually or as a collection of directory files with set frequency, concurrent connection and repeat counts, delays and timeouts.

## Quickstart

## Features

### Commands

### `sling request send -f myrequest.dat`
Send myrequest.dat file content to default endpont address set in SLINGCONFIG.

### `sling config view`
View SLINGCONFIG file content.

### `sling log view`
View sling log file.

### `sling log clean`
Clean sling log file, request and response log data.

### Config
Sling configuration path is set using SLINGCONFIG environment variable. For conveniance you can permanently set the path as in the bash shell example below.

```
~ vi ~/.bashrc
export SLINGCONFIG=/home/yourusername/yourpath/config.DV.yml
source ~/.bashrc
```

All sling settings are set using SLINGCONFIG file. These include send file settings such as the filename of the file, directory where the file is stored, and wildcard to match multiple files from the same directory. 

```
file: ""
dir: "/home/alexstol/Data/Current/tpf/request"
wildcard: "*.dat*"
```

The send settings include repeat count to send single file or multiple files, destination endpoint configuraiton, and options to save send requests and responses. In the example below ten (10) requests are send to type 2 HTTP POST endpoint to the address http://localhost:8080/TRAN. Each request is saved in /home/myaccount/Logs/Sling directory before sent, and responses are saved to /home/myaccount/Logs/Sling.

```
repeat: 10
endpointIndex: 1
endpoints:
- endpoint:
  address: "localhost"
  port: 8634
  type: 1
- endpoint:
  address: htt://localhost:8080/TRAN
  type: 2
saveReq: true
saveReqDir: "/home/myaccount/Logs/Sling/"
saveRes: true
saveResDir: "/home/myaccount/Logs/Sling/"
```
Throttle settings control the rate of requests using **rateSec** and **rateMin** and **cxtNum** limits the  rate of new requests sling sends to the destination by restriction buffer capacity that holds connection bursts. Internally sling prepares requests before enquing them to network clinet for transmission. Storing requests affects local computer resources; the **cxnLim** is used to control the number of prepared to send requests when set set to true, along with **cxtNum** limit i.e. only two requests read from the files and prepared to send when **cxnLim = true** and **cxtNum = 2** vs. total number of repeat requests prepared when **cxnLim = false**.

**tmoCxn**, **tmoSec** control network cling timeout when sending requests to destiantion. **tmoRdS** and **tmoWrS** set read and write timeouts respectively.

```
throttle:  
  cxnNum : 2
  cxnLim : true
  sleepMs : 0
  rateSec : 100
  rateMin : 6000
  tmoCxn : 10
  tmoSec : 43
  tmoRdS : 43
  tmoWrS : 10
```

```
log:
  level: 5
  logFile: "/home/alexstol/Logs/Sling/sling.log"
  disableColors: true
  fullTimestamp: true
  histogram: true
  timestampformat: "2019-01-02 15:04:05"
```

```
console:
  level: 6
  flat: false
  disableColors: true
  fullTimestamp: true
  histogram: true
  timestampformat: "2019-01-02 15:04:05"
```

### Flags
Flags are used to customize sling functionality by overriding SLINGCONFIG settings. The folling flags are used:

```
Flags:
  -a, --address string      endpoint IP, DNS name, or HTTP address (default "HTTP://localhost:9013/GEN/TPFA")
  -c, --cltType string      network client type, TCP or HttpPost (default "HTTPPost")
  -y, --conHis              write histogram to console (default true)
  -l, --cxnLim              limit the number of concurrent connections (default true)
  -n, --cxnNum uint         number of concurrent connections (default 2)
  -d, --dir string          directory to send files from (default "/home/alexstol/Data/Current/tpf/request")
  -i, --endpoint uint       active endpoint index in SLINGCONFIG, zero-based (default 1)
  -f, --file string         filepath or filename to send
  -h, --help                help for send
  -g, --logHis              write histogram to log file (default true)
  -p, --port uint           endpoint port number
  -m, --rateMin uint        send rate per minute (default 6000)
  -s, --rateSec uint        send rate per second (default 100)
  -r, --repeat uint         send repeat count (default 1)
  -q, --saveReq             save requests
  -k, --saveReqDir string   directory to save requests (default "/home/alexstol/Logs/Sling/")
  -o, --saveRes             save responses
  -j, --saveResDir string   directory to save response (default "/home/alexstol/Logs/Sling/")
  -e, --sleepMs uint        delay after each repeated request
  -u, --tmoCxn uint         network client dial timeout (default 10)
  -v, --tmoRdS uint         network client timeout for Read calls (default 43)
  -t, --tmoSec uint         network client timeout (default 43)
  -x, --tmoWrS uint         network client timeout for Write calls (default 10)
  -w, --wildcard string     filename matching wildcard (default "*.dat*")

Global Flags:
  --conFlat       set console flat output without timestamp and fields
  --conLvl uint   console output level (default 6) 
  --logLvl uint   log output level (default 5)
```

## Usage
fdsa

`sling log view`

## Setup

## Contributing

## Credits

## Contact

## Licence
Sling is licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) (the "License"); you may not use this software except in compliance with the License.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
