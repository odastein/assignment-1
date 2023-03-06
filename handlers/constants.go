package handlers

const versionPath = "v1"
const rootPath = "/unisearcher/" + versionPath + "/"
const UniInfoPath = rootPath + "uniinfo/"
const NeighbourUniPath = rootPath + "neighbourunis/"
const DiagPath = rootPath + "diag/"

const UniversityURL = "http://universities.hipolabs.com"
const SearchNameURL = "/search?name="
const SearchCountryURL = "&search?country="

const RestCountriesPath = "https://restcountries.com"
const RestCountriesNamePath = RestCountriesPath + "/v3.1/name/"
const RestCountriesAlphaPath = RestCountriesPath + "/v3.1/alpha/"
const RestCountriesTextPath = "?fullText=true"
const RestCountriesFieldsPath = "?fields=name,maps,languages,borders"

var countryNames = map[string]string{
	"Bolivia":                "Bolivia, Plurinational State of",
	"Syria":                  "Syrian Arab Republic",
	"Palestine":              "Palestine, State of",
	"Macau":                  "Macao",
	"Ivory Coast":            "Côte d'Ivoire",
	"Moldova":                "Moldova, Republic of",
	"Republic of the Congo":  "Congo",
	"Czechia":                "Czech Republic",
	"Vietnam":                "Viet Nam",
	"DR Congo":               "Congo, the Democratic Republic of the",
	"South Korea":            "Korea, Republic of",
	"North Korea":            "Korea, Democratic People's Republic of",
	"Russia":                 "Russian Federation",
	"Laos":                   "Lao People's Democratic Republic",
	"Eswatini":               "Swaziland",
	"Venezuela":              "Venezuela, Bolivarian Republic of",
	"Tanzania":               "Tanzania, United Republic of",
	"British Virgin Islands": "Virgin Islands, British",
	"Vatican City":           "Holy See (Vatican City State)",
	"Brunei":                 "Brunei Darussalam"}
