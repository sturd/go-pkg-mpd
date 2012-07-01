// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain
// Dedication license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package mpd

import "strconv"

type Args map[string]string

func (a Args) B(k string) bool  { return a.I(k) == 1 }
func (a Args) U8(k string) byte { return byte(a.I(k)) }

func (a Args) I(k string) int {
	if v, e := strconv.Atoi(a.read(k)); e == nil {
		return v
	}
	return 0
}

func (a Args) F32(k string) float32 {
	if v, e := strconv.ParseFloat(a.read(k), 64); e == nil {
		return float32(v)
	}
	return 0
}

func (a Args) F64(k string) float64 {
	if v, e := strconv.ParseFloat(a.read(k), 64); e == nil {
		return v
	}
	return 0
}

func (a Args) I64(k string) int64 {
	if v, e := strconv.ParseInt(a.read(k), 10, 64); e == nil {
		return v
	}
	return 0
}

func (a Args) S(k string) string {
	if v := a.read(k); v != "" {
		return v
	}
	return ""
}

func (a Args) read(key string) string {
	v, ok := a[key]
	if !ok {
		return ""
	}
	return v
}
