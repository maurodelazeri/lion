package common

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

// Const declarations for common.go operations
const (
	HashSHA1 = iota
	HashSHA256
	HashSHA512
	HashSHA512_384
	SatoshisPerBTC = 100000000
	SatoshisPerLTC = 100000000
	WeiPerEther    = 1000000000000000000
)

// GetMD5 returns a MD5 hash of a byte array
func GetMD5(input []byte) []byte {
	hash := md5.New()
	hash.Write(input)
	return hash.Sum(nil)
}

// GetSHA512 returns a SHA512 hash of a byte array
func GetSHA512(input []byte) []byte {
	sha := sha512.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// GetSHA256 returns a SHA256 hash of a byte array
func GetSHA256(input []byte) []byte {
	sha := sha256.New()
	sha.Write(input)
	return sha.Sum(nil)
}

// GetHMAC returns a keyed-hash message authentication code using the desired
// hashtype
func GetHMAC(hashType int, input, key []byte) []byte {
	var hash func() hash.Hash

	switch hashType {
	case HashSHA1:
		{
			hash = sha1.New
		}
	case HashSHA256:
		{
			hash = sha256.New
		}
	case HashSHA512:
		{
			hash = sha512.New
		}
	case HashSHA512_384:
		{
			hash = sha512.New384
		}
	}

	hmac := hmac.New(hash, []byte(key))
	hmac.Write(input)
	return hmac.Sum(nil)
}

// HexEncodeToString takes in a hexadecimal byte array and returns a string
func HexEncodeToString(input []byte) string {
	return hex.EncodeToString(input)
}

// Base64Decode takes in a Base64 string and returns a byte array and an error
func Base64Decode(input string) ([]byte, error) {
	result, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Base64Encode takes in a byte array then returns an encoded base64 string
func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

// StringSliceDifference concatenates slices together based on its index and
// returns an individual string array
func StringSliceDifference(slice1 []string, slice2 []string) []string {
	var diff []string
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}

// StringContains checks a substring if it contains your input then returns a
// bool
func StringContains(input, substring string) bool {
	return strings.Contains(input, substring)
}

// StringDataContains checks the substring array with an input and returns a bool
func StringDataContains(haystack []string, needle string) bool {
	data := strings.Join(haystack, ",")
	return strings.Contains(data, needle)
}

// StringDataCompare data checks the substring array with an input and returns a bool
func StringDataCompare(haystack []string, needle string) bool {
	for x := range haystack {
		if haystack[x] == needle {
			return true
		}
	}
	return false
}

// StringDataContainsUpper checks the substring array with an input and returns
// a bool irrespective of lower or upper case strings
func StringDataContainsUpper(haystack []string, needle string) bool {
	for _, data := range haystack {
		if strings.Contains(StringToUpper(data), StringToUpper(needle)) {
			return true
		}
	}
	return false
}

// JoinStrings joins an array together with the required separator and returns
// it as a string
func JoinStrings(input []string, separator string) string {
	return strings.Join(input, separator)
}

// SplitStrings splits blocks of strings from string into a string array using
// a separator ie "," or "_"
func SplitStrings(input, separator string) []string {
	return strings.Split(input, separator)
}

// TrimString trims unwanted prefixes or postfixes
func TrimString(input, cutset string) string {
	return strings.Trim(input, cutset)
}

// ReplaceString replaces a string with another
func ReplaceString(input, old, new string, n int) string {
	return strings.Replace(input, old, new, n)
}

// StringToUpper changes strings to uppercase
func StringToUpper(input string) string {
	return strings.ToUpper(input)
}

// StringToLower changes strings to lowercase
func StringToLower(input string) string {
	return strings.ToLower(input)
}

// RoundFloat rounds your floating point number to the desired decimal place
func RoundFloat(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	intermed += .5
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed--
	}
	if frac >= x {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

// IsEnabled takes in a boolean param  and returns a string if it is enabled
// or disabled
func IsEnabled(isEnabled bool) string {
	if isEnabled {
		return "Enabled"
	}
	return "Disabled"
}

// IsValidCryptoAddress validates your cryptocurrency address string using the
// regexp package // Validation issues occurring because "3" is contained in
// litecoin and Bitcoin addresses - non-fatal
func IsValidCryptoAddress(address, crypto string) (bool, error) {
	switch StringToLower(crypto) {
	case "btc":
		return regexp.MatchString("^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$", address)
	case "ltc":
		return regexp.MatchString("^[L3M][a-km-zA-HJ-NP-Z1-9]{25,34}$", address)
	case "eth":
		return regexp.MatchString("^0x[a-km-z0-9]{40}$", address)
	default:
		return false, errors.New("Invalid crypto currency")
	}
}

// YesOrNo returns a boolean variable to check if input is "y" or "yes"
func YesOrNo(input string) bool {
	if StringToLower(input) == "y" || StringToLower(input) == "yes" {
		return true
	}
	return false
}

// CalculateAmountWithFee returns a calculated fee included amount on fee
func CalculateAmountWithFee(amount, fee float64) float64 {
	return amount + CalculateFee(amount, fee)
}

// CalculateFee returns a simple fee on amount
func CalculateFee(amount, fee float64) float64 {
	return amount * (fee / 100)
}

// CalculatePercentageGainOrLoss returns the percentage rise over a certain
// period
func CalculatePercentageGainOrLoss(priceNow, priceThen float64) float64 {
	return (priceNow - priceThen) / priceThen * 100
}

// CalculatePercentageDifference returns the percentage of difference between
// multiple time periods
func CalculatePercentageDifference(amount, secondAmount float64) float64 {
	return (amount - secondAmount) / ((amount + secondAmount) / 2) * 100
}

// CalculateNetProfit returns net profit
func CalculateNetProfit(amount, priceThen, priceNow, costs float64) float64 {
	return (priceNow * amount) - (priceThen * amount) - costs
}

// SendHTTPRequest sends a request using the http package and returns a response
// as a string and an error
func SendHTTPRequest(method, path string, headers map[string]string, body io.Reader) (string, error) {
	result := strings.ToUpper(method)

	if result != "POST" && result != "GET" && result != "DELETE" {
		return "", errors.New("invalid HTTP method specified")
	}

	req, err := http.NewRequest(method, path, body)

	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	return string(contents), nil
}

// SendHTTPGetRequest sends a simple get request using a url string & JSON
// decodes the response into a struct pointer you have supplied. Returns an error
// on failure.
func SendHTTPGetRequest(url string, jsonDecode, isVerbose bool, result interface{}) error {
	if isVerbose {
		log.Println("Raw URL: ", url)
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("common.SendHTTPGetRequest() error: HTTP status code %d", res.StatusCode)
	}

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if isVerbose {
		log.Println("Raw Resp: ", string(contents[:]))
	}

	defer res.Body.Close()

	if jsonDecode {
		err := JSONDecode(contents, result)
		if err != nil {
			log.Println(string(contents[:]))
			return err
		}
	}

	return nil
}

// JSONEncode encodes structure data into JSON
func JSONEncode(v interface{}) ([]byte, error) {
	return ffjson.Marshal(v)
}

// JSONDecode decodes JSON data into a structure
func JSONDecode(data []byte, to interface{}) error {
	if !StringContains(reflect.ValueOf(to).Type().String(), "*") {
		return errors.New("json decode error - memory address not supplied")
	}
	return ffjson.Unmarshal(data, to)
}

// EncodeURLValues concatenates url values onto a url string and returns a
// string
func EncodeURLValues(url string, values url.Values) string {
	path := url
	if len(values) > 0 {
		path += "?" + values.Encode()
	}
	return path
}

// ExtractHost returns the hostname out of a string
func ExtractHost(address string) string {
	host := SplitStrings(address, ":")[0]
	if host == "" {
		return "localhost"
	}
	return host
}

// ExtractPort returns the port name out of a string
func ExtractPort(host string) int {
	portStr := SplitStrings(host, ":")[1]
	port, _ := strconv.Atoi(portStr)
	return port
}

// OutputCSV dumps data into a file as comma-separated values
func OutputCSV(path string, data [][]string) error {
	_, err := ReadFile(path)
	if err != nil {
		errTwo := WriteFile(path, nil)
		if errTwo != nil {
			return errTwo
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	err = writer.WriteAll(data)
	if err != nil {
		return err
	}

	writer.Flush()
	file.Close()
	return nil
}

// UnixTimestampToTime returns time.time
func UnixTimestampToTime(timeint64 int64) time.Time {
	return time.Unix(timeint64, 0)
}

// UnixTimestampStrToTime returns a time.time and an error
func UnixTimestampStrToTime(timeStr string) (time.Time, error) {
	i, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(i, 0), nil
}

// ReadFile reads a file and returns read data as byte array.
func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// WriteFile writes selected data to a file and returns an error
func WriteFile(file string, data []byte) error {
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// RemoveFile removes a file
func RemoveFile(file string) error {
	return os.Remove(file)
}

// GetURIPath returns the path of a URL given a URI
func GetURIPath(uri string) string {
	urip, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	if urip.RawQuery != "" {
		return fmt.Sprintf("%s?%s", urip.Path, urip.RawQuery)
	}
	return urip.Path
}

// GetExecutablePath returns the executables launch path
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// GetOSPathSlash returns the slash used by the operating systems
// file system
func GetOSPathSlash() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

// GetPort is deprecated, use GetFreePort instead
// Ask the kernel for a free open port that is ready to use
func GetPort() int {
	port, err := GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

// LogError ...
func LogError(format string, v ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, v...))
}

// CheckNanoSecondTimeDiff event must be provided as microsecond
func CheckNanoSecondTimeDiff(event int64) time.Duration {
	now := time.Now()
	timeFromTS := time.Unix(0, event*int64(time.Nanosecond)).UTC()
	return now.Sub(timeFromTS)
}

// CheckSecondTimeDiff event must be provided as seconds
func CheckSecondTimeDiff(event int64) time.Duration {
	now := time.Now()
	timeFromTS := time.Unix(0, event*int64(time.Second)).UTC()
	return now.Sub(timeFromTS)
}

// MakeTimestampFromInt64 create standart timestamp
func MakeTimestampFromInt64(ref int64) time.Time {
	return time.Unix(0, ref*int64(time.Nanosecond)).UTC()
}

// MakeTimestampFromTime create standart timestamp
func MakeTimestampFromTime(ref time.Time) int64 {
	return ref.UTC().UnixNano() / int64(time.Nanosecond)
}

// MakeTimestamp ...
func MakeTimestamp() int64 {
	return time.Now().UTC().UnixNano() / int64(time.Nanosecond)
}

// CompressLZW ...
// https://www.cs.duke.edu/csed/curious/compression/lzw.html
func CompressLZW(testStr string) []int {
	code := 256
	dictionary := make(map[string]int)
	for i := 0; i < 256; i++ {
		dictionary[string(i)] = i
	}
	currChar := ""
	result := make([]int, 0)
	for _, c := range []byte(testStr) {
		phrase := currChar + string(c)
		if _, isTrue := dictionary[phrase]; isTrue {
			currChar = phrase
		} else {
			result = append(result, dictionary[currChar])
			dictionary[phrase] = code
			code++
			currChar = string(c)
		}
	}
	if currChar != "" {
		result = append(result, dictionary[currChar])
	}
	return result
}

// DecompressLZW ...
func DecompressLZW(compressed []int) string {
	code := 256
	dictionary := make(map[int]string)
	for i := 0; i < 256; i++ {
		dictionary[i] = string(i)
	}
	currChar := string(compressed[0])
	result := currChar
	for _, element := range compressed[1:] {
		var word string
		if x, ok := dictionary[element]; ok {
			word = x
		} else if element == code {
			word = currChar + currChar[:1]
		} else {
			panic(fmt.Sprintf("Bad compressed element: %d", element))
		}
		result += word
		dictionary[code] = currChar + word[:1]
		code++

		currChar = word
	}
	return result
}

// CompressFlate ...
func CompressFlate(data []byte) []byte {
	buf := bytes.Buffer{}
	w, _ := flate.NewWriter(&buf, flate.BestCompression)
	w.Write(data)
	w.Close()
	return buf.Bytes()
}

// DecompressFlate ...
func DecompressFlate(obj []byte) []byte {
	data := bytes.NewReader(obj)
	r := flate.NewReader(data)
	enflate, _ := ioutil.ReadAll(r)
	return enflate
}

// CompressGzip ...
func CompressGzip(data []byte) []byte {
	buf := bytes.Buffer{}
	w, _ := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	w.Write(data)
	w.Close()
	return buf.Bytes()
}

// CompressLZW2 ...
func CompressLZW2(data []byte) []byte {
	buf := bytes.Buffer{}
	w := lzw.NewWriter(&buf, lzw.LSB, 8)
	w.Write(data)
	w.Close()
	return buf.Bytes()
}

// CompressZlib ...
func CompressZlib(data []byte) []byte {
	buf := bytes.Buffer{}
	w, _ := zlib.NewWriterLevel(&buf, zlib.BestCompression)
	w.Write(data)
	w.Close()
	return buf.Bytes()
}

// GzipDecode ...
func GzipDecode(in []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(in))
	defer reader.Close()
	return ioutil.ReadAll(reader)
}
