package main

import (
	"fmt"
	"regexp"
	"strings"
)

const idprefix = "oae"
const maxOrestarIdLen = 30

// https://github.com/hackoregon/openelections/blob/develop/api/models/converters/index.ts#L4
// https://github.com/hackoregon/openelections/blob/develop/api/models/converters/orestarContributionConverter.ts#L78
func rowToContribution(r Row) (c Contact, t Transaction) {
	longContactId := fmt.Sprintf(
		`%s-%s-%s-%s-%s`,
		idprefix,
		r.DonorZIP,
		r.DonorLastName,
		r.DonorFirstName,
		r.PaymentID,
	)
	longContactId = removeSpaces.ReplaceAllString(longContactId, "-")
	longContactId = strings.ToLower(longContactId)
	oaeContactId := Substr(longContactId, maxOrestarIdLen)

	c.Id = oaeContactId
	switch r.Kind {
	case "page":
		c.Type = CctIndividual
	default:
		panic(fmt.Errorf("unknown contact type: %+v", r))
	}
	switch c.Type {
	case CctIndividual:
		c.ContactName.IndividualName = &ContactNameIndividual{
			First: r.DonorFirstName,
			Last:  r.DonorLastName,
		}
	default:
		panic(fmt.Errorf("unknown contact type: %v", c.Type))
	}

	c.Address.Street1 = r.DonorAddr1
	c.Address.Street2 = r.DonorAddr2
	c.Address.City = r.DonorCity
	c.Address.State = r.DonorState
	c.Address.Zip = r.DonorZIP[:5]

	// Phone can be invalid so give it a default if so.
	if len(r.DonorPhone) >= 10 {
		c.Phone.Home = r.DonorPhone
	} else {
		c.Phone.Home = "555-555-5555"
	}

	c.Email = r.DonorEmail
	c.Occupation = r.DonorOccupation
	switch r.DonorEmployer {
	case "Self-employed":
		c.Employment.SelfEmployed = Ptr("")
	case "Not Employed":
		c.Employment.NotEmployed = Ptr("")
	default:
		c.Employment.EmployerName = r.DonorEmployer
		c.Employment.State = r.EmployerState
		c.Employment.City = r.EmployerCity
	}

	t.Id = Substr(fmt.Sprintf("%s-%s-%s", idprefix, r.ReceiptID, r.PaymentID), maxOrestarIdLen)
	t.ContactId = oaeContactId
	t.Operation.Add = true
	t.Type = TtContribution
	t.SubType = TstCash
	t.Amount = r.Amount
	t.Date = r.PaymentDate[:len("2000-01-01")]

	return
}

var removeSpaces = regexp.MustCompile(`\s+`)
