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
	"io"

	//"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/4alexey/sling/sio"
	"github.com/4alexey/sling/slog"
	"github.com/4alexey/sling/util"
	"github.com/pkg/errors" //"errors"

	"github.com/sirupsen/logrus"
)

// SaveToFileFunc status callback.
type SaveToFileFunc func(filepath string, buff *bytes.Buffer) (err error)

// TCPClient Client interface implementation.
type TCPClient struct {
	client Client
	logger slog.Logger
	filer  sio.Filer
	conn   net.Conn
}

// NewTCPClient creates new TCP Client implementation.
func NewTCPClient(slog slog.Logger, flr sio.Filer) (clt Client, err error) {
	tcpClt := &TCPClient{logger: slog, filer: flr}
	tcpClt.client = tcpClt
	return tcpClt.client, nil
}

// Write sends bytes to specified address.
func (clt TCPClient) Write(msg []byte, args *WriteArgs) (err error) {
	var wrLen int
	addr := strings.Join([]string{args.IPAddress, strconv.FormatUint(uint64(args.Port), 10)}, ":")

	// Connect to endpoint.
	if clt.conn, err = net.Dial("tcp", addr); err != nil {
		err = errors.Wrap(err, "net.Dial")
		return
	}
	defer clt.conn.Close()

	// Set connection timeouts is set in the config or passed explicitly.
	if args.TmoSec != util.MaxUint && args.TmoSec != 0 {
		clt.conn.SetDeadline(time.Now().Add(time.Duration(args.TmoSec) * time.Second))
	}
	if args.TmoRdS != util.MaxUint && args.TmoRdS != 0 {
		clt.conn.SetReadDeadline(time.Now().Add(time.Duration(args.TmoRdS) * time.Second))
	}
	if args.TmoWrS != util.MaxUint && args.TmoWrS != 0 {
		clt.conn.SetWriteDeadline(time.Now().Add(time.Duration(args.TmoWrS) * time.Second))
	}

	// Send request.
	wrLen, err = clt.conn.Write(msg)
	if err != nil {
		err = errors.Wrap(err, " clt.conn.Write")
		return
	}
	clt.logger.Out(logrus.InfoLevel, logrus.Fields{"wrLen": wrLen}, "Successfully sent msg to destination address.")

	// Send message terminator.
	if wrLen, err = clt.conn.Write([]byte(MsgEndSequence)); err != nil {
		err = errors.Wrap(err, "conn.Write")
		return
	}
	clt.logger.Out(logrus.InfoLevel, nil, "Successfully sent message terminator.")

	// Read response.
	buf := make([]byte, 0, MaxMsgSize) // big buffer
	tmp := make([]byte, BufferSize)    // using small buffer
	var n int
	for {
		n, err = clt.conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				err = errors.Wrap(err, "clt.conn.Read(tmp)")
			} else {
				// EOF, no error.
				err = nil
			}
			break
		}

		buf = append(buf, tmp[:n]...)
	}

	if err != nil {
		clt.logger.Out(logrus.InfoLevel, logrus.Fields{"rdLen": len(buf)}, "Successfully read the response.")
	} else {
		err = errors.Wrap(err, "Read response")
		return
	}

	// Save response to a file.
	if args.SaveRes {
		if args.SaveResFilepath, err = clt.filer.BuildFilePath(args.SaveResDir, fmt.Sprintf("%03d", args.ReqID)+".res"); err != nil {
			clt.logger.Out(logrus.ErrorLevel, logrus.Fields{"filepath": args.SaveResFilepath}, "Cannot save the request.")
			err = errors.Wrap(err, "os.Stat")
			return
		}

		go args.SaveResCallback(args.SaveResFilepath, bytes.NewBuffer(buf))
	}

	return
}

// MsgEndSequence signals message completion.
const (
	MsgEndSequence = "\r\n\r\n"
	MaxMsgSize     = 67108864 // 64Kb
	BufferSize     = 1024
)
