package geofire

import "math"

const (
	metersPerDegreeLatitude      = 110574
	earthMeridionalCircunference = 40007860
	earthEqRadius                = 6378137
	earthPolarRadius             = 6357852.3
	earthE2                      = 0.00669447819799
	epsilon                      = 1e-12
)

//GeoQuery ..
type GeoQuery struct {
	Center  locationGeoFire
	Radius  float64
	Geofire GeoFire
}

//Get ..
func (geoquery GeoQuery) Get() {
	latitudeDegrees := geoquery.Radius / metersPerDegreeLatitude
	latitudeNorth := math.Min(90, geoquery.Center.Latitude+latitudeDegrees)
	latitudeSouth := math.Max(-90, geoquery.Center.Latitude-latitudeDegrees)
	longitudeDeltaNorth := distanceToLongitudeDegrees(geoquery.Radius, latitudeNorth)
	longitudeDeltaSouth := distanceToLongitudeDegrees(geoquery.Radius, latitudeSouth)
	longitudeDelta := math.Max(longitudeDeltaNorth, longitudeDeltaSouth)
	_ = longitudeDelta //Solo para que no de error
	//""""https://meetizer.firebaseio.com/places.json?orderBy="g"&startAt="ezp8x64c8x"&endAt="sp3e3v4m0m""""
}

func distanceToLongitudeDegrees(distance float64, latitude float64) float64 {
	radians := toRadians(latitude)
	numerator := math.Cos(radians) * earthEqRadius * math.Pi / 180
	denominator := 1 / math.Sqrt(1-earthE2*math.Sin(radians)*math.Sin(radians))
	deltaDegrees := numerator * denominator
	if deltaDegrees < epsilon {
		if distance > 0 {
			return 360
		}
		return distance
	}
	return math.Min(360, distance/deltaDegrees)
}

func toRadians(number float64) float64 {
	return number / 180.0 * math.Pi
}

func wrapLongitude(longitude float64) float64 {
	if longitude >= -180 && longitude <= 180 {
		return longitude
	}
	adjusted := longitude + 180.0

	if adjusted > 0 {
		return math.Mod(adjusted, 360.0) - 180
	}
	return 180 - math.Mod(-adjusted, 360.0)
}
