package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	flagset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	help := flagset.Bool("help", false, "")
	filerId := flagset.String("filer-id", "", "")
	out := flagset.String("out", "", "")

	_ = flagset.Parse(os.Args[1:])

	if *filerId == "" {
		filerId = Ptr(os.Getenv("FILER_ID"))
	}

	if *help || *filerId == "" || len(flagset.Args()) == 0 {
		//fmt.Println(*help, *filerId, flagset.Args(), os.Args, flagset.NArg())
		fmt.Fprint(os.Stderr, usage)
		return
	}

	inpath := flagset.Args()[0]
	if err := run(inpath, *filerId, *out); err != nil {
		log.Fatal(err)
	}
}

func run(inpath, filerId, outpath string) error {
	inRdr, err := os.Open(inpath)
	if err != nil {
		return ErrWrap(err, "opening file")
	}
	defer inRdr.Close()
	cr := csv.NewReader(inRdr)
	rowSlices, err := cr.ReadAll()
	if err != nil {
		return ErrWrap(err, "reading csv")
	}

	rows, err := toRows(rowSlices[0], rowSlices[1:])
	if err != nil {
		return err
	}
	contacts, transactions, err := initForExpenses(rows)
	if err != nil {
		return err
	}
	for _, row := range rows {
		c, t := rowToContribution(row)
		contacts = append(contacts, c)
		transactions = append(transactions, t)
	}

	var outWriter io.Writer
	if outpath == "" || outpath == "-" {
		outWriter = os.Stdout
	} else {
		outfile, err := os.Create(outpath)
		if err != nil {
			return ErrWrap(err, "creating file")
		}
		defer outfile.Close()
		outWriter = outfile
	}

	if err := WriteTransactionsDoc(outWriter, filerId, contacts, transactions); err != nil {
		return ErrWrap(err, "writing file")
	}
	return nil
}

func toRows(headers []string, rows [][]string) ([]Row, error) {
	if len(rows) > MaxOrestarRows {
		return nil, fmt.Errorf("orestar only allows %d transactions per-upload, please break the file up (and please file an issue)", MaxOrestarRows)
	}
	m := make(map[string]any, len(headers))
	result := make([]Row, len(rows))
	for rowIdx, row := range rows {
		for cellIdx, h := range headers {
			m[h] = row[cellIdx]
		}
		b, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(b, &result[rowIdx]); err != nil {
			panic(err)
		}
	}
	return result, nil
}

func WriteTransactionsDoc(w io.Writer, filerId string, contacts []Contact, transactions []Transaction) error {
	if _, err := io.WriteString(
		w,
		fmt.Sprintf(`<campaign-finance-transactions xmlns="http://www.state.or.us/sos/ebs2/ce/dataobject" filer-id="%s">`, filerId),
	); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "\n"); err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("  ", "  ")
	for _, c := range contacts {
		if err := enc.Encode(c); err != nil {
			return err
		}
	}
	for _, t := range transactions {
		if err := enc.Encode(t); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "\n</campaign-finance-transactions>\n"); err != nil {
		return err
	}
	return nil
}

const MaxOrestarRows = 2000

// Create the ActBlue contact and expenditure.
func initForExpenses(rows []Row) ([]Contact, []Transaction, error) {
	row := rows[0]
	c := Contact{}
	c.Id = fmt.Sprintf("%s-actblue", idprefix)
	c.Type = ContactTypeBusiness
	c.ContactName.BusinessName = "ActBlue Technical Services"
	c.Address.City = "Somerville"
	c.Address.State = "MA"
	c.Address.Zip = "02144"

	t := Transaction{}
	t.Id = fmt.Sprintf("%s-AB-%s", idprefix, row.CheckNumber)
	t.Operation.Add = true
	t.ContactId = c.Id
	t.Type = ETExpenditure
	t.SubType = ETCashExpenditure
	t.Purpose = TPGeneralOperating
	t.PaymentMethod = TPMEtf
	t.Date = row.CheckDate
	t.Description = "Merchant Fees"

	fee := Money{}
	for _, r := range rows {
		if ramount, err := ParseMoney(r.Fee); err != nil {
			return nil, nil, ErrWrap(err, "parsing fee")
		} else {
			fee = fee.Add(ramount)
		}
	}
	t.Amount = fee.String()
	return []Contact{c}, []Transaction{t}, nil
}
