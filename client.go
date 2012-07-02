// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Used to test whether we are compatible with the MPD server.
var SupportedVersion = [3]int{0, 15, 0}

type Client struct {
	conn            net.Conn
	writer          *bufio.Writer
	reader          *bufio.Reader
	ProtocolVersion string
}

// Dial opens a new connection to the specified MPD server and optionally
// logs in with the given password.
func Dial(address, password string) (c *Client, err error) {
	c = new(Client)

	if c.conn, err = net.Dial("tcp", address); err != nil {
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
		return nil, errors.New("No valid handshake received.")
	}

	if data[0:3] == "ACK" {
		c.Close()
		return nil, errors.New(fmt.Sprintf("Handshake error: %s", data[4:]))
	}

	c.ProtocolVersion = data[3:]
	if !isSupportedVersion(c.ProtocolVersion) {
		c.Close()
		return nil, errors.New(fmt.Sprintf(
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
func (c *Client) Close() (err error) {
	if c.conn != nil {
		c.send("close")

		c.reader = nil
		c.writer = nil

		err = c.conn.Close()
		c.conn = nil
	}

	return
}

func (c *Client) parseError(line string) error {
	if strings.HasPrefix(line, "ACK ") {
		// sig: [errcode@token] {command} message
		//  ex: [2@0] {enableoutput} wrong number of arguments for "enableoutput"
		pos := strings.Index(line, "}")
		return errors.New(strings.TrimSpace(line[pos+1:]))
	}
	return errors.New(line)
}

func (c *Client) request(cmd string, arg ...interface{}) (args Args, err error) {
	if err = c.send(fmt.Sprintf(cmd, arg...)); err != nil {
		return
	}
	return c.receive()
}

func (c *Client) requestList(cmd string, arg ...interface{}) (args []Args, err error) {
	if err = c.send(fmt.Sprintf(cmd, arg...)); err != nil {
		return
	}
	return c.receiveList()
}

func (c *Client) receive() (data Args, err error) {
	if c.reader == nil {
		return nil, errors.New("Stream reader is closed.")
	}

	data = make(Args)
	var line string
	var pos int

	for {
		if line, err = c.reader.ReadString('\n'); err != nil {
			return nil, err
		}

		if line = strings.TrimSpace(line); len(line) > 0 {
			if line == "OK" {
				break
			}

			if strings.HasPrefix(line, "ACK ") {
				return nil, c.parseError(line)
			}

			if pos = strings.Index(line, ":"); line[0:pos] == "error" {
				return nil, c.parseError(line)
			}

			data[line[0:pos]] = strings.TrimSpace(line[pos+1:])
		}
	}
	return
}

func (c *Client) receiveList() (data []Args, err error) {
	var line string
	var pos int

	if c.reader == nil {
		return nil, errors.New("Stream reader is closed.")
	}

	a := make(Args)

	for {
		if line, err = c.reader.ReadString('\n'); err != nil {
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
				return nil, c.parseError(line)
			}

			if pos = strings.Index(line, ":"); line[0:pos] == "error" {
				return nil, c.parseError(line)
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

func (c *Client) send(msg string, args ...interface{}) (err error) {
	const max_retries = 3
	var tries, num int

	if c.writer == nil {
		return errors.New("Stream writer is closed.")
	}

	msg = fmt.Sprintf(msg, args...)
	msg += "\n"

	for tries = 0; tries < max_retries; tries++ {
		if num, err = c.writer.WriteString(msg); num < len(msg) {
			time.Sleep(300000000) // 0.3 seconds between retries
			continue
		}
		c.writer.Flush()
		break
	}

	return
}

// The following methods are here to ensure mpd.Client implements the net.Conn
// interface. They are not necessarily useful for c particular connection,
// but the calls will be passed to the underlying connection.

// Read reads data from the connection.
// Read can be made to time out and return a net.Error with Timeout() == true
// after a fixed time limit; see SetTimeout and SetReadTimeout.
func (c *Client) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}

// Write writes data to the connection.
// Write can be made to time out and return a net.Error with Timeout() == true
// after a fixed time limit; see SetTimeout and SetWriteTimeout.
func (c *Client) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

// LocalAddr returns the local network address.
func (c *Client) LocalAddr() net.Addr { return c.conn.LocalAddr() }

// RemoteAddr returns the remote network address.
func (c *Client) RemoteAddr() net.Addr { return c.conn.RemoteAddr() }

func isSupportedVersion(ver string) bool {
	var reg_version *regexp.Regexp
	var err error

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
