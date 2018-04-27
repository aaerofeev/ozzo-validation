// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package is

import (
	"testing"

	"github.com/aaerofeev/ozzo-validation"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	tests := []struct {
		tag            string
		rule           validation.Rule
		valid, invalid string
		err            string
	}{
		{"Email", Email, "test@example.com", "example.com", "email"},
		{"URL", URL, "http://example.com", "examplecom", "url"},
		{"RequestURL", RequestURL, "http://example.com", "examplecom", "request_url"},
		{"RequestURI", RequestURI, "http://example.com", "examplecom", "request_uri"},
		{"Alpha", Alpha, "abcd", "ab12", "alpha"},
		{"Digit", Digit, "123", "12ab", "digits"},
		{"Alphanumeric", Alphanumeric, "abc123", "abc.123", "alphanumeric"},
		{"UTFLetter", UTFLetter, "ａｂｃ", "１２３", "utf_alpha"},
		{"UTFDigit", UTFDigit, "１２３", "ａｂｃ", "utf_digits"},
		{"UTFNumeric", UTFNumeric, "１２３", "ａｂｃ.１２３", "utf_numeric"},
		{"UTFLetterNumeric", UTFLetterNumeric, "ａｂｃ１２３", "ａｂｃ.１２３", "utf_alphanumeric"},
		{"LowerCase", LowerCase, "ａｂc", "Aｂｃ", "lowercase"},
		{"UpperCase", UpperCase, "ABC", "ABｃ", "uppercase"},
		{"IP", IP, "74.125.19.99", "74.125.19.999", "ip"},
		{"IPv4", IPv4, "74.125.19.99", "2001:4860:0:2001::68", "ipv4"},
		{"IPv6", IPv6, "2001:4860:0:2001::68", "74.125.19.99", "ipv6"},
		{"MAC", MAC, "0123.4567.89ab", "74.125.19.99", "mac"},
		{"Subdomain", Subdomain, "example-subdomain", "example.com", "subdomain"},
		{"Domain", Domain, "example-domain.com", "localhost", "domain"},
		{"DNSName", DNSName, "example.com", "abc%", "dns"},
		{"Host", Host, "example.com", "abc%", "ip_or_dns"},
		{"Port", Port, "123", "99999", "port"},
		{"Latitude", Latitude, "23.123", "100", "latitude"},
		{"Longitude", Longitude, "123.123", "abc", "longitude"},
		{"SSN", SSN, "100-00-1000", "100-0001000", "ssn"},
		{"Semver", Semver, "1.0.0", "1.0.0.0", "semver"},
		{"ISBN", ISBN, "1-61729-085-8", "1-61729-085-81", "isbn"},
		{"ISBN10", ISBN10, "1-61729-085-8", "1-61729-085-81", "isbn10"},
		{"ISBN13", ISBN13, "978-4-87311-368-5", "978-4-87311-368-a", "isbn13"},
		{"UUID", UUID, "a987fbc9-4bed-3078-cf07-9141ba07c9f1", "a987fbc9-4bed-3078-cf07-9141ba07c9f3a", "uuid"},
		{"UUIDv3", UUIDv3, "b987fbc9-4bed-3078-cf07-9141ba07c9f3", "b987fbc9-4bed-4078-cf07-9141ba07c9f3", "uuid3"},
		{"UUIDv4", UUIDv4, "57b73598-8764-4ad0-a76a-679bb6640eb1", "b987fbc9-4bed-3078-cf07-9141ba07c9f3", "uuid4"},
		{"UUIDv5", UUIDv5, "987fbc97-4bed-5078-af07-9141ba07c9f3", "b987fbc9-4bed-3078-cf07-9141ba07c9f3", "uuid5"},
		{"MongoID", MongoID, "507f1f77bcf86cd799439011", "507f1f77bcf86cd79943901", "mongodb_object_id"},
		{"CreditCard", CreditCard, "375556917985515", "375556917985516", "credit_card"},
		{"JSON", JSON, "[1, 2]", "[1, 2,]", "json"},
		{"ASCII", ASCII, "abc", "ａabc", "ascii_chars"},
		{"PrintableASCII", PrintableASCII, "abc", "ａabc", "ascii_chars_print"},
		{"CountryCode2", CountryCode2, "US", "XY", "country_code2"},
		{"CountryCode3", CountryCode3, "USA", "XYZ", "country_code3"},
		{"DialString", DialString, "localhost.local:1", "localhost.loc:100000", "dial"},
		{"DataURI", DataURI, "data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", "image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", "base64_uri"},
		{"Base64", Base64, "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", "image", "base64"},
		{"Multibyte", Multibyte, "ａｂｃ", "abc", "multibyte_chars"},
		{"FullWidth", FullWidth, "３ー０", "abc", "full_width_chars"},
		{"HalfWidth", HalfWidth, "abc123い", "００１１", "half_width_chars"},
		{"VariableWidth", VariableWidth, "３ー０123", "abc", "both_width_chars"},
		{"Hexadecimal", Hexadecimal, "FEF", "FTF", "hex"},
		{"HexColor", HexColor, "F00", "FTF", "hex_color"},
		{"RGBColor", RGBColor, "rgb(100, 200, 1)", "abc", "rgb_color"},
		{"Int", Int, "100", "1.1", "number_integer"},
		{"Float", Float, "1.1", "a.1", "number_float"},
		{"VariableWidth", VariableWidth, "", "", ""},
	}

	for _, test := range tests {
		err := test.rule.Validate("")
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(test.valid)
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(&test.valid)
		assert.Nil(t, err, test.tag)
		err = test.rule.Validate(test.invalid)
		assertError(t, test.err, err, test.tag)
		err = test.rule.Validate(&test.invalid)
		assertError(t, test.err, err, test.tag)
	}
}

func assertError(t *testing.T, expected string, err error, tag string) {
	if expected == "" {
		assert.Nil(t, err, tag)
	} else if assert.NotNil(t, err, tag) {
		assert.Equal(t, expected, err.Error(), tag)
	}
}
