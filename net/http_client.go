// Copyright © 2019 Alexey Stolpovskikh <stolpovskikh@hotmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package net

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/4alexey/sling/conf"
	"github.com/4alexey/sling/sio"
	"github.com/4alexey/sling/slog"
	"github.com/4alexey/sling/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	//"errors"
)

// HTTPHeader returns header string.
func HTTPHeader(len int) []byte {
	header := "POST /GEN/TPFA HTTP/1.0"
	return []byte(header)
}

// HTTPClient Client interface implementation,
type HTTPClient struct {
	client Client
	logger slog.Logger
	filer  sio.Filer
	conn   net.Conn
}

// NewHTTPClient creates new HTTP Client instance.
func NewHTTPClient(slog slog.Logger, filer sio.Filer) (clt Client, err error) {
	httpClt := &HTTPClient{logger: slog}
	httpClt.client = httpClt
	httpClt.filer = filer
	return httpClt.client, nil
}

// Write sends bytes to specified address.
func (clt HTTPClient) Write(msg []byte, args *WriteArgs) (err error) {
	var readWriteTimeout uint
	if args.TmoRdS > args.TmoWrS {
		readWriteTimeout = args.TmoRdS
	} else {
		readWriteTimeout = args.TmoWrS
	}

	httpClt := &http.Client{
		Transport: &http.Transport{
			Dial: timeoutDialer(args.TmoCxn, readWriteTimeout),
		},
		Timeout: time.Duration(args.TmoSec) * time.Second,
	}

	var resp *http.Response
	if args.CltType == conf.HTTPPost {
		// Post request
		resp, err = httpClt.Post(args.IPAddress, "application/x-www-form-urlencoded", bytes.NewBuffer(msg))
		if err != nil {
			clt.logger.Out(logrus.ErrorLevel, logrus.Fields{"error": err}, "HTTP client failed to send the request.")
			return
		}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	clt.logger.Out(logrus.InfoLevel, logrus.Fields{"numBytes": len(body)}, "Successfully received msg reply.")

	// Save response to a file.
	if args.SaveRes {
		if args.SaveResFilepath, err = clt.filer.BuildFilePath(args.SaveResDir, fmt.Sprintf("%03d", args.ReqID)+".res"); err != nil {
			clt.logger.Out(logrus.ErrorLevel, logrus.Fields{"filepath": args.SaveResFilepath}, "Cannot save the request.")
			err = errors.Wrap(err, "os.Stat")
		}

		buf := bytes.NewBuffer(body)

		go args.SaveResCallback(args.SaveResFilepath, buf)
	}
	return
}

// Set connection timeouts.
func timeoutDialer(cTimeout uint, rwTimeout uint) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, time.Duration(cTimeout)*time.Second)
		if err != nil {
			return nil, err
		}

		if rwTimeout != 0 && rwTimeout != util.MaxUint {
			conn.SetDeadline(time.Now().Add(time.Duration(rwTimeout)))
		}

		return conn, nil
	}
}
