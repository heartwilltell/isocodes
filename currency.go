package isocodes

import (
	"strings"
)

// CurrencyCode represents an ISO currency code.
type CurrencyCode byte

// String returns string representation of the code.
func (c CurrencyCode) String() string { return currencyCodesDetails[c].Code }

// Name returns a currency name.
func (c CurrencyCode) Name() string { return currencyCodesDetails[c].Name }

// Number returns a currency number related to the CurrencyCode.
func (c CurrencyCode) Number() string { return currencyCodesDetails[c].Number }

// Flag returns an emoji flag for the country code.
func (c CurrencyCode) Flag() string { return currencyCodesDetails[c].Flag }

func (c *CurrencyCode) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return ErrUnmarshalJSON
	}

	code, ok := stringToCurrencyCode[string(b)]
	if !ok {
		return ErrUnmarshalJSON
	}

	*c = code

	return nil
}

func (c CurrencyCode) MarshalJSON() ([]byte, error) {
	code := c.String()
	if code == "" {
		return nil, ErrMarshalJSON
	}

	return []byte(`"` + code + `"`), nil
}

func (c *CurrencyCode) UnmarshalText(b []byte) error { return c.UnmarshalJSON(b) }

func (c CurrencyCode) MarshalText() ([]byte, error) { return c.MarshalJSON() }

// CurrencyCodeDetails represents detailed information related to the currency code.
type CurrencyCodeDetails struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Number   string `json:"number"`
	Flag     string `json:"flag"`
	Decimals int    `json:"decimals"`
}

// StringToCurrencyCode takes string representation of an ISO currency
// code and returns a CurrencyCode.
func StringToCurrencyCode(code string) (CurrencyCode, error) {
	if len(code) > 3 {
		return 0, ErrInvalidStringCode
	}

	c, ok := stringToCurrencyCode[strings.ToUpper(code)]
	if !ok {
		return 0, ErrInvalidStringCode
	}

	return c, nil
}

const (
	// AED represents ISO currency code of the United Arab Emirates dirham.
	AED CurrencyCode = iota + 1
	// AFN represents ISO currency code of the Afghan afghani.
	AFN
	// ALL represents ISO currency code of the Albanian lek.
	ALL
	// AMD represents ISO currency code of the Armenian dram.
	AMD
	// ANG represents ISO currency code of the Netherlands Antillean guilder.
	ANG
	// AOA represents ISO currency code of the Angolan kwanza.
	AOA
	// ARS represents ISO currency code of the Argentine peso.
	ARS
	// AUD represents ISO currency code of the Australian dollar.
	AUD
	// AWG represents ISO currency code of the Aruban florin.
	AWG
	// AZN represents ISO currency code of the Azerbaijani manat.
	AZN
	// BAM represents ISO currency code of the Bosnia and Herzegovina convertible mark.
	BAM
	// BBD represents ISO currency code of the Barbados dollar.
	BBD
	// BDT represents ISO currency code of the Bangladeshi taka.
	BDT
	// BGN represents ISO currency code of the Bulgarian lev.
	BGN
	// BHD represents ISO currency code of the Bahraini dinar.
	BHD
	// BIF represents ISO currency code of the Burundian franc.
	BIF
	// BMD represents ISO currency code of the Bermudian dollar (customarily known as Bermuda dollar).
	BMD
	// BND represents ISO currency code of the Brunei dollar.
	BND
	// BOB represents ISO currency code of the Boliviano.
	BOB
	// BOV represents ISO currency code of the Bolivian Mvdol (funds code).
	BOV
	// BRL represents ISO currency code of the Brazilian real.
	BRL
	// BSD represents ISO currency code of the Bahamian dollar.
	BSD
	// BTN represents ISO currency code of the Bhutanese ngultrum.
	BTN
	// BWP represents ISO currency code of the Botswana pula.
	BWP
	// BYR represents ISO currency code of the Belarusian ruble.
	BYR
	// BZD represents ISO currency code of the Belize dollar.
	BZD
	// CAD represents ISO currency code of the Canadian dollar.
	CAD
	// CDF represents ISO currency code of the Congolese franc.
	CDF
	// CHE represents ISO currency code of the WIR Euro (complementary currency).
	CHE
	// CHF represents ISO currency code of the Swiss franc.
	CHF
	// CHW represents ISO currency code of the WIR Franc (complementary currency).
	CHW
	// CLF represents ISO currency code of the Unidad de Fomento (funds code).
	CLF
	// CLP represents ISO currency code of the Chilean peso.
	CLP
	// CNY represents ISO currency code of the Chinese yuan.
	CNY
	// COP represents ISO currency code of the Colombian peso.
	COP
	// COU represents ISO currency code of the Unidad de Valor Real.
	COU
	// CRC represents ISO currency code of the Costa Rican colon.
	CRC
	// CUC represents ISO currency code of the Cuban convertible peso.
	CUC
	// CUP represents ISO currency code of the Cuban peso.
	CUP
	// CVE represents ISO currency code of the Cape Verde escudo.
	CVE
	// CZK represents ISO currency code of the Czech koruna.
	CZK
	// DJF represents ISO currency code of the Djiboutian franc.
	DJF
	// DKK represents ISO currency code of the Danish krone.
	DKK
	// DOP represents ISO currency code of the Dominican peso.
	DOP
	// DZD represents ISO currency code of the Algerian dinar.
	DZD
	// EGP represents ISO currency code of the Egyptian pound.
	EGP
	// ERN represents ISO currency code of the Eritrean nakfa.
	ERN
	// ETB represents ISO currency code of the Ethiopian birr.
	ETB
	// EUR represents ISO currency code of the Euro.
	EUR
	// FJD represents ISO currency code of the Fiji dollar.
	FJD
	// FKP represents ISO currency code of the Falkland Islands pound.
	FKP
	// GBP represents ISO currency code of the Pound sterling.
	GBP
	// GEL represents ISO currency code of the Georgian lari.
	GEL
	// GHS represents ISO currency code of the Ghanaian cedi.
	GHS
	// GIP represents ISO currency code of the Gibraltar pound.
	GIP
	// GMD represents ISO currency code of the Gambian dalasi.
	GMD
	// GNF represents ISO currency code of the Guinean franc.
	GNF
	// GTQ represents ISO currency code of the Guatemalan quetzal.
	GTQ
	// GYD represents ISO currency code of the Guyanese dollar.
	GYD
	// HKD represents ISO currency code of the Hong Kong dollar.
	HKD
	// HNL represents ISO currency code of the Honduran lempira.
	HNL
	// HRK represents ISO currency code of the Croatian kuna.
	HRK
	// HTG represents ISO currency code of the Haitian gourde.
	HTG
	// HUF represents ISO currency code of the Hungarian forint.
	HUF
	// IDR represents ISO currency code of the Indonesian rupiah.
	IDR
	// ILS represents ISO currency code of the Israeli new shekel.
	ILS
	// INR represents ISO currency code of the Indian rupee.
	INR
	// IQD represents ISO currency code of the Iraqi dinar.
	IQD
	// IRR represents ISO currency code of the Iranian rial.
	IRR
	// ISK represents ISO currency code of the Icelandic króna.
	ISK
	// JMD represents ISO currency code of the Jamaican dollar.
	JMD
	// JOD represents ISO currency code of the Jordanian dinar.
	JOD
	// JPY represents ISO currency code of the Japanese yen.
	JPY
	// KES represents ISO currency code of the Kenyan shilling.
	KES
	// KGS represents ISO currency code of the Kyrgyzstani som.
	KGS
	// KHR represents ISO currency code of the Cambodian riel.
	KHR
	// KMF represents ISO currency code of the Comoro franc.
	KMF
	// KPW represents ISO currency code of the North Korean won.
	KPW
	// KRW represents ISO currency code of the South Korean won.
	KRW
	// KWD represents ISO currency code of the Kuwaiti dinar.
	KWD
	// KYD represents ISO currency code of the Cayman Islands dollar.
	KYD
	// KZT represents ISO currency code of the Kazakhstani tenge.
	KZT
	// LAK represents ISO currency code of the Lao kip.
	LAK
	// LBP represents ISO currency code of the Lebanese pound.
	LBP
	// LKR represents ISO currency code of the Sri Lankan rupee.
	LKR
	// LRD represents ISO currency code of the Liberian dollar.
	LRD
	// LSL represents ISO currency code of the Lesotho loti.
	LSL
	// LTL represents ISO currency code of the Lithuanian litas.
	LTL
	// LVL represents ISO currency code of the Latvian lats.
	LVL
	// LYD represents ISO currency code of the Libyan dinar.
	LYD
	// MAD represents ISO currency code of the Moroccan dirham.
	MAD
	// MDL represents ISO currency code of the Moldovan leu.
	MDL
	// MGA represents ISO currency code of the Malagasy ariary.
	MGA
	// MKD represents ISO currency code of the Macedonian denar.
	MKD
	// MMK represents ISO currency code of the Myanma kyat.
	MMK
	// MNT represents ISO currency code of the Mongolian tugrik.
	MNT
	// MOP represents ISO currency code of the Macanese pataca.
	MOP
	// MRO represents ISO currency code of the Mauritanian ouguiya.
	MRO
	// MUR represents ISO currency code of the Mauritian rupee.
	MUR
	// MVR represents ISO currency code of the Maldivian rufiyaa.
	MVR
	// MWK represents ISO currency code of the Malawian kwacha.
	MWK
	// MXN represents ISO currency code of the Mexican peso.
	MXN
	// MXV represents ISO currency code of the Mexican Unidad de Inversion (UDI) (funds code).
	MXV
	// MYR represents ISO currency code of the Malaysian ringgit.
	MYR
	// MZN represents ISO currency code of the Mozambican metical.
	MZN
	// NAD represents ISO currency code of the Namibian dollar.
	NAD
	// NGN represents ISO currency code of the Nigerian naira.
	NGN
	// NIO represents ISO currency code of the Nicaraguan córdoba.
	NIO
	// NOK represents ISO currency code of the Norwegian krone.
	NOK
	// NPR represents ISO currency code of the Nepalese rupee.
	NPR
	// NZD represents ISO currency code of the New Zealand dollar.
	NZD
	// OMR represents ISO currency code of the Omani rial.
	OMR
	// PAB represents ISO currency code of the Panamanian balboa.
	PAB
	// PEN represents ISO currency code of the Peruvian nuevo sol.
	PEN
	// PGK represents ISO currency code of the Papua New Guinean kina.
	PGK
	// PHP represents ISO currency code of the Philippine peso.
	PHP
	// PKR represents ISO currency code of the Pakistani rupee.
	PKR
	// PLN represents ISO currency code of the Polish złoty.
	PLN
	// PYG represents ISO currency code of the Paraguayan guaraní.
	PYG
	// QAR represents ISO currency code of the Qatari riyal.
	QAR
	// RON represents ISO currency code of the Romanian new leu.
	RON
	// RSD represents ISO currency code of the Serbian dinar.
	RSD
	// RUB represents ISO currency code of the Russian rouble.
	RUB
	// RWF represents ISO currency code of the Rwandan franc.
	RWF
	// SAR represents ISO currency code of the Saudi riyal.
	SAR
	// SBD represents ISO currency code of the Solomon Islands dollar.
	SBD
	// SCR represents ISO currency code of the Seychelles rupee.
	SCR
	// SDG represents ISO currency code of the Sudanese pound.
	SDG
	// SEK represents ISO currency code of the Swedish krona/kronor.
	SEK
	// SGD represents ISO currency code of the Singapore dollar.
	SGD
	// SHP represents ISO currency code of the Saint Helena pound.
	SHP
	// SLL represents ISO currency code of the Sierra Leonean leone.
	SLL
	// SOS represents ISO currency code of the Somali shilling.
	SOS
	// SRD represents ISO currency code of the Surinamese dollar.
	SRD
	// SSP represents ISO currency code of the South Sudanese pound.
	SSP
	// STD represents ISO currency code of the São Tomé and Príncipe dobra.
	STD
	// SYP represents ISO currency code of the Syrian pound.
	SYP
	// SZL represents ISO currency code of the Swazi lilangeni.
	SZL
	// THB represents ISO currency code of the Thai baht.
	THB
	// TJS represents ISO currency code of the Tajikistani somoni.
	TJS
	// TMT represents ISO currency code of the Turkmenistani manat.
	TMT
	// TND represents ISO currency code of the Tunisian dinar.
	TND
	// TOP represents ISO currency code of the Tongan paʻanga.
	TOP
	// TRY represents ISO currency code of the Turkish lira.
	TRY
	// TTD represents ISO currency code of the Trinidad and Tobago dollar.
	TTD
	// TWD represents ISO currency code of the New Taiwan dollar.
	TWD
	// TZS represents ISO currency code of the Tanzanian shilling.
	TZS
	// UAH represents ISO currency code of the Ukrainian hryvnia.
	UAH
	// UGX represents ISO currency code of the Ugandan shilling.
	UGX
	// USD represents ISO currency code of the United States dollar.
	USD
	// USN represents ISO currency code of the United States dollar (next day) (funds code).
	USN
	// USS represents ISO currency code of the United States dollar (same day) (funds code).
	USS
	// UYI represents ISO currency code of the Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code).
	UYI
	// UYU represents ISO currency code of the Uruguayan peso.
	UYU
	// UZS represents ISO currency code of the Uzbekistan som.
	UZS
	// VEF represents ISO currency code of the Venezuelan bolívar fuerte.
	VEF
	// VND represents ISO currency code of the Vietnamese dong.
	VND
	// VUV represents ISO currency code of the Vanuatu vatu.
	VUV
	// WST represents ISO currency code of the Samoan tala.
	WST
	// XAF represents ISO currency code of the CFA franc BEAC.
	XAF
	// XAG represents ISO currency code of the Silver (one troy ounce).
	XAG
	// XAU represents ISO currency code of the Gold (one troy ounce).
	XAU
	// XBA represents ISO currency code of the European Composite Unit (EURCO) (bond market unit).
	XBA
	// XBB represents ISO currency code of the European Monetary Unit (E.M.U.-6) (bond market unit).
	XBB
	// XBC represents ISO currency code of the European Unit of Account 9 (E.U.A.-9) (bond market unit).
	XBC
	// XBD represents ISO currency code of the European Unit of Account 17 (E.U.A.-17) (bond market unit).
	XBD
	// XCD represents ISO currency code of the East Caribbean dollar.
	XCD
	// XDR represents ISO currency code of the Special drawing rights.
	XDR
	// XFU represents ISO currency code of the UIC franc (special settlement currency).
	XFU
	// XOF represents ISO currency code of the CFA franc BCEAO.
	XOF
	// XPD represents ISO currency code of the Palladium (one troy ounce).
	XPD
	// XPF represents ISO currency code of the CFP franc.
	XPF
	// XPT represents ISO currency code of the Platinum (one troy ounce).
	XPT
	// XTS represents ISO currency code of the CurrencyCode reserved for testing purposes.
	XTS
	// XXX represents ISO currency code of the No currency.
	XXX
	// YER represents ISO currency code of the Yemeni rial.
	YER
	// ZAR represents ISO currency code of the South African rand.
	ZAR
	// ZMW represents ISO currency code of the Zambian kwacha.
	ZMW
)

var currencyCodesDetails = map[CurrencyCode]CurrencyCodeDetails{
	AED: {Code: "AED", Name: "United Arab Emirates dirham", Number: "784", Flag: "🇦🇪", Decimals: 2},
	AFN: {Code: "AFN", Name: "Afghan afghani", Number: "971", Flag: "🇦🇫", Decimals: 2},
	ALL: {Code: "ALL", Name: "Albanian lek", Number: "8", Flag: "🇦🇱", Decimals: 2},
	AMD: {Code: "AMD", Name: "Armenian dram", Number: "51", Flag: "🇦🇲", Decimals: 2},
	ANG: {Code: "ANG", Name: "Netherlands Antillean guilder", Number: "532", Flag: "🇳🇱", Decimals: 2},
	AOA: {Code: "AOA", Name: "Angolan kwanza", Number: "973", Flag: "🇦🇴", Decimals: 2},
	ARS: {Code: "ARS", Name: "Argentine peso", Number: "32", Flag: "🇦🇷", Decimals: 2},
	AUD: {Code: "AUD", Name: "Australian dollar", Number: "36", Flag: "🇦🇺", Decimals: 2},
	AWG: {Code: "AWG", Name: "Aruban florin", Number: "533", Flag: "🇦🇼", Decimals: 2},
	AZN: {Code: "AZN", Name: "Azerbaijani manat", Number: "944", Flag: "🇦🇿", Decimals: 2},
	BAM: {Code: "BAM", Name: "Bosnia and Herzegovina convertible mark", Number: "977", Flag: "🇧🇦", Decimals: 2},
	BBD: {Code: "BBD", Name: "Barbados dollar", Number: "52", Flag: "🇧🇧", Decimals: 2},
	BDT: {Code: "BDT", Name: "Bangladeshi taka", Number: "50", Flag: "🇧🇩", Decimals: 2},
	BGN: {Code: "BGN", Name: "Bulgarian lev", Number: "975", Flag: "🇧🇬", Decimals: 2},
	BHD: {Code: "BHD", Name: "Bahraini dinar", Number: "48", Flag: "", Decimals: 3},
	BIF: {Code: "BIF", Name: "Burundian franc", Number: "108", Flag: "🇧🇮", Decimals: 0},
	BMD: {Code: "BMD", Name: "Bermudian dollar (customarily known as Bermuda dollar)", Number: "60", Flag: "🇧🇲", Decimals: 2},
	BND: {Code: "BND", Name: "Brunei dollar", Number: "96", Flag: "🇧🇳", Decimals: 2},
	BOB: {Code: "BOB", Name: "Boliviano", Number: "68", Flag: "🇧🇴", Decimals: 2},
	BOV: {Code: "BOV", Name: "Bolivian Mvdol (funds code)", Number: "984", Flag: "", Decimals: 2},
	BRL: {Code: "BRL", Name: "Brazilian real", Number: "986", Flag: "🇧🇷", Decimals: 2},
	BSD: {Code: "BSD", Name: "Bahamian dollar", Number: "44", Flag: "🇧🇸", Decimals: 2},
	BTN: {Code: "BTN", Name: "Bhutanese ngultrum", Number: "64", Flag: "", Decimals: 2},
	BWP: {Code: "BWP", Name: "Botswana pula", Number: "72", Flag: "🇧🇼", Decimals: 2},
	BYR: {Code: "BYR", Name: "Belarusian ruble", Number: "974", Flag: "", Decimals: 0},
	BZD: {Code: "BZD", Name: "Belize dollar", Number: "84", Flag: "🇧🇿", Decimals: 2},
	CAD: {Code: "CAD", Name: "Canadian dollar", Number: "124", Flag: "🇨🇦", Decimals: 2},
	CDF: {Code: "CDF", Name: "Congolese franc", Number: "976", Flag: "🇨🇩", Decimals: 2},
	CHE: {Code: "CHE", Name: "WIR Euro (complementary currency)", Number: "947", Flag: "", Decimals: 2},
	CHF: {Code: "CHF", Name: "Swiss franc", Number: "756", Flag: "🇨🇭", Decimals: 2},
	CHW: {Code: "CHW", Name: "WIR Franc (complementary currency)", Number: "948", Flag: "", Decimals: 2},
	CLF: {Code: "CLF", Name: "Unidad de Fomento (funds code)", Number: "990", Flag: "", Decimals: 0},
	CLP: {Code: "CLP", Name: "Chilean peso", Number: "152", Flag: "🇨🇱", Decimals: 0},
	CNY: {Code: "CNY", Name: "Chinese yuan", Number: "156", Flag: "🇨🇳", Decimals: 2},
	COP: {Code: "COP", Name: "Colombian peso", Number: "170", Flag: "🇨🇴", Decimals: 2},
	COU: {Code: "COU", Name: "Unidad de Valor Real", Number: "970", Flag: "", Decimals: 2},
	CRC: {Code: "CRC", Name: "Costa Rican colon", Number: "188", Flag: "🇨🇷", Decimals: 2},
	CUC: {Code: "CUC", Name: "Cuban convertible peso", Number: "931", Flag: "", Decimals: 2},
	CUP: {Code: "CUP", Name: "Cuban peso", Number: "192", Flag: "", Decimals: 2},
	CVE: {Code: "CVE", Name: "Cape Verde escudo", Number: "132", Flag: "🇨🇻", Decimals: 0},
	CZK: {Code: "CZK", Name: "Czech koruna", Number: "203", Flag: "🇨🇿", Decimals: 2},
	DJF: {Code: "DJF", Name: "Djiboutian franc", Number: "262", Flag: "🇩🇯", Decimals: 0},
	DKK: {Code: "DKK", Name: "Danish krone", Number: "208", Flag: "🇩🇰", Decimals: 2},
	DOP: {Code: "DOP", Name: "Dominican peso", Number: "214", Flag: "🇩🇴", Decimals: 2},
	DZD: {Code: "DZD", Name: "Algerian dinar", Number: "12", Flag: "🇩🇿", Decimals: 2},
	EGP: {Code: "EGP", Name: "Egyptian pound", Number: "818", Flag: "🇪🇬", Decimals: 2},
	ERN: {Code: "ERN", Name: "Eritrean nakfa", Number: "232", Flag: "", Decimals: 2},
	ETB: {Code: "ETB", Name: "Ethiopian birr", Number: "230", Flag: "🇪🇹", Decimals: 2},
	EUR: {Code: "EUR", Name: "Euro", Number: "978", Flag: "🇪🇺", Decimals: 2},
	FJD: {Code: "FJD", Name: "Fiji dollar", Number: "242", Flag: "🇫🇯", Decimals: 2},
	FKP: {Code: "FKP", Name: "Falkland Islands pound", Number: "238", Flag: "🇫🇰", Decimals: 2},
	GBP: {Code: "GBP", Name: "Pound sterling", Number: "826", Flag: "🇬🇧", Decimals: 2},
	GEL: {Code: "GEL", Name: "Georgian lari", Number: "981", Flag: "🇬🇪", Decimals: 2},
	GHS: {Code: "GHS", Name: "Ghanaian cedi", Number: "936", Flag: "", Decimals: 2},
	GIP: {Code: "GIP", Name: "Gibraltar pound", Number: "292", Flag: "🇬🇮", Decimals: 2},
	GMD: {Code: "GMD", Name: "Gambian dalasi", Number: "270", Flag: "🇬🇲", Decimals: 2},
	GNF: {Code: "GNF", Name: "Guinean franc", Number: "324", Flag: "🇬🇳", Decimals: 0},
	GTQ: {Code: "GTQ", Name: "Guatemalan quetzal", Number: "320", Flag: "🇬🇹", Decimals: 2},
	GYD: {Code: "GYD", Name: "Guyanese dollar", Number: "328", Flag: "🇬🇾", Decimals: 2},
	HKD: {Code: "HKD", Name: "Hong Kong dollar", Number: "344", Flag: "🇭🇰", Decimals: 2},
	HNL: {Code: "HNL", Name: "Honduran lempira", Number: "340", Flag: "🇭🇳", Decimals: 2},
	HRK: {Code: "HRK", Name: "Croatian kuna", Number: "191", Flag: "🇭🇷", Decimals: 2},
	HTG: {Code: "HTG", Name: "Haitian gourde", Number: "332", Flag: "🇭🇹", Decimals: 2},
	HUF: {Code: "HUF", Name: "Hungarian forint", Number: "348", Flag: "🇭🇺", Decimals: 2},
	IDR: {Code: "IDR", Name: "Indonesian rupiah", Number: "360", Flag: "🇮🇩", Decimals: 2},
	ILS: {Code: "ILS", Name: "Israeli new shekel", Number: "376", Flag: "🇮🇱", Decimals: 2},
	INR: {Code: "INR", Name: "Indian rupee", Number: "356", Flag: "🇮🇳", Decimals: 2},
	IQD: {Code: "IQD", Name: "Iraqi dinar", Number: "368", Flag: "", Decimals: 3},
	IRR: {Code: "IRR", Name: "Iranian rial", Number: "364", Flag: "", Decimals: 0},
	ISK: {Code: "ISK", Name: "Icelandic króna", Number: "352", Flag: "🇮🇸", Decimals: 0},
	JMD: {Code: "JMD", Name: "Jamaican dollar", Number: "388", Flag: "🇯🇲", Decimals: 2},
	JOD: {Code: "JOD", Name: "Jordanian dinar", Number: "400", Flag: "", Decimals: 3},
	JPY: {Code: "JPY", Name: "Japanese yen", Number: "392", Flag: "🇯🇵", Decimals: 0},
	KES: {Code: "KES", Name: "Kenyan shilling", Number: "404", Flag: "🇰🇪", Decimals: 2},
	KGS: {Code: "KGS", Name: "Kyrgyzstani som", Number: "417", Flag: "🇰🇬", Decimals: 2},
	KHR: {Code: "KHR", Name: "Cambodian riel", Number: "116", Flag: "🇰🇭", Decimals: 2},
	KMF: {Code: "KMF", Name: "Comoro franc", Number: "174", Flag: "🇰🇲", Decimals: 0},
	KPW: {Code: "KPW", Name: "North Korean won", Number: "408", Flag: "", Decimals: 0},
	KRW: {Code: "KRW", Name: "South Korean won", Number: "410", Flag: "🇰🇷", Decimals: 0},
	KWD: {Code: "KWD", Name: "Kuwaiti dinar", Number: "414", Flag: "", Decimals: 3},
	KYD: {Code: "KYD", Name: "Cayman Islands dollar", Number: "136", Flag: "🇰🇾", Decimals: 2},
	KZT: {Code: "KZT", Name: "Kazakhstani tenge", Number: "398", Flag: "🇰🇿", Decimals: 2},
	LAK: {Code: "LAK", Name: "Lao kip", Number: "418", Flag: "🇱🇦", Decimals: 0},
	LBP: {Code: "LBP", Name: "Lebanese pound", Number: "422", Flag: "🇱🇧", Decimals: 0},
	LKR: {Code: "LKR", Name: "Sri Lankan rupee", Number: "144", Flag: "🇱🇰", Decimals: 2},
	LRD: {Code: "LRD", Name: "Liberian dollar", Number: "430", Flag: "🇱🇷", Decimals: 2},
	LSL: {Code: "LSL", Name: "Lesotho loti", Number: "426", Flag: "🇱🇸", Decimals: 2},
	LTL: {Code: "LTL", Name: "Lithuanian litas", Number: "440", Flag: "", Decimals: 2},
	LVL: {Code: "LVL", Name: "Latvian lats", Number: "428", Flag: "", Decimals: 2},
	LYD: {Code: "LYD", Name: "Libyan dinar", Number: "434", Flag: "", Decimals: 3},
	MAD: {Code: "MAD", Name: "Moroccan dirham", Number: "504", Flag: "🇲🇦", Decimals: 2},
	MDL: {Code: "MDL", Name: "Moldovan leu", Number: "498", Flag: "🇲🇩", Decimals: 2},
	MGA: {Code: "MGA", Name: "Malagasy ariary", Number: "969", Flag: "🇲🇬", Decimals: 0},
	MKD: {Code: "MKD", Name: "Macedonian denar", Number: "807", Flag: "🇲🇰", Decimals: 0},
	MMK: {Code: "MMK", Name: "Myanma kyat", Number: "104", Flag: "🇲🇲", Decimals: 0},
	MNT: {Code: "MNT", Name: "Mongolian tugrik", Number: "496", Flag: "🇲🇳", Decimals: 2},
	MOP: {Code: "MOP", Name: "Macanese pataca", Number: "446", Flag: "🇲🇴", Decimals: 2},
	MRO: {Code: "MRO", Name: "Mauritanian ouguiya", Number: "478", Flag: "🇲🇷", Decimals: 0},
	MUR: {Code: "MUR", Name: "Mauritian rupee", Number: "480", Flag: "🇲🇺", Decimals: 2},
	MVR: {Code: "MVR", Name: "Maldivian rufiyaa", Number: "462", Flag: "🇲🇻", Decimals: 2},
	MWK: {Code: "MWK", Name: "Malawian kwacha", Number: "454", Flag: "🇲🇼", Decimals: 2},
	MXN: {Code: "MXN", Name: "Mexican peso", Number: "484", Flag: "🇲🇽", Decimals: 2},
	MXV: {Code: "MXV", Name: "Mexican Unidad de Inversion (UDI) (funds code)", Number: "979", Flag: "", Decimals: 2},
	MYR: {Code: "MYR", Name: "Malaysian ringgit", Number: "458", Flag: "🇲🇾", Decimals: 2},
	MZN: {Code: "MZN", Name: "Mozambican metical", Number: "943", Flag: "🇲🇿", Decimals: 2},
	NAD: {Code: "NAD", Name: "Namibian dollar", Number: "516", Flag: "🇳🇦", Decimals: 2},
	NGN: {Code: "NGN", Name: "Nigerian naira", Number: "566", Flag: "🇳🇬", Decimals: 2},
	NIO: {Code: "NIO", Name: "Nicaraguan córdoba", Number: "558", Flag: "🇳🇮", Decimals: 2},
	NOK: {Code: "NOK", Name: "Norwegian krone", Number: "578", Flag: "🇳🇴", Decimals: 2},
	NPR: {Code: "NPR", Name: "Nepalese rupee", Number: "524", Flag: "🇳🇵", Decimals: 2},
	NZD: {Code: "NZD", Name: "New Zealand dollar", Number: "554", Flag: "🇳🇿", Decimals: 2},
	OMR: {Code: "OMR", Name: "Omani rial", Number: "512", Flag: "", Decimals: 3},
	PAB: {Code: "PAB", Name: "Panamanian balboa", Number: "590", Flag: "🇵🇦", Decimals: 2},
	PEN: {Code: "PEN", Name: "Peruvian nuevo sol", Number: "604", Flag: "🇵🇪", Decimals: 2},
	PGK: {Code: "PGK", Name: "Papua New Guinean kina", Number: "598", Flag: "🇵🇬", Decimals: 2},
	PHP: {Code: "PHP", Name: "Philippine peso", Number: "608", Flag: "🇵🇭", Decimals: 2},
	PKR: {Code: "PKR", Name: "Pakistani rupee", Number: "586", Flag: "🇵🇰", Decimals: 2},
	PLN: {Code: "PLN", Name: "Polish złoty", Number: "985", Flag: "🇵🇱", Decimals: 2},
	PYG: {Code: "PYG", Name: "Paraguayan guaraní", Number: "600", Flag: "🇵🇾", Decimals: 0},
	QAR: {Code: "QAR", Name: "Qatari riyal", Number: "634", Flag: "🇶🇦", Decimals: 2},
	RON: {Code: "RON", Name: "Romanian new leu", Number: "946", Flag: "🇷🇴", Decimals: 2},
	RSD: {Code: "RSD", Name: "Serbian dinar", Number: "941", Flag: "🇷🇸", Decimals: 2},
	RUB: {Code: "RUB", Name: "Russian rouble", Number: "643", Flag: "🇷🇺", Decimals: 2},
	RWF: {Code: "RWF", Name: "Rwandan franc", Number: "646", Flag: "🇷🇼", Decimals: 0},
	SAR: {Code: "SAR", Name: "Saudi riyal", Number: "682", Flag: "🇸🇦", Decimals: 2},
	SBD: {Code: "SBD", Name: "Solomon Islands dollar", Number: "90", Flag: "🇸🇧", Decimals: 2},
	SCR: {Code: "SCR", Name: "Seychelles rupee", Number: "690", Flag: "🇸🇨", Decimals: 2},
	SDG: {Code: "SDG", Name: "Sudanese pound", Number: "938", Flag: "", Decimals: 2},
	SEK: {Code: "SEK", Name: "Swedish krona/kronor", Number: "752", Flag: "🇸🇪", Decimals: 2},
	SGD: {Code: "SGD", Name: "Singapore dollar", Number: "702", Flag: "🇸🇬", Decimals: 2},
	SHP: {Code: "SHP", Name: "Saint Helena pound", Number: "654", Flag: "🇸🇭", Decimals: 2},
	SLL: {Code: "SLL", Name: "Sierra Leonean leone", Number: "694", Flag: "🇸🇱", Decimals: 0},
	SOS: {Code: "SOS", Name: "Somali shilling", Number: "706", Flag: "🇸🇴", Decimals: 2},
	SRD: {Code: "SRD", Name: "Surinamese dollar", Number: "968", Flag: "🇸🇷", Decimals: 2},
	SSP: {Code: "SSP", Name: "South Sudanese pound", Number: "728", Flag: "", Decimals: 2},
	STD: {Code: "STD", Name: "São Tomé and Príncipe dobra", Number: "678", Flag: "🇸🇹", Decimals: 0},
	SYP: {Code: "SYP", Name: "Syrian pound", Number: "760", Flag: "", Decimals: 2},
	SZL: {Code: "SZL", Name: "Swazi lilangeni", Number: "748", Flag: "🇸🇿", Decimals: 2},
	THB: {Code: "THB", Name: "Thai baht", Number: "764", Flag: "🇹🇭", Decimals: 2},
	TJS: {Code: "TJS", Name: "Tajikistani somoni", Number: "972", Flag: "🇹🇯", Decimals: 2},
	TMT: {Code: "TMT", Name: "Turkmenistani manat", Number: "934", Flag: "", Decimals: 2},
	TND: {Code: "TND", Name: "Tunisian dinar", Number: "788", Flag: "", Decimals: 3},
	TOP: {Code: "TOP", Name: "Tongan paʻanga", Number: "776", Flag: "🇹🇴", Decimals: 2},
	TRY: {Code: "TRY", Name: "Turkish lira", Number: "949", Flag: "🇹🇷", Decimals: 2},
	TTD: {Code: "TTD", Name: "Trinidad and Tobago dollar", Number: "780", Flag: "🇹🇹", Decimals: 2},
	TWD: {Code: "TWD", Name: "New Taiwan dollar", Number: "901", Flag: "🇹🇼", Decimals: 2},
	TZS: {Code: "TZS", Name: "Tanzanian shilling", Number: "834", Flag: "🇹🇿", Decimals: 2},
	UAH: {Code: "UAH", Name: "Ukrainian hryvnia", Number: "980", Flag: "🇺🇦", Decimals: 2},
	UGX: {Code: "UGX", Name: "Ugandan shilling", Number: "800", Flag: "🇺🇬", Decimals: 2},
	USD: {Code: "USD", Name: "United States dollar", Number: "840", Flag: "🇺🇸", Decimals: 2},
	USN: {Code: "USN", Name: "United States dollar (next day) (funds code)", Number: "997", Flag: "", Decimals: 2},
	USS: {Code: "USS", Name: "United States dollar (same day) (funds code)", Number: "998", Flag: "", Decimals: 2},
	UYI: {Code: "UYI", Name: "Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)", Number: "940", Flag: "", Decimals: 0},
	UYU: {Code: "UYU", Name: "Uruguayan peso", Number: "858", Flag: "🇺🇾", Decimals: 2},
	UZS: {Code: "UZS", Name: "Uzbekistan som", Number: "860", Flag: "🇺🇿", Decimals: 2},
	VEF: {Code: "VEF", Name: "Venezuelan bolívar fuerte", Number: "937", Flag: "", Decimals: 2},
	VND: {Code: "VND", Name: "Vietnamese dong", Number: "704", Flag: "🇻🇳", Decimals: 0},
	VUV: {Code: "VUV", Name: "Vanuatu vatu", Number: "548", Flag: "🇻🇺", Decimals: 0},
	WST: {Code: "WST", Name: "Samoan tala", Number: "882", Flag: "🇼🇸", Decimals: 2},
	XAF: {Code: "XAF", Name: "CFA franc BEAC", Number: "950", Flag: "🇨🇲", Decimals: 0},
	XAG: {Code: "XAG", Name: "Silver (one troy ounce)", Number: "961", Flag: "", Decimals: 0},
	XAU: {Code: "XAU", Name: "Gold (one troy ounce)", Number: "959", Flag: "", Decimals: 0},
	XBA: {Code: "XBA", Name: "European Composite Unit (EURCO) (bond market unit)", Number: "955", Flag: "", Decimals: 0},
	XBB: {Code: "XBB", Name: "European Monetary Unit (E.M.U.-6) (bond market unit)", Number: "956", Flag: "", Decimals: 0},
	XBC: {Code: "XBC", Name: "European Unit of Account 9 (E.U.A.-9) (bond market unit)", Number: "957", Flag: "", Decimals: 0},
	XBD: {Code: "XBD", Name: "European Unit of Account 17 (E.U.A.-17) (bond market unit)", Number: "958", Flag: "", Decimals: 0},
	XCD: {Code: "XCD", Name: "East Caribbean dollar", Number: "951", Flag: "🇦🇮", Decimals: 2},
	XDR: {Code: "XDR", Name: "Special drawing rights", Number: "960", Flag: "", Decimals: 0},
	XFU: {Code: "XFU", Name: "UIC franc (special settlement currency)", Number: "Nil", Flag: "", Decimals: 0},
	XOF: {Code: "XOF", Name: "CFA franc BCEAO", Number: "952", Flag: "🇧🇯", Decimals: 0},
	XPD: {Code: "XPD", Name: "Palladium (one troy ounce)", Number: "964", Flag: "", Decimals: 0},
	XPF: {Code: "XPF", Name: "CFP franc", Number: "953", Flag: "🇵🇫", Decimals: 0},
	XPT: {Code: "XPT", Name: "Platinum (one troy ounce)", Number: "962", Flag: "", Decimals: 0},
	XTS: {Code: "XTS", Name: "Code reserved for testing purposes", Number: "963", Flag: "", Decimals: 0},
	XXX: {Code: "XXX", Name: "No currency", Number: "999", Flag: "", Decimals: 0},
	YER: {Code: "YER", Name: "Yemeni rial", Number: "886", Flag: "🇾🇪", Decimals: 2},
	ZAR: {Code: "ZAR", Name: "South African rand", Number: "710", Flag: "🇿🇦", Decimals: 2},
	ZMW: {Code: "ZMW", Name: "Zambian kwacha", Number: "967", Flag: "🇿🇲", Decimals: 2},
}

var stringToCurrencyCode = map[string]CurrencyCode{
	"AED": AED, "AFN": AFN, "ALL": ALL, "AMD": AMD, "ANG": ANG, "AOA": AOA, "ARS": ARS, "AUD": AUD, "AWG": AWG, "AZN": AZN,
	"BAM": BAM, "BBD": BBD, "BDT": BDT, "BGN": BGN, "BHD": BHD, "BIF": BIF, "BMD": BMD, "BND": BND, "BOB": BOB, "BOV": BOV,
	"BRL": BRL, "BSD": BSD, "BTN": BTN, "BWP": BWP, "BYR": BYR, "BZD": BZD, "CAD": CAD, "CDF": CDF, "CHE": CHE, "CHF": CHF,
	"CHW": CHW, "CLF": CLF, "CLP": CLP, "CNY": CNY, "COP": COP, "COU": COU, "CRC": CRC, "CUC": CUC, "CUP": CUP, "CVE": CVE,
	"CZK": CZK, "DJF": DJF, "DKK": DKK, "DOP": DOP, "DZD": DZD, "EGP": EGP, "ERN": ERN, "ETB": ETB, "EUR": EUR, "FJD": FJD,
	"FKP": FKP, "GBP": GBP, "GEL": GEL, "GHS": GHS, "GIP": GIP, "GMD": GMD, "GNF": GNF, "GTQ": GTQ, "GYD": GYD, "HKD": HKD,
	"HNL": HNL, "HRK": HRK, "HTG": HTG, "HUF": HUF, "IDR": IDR, "ILS": ILS, "INR": INR, "IQD": IQD, "IRR": IRR, "ISK": ISK,
	"JMD": JMD, "JOD": JOD, "JPY": JPY, "KES": KES, "KGS": KGS, "KHR": KHR, "KMF": KMF, "KPW": KPW, "KRW": KRW, "KWD": KWD,
	"KYD": KYD, "KZT": KZT, "LAK": LAK, "LBP": LBP, "LKR": LKR, "LRD": LRD, "LSL": LSL, "LTL": LTL, "LVL": LVL, "LYD": LYD,
	"MAD": MAD, "MDL": MDL, "MGA": MGA, "MKD": MKD, "MMK": MMK, "MNT": MNT, "MOP": MOP, "MRO": MRO, "MUR": MUR, "MVR": MVR,
	"MWK": MWK, "MXN": MXN, "MXV": MXV, "MYR": MYR, "MZN": MZN, "NAD": NAD, "NGN": NGN, "NIO": NIO, "NOK": NOK, "NPR": NPR,
	"NZD": NZD, "OMR": OMR, "PAB": PAB, "PEN": PEN, "PGK": PGK, "PHP": PHP, "PKR": PKR, "PLN": PLN, "PYG": PYG, "QAR": QAR,
	"RON": RON, "RSD": RSD, "RUB": RUB, "RWF": RWF, "SAR": SAR, "SBD": SBD, "SCR": SCR, "SDG": SDG, "SEK": SEK, "SGD": SGD,
	"SHP": SHP, "SLL": SLL, "SOS": SOS, "SRD": SRD, "SSP": SSP, "STD": STD, "SYP": SYP, "SZL": SZL, "THB": THB, "TJS": TJS,
	"TMT": TMT, "TND": TND, "TOP": TOP, "TRY": TRY, "TTD": TTD, "TWD": TWD, "TZS": TZS, "UAH": UAH, "UGX": UGX, "USD": USD,
	"USN": USN, "USS": USS, "UYI": UYI, "UYU": UYU, "UZS": UZS, "VEF": VEF, "VND": VND, "VUV": VUV, "WST": WST, "XAF": XAF,
	"XAG": XAG, "XAU": XAU, "XBA": XBA, "XBB": XBB, "XBC": XBC, "XBD": XBD, "XCD": XCD, "XDR": XDR, "XFU": XFU, "XOF": XOF,
	"XPD": XPD, "XPF": XPF, "XPT": XPT, "XTS": XTS, "XXX": XXX, "YER": YER, "ZAR": ZAR, "ZMW": ZMW,
}
