ActBlue CSV to ORESTAR XML Converter
===

**Developed by Rob Galanakis, [www.robgforpps.com](https://robgforpps.com).**

**Vote for Rob G in the May 2025 election for Portland School Board, Zone 6!**

# Usage

```
Usage: orestarblue [--filer-id=<id>] [--out=<xml filename>] <ActBlue check CSV export>

Convert an ActBlue check CSV download into an ORESTAR XML file.

---
NOTE: This software is still in Alpha. Campaign Finance is serious business,
and you MUST review your transactions on ORESTAR after uploading them,
before you file. If there are any issues, PLEASE file a bug report.
The author is not responsible for any issues this software causes.
---

To get the CSV for a check, from your ActBlue dashboard, go to Reporting, then do a CSV download for the check.

Pass that filename to orestarblue.

You will also need your Filer ID. You can get this from ORESTAR, it is the value next to your entity name.

For example, "Rob Galanakis for PPS (24090)" has a Filer ID of 24090.

--filer-id=<id>

	Pass your Filer ID (like 24090) as --filter-id, or set the FILER_ID env var.

--out=<xml filename>

	Pass the filename to write to. If not given, write to stdout.
```

## Notes

This uploader only handles adding transactions for cash contributions and actblue fees;
it does not (yet?) handle things like refunds.

**YOU MUST CHECK OVER THE OUTPUTS.**

We create a Contact for every row, based on their name and zipcode.
This ensures they stay consistent between uploads.

We create a Cash Contribution for every row in the CSV.

We create one Cash Expenditure for the total of all Fees in the CSV,
and add it as one General Expense line item. This is how I have seen it recorded in ORESTAR.

## Development

Code references ('G' means 'Transaction Puropose -> General') can be found here:
https://sos.oregon.gov/elections/Documents/xml_overview.pdf

More formal XML specs can be found here:
https://sos.oregon.gov/elections/Documents/xml_specs.pdf

The Hack Oregon converters provide some great working code. They can can be found here:
- https://github.com/hackoregon/openelections/blob/develop/api/models/converters/orestarContributionConverter.ts
- https://github.com/hackoregon/openelections/blob/develop/api/models/converters/orestarExpenditureConverter.ts
