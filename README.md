# ![sling](https://github.com/alexstov/sling/blob/master/img/Sling.png)

### Network traffic simulation, test automation tool for middleware developers, testers or anybody else in need to send file requests through the HTTP or TCP protocols, controlling rate frequency, number of concurrent connections, delays, and timeouts. It allows to collect the response time statistics, mean and percentiles.

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
Sling configuration path is set using SLINGCONFIG environment variable. For conveniance you can permanently safe the path, for bash shell in .bashrc, as below.

```
~ vi ~/.bashrc
export SLINGCONFIG=/home/yourusername/yourpath/config.DV.yml
source ~/.bashrc
```

All sling settings can be set using SLINGCONFIG file. These include send file settings such as the filename of the file, directory where the file is stored, and wildcard to match multiple files from the same directory. 

```
file: ""
dir: "/home/alexstol/Data/Current/tpf/request"
wildcard: "*.dat*"
```

The send settings include repeat count to send single file or multiple files from the same directory, destinationendpoint configuraiton, and options to save send requests and responses.

```
repeat: 1
endpointIndex: 1
endpoints:
- endpoint:
  address: "localhost"
  port: 9013
  type: 1
- endpoint:
  address: HTTP://localhost:9013/GEN/TPFA
  type: 2
- endpoint:
  address: "VHLDVBUAS008.TVLPORT.NET"
  port: 9000
  type: 1
- endpoint:
  address: HTTP://vhldvgfbu001.tvlport.net:9000/GEN/TPFA
  port: 9000
  type: 2
saveReq: false
saveReqDir: "/home/alexstol/Logs/Sling/"
saveRes: false
saveResDir: "/home/alexstol/Logs/Sling/"
```

### Flags

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
