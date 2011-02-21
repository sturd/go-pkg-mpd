// Copyright (c) 2010, Jim Teeuwen. All rights reserved.
// This code is subject to a 1-clause BSD license.
// See the LICENSE file for its contents.

package mpd

import (
	"os"
	"time"
	"regexp"
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"net"
)

// Used to test whether we are compatible with the MPD server.
var SupportedVersion = [3]int{0, 15, 0}

type Client struct {
	conn            net.Conn
	writer          *bufio.Writer
	reader          *bufio.Reader
	ProtocolVersion string
}

// This Opens a new connection to the specified MPD server and optionally
// logs in with the given password.
func Dial(address, password string) (c *Client, err os.Error) {
	c = new(Client)

	if c.conn, err = net.Dial("tcp", "", address); err != nil {
		return
	}

	c.reader = bufio.NewReader(c.conn)
	c.writer = bufio.NewWriter(c.conn)

	// Complete handshake. Server should send 'OK MPD 0.15.0'. This is the
	// protocol version, not the version of the MPD daemon itself. We can use it
	// to test if our program is compatible with the api exposed by the daemon.
	var data string
	if data, err = c.reader.ReadString('\n'); err != nil {
		c.Close()
		return nil, err
	}

	if data = strings.TrimSpace(data); len(data) == 0 {
		c.Close()
		return nil, os.NewError("No valid handshake received.")
	}

	if data[0:3] == "ACK" {
		c.Close()
		return nil, os.NewError(fmt.Sprintf("Handshake error: %s", data[4:]))
	}

	c.ProtocolVersion = data[3:]
	if !isSupportedVersion(c.ProtocolVersion) {
		c.Close()
		return nil, os.NewError(fmt.Sprintf(
			"Invalid protocol version. This library requires at least 'MPD %d.%d.%d'. Server sent '%s'.",
			SupportedVersion[0], SupportedVersion[1], SupportedVersion[2],
			c.ProtocolVersion,
		))
	}

	if len(password) > 0 {
		_, err = c.request("password \"%s\"", password)
	}

	return
}

// Close the open connection.
// The error returned is an os.Error to satisfy io.Closer;
func (this *Client) Close() (err os.Error) {
	if this.conn != nil {
		this.send("close")

		this.reader = nil
		this.writer = nil

		err = this.conn.Close()
		this.conn = nil
	}

	return
}

func (this *Client) parseError(line string) os.Error {
	if strings.HasPrefix(line, "ACK ") {
		// sig: [errcode@token] {command} message
		//  ex: [2@0] {enableoutput} wrong number of arguments for "enableoutput"
		pos := strings.Index(line, "}")
		return os.NewError(strings.TrimSpace(line[pos+1:]))
	}
	return os.NewError(line)
}

func (this *Client) request(cmd string, arg ...interface{}) (args Args, err os.Error) {
	if err = this.send(fmt.Sprintf(cmd, arg...)); err != nil {
		return
	}
	return this.receive()
}

func (this *Client) requestList(cmd string, arg ...interface{}) (args []Args, err os.Error) {
	if err = this.send(fmt.Sprintf(cmd, arg...)); err != nil {
		return
	}
	return this.receiveList()
}

func (this *Client) receive() (data Args, err os.Error) {
	if this.reader == nil {
		return nil, os.NewError("Stream reader is closed.")
	}

	data = make(Args)
	var line string
	var pos int

	for {
		if line, err = this.reader.ReadString('\n'); err != nil {
			return nil, err
		}

		if line = strings.TrimSpace(line); len(line) > 0 {
			if line == "OK" {
				break
			}

			if strings.HasPrefix(line, "ACK ") {
				return nil, this.parseError(line)
			}

			if pos = strings.Index(line, ":"); line[0:pos] == "error" {
				return nil, this.parseError(line)
			}

			data[line[0:pos]] = strings.TrimSpace(line[pos+1:])
		}
	}
	return
}

func (this *Client) receiveList() (data []Args, err os.Error) {
	var line string
	var pos int

	if this.reader == nil {
		return nil, os.NewError("Stream reader is closed.")
	}

	a := make(Args)

	for {
		if line, err = this.reader.ReadString('\n'); err != nil {
			return nil, err
		}

		if line = strings.TrimSpace(line); len(line) > 0 {
			if line == "OK" {
				if len(a) > 0 {
					data = append(data, a)
				}
				break
			}

			if strings.HasPrefix(line, "ACK ") {
				return nil, this.parseError(line)
			}

			if pos = strings.Index(line, ":"); line[0:pos] == "error" {
				return nil, this.parseError(line)
			}

			// Lists of entries are not delimited by a special token. We need
			// to tell them apart by checking for keys in the map which already
			// exist. If so, we are dealing with a new entry.
			if _, ok := a[line[0:pos]]; ok {
				if len(a) > 0 {
					data = append(data, a)
				}
				a = make(Args)
			}

			a[line[0:pos]] = strings.TrimSpace(line[pos+1:])
		}
	}
	return
}

func (this *Client) send(msg string, args ...interface{}) (err os.Error) {
	const max_retries = 3
	var tries, num int

	if this.writer == nil {
		return os.NewError("Stream writer is closed.")
	}

	msg = fmt.Sprintf(msg, args...)
	msg += "\n"

	for tries = 0; tries < max_retries; tries++ {
		if num, err = this.writer.WriteString(msg); num < len(msg) {
			time.Sleep(300000000) // 0.3 seconds between retries
			continue
		}
		this.writer.Flush()
		break
	}

	return
}

// The following methods are here to ensure mpd.Client implements the net.Conn
// interface. They are not necessarily useful for this particular connection,
// but the calls will be passed to the underlying connection.


// Read reads data from the connection.
// Read can be made to time out and return a net.Error with Timeout() == true
// after a fixed time limit; see SetTimeout and SetReadTimeout.
func (this *Client) Read(b []byte) (n int, err os.Error) {
	return this.conn.Read(b)
}

// Write writes data to the connection.
// Write can be made to time out and return a net.Error with Timeout() == true
// after a fixed time limit; see SetTimeout and SetWriteTimeout.
func (this *Client) Write(b []byte) (n int, err os.Error) {
	return this.conn.Write(b)
}

// LocalAddr returns the local network address.
func (this *Client) LocalAddr() net.Addr { return this.conn.LocalAddr() }

// RemoteAddr returns the remote network address.
func (this *Client) RemoteAddr() net.Addr { return this.conn.RemoteAddr() }

// SetTimeout sets the read and write deadlines associated
// with the connection.
func (this *Client) SetTimeout(nsec int64) os.Error { return this.conn.SetTimeout(nsec) }

// SetReadTimeout sets the time (in nanoseconds) that
// Read will wait for data before returning an error with Timeout() == true.
// Setting nsec == 0 (the default) disables the deadline.
func (this *Client) SetReadTimeout(nsec int64) os.Error { return this.conn.SetReadTimeout(nsec) }

// SetWriteTimeout sets the time (in nanoseconds) that
// Write will wait to send its data before returning an error with Timeout() == true.
// Setting nsec == 0 (the default) disables the deadline.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
func (this *Client) SetWriteTimeout(nsec int64) os.Error { return this.conn.SetWriteTimeout(nsec) }

func isSupportedVersion(ver string) bool {
	var reg_version *regexp.Regexp
	var err os.Error

	if reg_version, err = regexp.Compile(`^MPD ([0-9]+).([0-9]+).([0-9]+)$`); err != nil {
		return false
	}

	matches := reg_version.FindStringSubmatch(ver)
	if len(matches) == 0 {
		return false
	}

	version := [3]int{0, 0, 0}
	version[0], _ = strconv.Atoi(matches[1])
	version[1], _ = strconv.Atoi(matches[2])
	version[2], _ = strconv.Atoi(matches[3])

	if version[0] < SupportedVersion[0] {
		return false
	}

	if version[0] > SupportedVersion[0] {
		return true
	}

	if version[1] < SupportedVersion[1] {
		return false
	}

	if version[1] > SupportedVersion[1] {
		return true
	}

	if version[2] < SupportedVersion[2] {
		return false
	}

	return true
}
