package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CostByCat struct {
	C string `json:"category"`
	A float64 `json:"amount"`
	P float64 `json:"percent"`
	Av float64 `json:"average"`
}

type CostData struct {
	C []CostByCat `json:"individual"`
	T float64 `json:"total"`
}

type Total struct {
	S float64 `json:"services"`
	P float64 `json:"products"`
	T float64 `json:"total"`
	A float64 `json:"average"`
}

func dbConn() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "postgresql://postgres:password@localhost:5433/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func takings() {
	var s Total

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	dateFrom := "2020-07-03"
	dateTo := "2020-10-03"

	db.Table("takings").Select("sum(services) as s, sum(products) as p").Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&s)

	s.T = s.P + s.S
	s.A = s.T / 3

	fmt.Println(s)
}

func cstsByCat() {
	var d CostData
	var c CostByCat
	var e []CostByCat

	db := dbConn()
	db.LogMode(true)
	defer db.Close()

	dateFrom := "2020-07-03"
	dateTo := "2020-10-03"

	categories := GetCategories()

	for cat, _ := range categories {
		db.Table("costs").Select("sum(debit) as a").Where("category = ?", cat).Where("date >= ? AND date <= ?", dateFrom, dateTo).Scan(&c)

		e = append(e, CostByCat{cat, c.A, 10, 10})
	}

	fmt.Println(d)

}

func GetCategories() (c map[string][]string) {
	c = map[string][]string{
		"Wages": {
			"MR J SHARP",
			"JAMES SHARPE",
			"MISS L HALL",
			"GEORGIA LUTTON",
			"LAUREN THOMPSON",
			"LAYLA RELF",
			"ABBI GREEN",
			"MISS V ROWLAND",
			"KELLY REEDY",
			"JOANNE MAHONEY",
			"BRADLEY RYAN",
			"ABIGAIL CLARKE",
			"DAVID RANDLES",
			"LUCY WATSON",
			"LAUREN WATSON",
			"RUBY JOHNSON",
			"SOPHIE YOUDS",
			"BETH BROWN",
			"HARRISON DOOLEY",
			"SARAH CARTWRIGHT",
			"KATE O HALLORAN",
			"DOROTA SOKOLOWSKI",
			"VICTORIA NYLAND",
			"LOUISE BAILEY",
			"LILLY SMITH",
		},
		"PAYE": {
			"HMRC CUSTOMS AND E",
			"FREDRICKSON",
			"ADVANTIS",
		},
		"Pension": {
			"NEST",
		},
		"Freelance": {
			"NATALIE SHARPE",
			"MATHEW LANE",
			"MATTHEW LANE",
			"AMY WOODS",
			"MICHELLE RAILTON",
			"LEON PRITCHARD",
		},
		"VAT": {
			"HMRC VAT",
		},
		"Utilities": {
			"BRIT GAS",
			"BRITISH TELECOM",
			"SCOTTISH POWER",
			"EDF ENERGY",
			"WATER PLUS",
			"EE LIMITED",
			"BT GROUP",
			"BG BUSINESS",
			"ASH WASTE",
			"CATHEDRAL LEASING",
			"WWW.BRITISHGAS.CO.",
			"PASTDUE",
			"SCOTTISHPOWER",
			"O2 DEVICE PLAN",
			"O2 05056477/001",
			"EE & T-MOBILE",
			"E.ON",
		},
		"Building": {
			"MEREHALL ESTATES",
			"BETTERBOOZE LTD",
			"WBC NNDR",
			"JENSON INVESTMENTS",
			"WARRINGTON BOROUGH",
			"WARRINGTON CD",
			"WARRINGTON B.C.",
			"W.B.C MV INTERNET",
			"JENSEN INVESTMENTS",
		},
		"Base": {
			"NJS MAINTENANCE",
			"M SUTTON",
			"MODERN LIGHTING",
			"M A SUTTON",
			"JACOB INTERIORS",
		},
		"Stock": {
			"BEAUTY WORKS",
			"BEAUTYWORKS",
			"ICON CONSULTANCY",
			"SWEET SQUARED",
			"ALAN HOWARD",
			"HENKEL",
			"WWW.ASTONANDFINCHE",
			"ALAN HOWARD(STOCKP CD 6954",
			"SALLY SALON",
			"WWW.GHDHAIR.COM",
			"FEEL FOR HAIR",
			"SALONS DIRECT",
			"GOCARDLESS",
			"FA UK LIMITED",
			"JAMELLA",
			"WWW.SALONSDIRECT.C",
			"BALMAINHAIR.CO.UK",
			"WWW.FEEFORHAIR.CO.",
			"JEMELLA LTD",
			"THE WIGGINS",
			"SIMPLYHAIR",
			"SALONEASY",
			"JEMELLA",
			"CLOUD9",
			"BEAUTY WORX",
			"AMERICAN CREW",
		},
		"Marketing": {
			"RACKSPACE",
			"GOOGLE",
			"TEXTANYWHERE",
			"BUFFER",
			"HEROKU",
			"FACEBK",
			"ADOBE",
			"JetBrains",
			"DIGITALOCEAN.COM",
			"FORGE.LARAVEL.COM",
			"LARACASTS",
			"NDEVOR",
			"COSCHEDULE.COM",
			"COSCHEDULE",
			"THREE BEST RATED",
			"VUEMASTERY.COM",
			"DNSIMPLE",
			"123 REG",
			"WWW.DISCOUNTDISPLA",
			"WINDOWFILMS",
			"THE PRINTING PEOPL",
			"SG MANUFACTURING",
			"POST OFFICE SELF",
			"PENTANGLE CARDS",
			"GRAFENIA",
			"Evernote",
			"DISCOUNT DISPLAYS",
			"CARTRIDGEPEOPLE.CO",
		},
		"Insurance": {
			"CLOSE-COVERSURE",
			"BAUER CONSUMER MED",
			"VLS RE KLARNA",
			"VLS RE CLOSE BROS",
			"GROVE-DEAN.CO.UK",
			"CURRYS  3267567149",
		},
		"Tax": {
			"HMRC NDDS",
			"HMRC - ACCOUNTS OF",
			"HMRC 600000000562153302",
		},
		"Staff": {
			"TRAINLINE",
			"D WRIGHT",
			"PARAGON",
			"Trainline.com",
			"Village Hotel Warr",
			"VIRGIN TRAINS",
			"THE CUMBERLAND",
			"SUITES HOTEL KNOWS",
			"SALON PUNK",
			"PAPA JONES PIZZA",
			"MR LAU'S",
			"LAS RAMBLAS",
			"Just Eat",
			"GREAT CUMBERLAND",
			"FRIAR PENKETH",
			"FIRE EVENTS",
			"DMN/DIRTYMARTINIMA",
			"Circo",
			"AIRBNB",
		},
		"Sundries": {
			"Spotify",
			"PPLPRS",
			"VIMTO OUT",
			"DLT MEDIA",
			"WWW.GOMPELS.CO.UK",
			"DJ DRINK SOLUTIONS",
			"VIKING",
			"TILLROLLSDIRECT.CO",
			"WWW.COSTCO.CO.UK",
			"PPL PRS",
			"POUNDLAND",
			"POUND SUPER STORE",
			"MM NEWSAGENTS",
			"MARTIN MCCOLL",
			"COSTCO",
		},
		"Paypal": {
			"PAYPAL",
		},
		"Amazon": {
			"AMZNMKTPLACE ",
			"AMZNMktplace",
			"Amazon",
			"AMZN",
			"AMZ*BC",
			"AMAZON",
		},
		"Loans": {
			"KENNET",
			"INVESTEC",
			"QUANTUM",
			"JOHN LAMB",
			"D A CARTER",
		},
		"Bank": {
			"NON-GBP TRANS FEE",
			"O/DRAFT INTEREST",
			"SERVICE CHARGES",
			"RETAIL MERCHANT SE",
			"GLOBAL PAYMENTS",
			"EMS",
			"UNAUTH'D BORR. FEE",
			"RETURNED D/D",
			"OVERDRAFT FEE",
			"NON-STG TRANS FEE",
		},
		"Drawings": {
			"ADAM CARTER",
			"MISS I LAMB",
			"NETFLIX.COM",
			"APPLE.COM/BILL",
			"ITUNES.COM/BILL",
			"Audible.co.uk",
			"STARBUCKS",
			"LNK WARRINGTON",
			"Prime",
			"VISION DIRECT",
			"Kindle",
			"Amazon Prime*MB8HF",
			"WWW.ASOS.COM",
			"WWW.TWOSEASONS.CO.",
			"Audible UK",
			"WWW.MISSGUIDED.CO.",
			"VOLCOM SAS",
			"TOPPIK.CO.UK",
			"TICKETMASTER",
			"PropellerheadsCOM",
			"MCDONALDS",
			"GRUMPY MULE",
			"FGT*TICKETMASTER",
			"Etsy.com",
			"EASYJET",
			"DANIEL HANCOCK",
			"CHATURBIL",
			"APPLE.COM/UK",
			"ACOUSTIC CAFE",
		},
	}
	return
}

func main() {
	cstsByCat()
}
