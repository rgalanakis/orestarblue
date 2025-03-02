package main

type Contact struct {
	XMLName struct{}    `xml:"contact"`
	Id      string      `xml:"id,attr"`
	Type    ContactType `xml:"type"`

	ContactName ContactName `xml:"contact-name"`

	Address struct {
		Street1 string `xml:"street1"`
		Street2 string `xml:"street2"`
		City    string `xml:"city"`
		State   string `xml:"state"`
		Zip     string `xml:"zip"`
	} `xml:"address"`

	Phone struct {
		Home string `xml:"home"`
	} `xml:"phone"`
	Email      string `xml:"email,omitempty"`
	Occupation string `xml:"occupation"`

	Employment Employment `xml:"employment"`
}

type ContactName struct {
	Committee      *ContactNameCommittee  `xml:"committee,omitempty"`
	IndividualName *ContactNameIndividual `xml:"individual-name,omitempty"`
	BusinessName   string                 `xml:"business-name,omitempty"`
}

type ContactNameCommittee struct {
	Name string `xml:"name"`
}

type ContactNameIndividual struct {
	First string `xml:"first"`
	Last  string `xml:"last"`
}

type Employment struct {
	NotEmployed  *string `xml:"not-employed,omitempty"`
	SelfEmployed *string `xml:"self-employed,omitempty"`
	EmployerName string  `xml:"employer-name,omitempty"`
	City         string  `xml:"city,omitempty"`
	State        string  `xml:"state,omitempty"`
}

type ContactType string

//goland:noinspection GoUnusedConst
const (
	ContactTypeBusiness           ContactType = "B"
	ContactTypePoliticalCommittee ContactType = "C"
	ContactTypeFamily             ContactType = "F"
	ContactTypeIndividual         ContactType = "I"
	ContactTypeLabor              ContactType = "L"
	ContactTypeOther              ContactType = "O"
	ContactTypePoliticalParty     ContactType = "P"
	ContactTypeUnregistered       ContactType = "U"
)

type Transaction struct {
	XMLName   struct{} `xml:"transaction"`
	Id        string   `xml:"id,attr"`
	Operation struct {
		Add bool `xml:"add"`
	} `xml:"operation"`
	ContactId string             `xml:"contact-id"`
	Type      TransactionType    `xml:"type"`
	SubType   TransactionSubType `xml:"sub-type"`
	Amount    string             `xml:"amount"`
	Date      string             `xml:"date"`

	Purpose       TransactionPurpose       `xml:"tran-purpose,omitempty"`
	PaymentMethod TransactionPaymentMethod `xml:"payment-method,omitempty"`
	Description   string                   `xml:"description,omitempty"`
}

type TransactionType string

//goland:noinspection GoUnusedConst
const (
	CTContribution TransactionType = "C"
	CTOther        TransactionType = "OR"

	ETExpenditure       TransactionType = "E"
	ETOther             TransactionType = "O"
	ETOtherDisbursement TransactionType = "OD"
)

type TransactionSubType string

//goland:noinspection GoUnusedConst
const (
	TSTCash                   TransactionSubType = "CA"
	TSTInkindContribution     TransactionSubType = "IK"
	TSTInkindForgivenAccount  TransactionSubType = "IKA"
	TSTInkindForgivenPersonal TransactionSubType = "IKP"
	TSTItemSoldFairMarket     TransactionSubType = "FM"
	TSTItemReturnedCheck      TransactionSubType = "LC"
	TSTItemMisc               TransactionSubType = "OM"
	TSTItemRefund             TransactionSubType = "RF"

	ETAccountsPayable                TransactionSubType = "AP"
	ETCashExpenditure                TransactionSubType = "CE"
	ETPersonalExpenditure            TransactionSubType = "PE"
	ETAccountsPayableRescinded       TransactionSubType = "APR"
	ETCashBalanceAdjustment          TransactionSubType = "CBA"
	ETMiscellaneousOtherDisbursement TransactionSubType = "OMD"
	ETRefundOfContribution           TransactionSubType = "RF"
)

type TransactionPurpose string

//goland:noinspection GoUnusedConst
const (
	TPWages            TransactionPurpose = "W"
	TPCash             TransactionPurpose = "C"
	TPReimbursement    TransactionPurpose = "R"
	TPBroadcast        TransactionPurpose = "B"
	TPFundraising      TransactionPurpose = "F"
	TPGeneralOperating TransactionPurpose = "G"
	TPPrimting         TransactionPurpose = "L"
	TPManagement       TransactionPurpose = "M"
	TPNewspaper        TransactionPurpose = "N"
	TPOtherAd          TransactionPurpose = "O"
	TPPetition         TransactionPurpose = "Y"
	TPPostage          TransactionPurpose = "P"
	TPPrepAd           TransactionPurpose = "A"
	TPPolling          TransactionPurpose = "S"
	TPTravel           TransactionPurpose = "T"
	TPUtilities        TransactionPurpose = "U"
)

// https://sos.oregon.gov/elections/Documents/orestarTransFiling.pdf
// source on page 33-34:

type TransactionPaymentMethod string

//goland:noinspection GoUnusedConst
const (
	TPMCash             TransactionPaymentMethod = "CA"
	TPMCheck            TransactionPaymentMethod = "CHK"
	TPMMoneyOrder       TransactionPaymentMethod = "CHK"
	TPMCreditCardOnline TransactionPaymentMethod = "CC"
	TPMCreditCardPaper  TransactionPaymentMethod = "CC"
	TPMEtf              TransactionPaymentMethod = "EFT"
	TPMDebit            TransactionPaymentMethod = "DC"
)
