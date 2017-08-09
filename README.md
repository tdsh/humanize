

# humanize
`import "github.com/tdsh/humanize"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
(Exercise) Package humanize provides functions to convert number to various human
readable format.




## <a name="pkg-index">Index</a>
* [func APNumber(n int) string](#APNumber)
* [func FormatPercent(f float64) string](#FormatPercent)
* [func Fractional(f float64) string](#Fractional)
* [func IntComma(n int) string](#IntComma)
* [func IntWord(n int64) string](#IntWord)
* [func NaturalDay(t time.Time) string](#NaturalDay)
* [func NaturalDelta(t time.Time, relative bool) string](#NaturalDelta)
* [func NaturalSize(val int, options ...SetOption) string](#NaturalSize)
* [func Ordinal(n int) string](#Ordinal)
* [func SetLocale(l string)](#SetLocale)
* [type FormatOption](#FormatOption)
* [type SetOption](#SetOption)
  * [func Binary(on bool) SetOption](#Binary)
  * [func GNU(on bool) SetOption](#GNU)


#### <a name="pkg-files">Package files</a>
[filesize.go](/src/github.com/tdsh/humanize/filesize.go) [i18n.go](/src/github.com/tdsh/humanize/i18n.go) [number.go](/src/github.com/tdsh/humanize/number.go) [time.go](/src/github.com/tdsh/humanize/time.go) 





## <a name="APNumber">func</a> [APNumber](/src/target/number.go?s=2185:2212#L80)
``` go
func APNumber(n int) string
```
APNumber returns the number spelled out for 1-9. Otherwise, just returns
the number as string.



## <a name="FormatPercent">func</a> [FormatPercent](/src/target/number.go?s=4009:4045#L145)
``` go
func FormatPercent(f float64) string
```
FormatPercent converts float64 value f to a percentage.
It can accept value more than 1 or negative.



## <a name="Fractional">func</a> [Fractional](/src/target/number.go?s=2619:2652#L91)
``` go
func Fractional(f float64) string
```
Fractional converts float64 value f to fractinal number
as a string, in forms of fractions and mixed fractions.



## <a name="IntComma">func</a> [IntComma](/src/target/number.go?s=626:653#L22)
``` go
func IntComma(n int) string
```
IntComma converts the integer n to a string containing commas
every three digits. It accepts octal or hex number.
Returned value is decimal in those cases too.
If the integer is octal or hex, decimal is also returned.



## <a name="IntWord">func</a> [IntWord](/src/target/number.go?s=1250:1278#L44)
``` go
func IntWord(n int64) string
```
IntWord converts a large integer n to a friendly text representation.
For example, 1000000 becomes '1.0 million', and '1200000000' becomes
'1.2 billion'. If n is negative, the number is just returned as string.
It accepts octal or hex number. Returned value is decimal in those cases
too.



## <a name="NaturalDay">func</a> [NaturalDay](/src/target/time.go?s=2603:2638#L92)
``` go
func NaturalDay(t time.Time) string
```
NaturalDay compares t to present day and returns tomorrow, today or yesterday
if applicable. Otherwise, it returns date as a string.



## <a name="NaturalDelta">func</a> [NaturalDelta](/src/target/time.go?s=264:316#L3)
``` go
func NaturalDelta(t time.Time, relative bool) string
```
NaturalDelta computes the natural representation of the amount
of time elapsed from now. If relative is true, "ago" or "from now"
is also returned depending on whether t is past or future.



## <a name="NaturalSize">func</a> [NaturalSize](/src/target/filesize.go?s=1885:1939#L50)
``` go
func NaturalSize(val int, options ...SetOption) string
```
NaturalSize formats val as a number of byteslike a human readable
filesize (eg. 10 kB).
Returns decimal suffixes (kB, MB...) by default.
If option Binary(true) is given, binary suffixes (KiB, MiB...) are used
and the base will be 2**10 instead of 10**3.
If option GNU(true) is given, GNU-style prefixes (K, M...) are used.
GNU had higher priority than Binary so if both options are true, Binary
is ignored.

Here's some examples to call NaturalSize:


	NaturalSize(123)
	NaturalSize(1978, GNU(true))
	NaturalSize(1000000000, Binary(true))

Passing both GNU and Binary options is allowed. But in this case GNU is
given priority and Binary is ignored.


	NaturalSize(742617000027, GNU(true), Binary(true))



## <a name="Ordinal">func</a> [Ordinal](/src/target/number.go?s=148:174#L2)
``` go
func Ordinal(n int) string
```
Ordinal converts the integer n to the ordinal number as string.



## <a name="SetLocale">func</a> [SetLocale](/src/target/i18n.go?s=403:427#L7)
``` go
func SetLocale(l string)
```
SetLocale sets l to current locale. l must be specified in
"lang code (ISO 639)"-"country code (ISO 3166-1)". (en-us, ja-jp,
for example.) Available translations are located at translations
directory as JSON file. If unsupported language is selected,
English (en-us) is used.




## <a name="FormatOption">type</a> [FormatOption](/src/target/filesize.go?s=205:259#L2)
``` go
type FormatOption struct {
    // contains filtered or unexported fields
}
```
FormatOption is the options NaturalSize can take.










## <a name="SetOption">type</a> [SetOption](/src/target/filesize.go?s=316:350#L8)
``` go
type SetOption func(*FormatOption)
```
SetOption is method definition to set FormatOption.







### <a name="Binary">func</a> [Binary](/src/target/filesize.go?s=475:505#L12)
``` go
func Binary(on bool) SetOption
```
Binary enables binary suffixes (KiB, MiB...) for NaturalSize function.
It also sets the base 2**10 instead of 10**3.


### <a name="GNU">func</a> [GNU](/src/target/filesize.go?s=636:663#L19)
``` go
func GNU(on bool) SetOption
```
GNU enables GNU-style prefixes (K, M...) for NaturalSize function.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
