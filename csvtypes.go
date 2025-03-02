package main

type Row struct {
	ReceiptID                    string `json:"Receipt ID"`
	Date                         string `json:"Date"`
	Amount                       string `json:"Amount"`
	RecurringTotalMonths         string `json:"Recurring Total Months"`
	RecurrenceNumber             string `json:"Recurrence Number"`
	Recipient                    string `json:"Recipient"`
	FundraisingPage              string `json:"Fundraising Page"`
	FundraisingPartner           string `json:"Fundraising Partner"`
	ReferenceCode2               string `json:"Reference Code 2"`
	ReferenceCode                string `json:"Reference Code"`
	DonorFirstName               string `json:"Donor First Name"`
	DonorLastName                string `json:"Donor Last Name"`
	DonorAddr1                   string `json:"Donor Addr1"`
	DonorAddr2                   string `json:"Donor Addr2"`
	DonorCity                    string `json:"Donor City"`
	DonorState                   string `json:"Donor State"`
	DonorZIP                     string `json:"Donor ZIP"`
	DonorCountry                 string `json:"Donor Country"`
	DonorOccupation              string `json:"Donor Occupation"`
	DonorEmployer                string `json:"Donor Employer"`
	DonorEmail                   string `json:"Donor Email"`
	DonorPhone                   string `json:"Donor Phone"`
	NewExpressSignup             string `json:"New Express Signup"`
	Comments                     string `json:"Comments"`
	CheckNumber                  string `json:"Check Number"`
	CheckDate                    string `json:"Check Date"`
	EmployerAddr1                string `json:"Employer Addr1"`
	EmployerAddr2                string `json:"Employer Addr2"`
	EmployerCity                 string `json:"Employer City"`
	EmployerState                string `json:"Employer State"`
	EmployerZIP                  string `json:"Employer ZIP"`
	EmployerCountry              string `json:"Employer Country"`
	DonorID                      string `json:"Donor ID"`
	FundraiserID                 string `json:"Fundraiser ID"`
	FundraiserRecipientID        string `json:"Fundraiser Recipient ID"`
	FundraiserContactEmail       string `json:"Fundraiser Contact Email"`
	FundraiserContactFirstName   string `json:"Fundraiser Contact First Name"`
	FundraiserContactLastName    string `json:"Fundraiser Contact Last Name"`
	PartnerID                    string `json:"Partner ID"`
	PartnerContactEmail          string `json:"Partner Contact Email"`
	PartnerContactFirstName      string `json:"Partner Contact First Name"`
	PartnerContactLastName       string `json:"Partner Contact Last Name"`
	LineitemID                   string `json:"Lineitem ID"`
	ABTestName                   string `json:"AB Test Name"`
	ABVariation                  string `json:"AB Variation"`
	RecipientCommittee           string `json:"Recipient Committee"`
	RecipientID                  string `json:"Recipient ID"`
	RecipientGovID               string `json:"Recipient Gov ID"`
	RecipientElection            string `json:"Recipient Election"`
	PaymentID                    string `json:"Payment ID"`
	PaymentDate                  string `json:"Payment Date"`
	DisbursementID               string `json:"Disbursement ID"`
	DisbursementDate             string `json:"Disbursement Date"`
	RecoveryID                   string `json:"Recovery ID"`
	RecoveryDate                 string `json:"Recovery Date"`
	RefundID                     string `json:"Refund ID"`
	RefundDate                   string `json:"Refund Date"`
	Fee                          string `json:"Fee"`
	RecurWeekly                  string `json:"Recur Weekly"`
	ActBlueExpressLane           string `json:"ActBlue Express Lane"`
	CardType                     string `json:"Card Type"`
	Mobile                       string `json:"Mobile"`
	RecurringUpsellShown         string `json:"Recurring Upsell Shown"`
	RecurringUpsellSucceeded     string `json:"Recurring Upsell Succeeded"`
	DoubleDown                   string `json:"Double Down"`
	SmartRecurring               string `json:"Smart Recurring"`
	MonthlyRecurringAmount       string `json:"Monthly Recurring Amount"`
	ApplePay                     string `json:"Apple Pay"`
	CardReplacedByAccountUpdater string `json:"Card Replaced by Account Updater"`
	ActBlueExpressDonor          string `json:"ActBlue Express Donor"`
	CustomField1Label            string `json:"Custom Field 1 Label"`
	CustomField1Value            string `json:"Custom Field 1 Value"`
	DonorUSPassportNumber        string `json:"Donor US Passport Number"`
	TextMessageOptIn             string `json:"Text Message Opt In"`
	GiftIdentifier               string `json:"Gift Identifier"`
	GiftDeclined                 string `json:"Gift Declined"`
	ShippingAddr1                string `json:"Shipping Addr1"`
	ShippingCity                 string `json:"Shipping City"`
	ShippingState                string `json:"Shipping State"`
	ShippingZip                  string `json:"Shipping Zip"`
	ShippingCountry              string `json:"Shipping Country"`
	WeeklyRecurringAmount        string `json:"Weekly Recurring Amount"`
	SmartBoostAmount             string `json:"Smart Boost Amount"`
	SmartBoostShown              string `json:"Smart Boost Shown"`
	BumpRecurringSeen            string `json:"Bump Recurring Seen"`
	BumpRecurringSucceeded       string `json:"Bump Recurring Succeeded"`
	WeeklyToMonthlyRolloverDate  string `json:"Weekly to Monthly Rollover Date"`
	WeeklyRecurringSunset        string `json:"Weekly Recurring Sunset"`
	RecurringType                string `json:"Recurring Type"`
	RecurringPledged             string `json:"Recurring Pledged"`
	Paypal                       string `json:"Paypal"`
	Kind                         string `json:"Kind"`
	ManagedEntityName            string `json:"Managed Entity Name"`
	ManagedEntityCommitteeName   string `json:"Managed Entity Committee Name"`
}
