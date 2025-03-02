package main

const usage = `Usage: orestarblue [--filer-id=<id>] [--out=<xml filename>] <ActBlue check CSV export>

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
`
