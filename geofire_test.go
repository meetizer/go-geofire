package geofire

import "testing"

func TestEncodeGeoHash(t *testing.T) {
	input := Location{39.476456907107, -0.37986040048598}
	expected := "ezp8x67rh8"
	output, err := encodeGeoHash(&input)
	if err != nil {
		t.Errorf("encodeGeoHash(%f, %f) returned an unexpected error: %s", input.Lat, input.Lng, err.Error())
	} else if expected != output {
		t.Errorf("encodeGeoHash(%f, %f) != %s", input.Lat, input.Lng, expected)
	}
	input = Location{0, 0}
	expected = "7zzzzzzzzz"
	output, err = encodeGeoHash(&input)
	if err != nil {
		t.Errorf("encodeGeoHash(%f, %f) returned an unexpected error: %s", input.Lat, input.Lng, err.Error())
	} else if expected != output {
		t.Errorf("encodeGeoHash(%f, %f) != %s", input.Lat, input.Lng, expected)
	}
	/*input = Location{500, 0}
	expected = ""
	output, err = encodeGeoHash(&input)
	if err != nil {
		t.Errorf("encodeGeoHash(%f, %f) returned an unexpected error: %s", input.Lat, input.Lng, err.Error())
	} else if expected != output {
		t.Errorf("encodeGeoHash(%f, %f) != %s", input.Lat, input.Lng, expected)
	}
	input = Location{0, 500}
	expected = ""
	output, err = encodeGeoHash(&input)
	if err != nil {
		t.Errorf("encodeGeoHash(%f, %f) returned an unexpected error: %s", input.Lat, input.Lng, err.Error())
	} else if expected != output {
		t.Errorf("encodeGeoHash(%f, %f) != %s", input.Lat, input.Lng, expected)
	}*/
	input = Location{90, 180}
	expected = "zzzzzzzzzz"
	output, err = encodeGeoHash(&input)
	if err != nil {
		t.Errorf("encodeGeoHash(%f, %f) returned an unexpected error: %s", input.Lat, input.Lng, err.Error())
	} else if expected != output {
		t.Errorf("encodeGeoHash(%f, %f) != %s", input.Lat, input.Lng, expected)
	}
	input = Location{-90, -180}
	expected = "0000000000"
	output, err = encodeGeoHash(&input)
	if err != nil {
		t.Errorf("encodeGeoHash(%f, %f) returned an unexpected error: %s", input.Lat, input.Lng, err.Error())
	} else if expected != output {
		t.Errorf("encodeGeoHash(%f, %f) != %s", input.Lat, input.Lng, expected)
	}
}
