# Copyright (c) 2010, Jim Teeuwen. All rights reserved.
# This code is subject to a 1-clause BSD license.
# See the LICENSE file for its contents.

include $(GOROOT)/src/Make.inc

TARG = github.com/jteeuwen/go-pkg-mpd
GOFILES = client.go args.go api_*.go song.go status.go stats.go output.go

include $(GOROOT)/src/Make.pkg
