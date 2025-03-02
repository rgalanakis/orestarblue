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
	CctBusiness           ContactType = "B"
	CctPoliticalCommittee ContactType = "C"
	CctFamily             ContactType = "F"
	CctIndividual         ContactType = "I"
	CctLabor              ContactType = "L"
	CctOther              ContactType = "O"
	CctPoliticalParty     ContactType = "P"
	CctUnregistered       ContactType = "U"
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
}

type TransactionType string

//goland:noinspection GoUnusedConst
const (
	TtContribution TransactionType = "C"
	TtOther        TransactionType = "OR"
)

type TransactionSubType string

//goland:noinspection GoUnusedConst
const (
	TstCash                   TransactionSubType = "CA"
	TstInkindContribution     TransactionSubType = "IK"
	TstInkindForgivenAccount  TransactionSubType = "IKA"
	TstInkindForgivenPersonal TransactionSubType = "IKP"
	TstItemSoldFairMarket     TransactionSubType = "FM"
	TstItemReturnedCheck      TransactionSubType = "LC"
	TstItemMisc               TransactionSubType = "OM"
	TstItemRefund             TransactionSubType = "RF"
)
