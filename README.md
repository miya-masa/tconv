# tmconv (TiMe CONVerter)
`tmconv` is a tool of fuzzy time conversion.

# Installation

`go get -u github.com/miya-masa/tmconv`

# Usage

```
#####  unix seconds to rfc3339
> tmconv 1551942249
{
  "utc": "2019-03-07T07:04:09Z",
  "local": "2019-03-07T16:04:09+09:00"
}

##### rfc3339 to unix seconds
> tmconv 2019-03-07T07:04:09Z
{
  "unix_seconds": 1551942249,
  "unix_millis": 1551942249000,
  "unix_micros": 1551942249000000,
  "unix_nanos": 1551942249000000000
}

##### ansic to unix seconds
> tmconv Mon Jan 2 15:04:05 2006
{
  "unix_seconds": 1136214245,
  "unix_millis": 1136214245000,
  "unix_micros": 1136214245000000,
  "unix_nanos": 1136214245000000000
}

##### pipe
> date | tmconv
{
  "unix_seconds": 1551942377,
  "unix_millis": 1551942377000,
  "unix_micros": 1551942377000000,
  "unix_nanos": 1551942377000000000
}
```

# Support format

```
	ANSIC       ex. "Mon Jan 2 15:04:05 2006"
	UnixDate    ex. "Mon Jan 2 15:04:05 MST 2006"
	RubyDate    ex. "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      ex. "02 Jan 06 15:04 MST"
	RFC822Z     ex. "02 Jan 06 15:04 -0700"
	RFC850      ex. "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     ex. "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    ex. "Mon, 02 Jan 2006 15:04:05 -0700"
	RFC3339Nano ex. "2006-01-02T15:04:05.999999999Z07:00"
	RFC3339     ex. "2006-01-02T15:04:05Z07:00"
```
