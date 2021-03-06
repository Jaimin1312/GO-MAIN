package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DataStruct struct {
	Vehicle struct {
		Age                           string        `json:"Age"`
		BodyType                      string        `json:"BodyType"`
		Color                         string        `json:"Color"`
		DateFirstRegistered           string        `json:"DateFirstRegistered"`
		DefaultPercentOfMarket        string        `json:"DefaultPercentOfMarket"`
		DrivetrainType                string        `json:"DrivetrainType"`
		Engine                        string        `json:"Engine"`
		Transmission                  string        `json:"Transmission"`
		FLIP                          int           `json:"FLIP"`
		GotFirstRegFromServiceArizona bool          `json:"GotFirstRegFromServiceArizona"`
		HasBuybackForm                int           `json:"HasBuybackForm"`
		ID                            string        `json:"Id"`
		InteriorColor                 string        `json:"InteriorColor"`
		InventoryDate                 string        `json:"InventoryDate"`
		KbbLending                    int           `json:"KbbLending"`
		KbbLendingFormatted           string        `json:"KbbLendingFormatted"`
		KbbRetail                     int           `json:"KbbRetail"`
		KbbRetailFormatted            string        `json:"KbbRetailFormatted"`
		Labels                        []string      `json:"Labels"`
		LastDollarChange              string        `json:"LastDollarChange"`
		MPGCity                       string        `json:"MPGCity"`
		MPGHwy                        string        `json:"MPGHwy"`
		Make                          string        `json:"Make"`
		ManuallyAddedLabels           []interface{} `json:"ManuallyAddedLabels"`
		Mileage                       string        `json:"Mileage"`
		MktAvgPrice                   string        `json:"MktAvgPrice"`
		Model                         string        `json:"Model"`
		OriginalPrice                 string        `json:"OriginalPrice"`
		Checkout                      string        `json:"Checkout"`
		PhotoAlbum                    []string      `json:"PhotoAlbum"`
		PhotoSmall                    string        `json:"PhotoSmall"`
		VehicleImage                  string        `json:"VehicleImage"`
		PlateDuration                 string        `json:"PlateDuration"`
		PlateDuration2                string        `json:"PlateDuration2"`
		PlateFee                      string        `json:"PlateFee"`
		PlateFee2                     string        `json:"PlateFee2"`
		Price                         int           `json:"Price"`
		SearchText                    string        `json:"SearchText"`
		Sold                          bool          `json:"Sold"`
		Status                        string        `json:"Status"`
		StatusAuthor                  string        `json:"StatusAuthor"`
		StatusShow                    bool          `json:"StatusShow"`
		Trim                          string        `json:"Trim"`
		URL                           string        `json:"Url"`
		Vin                           string        `json:"Vin"`
		Year                          int           `json:"Year"`
		Notes                         interface{}   `json:"Notes"`
		TitleNumber                   string        `json:"TitleNumber"`
		Cost                          string        `json:"Cost"`
		PurchasedDate                 string        `json:"PurchasedDate"`
		WherePurchased                string        `json:"WherePurchased"`
		BoughtUnder                   string        `json:"BoughtUnder"`
		BoughtFrom                    string        `json:"BoughtFrom"`
		AddedFrom                     string        `json:"AddedFrom"`
		CreatedDate                   string        `json:"CreatedDate"`
		IsFullyElectric               bool          `json:"IsFullyElectric"`
		Stock                         string        `json:"Stock"`
		Cylinders                     string        `json:"Cylinders"`
	} `json:"Vehicle"`
	PriceTable struct {
		DeltaPrice                    int     `json:"DeltaPrice"`
		Price                         int     `json:"Price"`
		DocFee                        int     `json:"DocFee"`
		MVD                           float64 `json:"MVD"`
		Plates                        float64 `json:"Plates"`
		PlateLessCredit               float64 `json:"PlateLessCredit"`
		PlateCredit                   int     `json:"PlateCredit"`
		FactoryPrice                  int     `json:"FactoryPrice"`
		PlatesExp                     int     `json:"PlatesExp"`
		TitleReg                      float64 `json:"TitleReg"`
		TaxRate                       float64 `json:"TaxRate"`
		CalcTax                       float64 `json:"CalcTax"`
		TradeIn                       int     `json:"TradeIn"`
		Lease                         string  `json:"Lease"`
		Payoff                        int     `json:"Payoff"`
		Shipping                      int     `json:"Shipping"`
		Warranty                      int     `json:"Warranty"`
		Gap                           int     `json:"Gap"`
		Paint                         int     `json:"Paint"`
		TnW                           int     `json:"TnW"`
		CustomInput                   string  `json:"CustomInput"`
		CustomInputValue              int     `json:"CustomInputValue"`
		Total                         int     `json:"Total"`
		DownPayment                   int     `json:"DownPayment"`
		Financed                      int     `json:"Financed"`
		InterestRate                  float64 `json:"InterestRate"`
		Months                        int     `json:"Months"`
		Payment                       float64 `json:"Payment"`
		ContractDate                  string  `json:"ContractDate"`
		DaysTillPayment               int     `json:"DaysTillPayment"`
		FirstPaymentDue               string  `json:"FirstPaymentDue"`
		TotalInterestPaid             float64 `json:"TotalInterestPaid"`
		TotalOfPayments               float64 `json:"TotalOfPayments"`
		Notes                         string  `json:"Notes"`
		LtvKbbRetail                  float64 `json:"LtvKbbRetail"`
		LtvKbbLending                 float64 `json:"LtvKbbLending"`
		TaxRateState                  int     `json:"TaxRateState"`
		DeltaWarranty                 int     `json:"DeltaWarranty"`
		DeltaGap                      int     `json:"DeltaGap"`
		DeltaTnW                      int     `json:"DeltaTnW"`
		DeltaPaint                    int     `json:"DeltaPaint"`
		ShowTradeIn                   bool    `json:"ShowTradeIn"`
		ShowGap                       bool    `json:"ShowGap"`
		ShowTnW                       bool    `json:"ShowTnW"`
		ShowShipping                  bool    `json:"ShowShipping"`
		ShowWarranty                  bool    `json:"ShowWarranty"`
		ShowCustomInput               bool    `json:"ShowCustomInput"`
		ShowDownPayment               bool    `json:"ShowDownPayment"`
		ShowPaymentBreakdown          bool    `json:"ShowPaymentBreakdown"`
		ShowNote                      bool    `json:"ShowNote"`
		ShowMVDBreakdown              bool    `json:"ShowMVDBreakdown"`
		ShowPaint                     bool    `json:"ShowPaint"`
		ShowDiff                      bool    `json:"ShowDiff"`
		ShowPlateCredit               bool    `json:"ShowPlateCredit"`
		TitleRegSelected              string  `json:"TitleRegSelected"`
		PaymentDeltaAppliedTo         string  `json:"PaymentDeltaAppliedTo"`
		RollPaymentDifferenceDown     int     `json:"RollPaymentDifferenceDown"`
		RollPaymentDifferencePrice    int     `json:"RollPaymentDifferencePrice"`
		RollPaymentDifferenceTrade    int     `json:"RollPaymentDifferenceTrade"`
		RollPaymentDifferenceWarranty int     `json:"RollPaymentDifferenceWarranty"`
		IsFullyElectric               bool    `json:"IsFullyElectric"`
	} `json:"PriceTable"`
	PriceTablePrevVal struct {
		DeltaPrice                    int       `json:"DeltaPrice"`
		Price                         int       `json:"Price"`
		DocFee                        int       `json:"DocFee"`
		MVD                           int       `json:"MVD"`
		Plates                        int       `json:"Plates"`
		PlateLessCredit               int       `json:"PlateLessCredit"`
		PlateCredit                   int       `json:"PlateCredit"`
		FactoryPrice                  int       `json:"FactoryPrice"`
		PlatesExp                     int       `json:"PlatesExp"`
		TitleReg                      float64   `json:"TitleReg"`
		TaxRate                       float64   `json:"TaxRate"`
		CalcTax                       int       `json:"CalcTax"`
		TradeIn                       int       `json:"TradeIn"`
		Lease                         string    `json:"Lease"`
		Payoff                        int       `json:"Payoff"`
		Shipping                      int       `json:"Shipping"`
		Warranty                      int       `json:"Warranty"`
		Gap                           int       `json:"Gap"`
		Paint                         int       `json:"Paint"`
		TnW                           int       `json:"TnW"`
		CustomInput                   string    `json:"CustomInput"`
		CustomInputValue              int       `json:"CustomInputValue"`
		Total                         int       `json:"Total"`
		DownPayment                   int       `json:"DownPayment"`
		Financed                      int       `json:"Financed"`
		InterestRate                  float64   `json:"InterestRate"`
		Months                        int       `json:"Months"`
		Payment                       int       `json:"Payment"`
		ContractDate                  string    `json:"ContractDate"`
		DaysTillPayment               int       `json:"DaysTillPayment"`
		FirstPaymentDue               time.Time `json:"FirstPaymentDue"`
		TotalInterestPaid             int       `json:"TotalInterestPaid"`
		TotalOfPayments               int       `json:"TotalOfPayments"`
		Notes                         string    `json:"Notes"`
		LtvKbbRetail                  int       `json:"LtvKbbRetail"`
		LtvKbbLending                 int       `json:"LtvKbbLending"`
		TaxRateState                  int       `json:"TaxRateState"`
		DeltaWarranty                 int       `json:"DeltaWarranty"`
		DeltaGap                      int       `json:"DeltaGap"`
		DeltaTnW                      int       `json:"DeltaTnW"`
		DeltaPaint                    int       `json:"DeltaPaint"`
		ShowTradeIn                   bool      `json:"ShowTradeIn"`
		ShowGap                       bool      `json:"ShowGap"`
		ShowTnW                       bool      `json:"ShowTnW"`
		ShowShipping                  bool      `json:"ShowShipping"`
		ShowWarranty                  bool      `json:"ShowWarranty"`
		ShowCustomInput               bool      `json:"ShowCustomInput"`
		ShowDownPayment               bool      `json:"ShowDownPayment"`
		ShowPaymentBreakdown          bool      `json:"ShowPaymentBreakdown"`
		ShowNote                      bool      `json:"ShowNote"`
		ShowMVDBreakdown              bool      `json:"ShowMVDBreakdown"`
		ShowPaint                     bool      `json:"ShowPaint"`
		ShowDiff                      bool      `json:"ShowDiff"`
		ShowPlateCredit               bool      `json:"ShowPlateCredit"`
		TitleRegSelected              string    `json:"TitleRegSelected"`
		PaymentDeltaAppliedTo         string    `json:"PaymentDeltaAppliedTo"`
		RollPaymentDifferenceDown     int       `json:"RollPaymentDifferenceDown"`
		RollPaymentDifferencePrice    int       `json:"RollPaymentDifferencePrice"`
		RollPaymentDifferenceTrade    int       `json:"RollPaymentDifferenceTrade"`
		RollPaymentDifferenceWarranty int       `json:"RollPaymentDifferenceWarranty"`
		IsFullyElectric               bool      `json:"IsFullyElectric"`
	} `json:"PriceTablePrevVal"`
	PurchaseForm struct {
		BuyerID                       string `json:"BuyerId"`
		BuyerName                     string `json:"BuyerName"`
		BuyerNameLF                   string `json:"BuyerNameLF"`
		BuyerDob                      string `json:"BuyerDob"`
		BuyerDl                       string `json:"BuyerDl"`
		BuyerDlIssueDate              string `json:"BuyerDlIssueDate"`
		BuyerDlExpireDate             string `json:"BuyerDlExpireDate"`
		BuyerSocial                   string `json:"BuyerSocial"`
		BuyerMaritalStatus            string `json:"BuyerMaritalStatus"`
		BuyerDLImage                  string `json:"BuyerDLImage"`
		BuyerFaceImage                string `json:"BuyerFaceImage"`
		BuyerPersonal                 string `json:"BuyerPersonal"`
		MailingAddressBuyer           string `json:"MailingAddressBuyer"`
		MailingCityStateZipBuyer      string `json:"MailingCityStateZipBuyer"`
		MailingCityBuyer              string `json:"MailingCityBuyer"`
		MailingStateBuyer             string `json:"MailingStateBuyer"`
		MailingZipBuyer               int    `json:"MailingZipBuyer"`
		BuyerDlStreetAdd              string `json:"BuyerDlStreetAdd"`
		BuyerDlCity                   string `json:"BuyerDlCity"`
		BuyerDlState                  string `json:"BuyerDlState"`
		BuyerDlZip                    int    `json:"BuyerDlZip"`
		BuyerDlCityStateZip           string `json:"BuyerDlCityStateZip"`
		BuyerStreetAdd                string `json:"BuyerStreetAdd"`
		BuyerCity                     string `json:"BuyerCity"`
		BuyerState                    string `json:"BuyerState"`
		BuyerZip                      int    `json:"BuyerZip"`
		BuyerCityStateZip             string `json:"BuyerCityStateZip"`
		BuyerMonthlyHousingExpense    string `json:"BuyerMonthlyHousingExpense"`
		BuyerRentOrMortgageOrNone     string `json:"BuyerRentOrMortgageOrNone"`
		BuyerTimeAtAddress            string `json:"BuyerTimeAtAddress"`
		BuyerYearsatAddress           int    `json:"BuyerYearsatAddress"`
		BuyerMonthsatAddress          int    `json:"BuyerMonthsatAddress"`
		BuyerPrevAddStreetAdd         string `json:"BuyerPrevAddStreetAdd"`
		BuyerPrevAddCity              string `json:"BuyerPrevAddCity"`
		BuyerPrevAddState             string `json:"BuyerPrevAddState"`
		BuyerPrevAddZip               int    `json:"BuyerPrevAddZip"`
		BuyerPrevAddCityStateZip      string `json:"BuyerPrevAddCityStateZip"`
		BuyerPrevAddTimeAtAddress     string `json:"BuyerPrevAddTimeAtAddress"`
		BuyerPrevAddYearsatAddress    string `json:"BuyerPrevAddYearsatAddress"`
		BuyerPrevAddMonthsatAddress   string `json:"BuyerPrevAddMonthsatAddress"`
		BuyerCell                     string `json:"BuyerCell"`
		BuyerWork                     string `json:"BuyerWork"`
		BuyerHome                     string `json:"BuyerHome"`
		BuyerEmail                    string `json:"BuyerEmail"`
		BuyerJobEmployer              string `json:"BuyerJobEmployer"`
		BuyerJobTitle                 string `json:"BuyerJobTitle"`
		BuyerJobIncome                string `json:"BuyerJobIncome"`
		BuyerJobIncomeFrequency       string `json:"BuyerJobIncomeFrequency"`
		BuyerJobEmployementStatus     string `json:"BuyerJobEmployementStatus"`
		BuyerJobTimeAtJob             string `json:"BuyerJobTimeAtJob"`
		BuyerJobSourceOfOtherIncome   string `json:"BuyerJobSourceOfOtherIncome"`
		BuyerJobOtherIncomeMonthly    string `json:"BuyerJobOtherIncomeMonthly"`
		BuyerJobOtherIncomeTitle      string `json:"BuyerJobOtherIncomeTitle"`
		BuyerJobOtherIncomeTime       string `json:"BuyerJobOtherIncomeTime"`
		BuyerJobOtherIncomePhone      string `json:"BuyerJobOtherIncomePhone"`
		BuyerPrevJobEmployer          string `json:"BuyerPrevJobEmployer"`
		BuyerPrevJobTitle             string `json:"BuyerPrevJobTitle"`
		BuyerPrevJobTimeAtJob         string `json:"BuyerPrevJobTimeAtJob"`
		VehicleYear                   int    `json:"VehicleYear"`
		VehicleMake                   string `json:"VehicleMake"`
		VehicleModel                  string `json:"VehicleModel"`
		VehicleYearMakeModel          string `json:"VehicleYearMakeModel"`
		VehicleTrim                   string `json:"VehicleTrim"`
		VehicleVin                    string `json:"VehicleVin"`
		VehicleMiles                  string `json:"VehicleMiles"`
		VehicleDrivetrainType         string `json:"VehicleDrivetrainType"`
		VehicleEngine                 string `json:"VehicleEngine"`
		VehicleAge                    string `json:"VehicleAge"`
		VehicleInventoryDate          string `json:"VehicleInventoryDate"`
		VehiclePercentOfMarket        string `json:"VehiclePercentOfMarket"`
		VehicleMktAvgPrice            string `json:"VehicleMktAvgPrice"`
		VehicleLastDollarChange       string `json:"VehicleLastDollarChange"`
		VehicleColor                  string `json:"VehicleColor"`
		VehicleInteriorColor          string `json:"VehicleInteriorColor"`
		VehicleBodyType               string `json:"VehicleBodyType"`
		VehicleManufWarranty          string `json:"VehicleManufWarranty"`
		VehicleTitleNumber            string `json:"VehicleTitleNumber"`
		VehicleSource                 string `json:"VehicleSource"`
		VehicleCost                   string `json:"VehicleCost"`
		VehicleKbbRetail              string `json:"VehicleKbbRetail"`
		VehicleKbbLending             string `json:"VehicleKbbLending"`
		VehicleKbbWholeSale           string `json:"VehicleKbbWholeSale"`
		VehicleCylinders              string `json:"VehicleCylinders"`
		VehicleServiceContract        string `json:"VehicleServiceContract"`
		VehicleOdometerStatus         string `json:"VehicleOdometerStatus"`
		DealershipInfoNumber          string `json:"DealershipInfoNumber"`
		DealershipInfoEIN             string `json:"DealershipInfoEIN"`
		DealershipInfoName            string `json:"DealershipInfoName"`
		DealershipInfoAddress         string `json:"DealershipInfoAddress"`
		DealershipInfoCityStateZip    string `json:"DealershipInfoCityStateZip"`
		DealershipInfoCity            string `json:"DealershipInfoCity"`
		DealershipInfoState           string `json:"DealershipInfoState"`
		DealershipInfoZip             int    `json:"DealershipInfoZip"`
		DealershipInfoPhone           string `json:"DealershipInfoPhone"`
		DealershipInfoEmail           string `json:"DealershipInfoEmail"`
		TradeYear                     int    `json:"TradeYear"`
		TradeMake                     string `json:"TradeMake"`
		TradeModel                    string `json:"TradeModel"`
		TradeYearMakeModel            string `json:"TradeYearMakeModel"`
		TradeTrim                     string `json:"TradeTrim"`
		TradeVIN                      string `json:"TradeVIN"`
		TradeMiles                    string `json:"TradeMiles"`
		TradeColor                    string `json:"TradeColor"`
		TradeInteriorColor            string `json:"TradeInteriorColor"`
		TradeBodyType                 string `json:"TradeBodyType"`
		TradeTitleNumber              string `json:"TradeTitleNumber"`
		TradeTransferPlate            string `json:"TradeTransferPlate"`
		TradeNameOnTitle              string `json:"TradeNameOnTitle"`
		TradeCylinders                int    `json:"TradeCylinders"`
		TradeLienHolder               string `json:"TradeLienHolder"`
		TradeLienHolderAbbreviation   string `json:"TradeLienHolderAbbreviation"`
		TradeLienHolderContact        string `json:"TradeLienHolderContact"`
		TradeLienHolderAddress        string `json:"TradeLienHolderAddress"`
		TradeLienHolderCityStateZip   string `json:"TradeLienHolderCityStateZip"`
		TradeLienHolderCity           string `json:"TradeLienHolderCity"`
		TradeLienHolderState          string `json:"TradeLienHolderState"`
		TradeLienHolderZip            int    `json:"TradeLienHolderZip"`
		TradeOwner                    string `json:"TradeOwner"`
		TradeOwnerDob                 string `json:"TradeOwnerDob"`
		TradeOwnerDl                  string `json:"TradeOwnerDl"`
		TradeOwnerAddress             string `json:"TradeOwnerAddress"`
		TradeOwnerCityStateZip        string `json:"TradeOwnerCityStateZip"`
		TradeOwnerCity                string `json:"TradeOwnerCity"`
		TradeOwnerState               string `json:"TradeOwnerState"`
		TradeOwnerZip                 int    `json:"TradeOwnerZip"`
		SalesInfoDate                 string `json:"SalesInfoDate"`
		SalesInfoDateOfSale           int    `json:"SalesInfoDateOfSale"`
		SalesInfoMonthOfSale          string `json:"SalesInfoMonthOfSale"`
		SalesInfoYearOfSale           int    `json:"SalesInfoYearOfSale"`
		SalesInfoSalesPerson          string `json:"SalesInfoSalesPerson"`
		SalesInfoEmail                string `json:"SalesInfoEmail"`
		SalesInfoWarrantyCompany      string `json:"SalesInfoWarrantyCompany"`
		SalesInfoWarrantyDeductible   string `json:"SalesInfoWarrantyDeductible"`
		SalesInfoWarrantyTerm         string `json:"SalesInfoWarrantyTerm"`
		SalesInfoWarrantyMiles        string `json:"SalesInfoWarrantyMiles"`
		SalesInfoGapCompany           string `json:"SalesInfoGapCompany"`
		SalesInfoTnWCompany           string `json:"SalesInfoTnWCompany"`
		SalesInfoPaintCompany         string `json:"SalesInfoPaintCompany"`
		SalesInfoAfterSaleCompany     string `json:"SalesInfoAfterSaleCompany"`
		SalesInfoNotes                string `json:"SalesInfoNotes"`
		SalesInfoSalesPersonJobTitle  string `json:"SalesInfoSalesPersonJobTitle"`
		LienHolderCity                string `json:"LienHolderCity"`
		LienHolderContact             string `json:"LienHolderContact"`
		LienHolderEIN                 string `json:"LienHolderEIN"`
		LienHolder                    string `json:"LienHolder"`
		LienHolderAddress             string `json:"LienHolderAddress"`
		LienHolderCityStateZip        string `json:"LienHolderCityStateZip"`
		LienHolderState               string `json:"LienHolderState"`
		LienHolderZip                 int    `json:"LienHolderZip"`
		CoBuyerID                     string `json:"CoBuyerId"`
		CoBuyerName                   string `json:"CoBuyerName"`
		CoBuyerNameLF                 string `json:"CoBuyerNameLF"`
		CoBuyerDob                    string `json:"CoBuyerDob"`
		CoBuyerDl                     string `json:"CoBuyerDl"`
		CoBuyerDlIssueDate            string `json:"CoBuyerDlIssueDate"`
		CoBuyerDlExpireDate           string `json:"CoBuyerDlExpireDate"`
		CoBuyerSocial                 string `json:"CoBuyerSocial"`
		CoBuyerMaritalStatus          string `json:"CoBuyerMaritalStatus"`
		CoBuyerRelationToBuyer        string `json:"CoBuyerRelationToBuyer"`
		CoBuyerDLImage                string `json:"CoBuyerDLImage"`
		CoBuyerFaceImage              string `json:"CoBuyerFaceImage"`
		CoBuyerPersonal               string `json:"CoBuyerPersonal"`
		CoBuyerDlStreetAdd            string `json:"CoBuyerDlStreetAdd"`
		CoBuyerDlCity                 string `json:"CoBuyerDlCity"`
		CoBuyerDlState                string `json:"CoBuyerDlState"`
		CoBuyerDlZip                  int    `json:"CoBuyerDlZip"`
		CoBuyerDlCityStateZip         string `json:"CoBuyerDlCityStateZip"`
		CoBuyerStreetAdd              string `json:"CoBuyerStreetAdd"`
		CoBuyerCity                   string `json:"CoBuyerCity"`
		CoBuyerState                  string `json:"CoBuyerState"`
		CoBuyerZip                    int    `json:"CoBuyerZip"`
		CoBuyerCityStateZip           string `json:"CoBuyerCityStateZip"`
		CoBuyerMonthlyHousingExpense  string `json:"CoBuyerMonthlyHousingExpense"`
		CoBuyerRentOrMortgageOrNone   string `json:"CoBuyerRentOrMortgageOrNone"`
		CoBuyerTimeAtAddress          string `json:"CoBuyerTimeAtAddress"`
		CoBuyerYearsatAddress         int    `json:"CoBuyerYearsatAddress"`
		CoBuyerMonthsatAddress        int    `json:"CoBuyerMonthsatAddress"`
		MailingAddressCoBuyer         string `json:"MailingAddressCoBuyer"`
		MailingCityStateZipCoBuyer    string `json:"MailingCityStateZipCoBuyer"`
		MailingCityCoBuyer            string `json:"MailingCityCoBuyer"`
		MailingStateCoBuyer           string `json:"MailingStateCoBuyer"`
		MailingZipCoBuyer             int    `json:"MailingZipCoBuyer"`
		CoBuyerPrevAddStreetAdd       string `json:"CoBuyerPrevAddStreetAdd"`
		CoBuyerPrevAddCity            string `json:"CoBuyerPrevAddCity"`
		CoBuyerPrevAddState           string `json:"CoBuyerPrevAddState"`
		CoBuyerPrevAddZip             int    `json:"CoBuyerPrevAddZip"`
		CoBuyerPrevAddCityStateZip    string `json:"CoBuyerPrevAddCityStateZip"`
		CoBuyerPrevAddTimeAtAddress   string `json:"CoBuyerPrevAddTimeAtAddress"`
		CoBuyerPrevAddYearsatAddress  string `json:"CoBuyerPrevAddYearsatAddress"`
		CoBuyerPrevAddMonthsatAddress string `json:"CoBuyerPrevAddMonthsatAddress"`
		CoBuyerCell                   string `json:"CoBuyerCell"`
		CoBuyerWork                   string `json:"CoBuyerWork"`
		CoBuyerHome                   string `json:"CoBuyerHome"`
		CoBuyerEmail                  string `json:"CoBuyerEmail"`
		CoBuyerJobEmployer            string `json:"CoBuyerJobEmployer"`
		CoBuyerJobTitle               string `json:"CoBuyerJobTitle"`
		CoBuyerJobIncome              string `json:"CoBuyerJobIncome"`
		CoBuyerJobIncomeFrequency     string `json:"CoBuyerJobIncomeFrequency"`
		CoBuyerJobEmployementStatus   string `json:"CoBuyerJobEmployementStatus"`
		CoBuyerJobTimeAtJob           string `json:"CoBuyerJobTimeAtJob"`
		CoBuyerJobSourceOfOtherIncome string `json:"CoBuyerJobSourceOfOtherIncome"`
		CoBuyerJobOtherIncomeMonthly  string `json:"CoBuyerJobOtherIncomeMonthly"`
		CoBuyerJobOtherIncomeTitle    string `json:"CoBuyerJobOtherIncomeTitle"`
		CoBuyerJobOtherIncomeTime     string `json:"CoBuyerJobOtherIncomeTime"`
		CoBuyerJobOtherIncomePhone    string `json:"CoBuyerJobOtherIncomePhone"`
		CoBuyerPrevJobEmployer        string `json:"CoBuyerPrevJobEmployer"`
		CoBuyerPrevJobTitle           string `json:"CoBuyerPrevJobTitle"`
		CoBuyerPrevJobTimeAtJob       string `json:"CoBuyerPrevJobTimeAtJob"`
	} `json:"PurchaseForm"`
	AulWarranty struct {
		Session                 string `json:"Session"`
		Pricer                  string `json:"Pricer"`
		DealerCost              int    `json:"DealerCost"`
		TermFilterApplied       int    `json:"TermFilterApplied"`
		MilesFilterApplied      int    `json:"MilesFilterApplied"`
		DeductibleFilterApplied string `json:"DeductibleFilterApplied"`
		Deductible              int    `json:"Deductible"`
		DisappearingDeductible  bool   `json:"DisappearingDeductible"`
		WarrantyNotAvailable    bool   `json:"WarrantyNotAvailable"`
		Coverage                string `json:"Coverage"`
		ProgramText             string `json:"ProgramText"`
	} `json:"AulWarranty"`
	CnaWarranty struct {
		TermFilterApplied       int    `json:"TermFilterApplied"`
		MilesFilterApplied      int    `json:"MilesFilterApplied"`
		DeductibleFilterApplied string `json:"DeductibleFilterApplied"`
		RateGUID                string `json:"RateGUID"`
		RateID                  string `json:"RateID"`
		ProductID               string `json:"ProductID"`
		DealerCost              int    `json:"DealerCost"`
	} `json:"CnaWarranty"`
	CnaGap struct {
		RateGUID   string `json:"RateGUID"`
		RateID     string `json:"RateID"`
		ProductID  string `json:"ProductID"`
		DealerCost int    `json:"DealerCost"`
	} `json:"CnaGap"`
	CnaTnW struct {
		TermFilterApplied       int    `json:"TermFilterApplied"`
		MilesFilterApplied      int    `json:"MilesFilterApplied"`
		DeductibleFilterApplied string `json:"DeductibleFilterApplied"`
		RateGUID                string `json:"RateGUID"`
		RateID                  string `json:"RateID"`
		ProductID               string `json:"ProductID"`
		DealerCost              int    `json:"DealerCost"`
	} `json:"CnaTnW"`
	Paint struct {
		TermFilterApplied int    `json:"TermFilterApplied"`
		DealerCost        int    `json:"DealerCost"`
		ContractNumber    string `json:"ContractNumber"`
	} `json:"Paint"`
	ProductLists []struct {
		Products []struct {
			AUL struct {
				Coverage               string `json:"Coverage"`
				DealerCost             int    `json:"DealerCost"`
				Deductible             int    `json:"Deductible"`
				DisappearingDeductible bool   `json:"DisappearingDeductible"`
				PlanDeductible         string `json:"PlanDeductible"`
				PlanMiles              int    `json:"PlanMiles"`
				PlanTerm               int    `json:"PlanTerm"`
				Pricer                 string `json:"Pricer"`
				ProgramText            string `json:"ProgramText"`
				Rate                   int    `json:"Rate"`
				Session                string `json:"Session"`
			} `json:"AUL,omitempty"`
			CNA struct {
				DealerCost     int    `json:"DealerCost"`
				PlanDeductible string `json:"PlanDeductible"`
				PlanMiles      int    `json:"PlanMiles"`
				PlanTerm       int    `json:"PlanTerm"`
				ProgramID      string `json:"ProgramId"`
				Rate           int    `json:"Rate"`
				RateID         string `json:"RateId"`
			} `json:"CNA,omitempty"`
			DeltaPrice           int    `json:"DeltaPrice"`
			FormattedDeltaPrice  string `json:"FormattedDeltaPrice"`
			FormattedOfferedRate string `json:"FormattedOfferedRate"`
			ID                   string `json:"Id"`
			OfferedRate          int    `json:"OfferedRate"`
			ProductType          string `json:"ProductType,omitempty"`
			Rate                 int    `json:"Rate"`
			ShowPrice            bool   `json:"ShowPrice"`
			DealerCost           int    `json:"DealerCost,omitempty"`
			PlanTerm             int    `json:"PlanTerm,omitempty"`
			ProgramID            string `json:"ProgramId,omitempty"`
			RateID               string `json:"RateId,omitempty"`
		} `json:"Products"`
		ProductName           string `json:"ProductName"`
		ProductColor          string `json:"ProductColor"`
		SelectedPaymentOption int    `json:"SelectedPaymentOption"`
		PaymentOptions        []struct {
			SelectedProductMonth    int  `json:"SelectedProductMonth"`
			SelectedProductRate     int  `json:"SelectedProductRate"`
			FinancedWithNewProducts int  `json:"FinancedWithNewProducts"`
			ProductTotal            int  `json:"ProductTotal"`
			ProductPayment          int  `json:"ProductPayment"`
			ShowPaymentDropdown     bool `json:"ShowPaymentDropdown"`
			ShowPriceOptionDropdown bool `json:"ShowPriceOptionDropdown"`
			GapObj                  struct {
				PlanTerm    int    `json:"PlanTerm"`
				OfferedRate int    `json:"OfferedRate"`
				DealerCost  int    `json:"DealerCost"`
				DeltaPrice  int    `json:"DeltaPrice"`
				Rate        int    `json:"Rate"`
				RateID      string `json:"RateId"`
				ProgramID   string `json:"ProgramId"`
			} `json:"GapObj"`
			MonthsAndRates []struct {
				Months        string `json:"Months"`
				Rate          int    `json:"Rate"`
				FormattedRate string `json:"FormattedRate"`
			} `json:"MonthsAndRates"`
		} `json:"PaymentOptions"`
		ShowDropDownLists    []string `json:"ShowDropDownLists"`
		AddProductDropdown   bool     `json:"AddProductDropdown"`
		ShowWarrantyDropdown bool     `json:"ShowWarrantyDropdown"`
		ShowPaintDropdown    bool     `json:"ShowPaintDropdown"`
	} `json:"ProductLists"`
	ID                      string        `json:"Id"`
	Created                 string        `json:"Created"`
	Seller                  string        `json:"Seller"`
	Sold                    bool          `json:"Sold"`
	DealPrinted             bool          `json:"DealPrinted"`
	IsCna                   bool          `json:"IsCna"`
	SignedAndPrintedDate    string        `json:"SignedAndPrintedDate"`
	PaymentOptions          string        `json:"PaymentOptions"`
	LeadSource              string        `json:"LeadSource"`
	TradeInStatusOptions    string        `json:"TradeInStatusOptions"`
	SelectedForms           []string      `json:"SelectedForms"`
	CustomizeSelectedForms  []string      `json:"CustomizeSelectedForms"`
	PrintPdfSelectedForms   []interface{} `json:"PrintPdfSelectedForms"`
	EmailEsignSelectedForms []string      `json:"EmailEsignSelectedForms"`
	Notes                   []struct {
		Note            string `json:"Note"`
		SalesPersonName string `json:"SalesPersonName"`
		Timestamp       string `json:"Timestamp"`
	} `json:"Notes"`
	SignedDealPdf          string        `json:"SignedDealPdf"`
	BlankFormsPdf          string        `json:"BlankFormsPdf"`
	BuyerCopiesPdf         string        `json:"BuyerCopiesPdf"`
	LenderCopiesPdf        string        `json:"LenderCopiesPdf"`
	DealerCopiesPdf        string        `json:"DealerCopiesPdf"`
	OldSignedDealPdf       []string      `json:"OldSignedDealPdf"`
	OldBuyerCopiesPdf      []string      `json:"OldBuyerCopiesPdf"`
	OldLenderCopiesPdf     []string      `json:"OldLenderCopiesPdf"`
	OldDealerCopiesPdf     []string      `json:"OldDealerCopiesPdf"`
	OldEmailPdfCopy        []interface{} `json:"OldEmailPdfCopy"`
	OldPrintPdfCopy        []interface{} `json:"OldPrintPdfCopy"`
	UploadedPdfs           []interface{} `json:"UploadedPdfs"`
	CarfaxPdf              string        `json:"CarfaxPdf"`
	TRPPdf                 string        `json:"TRPPdf"`
	NintyDayPdf            string        `json:"NintyDayPdf"`
	NintyDayAffedavite     string        `json:"NintyDayAffedavite"`
	WarrantyPdf            []interface{} `json:"WarrantyPdf"`
	GapPdf                 []interface{} `json:"GapPdf"`
	TWPdf                  []interface{} `json:"TWPdf"`
	PaintPdf               []interface{} `json:"PaintPdf"`
	EsignProcessStatus     []interface{} `json:"EsignProcessStatus"`
	EsignLink              string        `json:"EsignLink"`
	CustomizeEsignLink     string        `json:"CustomizeEsignLink"`
	BuyerEsignLink         string        `json:"BuyerEsignLink"`
	EmailEsignBuyerEmail   string        `json:"EmailEsignBuyerEmail"`
	CoBuyerEsignLink       string        `json:"CoBuyerEsignLink"`
	EmailEsignCoBuyerEmail string        `json:"EmailEsignCoBuyerEmail"`
	EmailPdfCopy           string        `json:"EmailPdfCopy"`
	EmailPdfAddress        string        `json:"EmailPdfAddress"`
	PrintPdfCopy           string        `json:"PrintPdfCopy"`
	PrintPurchaseOrder     []interface{} `json:"PrintPurchaseOrder"`
	PrintTTBOS             []string      `json:"PrintTTBOS"`
	InsuranceCard          []string      `json:"InsuranceCard"`
	PaymentReceipt         []string      `json:"PaymentReceipt"`
	BBContractNumber       string        `json:"BBContractNumber"`
	BuyNowClicked          bool          `json:"BuyNowClicked"`
	LogText                string        `json:"LogText"`
	SalesInfoSalesPerson   string        `json:"SalesInfoSalesPerson"`
	SnapshotID             string        `json:"SnapshotId"`
	SignMode               string        `json:"SignMode"`
	Note                   string        `json:"Note"`
	TradeNote              string        `json:"TradeNote"`
	TradeSalesPerson       string        `json:"TradeSalesPerson"`
	IsDeleted              bool          `json:"IsDeleted"`
}

func showdetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This home page")

	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}

	var autogen DataStruct
	err = json.Unmarshal([]byte(data), &autogen)
	if err != nil {
		fmt.Println("Error JSON Unmarshling for user file")
		fmt.Println(err)

	}

	w.Write([]byte(data))
}

func main() {
	fmt.Println("Json Marshal and Unmarshal Code")
	http.HandleFunc("/", showdetails)
	fmt.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)

}
