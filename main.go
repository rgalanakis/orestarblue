package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"errors"
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
	if len(rows) > MaxOrestarRows {
		return errors.New("orestar only allows 2000 transactions per-upload, please break the file up or add support to this project")
	}
	var contacts []Contact
	var transactions []Transaction
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
