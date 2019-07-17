# ![sling](https://github.com/alexstov/sling/blob/master/img/Sling.png)

### Network traffic simulator and test automation tool for software developers, testers, or anybody else in need to send file requests through the HTTP or TCP protocol, control rate frequency, number of concurrent connections, delays, timeouts, and collect the response time statistics, means, and percentiles.

## Overview
Sling is a lightweight CUI alternative to network test automation tools like Postman with the set of features required to send file requests to network endpoints and collect performance statistics. The requests are stored in files and sent individually or as a collection of directory files with configurable frequency, number of concurrent connections and repeat counts, delays and timeouts.

## Quickstart

## Features

### Commands

### `sling request send -f myrequest.dat`
Send myrequest.dat file content to default endpoint address set in SLINGCONFIG.

### `sling config view`
View SLINGCONFIG file content in the console.

### `sling log view`
View sling log file in the console.

### `sling log clean`
Clean sling log file, request and response log data.

### Config
Sling configuration path is set using SLINGCONFIG environment variable. For convenience you can permanently set the configuration file path as in the bash shell example below.

```
~ vi ~/.bashrc
export SLINGCONFIG=/home/alexstov/sling/config.DV.yml
source ~/.bashrc
```

All sling settings are set using SLINGCONFIG file. These include send file settings such as the filename of the file, directory where the file is stored, and wildcard to match multiple files from the same directory. 

```
file: ""
dir: "/home/alexstov/sling/data"
wildcard: "*.dat*"
```
Send settings include send repeat count for single file or multiple files, destination endpoint configuration, and options to save requests and responses. In the example below ten (10) requests are send to **type 2 HTTP POST** endpoint to the address http://<i><i>localhost:8080/TR. Each request is saved in /home/alexstov/Logs/Sling directory before sending; the responses are saved in /home/alexstov/Logs/Sling upon completion.

NOTE: The first endpoint is in the configuration below is of **type1 TCP**.

```
repeat: 10
endpointIndex: 1
endpoints:
- endpoint:
  address: "localhost"
  port: 8634
  type: 1 # TCP
- endpoint:
  address: http://localhost:8080/TR
  type: 2 # HTTPPost
saveReq: true
saveReqDir: "/home/alexstov/Logs/Sling/"
saveRes: true
saveResDir: "/home/alexstov/Logs/Sling/"
```
Throttle settings control the rate of requests using **rateSec** and **rateMin**. **cxtNum** sets tee burst rate to limit the rate of the requests by restricting buffer capacity of connection bursts. Internally sling prepares requests before enqueuing them to network client for transmission. Enqueued requests affects local resource consumption; this can be controlled with **cxnLim** flag to limit the number of prepared requests. When **cxnLim** is set to true, the number of enqueued requests will not exceed **cxnNum** limit. When **cxnLim** is set to false sling will enqueue as many as repeat count of requests. **sleepMs** sets the number of milliseconds to sleep after sending each request  before pulling another request from the queue.

**tmoCxn**, **tmoSec** control network client timeout for sending requests to destiantion. **tmoRdS** and **tmoWrS** set read and write timeouts respectively. A Zero value for Tmo settings mean the request will not time out.

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

Log settings control the parameters of sling logging. **histogram** enables metrics output in the log file.

```
log:
  level: 5
  logFile: "/home/alexstov/Logs/Sling/sling.log"
  disableColors: true
  fullTimestamp: true
  histogram: true
  timestampformat: "2019-01-02 15:04:05"
```
Console settings are similar to Log settings above but intended for sling console output.

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
Flags are used to customize sling functionality by overriding SLINGCONFIG settings. The following flags are used:

```
Flags:
  -a, --address string      endpoint IP, DNS name, or HTTP address (default "HTTP://localhost:9013/GEN/TPFA")
  -c, --cltType string      network client type, TCP or HttpPost (default "HTTPPost")
  -y, --conHis              write histogram to console (default true)
  -l, --cxnLim              limit the number of concurrent connections (default true)
  -n, --cxnNum uint         number of concurrent connections (default 2)
  -d, --dir string          directory to send files from (default "/home/alexstov/sling/data")
  -i, --endpoint uint       active endpoint index in SLINGCONFIG, zero-based (default 1)
  -f, --file string         filepath or filename to send
  -h, --help                help for send
  -g, --logHis              write histogram to log file (default true)
  -p, --port uint           endpoint port number
  -m, --rateMin uint        send rate per minute (default 6000)
  -s, --rateSec uint        send rate per second (default 100)
  -r, --repeat uint         send repeat count (default 1)
  -q, --saveReq             save requests
  -k, --saveReqDir string   directory to save requests (default "/home/alexstov/Logs/Sling/")
  -o, --saveRes             save responses
  -j, --saveResDir string   directory to save response (default "/home/alexstov/Logs/Sling/")
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

### sling request send -r 10 -f my_http_request.dat

Send my_http_request.dat from SLINGCONFIG directory using default settings, repeat the request 10 times.

```
sling request send -r 10 -f my_http_request.dat
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:01:45]  INFO histogram Client

[2019-07-15 09:01:45]  INFO   count:              10

[2019-07-15 09:01:45]  INFO   min:              3211

[2019-07-15 09:01:45]  INFO   max:              3294

[2019-07-15 09:01:45]  INFO   mean:             3249.80

[2019-07-15 09:01:45]  INFO   stddev:             29.40

[2019-07-15 09:01:45]  INFO   median:           3243.00

[2019-07-15 09:01:45]  INFO   75%:              3288.00

[2019-07-15 09:01:45]  INFO   95%:              3294.00

[2019-07-15 09:01:45]  INFO   99%:              3294.00

[2019-07-15 09:01:45]  INFO   99.9%:            3294.00
```

### sling request send -r 5 -a http://<i></i>localhost:8080/TR --file=my_http_request.dat -d /tmp -m 2

Send my_http_request.dat from /tmp directory to http://<i><i>localhost:8080/TR, repeat 5 times, limit the rate to 2 requests per minute.

```
sling request send -r 5 -a http://localhost:8080/TR --file=my_http_request.dat -d /tmp -m 2
[2019-07-15 09:49:51]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:49:51]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:50:21]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:50:51]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:51:21]  INFO Request sent successfully. FilePath=my_http_request.dat Length=160137
[2019-07-15 09:51:21]  INFO histogram Client

[2019-07-15 09:51:21]  INFO   count:               5

[2019-07-15 09:51:21]  INFO   min:              3108

[2019-07-15 09:51:21]  INFO   max:              3498

[2019-07-15 09:51:21]  INFO   mean:             3328.40

[2019-07-15 09:51:21]  INFO   stddev:            149.43

[2019-07-15 09:51:21]  INFO   median:           3419.00

[2019-07-15 09:51:21]  INFO   75%:              3459.50

[2019-07-15 09:51:21]  INFO   95%:              3498.00

[2019-07-15 09:51:21]  INFO   99%:              3498.00

[2019-07-15 09:51:21]  INFO   99.9%:            3498.00
```
### sling request send -i 0 -f my_tcp_request.dat

Send TCP request from the file to the endpoint with **index 0** in SLINGCONFIG. The endpoint is configured below with **type 1 TCP**.

```
endpoints:
- endpoint:
  address: "localhost"
  port: 8634
  type: 1
```

## Setup

## Contributing
1. Fork it
2. Download your fork to your PC (git clone https://github.com/your_username/sling && cd sling)
3. Create your feature branch (git checkout -b my-new-feature)
4. Make changes and add them (git add .)
5. Commit your changes (git commit -m 'Add some feature')
6. Push to the branch (git push origin my-new-feature)
7. Create new pull request

## Credits
Sling is powered by [spf13/cobra](https://github.com/spf13/cobra)

## Contact
Created by Alexey Stolpovskikh (alexstov@gmail.gom; stolpovskikh@hotmail.com)

## Licence
Sling is licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0) (the "License"); you may not use this software except in compliance with the License.

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
