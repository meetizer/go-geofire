package geofire

import (
	"errors"
	"time"

	"github.com/alvivi/firego"
)

const (
	precision int    = 10
	gBASE32   string = "0123456789bcdefghjkmnpqrstuvwxyz"
)

type geoFireObject struct {
	Hash      string          `json:"g"`
	Location  locationGeoFire `json:"l"`
	Name      string          `json:"name"`
	TimeStamp int64           `json:"timestamp"`
}

type locationGeoFire struct {
	Latitude  float64 `json:"0"`
	Longitude float64 `json:"1"`
}

// Location represents a geographical position
type Location struct {
	// Latitude and longitude of the position
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

//Place represents a place with a defined name and the location of the place
type Place struct {
	Name     string    `json:"name"`
	Location *Location `json:"location"`
}

//GeoFire ..
type GeoFire struct {
	firebaseRef *firego.Firebase
}

func (geoFire GeoFire) queryAtLocation(center locationGeoFire, radius float64) GeoQuery {
	var newGeoQuery GeoQuery
	newGeoQuery.Center = center
	newGeoQuery.Radius = radius
	newGeoQuery.Geofire = geoFire
	return newGeoQuery
}

/**
 * Generates a geohash of the specified precision/string length from the  [latitude, longitude]
 * pair, specified as an array.
 *
 * @param {Array.<number>} location The [latitude, longitude] pair to encode into a geohash.
 * @param {number=} precision The length of the geohash to create. If no precision is
 * specified, the global default is used.
 * @return {string} The geohash of the inputted location.
 */

//SetLocation ..
func SetLocation(place *Place, key string, ref *firego.Firebase) error {
	hash, err := encodeGeoHash(place.Location)
	if err != nil {
		return err
	}
	var auxlocation locationGeoFire
	var newGeofireObject geoFireObject
	auxlocation.Latitude = place.Location.Lat
	auxlocation.Longitude = place.Location.Lng
	newGeofireObject.Hash = hash
	newGeofireObject.Location = auxlocation
	newGeofireObject.Name = place.Name
	newGeofireObject.TimeStamp = time.Now().Unix()
	placeRef := ref.Child("places").Child(key)
	placeRef.Update(newGeofireObject)

	return nil
}

func encodeGeoHash(location *Location) (string, error) {
	err := validateLocation(location)
	if err != nil {
		return "", err
	}

	var hash string
	var hashVal = 0
	var bits = 0
	var even = true
	var latitudeRange = map[string]float64{"min": -90, "max": 90}
	var longitudeRange = map[string]float64{"min": -180, "max": 180}
	for len(hash) < precision {
		var nrange map[string]float64
		var val float64
		if even {
			val = location.Lng
			nrange = longitudeRange
		} else {
			val = location.Lat
			nrange = latitudeRange
		}
		mid := float64((nrange["min"] + nrange["max"]) / 2)
		/* jshint -W016 */
		if val > mid {
			hashVal = (hashVal << 1) + 1
			nrange["min"] = mid
		} else {
			hashVal = (hashVal << 1) + 0
			nrange["max"] = mid
		}
		/* jshint +W016 */
		even = !even
		if bits < 4 {
			bits++
		} else {
			bits = 0
			hash += string(gBASE32[hashVal])
			hashVal = 0
		}
	}

	return hash, nil
}

func validateLocation(location *Location) error {
	if location.Lat < -90 || location.Lat > 90 {
		return errors.New("Latitude must be between [-90, 90]")
	}
	if location.Lng < -180 || location.Lng > 180 {
		return errors.New("Longitude must be between [-180, 180]")
	}
	return nil
}
