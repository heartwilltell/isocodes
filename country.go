package isocodes

import (
	"sort"
	"strings"
)

// CountryCode represents an ISO 3166-1 Alpha2 country code.
type CountryCode byte

// String returns an Alpha2 string representation of the code.
func (c CountryCode) String() string { return countryCodesDetails[c].Alpha2 }

// Alpha3 returns an Alpha3 string representation of the code.
func (c CountryCode) Alpha3() string { return countryCodesDetails[c].Alpha3 }

// Name returns a country name related to CountryCode.
func (c CountryCode) Name() string { return countryCodesDetails[c].Name }

// Number returns a country number related to the CountryCode.
func (c CountryCode) Number() string { return countryCodesDetails[c].Number }

// Flag returns an emoji flag for the country code.
func (c CountryCode) Flag() string { return countryCodesDetails[c].Flag }

func (c *CountryCode) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return ErrUnmarshalJSON
	}

	code, ok := stringToCountryCode[string(b)]
	if !ok {
		return ErrUnmarshalJSON
	}

	*c = code

	return nil
}

func (c CountryCode) MarshalJSON() ([]byte, error) {
	code := c.String()
	if code == "" {
		return nil, ErrMarshalJSON
	}

	return []byte(`"` + code + `"`), nil
}

func (c *CountryCode) UnmarshalText(b []byte) error { return c.UnmarshalJSON(b) }

func (c CountryCode) MarshalText() ([]byte, error) { return c.MarshalJSON() }

// CountryCodeDetails represents detailed information related to country code.
type CountryCodeDetails struct {
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Flag   string `json:"flag"`
	Number string `json:"number"`
	Name   string `json:"name"`
}

// StringToCountryCode takes string representation of ISO 3166-1 Alpha2 country
// code and returns a CountryCode.
func StringToCountryCode(code string) (CountryCode, error) {
	if len(code) > 2 {
		return 0, ErrInvalidStringCode
	}

	c, ok := stringToCountryCode[strings.ToUpper(code)]
	if !ok {
		return 0, ErrInvalidStringCode
	}

	return c, nil
}

// ListCountryCodes returns a list of CountryCode.
func ListCountryCodes() []CountryCode {
	codes := make([]CountryCode, 0, len(stringToCountryCode))

	for _, c := range stringToCountryCode {
		codes = append(codes, c)
	}

	sort.Slice(codes, func(i, j int) bool {
		return codes[i].String() < codes[j].String()
	})

	return codes
}

// Enumeration of ISO 3166-1 country codes.
const (
	// AD represents the ISO 3166-1 Alpha-2 country code of Andorra.
	AD CountryCode = iota + 1
	// AE represents the ISO 3166-1 Alpha-2 country code of the United Arab Emirates.
	AE
	// AF represents the ISO 3166-1 Alpha-2 country code of Afghanistan.
	AF
	// AG represents the ISO 3166-1 Alpha-2 country code of Antigua and Barbuda.
	AG
	// AI represents the ISO 3166-1 Alpha-2 country code of Anguilla.
	AI
	// AL represents the ISO 3166-1 Alpha-2 country code of Albania.
	AL
	// AM represents the ISO 3166-1 Alpha-2 country code of Armenia.
	AM
	// AO represents the ISO 3166-1 Alpha-2 country code of Angola.
	AO
	// AQ represents the ISO 3166-1 Alpha-2 country code of Antarctica.
	AQ
	// AR represents the ISO 3166-1 Alpha-2 country code of Argentina.
	AR
	// AS represents the ISO 3166-1 Alpha-2 country code of American Samoa.
	AS
	// AT represents the ISO 3166-1 Alpha-2 country code of Austria.
	AT
	// AU represents the ISO 3166-1 Alpha-2 country code of Australia.
	AU
	// AW represents the ISO 3166-1 Alpha-2 country code of Aruba.
	AW
	// AX represents the ISO 3166-1 Alpha-2 country code of Åland Islands.
	AX
	// AZ represents the ISO 3166-1 Alpha-2 country code of Azerbaijan.
	AZ
	// BA represents the ISO 3166-1 Alpha-2 country code of Bosnia and Herzegovina.
	BA
	// BB represents the ISO 3166-1 Alpha-2 country code of Barbados.
	BB
	// BD represents the ISO 3166-1 Alpha-2 country code of Bangladesh.
	BD
	// BE represents the ISO 3166-1 Alpha-2 country code of Belgium.
	BE
	// BF represents the ISO 3166-1 Alpha-2 country code of Burkina Faso.
	BF
	// BG represents the ISO 3166-1 Alpha-2 country code of Bulgaria.
	BG
	// BH represents the ISO 3166-1 Alpha-2 country code of Bahrain.
	BH
	// BI represents the ISO 3166-1 Alpha-2 country code of Burundi.
	BI
	// BJ represents the ISO 3166-1 Alpha-2 country code of Benin.
	BJ
	// BL represents the ISO 3166-1 Alpha-2 country code of Saint Barthélemy.
	BL
	// BM represents the ISO 3166-1 Alpha-2 country code of Bermuda.
	BM
	// BN represents the ISO 3166-1 Alpha-2 country code of Brunei Darussalam.
	BN
	// BO represents the ISO 3166-1 Alpha-2 country code of Bolivia (Plurinational State of).
	BO
	// BQ represents the ISO 3166-1 Alpha-2 country code of Bonaire, Sint Eustatius and Saba.
	BQ
	// BR represents the ISO 3166-1 Alpha-2 country code of Brazil.
	BR
	// BS represents the ISO 3166-1 Alpha-2 country code of the Bahamas.
	BS
	// BT represents the ISO 3166-1 Alpha-2 country code of Bhutan.
	BT
	// BV represents the ISO 3166-1 Alpha-2 country code of Bouvet Island.
	BV
	// BW represents the ISO 3166-1 Alpha-2 country code of Botswana.
	BW
	// BY represents the ISO 3166-1 Alpha-2 country code of Belarus.
	BY
	// BZ represents the ISO 3166-1 Alpha-2 country code of Belize.
	BZ
	// CA represents the ISO 3166-1 Alpha-2 country code of Canada.
	CA
	// CC represents the ISO 3166-1 Alpha-2 country code of Cocos (Keeling) Islands.
	CC
	// CD represents the ISO 3166-1 Alpha-2 country code of Democratic Republic of the Congo.
	CD
	// CF represents the ISO 3166-1 Alpha-2 country code of the Central African Republic.
	CF
	// CG represents the ISO 3166-1 Alpha-2 country code of Congo.
	CG
	// CH represents the ISO 3166-1 Alpha-2 country code of Switzerland.
	CH
	// CI represents the ISO 3166-1 Alpha-2 country code of Côte d'Ivoire.
	CI
	// CK represents the ISO 3166-1 Alpha-2 country code of Cook Islands.
	CK
	// CL represents the ISO 3166-1 Alpha-2 country code of Chile.
	CL
	// CM represents the ISO 3166-1 Alpha-2 country code of Cameroon.
	CM
	// CN represents the ISO 3166-1 Alpha-2 country code of China.
	CN
	// CO represents the ISO 3166-1 Alpha-2 country code of Colombia.
	CO
	// CR represents the ISO 3166-1 Alpha-2 country code of Costa Rica.
	CR
	// CU represents the ISO 3166-1 Alpha-2 country code of Cuba.
	CU
	// CV represents the ISO 3166-1 Alpha-2 country code of Cabo Verde.
	CV
	// CW represents the ISO 3166-1 Alpha-2 country code of Curaçao.
	CW
	// CX represents the ISO 3166-1 Alpha-2 country code of Christmas Island.
	CX
	// CY represents the ISO 3166-1 Alpha-2 country code of Cyprus.
	CY
	// CZ represents the ISO 3166-1 Alpha-2 country code of Czechia.
	CZ
	// DE represents the ISO 3166-1 Alpha-2 country code of Germany.
	DE
	// DJ represents the ISO 3166-1 Alpha-2 country code of Djibouti.
	DJ
	// DK represents the ISO 3166-1 Alpha-2 country code of Denmark.
	DK
	// DM represents the ISO 3166-1 Alpha-2 country code of Dominica.
	DM
	// DO represents the ISO 3166-1 Alpha-2 country code of the Dominican Republic.
	DO
	// DZ represents the ISO 3166-1 Alpha-2 country code of Algeria.
	DZ
	// EC represents the ISO 3166-1 Alpha-2 country code of Ecuador.
	EC
	// EE represents the ISO 3166-1 Alpha-2 country code of Estonia.
	EE
	// EG represents the ISO 3166-1 Alpha-2 country code of Egypt.
	EG
	// EH represents the ISO 3166-1 Alpha-2 country code of Western Sahara.
	EH
	// ER represents the ISO 3166-1 Alpha-2 country code of Eritrea.
	ER
	// ES represents the ISO 3166-1 Alpha-2 country code of Spain.
	ES
	// ET represents the ISO 3166-1 Alpha-2 country code of Ethiopia.
	ET
	// FI represents the ISO 3166-1 Alpha-2 country code of Finland.
	FI
	// FJ represents the ISO 3166-1 Alpha-2 country code of Fiji.
	FJ
	// FK represents the ISO 3166-1 Alpha-2 country code of the Falkland Islands (Malvinas).
	FK
	// FM represents the ISO 3166-1 Alpha-2 country code of Micronesia (Federated States of).
	FM
	// FO represents the ISO 3166-1 Alpha-2 country code of Faroe Islands.
	FO
	// FR represents the ISO 3166-1 Alpha-2 country code of France.
	FR
	// GA represents the ISO 3166-1 Alpha-2 country code of Gabon.
	GA
	// GB represents the ISO 3166-1 Alpha-2 country code of the United Kingdom of Great Britain and Northern Ireland.
	GB
	// GD represents the ISO 3166-1 Alpha-2 country code of Grenada.
	GD
	// GE represents the ISO 3166-1 Alpha-2 country code of Georgia.
	GE
	// GF represents the ISO 3166-1 Alpha-2 country code of French Guiana.
	GF
	// GG represents the ISO 3166-1 Alpha-2 country code of Guernsey.
	GG
	// GH represents the ISO 3166-1 Alpha-2 country code of Ghana.
	GH
	// GI represents the ISO 3166-1 Alpha-2 country code of Gibraltar.
	GI
	// GL represents the ISO 3166-1 Alpha-2 country code of Greenland.
	GL
	// GM represents the ISO 3166-1 Alpha-2 country code of the Gambia.
	GM
	// GN represents the ISO 3166-1 Alpha-2 country code of Guinea.
	GN
	// GP represents the ISO 3166-1 Alpha-2 country code of Guadeloupe.
	GP
	// GQ represents the ISO 3166-1 Alpha-2 country code of Equatorial Guinea.
	GQ
	// GR represents the ISO 3166-1 Alpha-2 country code of Greece.
	GR
	// GS represents the ISO 3166-1 Alpha-2 country code of South Georgia and the South Sandwich Islands.
	GS
	// GT represents the ISO 3166-1 Alpha-2 country code of Guatemala.
	GT
	// GU represents the ISO 3166-1 Alpha-2 country code of Guam.
	GU
	// GW represents the ISO 3166-1 Alpha-2 country code of Guinea-Bissau.
	GW
	// GY represents the ISO 3166-1 Alpha-2 country code of Guyana.
	GY
	// HK represents the ISO 3166-1 Alpha-2 country code of Hong Kong.
	HK
	// HM represents the ISO 3166-1 Alpha-2 country code of Heard Island and McDonald Islands.
	HM
	// HN represents the ISO 3166-1 Alpha-2 country code of Honduras.
	HN
	// HR represents the ISO 3166-1 Alpha-2 country code of Croatia.
	HR
	// HT represents the ISO 3166-1 Alpha-2 country code of Haiti.
	HT
	// HU represents the ISO 3166-1 Alpha-2 country code of Hungary.
	HU
	// ID represents the ISO 3166-1 Alpha-2 country code of Indonesia.
	ID
	// IE represents the ISO 3166-1 Alpha-2 country code of Ireland.
	IE
	// IL represents the ISO 3166-1 Alpha-2 country code of Israel.
	IL
	// IM represents the ISO 3166-1 Alpha-2 country code of the Isle of Man.
	IM
	// IN represents the ISO 3166-1 Alpha-2 country code of India.
	IN
	// IO represents the ISO 3166-1 Alpha-2 country code of British Indian Ocean Territory.
	IO
	// IQ represents the ISO 3166-1 Alpha-2 country code of Iraq.
	IQ
	// IR represents the ISO 3166-1 Alpha-2 country code of Iran (Islamic Republic of).
	IR
	// IS represents the ISO 3166-1 Alpha-2 country code of Iceland.
	IS
	// IT represents the ISO 3166-1 Alpha-2 country code of Italy.
	IT
	// JE represents the ISO 3166-1 Alpha-2 country code of Jersey.
	JE
	// JM represents the ISO 3166-1 Alpha-2 country code of Jamaica.
	JM
	// JO represents the ISO 3166-1 Alpha-2 country code of Jordan.
	JO
	// JP represents the ISO 3166-1 Alpha-2 country code of Japan.
	JP
	// KE represents the ISO 3166-1 Alpha-2 country code of Kenya.
	KE
	// KG represents the ISO 3166-1 Alpha-2 country code of Kyrgyzstan.
	KG
	// KH represents the ISO 3166-1 Alpha-2 country code of Cambodia.
	KH
	// KI represents the ISO 3166-1 Alpha-2 country code of Kiribati.
	KI
	// KM represents the ISO 3166-1 Alpha-2 country code of the Comoros.
	KM
	// KN represents the ISO 3166-1 Alpha-2 country code of Saint Kitts and Nevis.
	KN
	// KP represents the ISO 3166-1 Alpha-2 country code of Korea (Democratic People's Republic of).
	KP
	// KR represents the ISO 3166-1 Alpha-2 country code of Korea, Republic of.
	KR
	// KW represents the ISO 3166-1 Alpha-2 country code of Kuwait.
	KW
	// KY represents the ISO 3166-1 Alpha-2 country code of the Cayman Islands.
	KY
	// KZ represents the ISO 3166-1 Alpha-2 country code of Kazakhstan.
	KZ
	// LA represents the ISO 3166-1 Alpha-2 country code of Lao People's Democratic Republic.
	LA
	// LB represents the ISO 3166-1 Alpha-2 country code of Lebanon.
	LB
	// LC represents the ISO 3166-1 Alpha-2 country code of Saint Lucia.
	LC
	// LI represents the ISO 3166-1 Alpha-2 country code of Liechtenstein.
	LI
	// LK represents the ISO 3166-1 Alpha-2 country code of Sri Lanka.
	LK
	// LR represents the ISO 3166-1 Alpha-2 country code of Liberia.
	LR
	// LS represents the ISO 3166-1 Alpha-2 country code of Lesotho.
	LS
	// LT represents the ISO 3166-1 Alpha-2 country code of Lithuania.
	LT
	// LU represents the ISO 3166-1 Alpha-2 country code of Luxembourg.
	LU
	// LV represents the ISO 3166-1 Alpha-2 country code of Latvia.
	LV
	// LY represents the ISO 3166-1 Alpha-2 country code of Libya.
	LY
	// MA represents the ISO 3166-1 Alpha-2 country code of Morocco.
	MA
	// MC represents the ISO 3166-1 Alpha-2 country code of Monaco.
	MC
	// MD represents the ISO 3166-1 Alpha-2 country code of Moldova, Republic of.
	MD
	// ME represents the ISO 3166-1 Alpha-2 country code of Montenegro.
	ME
	// MF represents the ISO 3166-1 Alpha-2 country code of Saint Martin (French part).
	MF
	// MG represents the ISO 3166-1 Alpha-2 country code of Madagascar.
	MG
	// MH represents the ISO 3166-1 Alpha-2 country code of the Marshall Islands.
	MH
	// MK represents the ISO 3166-1 Alpha-2 country code of North Macedonia.
	MK
	// ML represents the ISO 3166-1 Alpha-2 country code of Mali.
	ML
	// MM represents the ISO 3166-1 Alpha-2 country code of Myanmar.
	MM
	// MN represents the ISO 3166-1 Alpha-2 country code of Mongolia.
	MN
	// MO represents the ISO 3166-1 Alpha-2 country code of Macao.
	MO
	// MP represents the ISO 3166-1 Alpha-2 country code of Northern Mariana Islands.
	MP
	// MQ represents the ISO 3166-1 Alpha-2 country code of Martinique.
	MQ
	// MR represents the ISO 3166-1 Alpha-2 country code of Mauritania.
	MR
	// MS represents the ISO 3166-1 Alpha-2 country code of Montserrat.
	MS
	// MT represents the ISO 3166-1 Alpha-2 country code of Malta.
	MT
	// MU represents the ISO 3166-1 Alpha-2 country code of Mauritius.
	MU
	// MV represents the ISO 3166-1 Alpha-2 country code of the Maldives.
	MV
	// MW represents the ISO 3166-1 Alpha-2 country code of Malawi.
	MW
	// MX represents the ISO 3166-1 Alpha-2 country code of Mexico.
	MX
	// MY represents the ISO 3166-1 Alpha-2 country code of Malaysia.
	MY
	// MZ represents the ISO 3166-1 Alpha-2 country code of Mozambique.
	MZ
	// NA represents the ISO 3166-1 Alpha-2 country code of Namibia.
	NA
	// NC represents the ISO 3166-1 Alpha-2 country code of New Caledonia.
	NC
	// NE represents the ISO 3166-1 Alpha-2 country code of Niger.
	NE
	// NF represents the ISO 3166-1 Alpha-2 country code of Norfolk Island.
	NF
	// NG represents the ISO 3166-1 Alpha-2 country code of Nigeria.
	NG
	// NI represents the ISO 3166-1 Alpha-2 country code of Nicaragua.
	NI
	// NL represents the ISO 3166-1 Alpha-2 country code of the Netherlands.
	NL
	// NO represents the ISO 3166-1 Alpha-2 country code of Norway.
	NO
	// NP represents the ISO 3166-1 Alpha-2 country code of Nepal.
	NP
	// NR represents the ISO 3166-1 Alpha-2 country code of Nauru.
	NR
	// NU represents the ISO 3166-1 Alpha-2 country code of Niue.
	NU
	// NZ represents the ISO 3166-1 Alpha-2 country code of New Zealand.
	NZ
	// OM represents the ISO 3166-1 Alpha-2 country code of Oman.
	OM
	// PA represents the ISO 3166-1 Alpha-2 country code of Panama.
	PA
	// PE represents the ISO 3166-1 Alpha-2 country code of Peru.
	PE
	// PF represents the ISO 3166-1 Alpha-2 country code of French Polynesia.
	PF
	// PG represents the ISO 3166-1 Alpha-2 country code of Papua New Guinea.
	PG
	// PH represents the ISO 3166-1 Alpha-2 country code of the Philippines.
	PH
	// PK represents the ISO 3166-1 Alpha-2 country code of Pakistan.
	PK
	// PL represents the ISO 3166-1 Alpha-2 country code of Poland.
	PL
	// PM represents the ISO 3166-1 Alpha-2 country code of Saint Pierre and Miquelon.
	PM
	// PN represents the ISO 3166-1 Alpha-2 country code of Pitcairn.
	PN
	// PR represents the ISO 3166-1 Alpha-2 country code of Puerto Rico.
	PR
	// PS represents the ISO 3166-1 Alpha-2 country code of Palestine, State of.
	PS
	// PT represents the ISO 3166-1 Alpha-2 country code of Portugal.
	PT
	// PW represents the ISO 3166-1 Alpha-2 country code of Palau.
	PW
	// PY represents the ISO 3166-1 Alpha-2 country code of Paraguay.
	PY
	// QA represents the ISO 3166-1 Alpha-2 country code of Qatar.
	QA
	// RE represents the ISO 3166-1 Alpha-2 country code of Réunion.
	RE
	// RO represents the ISO 3166-1 Alpha-2 country code of Romania.
	RO
	// RS represents the ISO 3166-1 Alpha-2 country code of Serbia.
	RS
	// RU represents the ISO 3166-1 Alpha-2 country code of Russian Federation.
	RU
	// RW represents the ISO 3166-1 Alpha-2 country code of Rwanda.
	RW
	// SA represents the ISO 3166-1 Alpha-2 country code of Saudi Arabia.
	SA
	// SB represents the ISO 3166-1 Alpha-2 country code of the Solomon Islands.
	SB
	// SC represents the ISO 3166-1 Alpha-2 country code of Seychelles.
	SC
	// SD represents the ISO 3166-1 Alpha-2 country code of Sudan.
	SD
	// SE represents the ISO 3166-1 Alpha-2 country code of Sweden.
	SE
	// SG represents the ISO 3166-1 Alpha-2 country code of Singapore.
	SG
	// SH represents the ISO 3166-1 Alpha-2 country code of Saint Helena, Ascension and Tristan da Cunha.
	SH
	// SI represents the ISO 3166-1 Alpha-2 country code of Slovenia.
	SI
	// SJ represents the ISO 3166-1 Alpha-2 country code of Svalbard and Jan Mayen.
	SJ
	// SK represents the ISO 3166-1 Alpha-2 country code of Slovakia.
	SK
	// SL represents the ISO 3166-1 Alpha-2 country code of Sierra Leone.
	SL
	// SM represents the ISO 3166-1 Alpha-2 country code of San Marino.
	SM
	// SN represents the ISO 3166-1 Alpha-2 country code of Senegal.
	SN
	// SO represents the ISO 3166-1 Alpha-2 country code of Somalia.
	SO
	// SR represents the ISO 3166-1 Alpha-2 country code of Suriname.
	SR
	// SS represents the ISO 3166-1 Alpha-2 country code of South Sudan.
	SS
	// ST represents the ISO 3166-1 Alpha-2 country code of São Tomé and Príncipe.
	ST
	// SV represents the ISO 3166-1 Alpha-2 country code of El Salvador.
	SV
	// SX represents the ISO 3166-1 Alpha-2 country code of Sint Maarten (Dutch part).
	SX
	// SY represents the ISO 3166-1 Alpha-2 country code of Syrian Arab Republic.
	SY
	// SZ represents the ISO 3166-1 Alpha-2 country code of Eswatini.
	SZ
	// TC represents the ISO 3166-1 Alpha-2 country code of the Turks and Caicos Islands.
	TC
	// TD represents the ISO 3166-1 Alpha-2 country code of Chad.
	TD
	// TF represents the ISO 3166-1 Alpha-2 country code of French Southern Territories.
	TF
	// TG represents the ISO 3166-1 Alpha-2 country code of Togo.
	TG
	// TH represents the ISO 3166-1 Alpha-2 country code of Thailand.
	TH
	// TJ represents the ISO 3166-1 Alpha-2 country code of Tajikistan.
	TJ
	// TK represents the ISO 3166-1 Alpha-2 country code of Tokelau.
	TK
	// TL represents the ISO 3166-1 Alpha-2 country code of Timor-Leste.
	TL
	// TM represents the ISO 3166-1 Alpha-2 country code of Turkmenistan.
	TM
	// TN represents the ISO 3166-1 Alpha-2 country code of Tunisia.
	TN
	// TO represents the ISO 3166-1 Alpha-2 country code of Tonga.
	TO
	// TR represents the ISO 3166-1 Alpha-2 country code of Turkey.
	TR
	// TT represents the ISO 3166-1 Alpha-2 country code of Trinidad and Tobago.
	TT
	// TV represents the ISO 3166-1 Alpha-2 country code of Tuvalu.
	TV
	// TW represents the ISO 3166-1 Alpha-2 country code of Taiwan, Province of China.
	TW
	// TZ represents the ISO 3166-1 Alpha-2 country code of Tanzania, United Republic of.
	TZ
	// UA represents the ISO 3166-1 Alpha-2 country code of Ukraine.
	UA
	// UG represents the ISO 3166-1 Alpha-2 country code of Uganda.
	UG
	// UM represents the ISO 3166-1 Alpha-2 country code of United States Minor Outlying Islands.
	UM
	// US represents the ISO 3166-1 Alpha-2 country code of the United States of America.
	US
	// UY represents the ISO 3166-1 Alpha-2 country code of Uruguay.
	UY
	// UZ represents the ISO 3166-1 Alpha-2 country code of Uzbekistan.
	UZ
	// VA represents the ISO 3166-1 Alpha-2 country code of Holy See.
	VA
	// VC represents the ISO 3166-1 Alpha-2 country code of Saint Vincent and the Grenadines.
	VC
	// VE represents the ISO 3166-1 Alpha-2 country code of Venezuela (Bolivarian Republic of).
	VE
	// VG represents the ISO 3166-1 Alpha-2 country code of the Virgin Islands (British).
	VG
	// VI represents the ISO 3166-1 Alpha-2 country code of Virgin Islands (U.S.).
	VI
	// VN represents the ISO 3166-1 Alpha-2 country code of Vietnam.
	VN
	// VU represents the ISO 3166-1 Alpha-2 country code of Vanuatu.
	VU
	// WF represents the ISO 3166-1 Alpha-2 country code of Wallis and Futuna.
	WF
	// WS represents the ISO 3166-1 Alpha-2 country code of Samoa.
	WS
	// YE represents the ISO 3166-1 Alpha-2 country code of Yemen.
	YE
	// YT represents the ISO 3166-1 Alpha-2 country code of Mayotte.
	YT
	// ZA represents the ISO 3166-1 Alpha-2 country code of South Africa.
	ZA
	// ZM represents the ISO 3166-1 Alpha-2 country code of Zambia.
	ZM
	// ZW represents the ISO 3166-1 Alpha-2 country code of Zimbabwe.
	ZW
)

var countryCodesDetails = map[CountryCode]CountryCodeDetails{
	AD: {Alpha2: "AD", Alpha3: "AND", Flag: "🇦🇩", Number: "020", Name: "Andorra"},
	AE: {Alpha2: "AE", Alpha3: "ARE", Flag: "🇦🇪", Number: "784", Name: "United Arab Emirates"},
	AF: {Alpha2: "AF", Alpha3: "AFG", Flag: "🇦🇫", Number: "004", Name: "Afghanistan"},
	AG: {Alpha2: "AG", Alpha3: "ATG", Flag: "🇦🇬", Number: "028", Name: "Antigua and Barbuda"},
	AI: {Alpha2: "AI", Alpha3: "AIA", Flag: "🇦🇮", Number: "660", Name: "Anguilla"},
	AL: {Alpha2: "AL", Alpha3: "ALB", Flag: "🇦🇱", Number: "008", Name: "Albania"},
	AM: {Alpha2: "AM", Alpha3: "ARM", Flag: "🇦🇲", Number: "051", Name: "Armenia"},
	AO: {Alpha2: "AO", Alpha3: "AGO", Flag: "🇦🇴", Number: "024", Name: "Angola"},
	AQ: {Alpha2: "AQ", Alpha3: "ATA", Flag: "🇦🇶", Number: "010", Name: "Antarctica"},
	AR: {Alpha2: "AR", Alpha3: "ARG", Flag: "🇦🇷", Number: "032", Name: "Argentina"},
	AS: {Alpha2: "AS", Alpha3: "ASM", Flag: "🇦🇸", Number: "016", Name: "American Samoa"},
	AT: {Alpha2: "AT", Alpha3: "AUT", Flag: "🇦🇹", Number: "040", Name: "Austria"},
	AU: {Alpha2: "AU", Alpha3: "AUS", Flag: "🇦🇺", Number: "036", Name: "Australia"},
	AW: {Alpha2: "AW", Alpha3: "ABW", Flag: "🇦🇼", Number: "533", Name: "Aruba"},
	AX: {Alpha2: "AX", Alpha3: "ALA", Flag: "🇦🇽", Number: "248", Name: "Åland Islands"},
	AZ: {Alpha2: "AZ", Alpha3: "AZE", Flag: "🇦🇿", Number: "031", Name: "Azerbaijan"},
	BA: {Alpha2: "BA", Alpha3: "BIH", Flag: "🇧🇦", Number: "070", Name: "Bosnia and Herzegovina"},
	BB: {Alpha2: "BB", Alpha3: "BRB", Flag: "🇧🇧", Number: "052", Name: "Barbados"},
	BD: {Alpha2: "BD", Alpha3: "BGD", Flag: "🇧🇩", Number: "050", Name: "Bangladesh"},
	BE: {Alpha2: "BE", Alpha3: "BEL", Flag: "🇧🇪", Number: "056", Name: "Belgium"},
	BF: {Alpha2: "BF", Alpha3: "BFA", Flag: "🇧🇫", Number: "854", Name: "Burkina Faso"},
	BG: {Alpha2: "BG", Alpha3: "BGR", Flag: "🇧🇬", Number: "100", Name: "Bulgaria"},
	BH: {Alpha2: "BH", Alpha3: "BHR", Flag: "🇧🇭", Number: "048", Name: "Bahrain"},
	BI: {Alpha2: "BI", Alpha3: "BDI", Flag: "🇧🇮", Number: "108", Name: "Burundi"},
	BJ: {Alpha2: "BJ", Alpha3: "BEN", Flag: "🇧🇯", Number: "204", Name: "Benin"},
	BL: {Alpha2: "BL", Alpha3: "BLM", Flag: "🇧🇱", Number: "652", Name: "Saint Barthélemy"},
	BM: {Alpha2: "BM", Alpha3: "BMU", Flag: "🇧🇲", Number: "060", Name: "Bermuda"},
	BN: {Alpha2: "BN", Alpha3: "BRN", Flag: "🇧🇳", Number: "096", Name: "Brunei Darussalam"},
	BO: {Alpha2: "BO", Alpha3: "BOL", Flag: "🇧🇴", Number: "068", Name: "Bolivia (Plurinational State of)"},
	BQ: {Alpha2: "BQ", Alpha3: "BES", Flag: "🇧🇶", Number: "535", Name: "Bonaire, Sint Eustatius and Saba"},
	BR: {Alpha2: "BR", Alpha3: "BRA", Flag: "🇧🇷", Number: "076", Name: "Brazil"},
	BS: {Alpha2: "BS", Alpha3: "BHS", Flag: "🇧🇸", Number: "044", Name: "Bahamas"},
	BT: {Alpha2: "BT", Alpha3: "BTN", Flag: "🇧🇹", Number: "064", Name: "Bhutan"},
	BV: {Alpha2: "BV", Alpha3: "BVT", Flag: "🇧🇻", Number: "074", Name: "Bouvet Island"},
	BW: {Alpha2: "BW", Alpha3: "BWA", Flag: "🇧🇼", Number: "072", Name: "Botswana"},
	BY: {Alpha2: "BY", Alpha3: "BLR", Flag: "🇧🇾", Number: "112", Name: "Belarus"},
	BZ: {Alpha2: "BZ", Alpha3: "BLZ", Flag: "🇧🇿", Number: "084", Name: "Belize"},
	CA: {Alpha2: "CA", Alpha3: "CAN", Flag: "🇨🇦", Number: "124", Name: "Canada"},
	CC: {Alpha2: "CC", Alpha3: "CCK", Flag: "🇨🇨", Number: "166", Name: "Cocos (Keeling) Islands"},
	CD: {Alpha2: "CD", Alpha3: "COD", Flag: "🇨🇩", Number: "180", Name: "Congo, Democratic Republic of the"},
	CF: {Alpha2: "CF", Alpha3: "CAF", Flag: "🇨🇫", Number: "140", Name: "Central African Republic"},
	CG: {Alpha2: "CG", Alpha3: "COG", Flag: "🇨🇬", Number: "178", Name: "Congo"},
	CH: {Alpha2: "CH", Alpha3: "CHE", Flag: "🇨🇭", Number: "756", Name: "Switzerland"},
	CI: {Alpha2: "CI", Alpha3: "CIV", Flag: "🇨🇮", Number: "384", Name: "Côte d'Ivoire"},
	CK: {Alpha2: "CK", Alpha3: "COK", Flag: "🇨🇰", Number: "184", Name: "Cook Islands"},
	CL: {Alpha2: "CL", Alpha3: "CHL", Flag: "🇨🇱", Number: "152", Name: "Chile"},
	CM: {Alpha2: "CM", Alpha3: "CMR", Flag: "🇨🇲", Number: "120", Name: "Cameroon"},
	CN: {Alpha2: "CN", Alpha3: "CHN", Flag: "🇨🇳", Number: "156", Name: "China"},
	CO: {Alpha2: "CO", Alpha3: "COL", Flag: "🇨🇴", Number: "170", Name: "Colombia"},
	CR: {Alpha2: "CR", Alpha3: "CRI", Flag: "🇨🇷", Number: "188", Name: "Costa Rica"},
	CU: {Alpha2: "CU", Alpha3: "CUB", Flag: "🇨🇺", Number: "192", Name: "Cuba"},
	CV: {Alpha2: "CV", Alpha3: "CPV", Flag: "🇨🇻", Number: "132", Name: "Cabo Verde"},
	CW: {Alpha2: "CW", Alpha3: "CUW", Flag: "🇨🇼", Number: "531", Name: "Curaçao"},
	CX: {Alpha2: "CX", Alpha3: "CXR", Flag: "🇨🇽", Number: "162", Name: "Christmas Island"},
	CY: {Alpha2: "CY", Alpha3: "CYP", Flag: "🇨🇾", Number: "196", Name: "Cyprus"},
	CZ: {Alpha2: "CZ", Alpha3: "CZE", Flag: "🇨🇿", Number: "203", Name: "Czechia"},
	DE: {Alpha2: "DE", Alpha3: "DEU", Flag: "🇩🇪", Number: "276", Name: "Germany"},
	DJ: {Alpha2: "DJ", Alpha3: "DJI", Flag: "🇩🇯", Number: "262", Name: "Djibouti"},
	DK: {Alpha2: "DK", Alpha3: "DNK", Flag: "🇩🇰", Number: "208", Name: "Denmark"},
	DM: {Alpha2: "DM", Alpha3: "DMA", Flag: "🇩🇲", Number: "212", Name: "Dominica"},
	DO: {Alpha2: "DO", Alpha3: "DOM", Flag: "🇩🇴", Number: "214", Name: "Dominican Republic"},
	DZ: {Alpha2: "DZ", Alpha3: "DZA", Flag: "🇩🇿", Number: "012", Name: "Algeria"},
	EC: {Alpha2: "EC", Alpha3: "ECU", Flag: "🇪🇨", Number: "218", Name: "Ecuador"},
	EE: {Alpha2: "EE", Alpha3: "EST", Flag: "🇪🇪", Number: "233", Name: "Estonia"},
	EG: {Alpha2: "EG", Alpha3: "EGY", Flag: "🇪🇬", Number: "818", Name: "Egypt"},
	EH: {Alpha2: "EH", Alpha3: "ESH", Flag: "🇪🇭", Number: "732", Name: "Western Sahara"},
	ER: {Alpha2: "ER", Alpha3: "ERI", Flag: "🇪🇷", Number: "232", Name: "Eritrea"},
	ES: {Alpha2: "ES", Alpha3: "ESP", Flag: "🇪🇸", Number: "724", Name: "Spain"},
	ET: {Alpha2: "ET", Alpha3: "ETH", Flag: "🇪🇹", Number: "231", Name: "Ethiopia"},
	FI: {Alpha2: "FI", Alpha3: "FIN", Flag: "🇫🇮", Number: "246", Name: "Finland"},
	FJ: {Alpha2: "FJ", Alpha3: "FJI", Flag: "🇫🇯", Number: "242", Name: "Fiji"},
	FK: {Alpha2: "FK", Alpha3: "FLK", Flag: "🇫🇰", Number: "238", Name: "Falkland Islands (Malvinas)"},
	FM: {Alpha2: "FM", Alpha3: "FSM", Flag: "🇫🇲", Number: "583", Name: "Micronesia (Federated States of)"},
	FO: {Alpha2: "FO", Alpha3: "FRO", Flag: "🇫🇴", Number: "234", Name: "Faroe Islands"},
	FR: {Alpha2: "FR", Alpha3: "FRA", Flag: "🇫🇷", Number: "250", Name: "France"},
	GA: {Alpha2: "GA", Alpha3: "GAB", Flag: "🇬🇦", Number: "266", Name: "Gabon"},
	GB: {Alpha2: "GB", Alpha3: "GBR", Flag: "🇬🇧", Number: "826", Name: "United Kingdom of Great Britain and Northern Ireland"},
	GD: {Alpha2: "GD", Alpha3: "GRD", Flag: "🇬🇩", Number: "308", Name: "Grenada"},
	GE: {Alpha2: "GE", Alpha3: "GEO", Flag: "🇬🇪", Number: "268", Name: "Georgia"},
	GF: {Alpha2: "GF", Alpha3: "GUF", Flag: "🇬🇫", Number: "254", Name: "French Guiana"},
	GG: {Alpha2: "GG", Alpha3: "GGY", Flag: "🇬🇬", Number: "831", Name: "Guernsey"},
	GH: {Alpha2: "GH", Alpha3: "GHA", Flag: "🇬🇭", Number: "288", Name: "Ghana"},
	GI: {Alpha2: "GI", Alpha3: "GIB", Flag: "🇬🇮", Number: "292", Name: "Gibraltar"},
	GL: {Alpha2: "GL", Alpha3: "GRL", Flag: "🇬🇱", Number: "304", Name: "Greenland"},
	GM: {Alpha2: "GM", Alpha3: "GMB", Flag: "🇬🇲", Number: "270", Name: "Gambia"},
	GN: {Alpha2: "GN", Alpha3: "GIN", Flag: "🇬🇳", Number: "324", Name: "Guinea"},
	GP: {Alpha2: "GP", Alpha3: "GLP", Flag: "🇬🇵", Number: "312", Name: "Guadeloupe"},
	GQ: {Alpha2: "GQ", Alpha3: "GNQ", Flag: "🇬🇶", Number: "226", Name: "Equatorial Guinea"},
	GR: {Alpha2: "GR", Alpha3: "GRC", Flag: "🇬🇷", Number: "300", Name: "Greece"},
	GS: {Alpha2: "GS", Alpha3: "SGS", Flag: "🇬🇸", Number: "239", Name: "South Georgia and the South Sandwich Islands"},
	GT: {Alpha2: "GT", Alpha3: "GTM", Flag: "🇬🇹", Number: "320", Name: "Guatemala"},
	GU: {Alpha2: "GU", Alpha3: "GUM", Flag: "🇬🇺", Number: "316", Name: "Guam"},
	GW: {Alpha2: "GW", Alpha3: "GNB", Flag: "🇬🇼", Number: "624", Name: "Guinea-Bissau"},
	GY: {Alpha2: "GY", Alpha3: "GUY", Flag: "🇬🇾", Number: "328", Name: "Guyana"},
	HK: {Alpha2: "HK", Alpha3: "HKG", Flag: "🇭🇰", Number: "344", Name: "Hong Kong"},
	HM: {Alpha2: "HM", Alpha3: "HMD", Flag: "🇭🇲", Number: "334", Name: "Heard Island and McDonald Islands"},
	HN: {Alpha2: "HN", Alpha3: "HND", Flag: "🇭🇳", Number: "340", Name: "Honduras"},
	HR: {Alpha2: "HR", Alpha3: "HRV", Flag: "🇭🇷", Number: "191", Name: "Croatia"},
	HT: {Alpha2: "HT", Alpha3: "HTI", Flag: "🇭🇹", Number: "332", Name: "Haiti"},
	HU: {Alpha2: "HU", Alpha3: "HUN", Flag: "🇭🇺", Number: "348", Name: "Hungary"},
	ID: {Alpha2: "ID", Alpha3: "IDN", Flag: "🇮🇩", Number: "360", Name: "Indonesia"},
	IE: {Alpha2: "IE", Alpha3: "IRL", Flag: "🇮🇪", Number: "372", Name: "Ireland"},
	IL: {Alpha2: "IL", Alpha3: "ISR", Flag: "🇮🇱", Number: "376", Name: "Israel"},
	IM: {Alpha2: "IM", Alpha3: "IMN", Flag: "🇮🇲", Number: "833", Name: "Isle of Man"},
	IN: {Alpha2: "IN", Alpha3: "IND", Flag: "🇮🇳", Number: "356", Name: "India"},
	IO: {Alpha2: "IO", Alpha3: "IOT", Flag: "🇮🇴", Number: "086", Name: "British Indian Ocean Territory"},
	IQ: {Alpha2: "IQ", Alpha3: "IRQ", Flag: "🇮🇶", Number: "368", Name: "Iraq"},
	IR: {Alpha2: "IR", Alpha3: "IRN", Flag: "🇮🇷", Number: "364", Name: "Iran (Islamic Republic of)"},
	IS: {Alpha2: "IS", Alpha3: "ISL", Flag: "🇮🇸", Number: "352", Name: "Iceland"},
	IT: {Alpha2: "IT", Alpha3: "ITA", Flag: "🇮🇹", Number: "380", Name: "Italy"},
	JE: {Alpha2: "JE", Alpha3: "JEY", Flag: "🇯🇪", Number: "832", Name: "Jersey"},
	JM: {Alpha2: "JM", Alpha3: "JAM", Flag: "🇯🇲", Number: "388", Name: "Jamaica"},
	JO: {Alpha2: "JO", Alpha3: "JOR", Flag: "🇯🇴", Number: "400", Name: "Jordan"},
	JP: {Alpha2: "JP", Alpha3: "JPN", Flag: "🇯🇵", Number: "392", Name: "Japan"},
	KE: {Alpha2: "KE", Alpha3: "KEN", Flag: "🇰🇪", Number: "404", Name: "Kenya"},
	KG: {Alpha2: "KG", Alpha3: "KGZ", Flag: "🇰🇬", Number: "417", Name: "Kyrgyzstan"},
	KH: {Alpha2: "KH", Alpha3: "KHM", Flag: "🇰🇭", Number: "116", Name: "Cambodia"},
	KI: {Alpha2: "KI", Alpha3: "KIR", Flag: "🇰🇮", Number: "296", Name: "Kiribati"},
	KM: {Alpha2: "KM", Alpha3: "COM", Flag: "🇰🇲", Number: "174", Name: "Comoros"},
	KN: {Alpha2: "KN", Alpha3: "KNA", Flag: "🇰🇳", Number: "659", Name: "Saint Kitts and Nevis"},
	KP: {Alpha2: "KP", Alpha3: "PRK", Flag: "🇰🇵", Number: "408", Name: "Korea (Democratic People's Republic of)"},
	KR: {Alpha2: "KR", Alpha3: "KOR", Flag: "🇰🇷", Number: "410", Name: "Korea, Republic of"},
	KW: {Alpha2: "KW", Alpha3: "KWT", Flag: "🇰🇼", Number: "414", Name: "Kuwait"},
	KY: {Alpha2: "KY", Alpha3: "CYM", Flag: "🇰🇾", Number: "136", Name: "Cayman Islands"},
	KZ: {Alpha2: "KZ", Alpha3: "KAZ", Flag: "🇰🇿", Number: "398", Name: "Kazakhstan"},
	LA: {Alpha2: "LA", Alpha3: "LAO", Flag: "🇱🇦", Number: "418", Name: "Lao People's Democratic Republic"},
	LB: {Alpha2: "LB", Alpha3: "LBN", Flag: "🇱🇧", Number: "422", Name: "Lebanon"},
	LC: {Alpha2: "LC", Alpha3: "LCA", Flag: "🇱🇨", Number: "662", Name: "Saint Lucia"},
	LI: {Alpha2: "LI", Alpha3: "LIE", Flag: "🇱🇮", Number: "438", Name: "Liechtenstein"},
	LK: {Alpha2: "LK", Alpha3: "LKA", Flag: "🇱🇰", Number: "144", Name: "Sri Lanka"},
	LR: {Alpha2: "LR", Alpha3: "LBR", Flag: "🇱🇷", Number: "430", Name: "Liberia"},
	LS: {Alpha2: "LS", Alpha3: "LSO", Flag: "🇱🇸", Number: "426", Name: "Lesotho"},
	LT: {Alpha2: "LT", Alpha3: "LTU", Flag: "🇱🇹", Number: "440", Name: "Lithuania"},
	LU: {Alpha2: "LU", Alpha3: "LUX", Flag: "🇱🇺", Number: "442", Name: "Luxembourg"},
	LV: {Alpha2: "LV", Alpha3: "LVA", Flag: "🇱🇻", Number: "428", Name: "Latvia"},
	LY: {Alpha2: "LY", Alpha3: "LBY", Flag: "🇱🇾", Number: "434", Name: "Libya"},
	MA: {Alpha2: "MA", Alpha3: "MAR", Flag: "🇲🇦", Number: "504", Name: "Morocco"},
	MC: {Alpha2: "MC", Alpha3: "MCO", Flag: "🇲🇨", Number: "492", Name: "Monaco"},
	MD: {Alpha2: "MD", Alpha3: "MDA", Flag: "🇲🇩", Number: "498", Name: "Moldova, Republic of"},
	ME: {Alpha2: "ME", Alpha3: "MNE", Flag: "🇲🇪", Number: "499", Name: "Montenegro"},
	MF: {Alpha2: "MF", Alpha3: "MAF", Flag: "🇲🇫", Number: "663", Name: "Saint Martin (French part)"},
	MG: {Alpha2: "MG", Alpha3: "MDG", Flag: "🇲🇬", Number: "450", Name: "Madagascar"},
	MH: {Alpha2: "MH", Alpha3: "MHL", Flag: "🇲🇭", Number: "584", Name: "Marshall Islands"},
	MK: {Alpha2: "MK", Alpha3: "MKD", Flag: "🇲🇰", Number: "807", Name: "North Macedonia"},
	ML: {Alpha2: "ML", Alpha3: "MLI", Flag: "🇲🇱", Number: "466", Name: "Mali"},
	MM: {Alpha2: "MM", Alpha3: "MMR", Flag: "🇲🇲", Number: "104", Name: "Myanmar"},
	MN: {Alpha2: "MN", Alpha3: "MNG", Flag: "🇲🇳", Number: "496", Name: "Mongolia"},
	MO: {Alpha2: "MO", Alpha3: "MAC", Flag: "🇲🇴", Number: "446", Name: "Macao"},
	MP: {Alpha2: "MP", Alpha3: "MNP", Flag: "🇲🇵", Number: "580", Name: "Northern Mariana Islands"},
	MQ: {Alpha2: "MQ", Alpha3: "MTQ", Flag: "🇲🇶", Number: "474", Name: "Martinique"},
	MR: {Alpha2: "MR", Alpha3: "MRT", Flag: "🇲🇷", Number: "478", Name: "Mauritania"},
	MS: {Alpha2: "MS", Alpha3: "MSR", Flag: "🇲🇸", Number: "500", Name: "Montserrat"},
	MT: {Alpha2: "MT", Alpha3: "MLT", Flag: "🇲🇹", Number: "470", Name: "Malta"},
	MU: {Alpha2: "MU", Alpha3: "MUS", Flag: "🇲🇺", Number: "480", Name: "Mauritius"},
	MV: {Alpha2: "MV", Alpha3: "MDV", Flag: "🇲🇻", Number: "462", Name: "Maldives"},
	MW: {Alpha2: "MW", Alpha3: "MWI", Flag: "🇲🇼", Number: "454", Name: "Malawi"},
	MX: {Alpha2: "MX", Alpha3: "MEX", Flag: "🇲🇽", Number: "484", Name: "Mexico"},
	MY: {Alpha2: "MY", Alpha3: "MYS", Flag: "🇲🇾", Number: "458", Name: "Malaysia"},
	MZ: {Alpha2: "MZ", Alpha3: "MOZ", Flag: "🇲🇿", Number: "508", Name: "Mozambique"},
	NA: {Alpha2: "NA", Alpha3: "NAM", Flag: "🇳🇦", Number: "516", Name: "Namibia"},
	NC: {Alpha2: "NC", Alpha3: "NCL", Flag: "🇳🇨", Number: "540", Name: "New Caledonia"},
	NE: {Alpha2: "NE", Alpha3: "NER", Flag: "🇳🇪", Number: "562", Name: "Niger"},
	NF: {Alpha2: "NF", Alpha3: "NFK", Flag: "🇳🇫", Number: "574", Name: "Norfolk Island"},
	NG: {Alpha2: "NG", Alpha3: "NGA", Flag: "🇳🇬", Number: "566", Name: "Nigeria"},
	NI: {Alpha2: "NI", Alpha3: "NIC", Flag: "🇳🇮", Number: "558", Name: "Nicaragua"},
	NL: {Alpha2: "NL", Alpha3: "NLD", Flag: "🇳🇱", Number: "528", Name: "Netherlands"},
	NO: {Alpha2: "NO", Alpha3: "NOR", Flag: "🇳🇴", Number: "578", Name: "Norway"},
	NP: {Alpha2: "NP", Alpha3: "NPL", Flag: "🇳🇵", Number: "524", Name: "Nepal"},
	NR: {Alpha2: "NR", Alpha3: "NRU", Flag: "🇳🇷", Number: "520", Name: "Nauru"},
	NU: {Alpha2: "NU", Alpha3: "NIU", Flag: "🇳🇺", Number: "570", Name: "Niue"},
	NZ: {Alpha2: "NZ", Alpha3: "NZL", Flag: "🇳🇿", Number: "554", Name: "New Zealand"},
	OM: {Alpha2: "OM", Alpha3: "OMN", Flag: "🇴🇲", Number: "512", Name: "Oman"},
	PA: {Alpha2: "PA", Alpha3: "PAN", Flag: "🇵🇦", Number: "591", Name: "Panama"},
	PE: {Alpha2: "PE", Alpha3: "PER", Flag: "🇵🇪", Number: "604", Name: "Peru"},
	PF: {Alpha2: "PF", Alpha3: "PYF", Flag: "🇵🇫", Number: "258", Name: "French Polynesia"},
	PG: {Alpha2: "PG", Alpha3: "PNG", Flag: "🇵🇬", Number: "598", Name: "Papua New Guinea"},
	PH: {Alpha2: "PH", Alpha3: "PHL", Flag: "🇵🇭", Number: "608", Name: "Philippines"},
	PK: {Alpha2: "PK", Alpha3: "PAK", Flag: "🇵🇰", Number: "586", Name: "Pakistan"},
	PL: {Alpha2: "PL", Alpha3: "POL", Flag: "🇵🇱", Number: "616", Name: "Poland"},
	PM: {Alpha2: "PM", Alpha3: "SPM", Flag: "🇵🇲", Number: "666", Name: "Saint Pierre and Miquelon"},
	PN: {Alpha2: "PN", Alpha3: "PCN", Flag: "🇵🇳", Number: "612", Name: "Pitcairn"},
	PR: {Alpha2: "PR", Alpha3: "PRI", Flag: "🇵🇷", Number: "630", Name: "Puerto Rico"},
	PS: {Alpha2: "PS", Alpha3: "PSE", Flag: "🇵🇸", Number: "275", Name: "Palestine, State of"},
	PT: {Alpha2: "PT", Alpha3: "PRT", Flag: "🇵🇹", Number: "620", Name: "Portugal"},
	PW: {Alpha2: "PW", Alpha3: "PLW", Flag: "🇵🇼", Number: "585", Name: "Palau"},
	PY: {Alpha2: "PY", Alpha3: "PRY", Flag: "🇵🇾", Number: "600", Name: "Paraguay"},
	QA: {Alpha2: "QA", Alpha3: "QAT", Flag: "🇶🇦", Number: "634", Name: "Qatar"},
	RE: {Alpha2: "RE", Alpha3: "REU", Flag: "🇷🇪", Number: "638", Name: "Réunion"},
	RO: {Alpha2: "RO", Alpha3: "ROU", Flag: "🇷🇴", Number: "642", Name: "Romania"},
	RS: {Alpha2: "RS", Alpha3: "SRB", Flag: "🇷🇸", Number: "688", Name: "Serbia"},
	RU: {Alpha2: "RU", Alpha3: "RUS", Flag: "🇷🇺", Number: "643", Name: "Russian Federation"},
	RW: {Alpha2: "RW", Alpha3: "RWA", Flag: "🇷🇼", Number: "646", Name: "Rwanda"},
	SA: {Alpha2: "SA", Alpha3: "SAU", Flag: "🇸🇦", Number: "682", Name: "Saudi Arabia"},
	SB: {Alpha2: "SB", Alpha3: "SLB", Flag: "🇸🇧", Number: "090", Name: "Solomon Islands"},
	SC: {Alpha2: "SC", Alpha3: "SYC", Flag: "🇸🇨", Number: "690", Name: "Seychelles"},
	SD: {Alpha2: "SD", Alpha3: "SDN", Flag: "🇸🇩", Number: "729", Name: "Sudan"},
	SE: {Alpha2: "SE", Alpha3: "SWE", Flag: "🇸🇪", Number: "752", Name: "Sweden"},
	SG: {Alpha2: "SG", Alpha3: "SGP", Flag: "🇸🇬", Number: "702", Name: "Singapore"},
	SH: {Alpha2: "SH", Alpha3: "SHN", Flag: "🇸🇭", Number: "654", Name: "Saint Helena, Ascension and Tristan da Cunha"},
	SI: {Alpha2: "SI", Alpha3: "SVN", Flag: "🇸🇮", Number: "705", Name: "Slovenia"},
	SJ: {Alpha2: "SJ", Alpha3: "SJM", Flag: "🇸🇯", Number: "744", Name: "Svalbard and Jan Mayen"},
	SK: {Alpha2: "SK", Alpha3: "SVK", Flag: "🇸🇰", Number: "703", Name: "Slovakia"},
	SL: {Alpha2: "SL", Alpha3: "SLE", Flag: "🇸🇱", Number: "694", Name: "Sierra Leone"},
	SM: {Alpha2: "SM", Alpha3: "SMR", Flag: "🇸🇲", Number: "674", Name: "San Marino"},
	SN: {Alpha2: "SN", Alpha3: "SEN", Flag: "🇸🇳", Number: "686", Name: "Senegal"},
	SO: {Alpha2: "SO", Alpha3: "SOM", Flag: "🇸🇴", Number: "706", Name: "Somalia"},
	SR: {Alpha2: "SR", Alpha3: "SUR", Flag: "🇸🇷", Number: "740", Name: "Suriname"},
	SS: {Alpha2: "SS", Alpha3: "SSD", Flag: "🇸🇸", Number: "728", Name: "South Sudan"},
	ST: {Alpha2: "ST", Alpha3: "STP", Flag: "🇸🇹", Number: "678", Name: "Sao Tome and Principe"},
	SV: {Alpha2: "SV", Alpha3: "SLV", Flag: "🇸🇻", Number: "222", Name: "El Salvador"},
	SX: {Alpha2: "SX", Alpha3: "SXM", Flag: "🇸🇽", Number: "534", Name: "Sint Maarten (Dutch part)"},
	SY: {Alpha2: "SY", Alpha3: "SYR", Flag: "🇸🇾", Number: "760", Name: "Syrian Arab Republic"},
	SZ: {Alpha2: "SZ", Alpha3: "SWZ", Flag: "🇸🇿", Number: "748", Name: "Eswatini"},
	TC: {Alpha2: "TC", Alpha3: "TCA", Flag: "🇹🇨", Number: "796", Name: "Turks and Caicos Islands"},
	TD: {Alpha2: "TD", Alpha3: "TCD", Flag: "🇹🇩", Number: "148", Name: "Chad"},
	TF: {Alpha2: "TF", Alpha3: "ATF", Flag: "🇹🇫", Number: "260", Name: "French Southern Territories"},
	TG: {Alpha2: "TG", Alpha3: "TGO", Flag: "🇹🇬", Number: "768", Name: "Togo"},
	TH: {Alpha2: "TH", Alpha3: "THA", Flag: "🇹🇭", Number: "764", Name: "Thailand"},
	TJ: {Alpha2: "TJ", Alpha3: "TJK", Flag: "🇹🇯", Number: "762", Name: "Tajikistan"},
	TK: {Alpha2: "TK", Alpha3: "TKL", Flag: "🇹🇰", Number: "772", Name: "Tokelau"},
	TL: {Alpha2: "TL", Alpha3: "TLS", Flag: "🇹🇱", Number: "626", Name: "Timor-Leste"},
	TM: {Alpha2: "TM", Alpha3: "TKM", Flag: "🇹🇲", Number: "795", Name: "Turkmenistan"},
	TN: {Alpha2: "TN", Alpha3: "TUN", Flag: "🇹🇳", Number: "788", Name: "Tunisia"},
	TO: {Alpha2: "TO", Alpha3: "TON", Flag: "🇹🇴", Number: "776", Name: "Tonga"},
	TR: {Alpha2: "TR", Alpha3: "TUR", Flag: "🇹🇷", Number: "792", Name: "Turkey"},
	TT: {Alpha2: "TT", Alpha3: "TTO", Flag: "🇹🇹", Number: "780", Name: "Trinidad and Tobago"},
	TV: {Alpha2: "TV", Alpha3: "TUV", Flag: "🇹🇻", Number: "798", Name: "Tuvalu"},
	TW: {Alpha2: "TW", Alpha3: "TWN", Flag: "🇹🇼", Number: "158", Name: "Taiwan, Province of China"},
	TZ: {Alpha2: "TZ", Alpha3: "TZA", Flag: "🇹🇿", Number: "834", Name: "Tanzania, United Republic of"},
	UA: {Alpha2: "UA", Alpha3: "UKR", Flag: "🇺🇦", Number: "804", Name: "Ukraine"},
	UG: {Alpha2: "UG", Alpha3: "UGA", Flag: "🇺🇬", Number: "800", Name: "Uganda"},
	UM: {Alpha2: "UM", Alpha3: "UMI", Flag: "🇺🇲", Number: "581", Name: "United States Minor Outlying Islands"},
	US: {Alpha2: "US", Alpha3: "USA", Flag: "🇺🇸", Number: "840", Name: "United States of America"},
	UY: {Alpha2: "UY", Alpha3: "URY", Flag: "🇺🇾", Number: "858", Name: "Uruguay"},
	UZ: {Alpha2: "UZ", Alpha3: "UZB", Flag: "🇺🇿", Number: "860", Name: "Uzbekistan"},
	VA: {Alpha2: "VA", Alpha3: "VAT", Flag: "🇻🇦", Number: "336", Name: "Holy See"},
	VC: {Alpha2: "VC", Alpha3: "VCT", Flag: "🇻🇨", Number: "670", Name: "Saint Vincent and the Grenadines"},
	VE: {Alpha2: "VE", Alpha3: "VEN", Flag: "🇻🇪", Number: "862", Name: "Venezuela (Bolivarian Republic of)"},
	VG: {Alpha2: "VG", Alpha3: "VGB", Flag: "🇻🇬", Number: "092", Name: "Virgin Islands (British)"},
	VI: {Alpha2: "VI", Alpha3: "VIR", Flag: "🇻🇮", Number: "850", Name: "Virgin Islands (U.S.)"},
	VN: {Alpha2: "VN", Alpha3: "VNM", Flag: "🇻🇳", Number: "704", Name: "Viet Nam"},
	VU: {Alpha2: "VU", Alpha3: "VUT", Flag: "🇻🇺", Number: "548", Name: "Vanuatu"},
	WF: {Alpha2: "WF", Alpha3: "WLF", Flag: "🇼🇫", Number: "876", Name: "Wallis and Futuna"},
	WS: {Alpha2: "WS", Alpha3: "WSM", Flag: "🇼🇸", Number: "882", Name: "Samoa"},
	YE: {Alpha2: "YE", Alpha3: "YEM", Flag: "🇾🇪", Number: "887", Name: "Yemen"},
	YT: {Alpha2: "YT", Alpha3: "MYT", Flag: "🇾🇹", Number: "175", Name: "Mayotte"},
	ZA: {Alpha2: "ZA", Alpha3: "ZAF", Flag: "🇿🇦", Number: "710", Name: "South Africa"},
	ZM: {Alpha2: "ZM", Alpha3: "ZMB", Flag: "🇿🇲", Number: "894", Name: "Zambia"},
	ZW: {Alpha2: "ZW", Alpha3: "ZWE", Flag: "🇿🇼", Number: "716", Name: "Zimbabwe"},
}

var stringToCountryCode = map[string]CountryCode{
	"AD": AD, "AE": AE, "AF": AF, "AG": AG, "AI": AI, "AL": AL, "AM": AM, "AO": AO, "AQ": AQ, "AR": AR,
	"AS": AS, "AT": AT, "AU": AU, "AW": AW, "AX": AX, "AZ": AZ, "BA": BA, "BB": BB, "BD": BD, "BE": BE,
	"BF": BF, "BG": BG, "BH": BH, "BI": BI, "BJ": BJ, "BL": BL, "BM": BM, "BN": BN, "BO": BO, "BQ": BQ,
	"BR": BR, "BS": BS, "BT": BT, "BV": BV, "BW": BW, "BY": BY, "BZ": BZ, "CA": CA, "CC": CC, "CD": CD,
	"CF": CF, "CG": CG, "CH": CH, "CI": CI, "CK": CK, "CL": CL, "CM": CM, "CN": CN, "CO": CO, "CR": CR,
	"CU": CU, "CV": CV, "CW": CW, "CX": CX, "CY": CY, "CZ": CZ, "DE": DE, "DJ": DJ, "DK": DK, "DM": DM,
	"DO": DO, "DZ": DZ, "EC": EC, "EE": EE, "EG": EG, "EH": EH, "ER": ER, "ES": ES, "ET": ET, "FI": FI,
	"FJ": FJ, "FK": FK, "FM": FM, "FO": FO, "FR": FR, "GA": GA, "GB": GB, "GD": GD, "GE": GE, "GF": GF,
	"GG": GG, "GH": GH, "GI": GI, "GL": GL, "GM": GM, "GN": GN, "GP": GP, "GQ": GQ, "GR": GR, "GS": GS,
	"GT": GT, "GU": GU, "GW": GW, "GY": GY, "HK": HK, "HM": HM, "HN": HN, "HR": HR, "HT": HT, "HU": HU,
	"ID": ID, "IE": IE, "IL": IL, "IM": IM, "IN": IN, "IO": IO, "IQ": IQ, "IR": IR, "IS": IS, "IT": IT,
	"JE": JE, "JM": JM, "JO": JO, "JP": JP, "KE": KE, "KG": KG, "KH": KH, "KI": KI, "KM": KM, "KN": KN,
	"KP": KP, "KR": KR, "KW": KW, "KY": KY, "KZ": KZ, "LA": LA, "LB": LB, "LC": LC, "LI": LI, "LK": LK,
	"LR": LR, "LS": LS, "LT": LT, "LU": LU, "LV": LV, "LY": LY, "MA": MA, "MC": MC, "MD": MD, "ME": ME,
	"MF": MF, "MG": MG, "MH": MH, "MK": MK, "ML": ML, "MM": MM, "MN": MN, "MO": MO, "MP": MP, "MQ": MQ,
	"MR": MR, "MS": MS, "MT": MT, "MU": MU, "MV": MV, "MW": MW, "MX": MX, "MY": MY, "MZ": MZ, "NA": NA,
	"NC": NC, "NE": NE, "NF": NF, "NG": NG, "NI": NI, "NL": NL, "NO": NO, "NP": NP, "NR": NR, "NU": NU,
	"NZ": NZ, "OM": OM, "PA": PA, "PE": PE, "PF": PF, "PG": PG, "PH": PH, "PK": PK, "PL": PL, "PM": PM,
	"PN": PN, "PR": PR, "PS": PS, "PT": PT, "PW": PW, "PY": PY, "QA": QA, "RE": RE, "RO": RO, "RS": RS,
	"RU": RU, "RW": RW, "SA": SA, "SB": SB, "SC": SC, "SD": SD, "SE": SE, "SG": SG, "SH": SH, "SI": SI,
	"SJ": SJ, "SK": SK, "SL": SL, "SM": SM, "SN": SN, "SO": SO, "SR": SR, "SS": SS, "ST": ST, "SV": SV,
	"SX": SX, "SY": SY, "SZ": SZ, "TC": TC, "TD": TD, "TF": TF, "TG": TG, "TH": TH, "TJ": TJ, "TK": TK,
	"TL": TL, "TM": TM, "TN": TN, "TO": TO, "TR": TR, "TT": TT, "TV": TV, "TW": TW, "TZ": TZ, "UA": UA,
	"UG": UG, "UM": UM, "US": US, "UY": UY, "UZ": UZ, "VA": VA, "VC": VC, "VE": VE, "VG": VG, "VI": VI,
	"VN": VN, "VU": VU, "WF": WF, "WS": WS, "YE": YE, "YT": YT, "ZA": ZA, "ZM": ZM, "ZW": ZW,
}
