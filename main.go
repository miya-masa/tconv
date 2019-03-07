package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	Version string
	Build   string
)

func main() {
	fs := flag.NewFlagSet("tmconv", flag.ExitOnError)
	if err := fs.Parse(os.Args); err != nil {
		log.Fatal(err)
		return
	}

	str := fs.Arg(1)
	if str == "" {
		bs, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
			return
		}
		str = strings.TrimSpace(string(bs))
	}

	val, err := tmconv(str)
	if err != nil {
		log.Fatal(err)
		return
	}
	bs, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s\n", bs)
}

var regUnixTime = regexp.MustCompile("(\\d{10})|(\\d{13})|(\\d{16})|(\\d{19})")

var timeFormats = []string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822Z,
	time.RFC822,
	time.RFC850,
	time.RFC1123Z,
	time.RFC1123,
	time.RFC3339Nano,
	time.RFC3339,
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

func tmconv(str string) (interface{}, error) {
	if regUnixTime.MatchString(str) {
		return unixToTimeSet(regUnixTime.FindStringSubmatch(str))
	}
	for _, f := range timeFormats {
		t, err := time.Parse(f, str)
		if err != nil {
			continue
		}
		return &UnixSet{
			UnixSeconds: t.Unix(),
			UnixMillis:  t.UnixNano() / 1e6,
			UnixMicros:  t.UnixNano() / 1e3,
			UnixNanos:   t.UnixNano(),
		}, nil
	}
	return nil, fmt.Errorf("unexpected error. input: %+v ", str)
}

func unixToTimeSet(str []string) (interface{}, error) {
	if s := str[1]; s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		t := time.Unix(int64(i), 0)
		return &TimeSet{
			UTC:   t.UTC().Format(time.RFC3339Nano),
			Local: t.Local().Format(time.RFC3339Nano),
		}, nil
	}
	if s := str[2]; s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		t := time.Unix(0, int64(i)*1e3)
		return &TimeSet{
			UTC:   t.UTC().Format(time.RFC3339Nano),
			Local: t.Local().Format(time.RFC3339Nano),
		}, nil
	}
	if s := str[3]; s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		t := time.Unix(0, int64(i)*1e6)
		return &TimeSet{
			UTC:   t.UTC().Format(time.RFC3339Nano),
			Local: t.Local().Format(time.RFC3339Nano),
		}, nil
	}
	if s := str[4]; s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		t := time.Unix(0, int64(i)*1e9)
		return &TimeSet{
			UTC:   t.UTC().Format(time.RFC3339Nano),
			Local: t.Local().Format(time.RFC3339Nano),
		}, nil
	}
	return nil, fmt.Errorf("unexpected unixToTimeSet error. input: %+v ", str)
}

type UnixSet struct {
	UnixSeconds int64 `json:"unix_seconds"`
	UnixMillis  int64 `json:"unix_millis"`
	UnixMicros  int64 `json:"unix_micros"`
	UnixNanos   int64 `json:"unix_nanos"`
}

type TimeSet struct {
	UTC   string `json:"utc"`
	Local string `json:"local"`
}
