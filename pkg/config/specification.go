package config

import "sort"

type SpecField struct {
	Start    int
	Length   int
	Type     int
	Required string
}

type SpecRecord struct {
	Key   int
	Name  string
	Field SpecField
}

// field properties
const (
	Nullable   = ""
	Required   = "Y"
	Applicable = "A"
	Expandable = "E"
	Omitted    = "O"
)

// field types
const (
	Alphanumeric = 1 << iota
	AlphanumericRightAlign
	Numeric
	ZeroNumeric
	TelephoneNumber
	Email
	DateYear
	Date
)

var (
	// Transmitter “T” Record
	TRecordLayout = map[string]SpecField{
		"RecordType":                   {0, 1, Alphanumeric, Required},
		"PaymentYear":                  {1, 4, DateYear, Required},
		"PriorYearDataIndicator":       {5, 1, Alphanumeric, Required},
		"TIN":                          {6, 9, Numeric, Required},
		"TCC":                          {15, 5, Alphanumeric, Required},
		"Blank1":                       {20, 7, Alphanumeric, Nullable},
		"TestFileIndicator":            {27, 1, Alphanumeric, Applicable},
		"ForeignEntityIndicator":       {28, 1, Alphanumeric, Applicable},
		"TransmitterName":              {29, 40, Alphanumeric, Required},
		"TransmitterNameContinuation":  {69, 40, Alphanumeric, Applicable},
		"CompanyName":                  {109, 40, Alphanumeric, Required},
		"CompanyNameContinuation":      {149, 40, Alphanumeric, Applicable},
		"CompanyMailingAddress":        {189, 40, Alphanumeric, Required},
		"CompanyCity":                  {229, 40, Alphanumeric, Required},
		"CompanyState":                 {269, 2, Alphanumeric, Required},
		"CompanyZipCode":               {271, 9, Numeric, Required},
		"Blank2":                       {280, 15, Alphanumeric, Nullable},
		"TotalNumberPayees":            {295, 8, ZeroNumeric, Applicable},
		"ContactName":                  {303, 40, Alphanumeric, Required},
		"ContactTelephoneNumber":       {343, 15, TelephoneNumber, Required},
		"ContactEmailAddress":          {358, 50, Email, Applicable},
		"Blank3":                       {408, 91, Alphanumeric, Nullable},
		"RecordSequenceNumber":         {499, 8, ZeroNumeric, Required},
		"Blank4":                       {507, 10, Alphanumeric, Nullable},
		"VendorIndicator":              {517, 1, Alphanumeric, Required},
		"VendorName":                   {518, 40, Alphanumeric, Required},
		"VendorMailingAddress":         {558, 40, Alphanumeric, Required},
		"VendorCity":                   {598, 40, Alphanumeric, Required},
		"VendorState":                  {638, 2, Alphanumeric, Required},
		"VendorZipCode":                {640, 9, Numeric, Required},
		"VendorContactName":            {649, 40, Alphanumeric, Required},
		"VendorContactTelephoneNumber": {689, 15, TelephoneNumber, Required},
		"Blank5":                       {704, 35, Alphanumeric, Nullable},
		"VendorForeignEntityIndicator": {739, 1, Alphanumeric, Applicable},
		"Blank6":                       {740, 8, Alphanumeric, Nullable},
		"Blank7":                       {748, 2, Alphanumeric, Nullable},
	}
	// Payer “A” Record
	ARecordLayout = map[string]SpecField{
		"RecordType":              {0, 1, Alphanumeric, Required},
		"PaymentYear":             {1, 4, DateYear, Required},
		"CombinedFSFilingProgram": {5, 1, Alphanumeric, Applicable},
		"Blank1":                  {6, 5, Alphanumeric, Nullable},
		"TIN":                     {11, 9, Numeric, Required},
		"PayerNameControl":        {20, 4, Alphanumeric, Applicable},
		"LastFilingIndicator":     {24, 1, Alphanumeric, Applicable},
		"TypeOfReturn":            {25, 2, Alphanumeric, Required},
		"AmountCodes":             {27, 16, Alphanumeric, Required},
		"Blank2":                  {43, 8, Alphanumeric, Nullable},
		"ForeignEntityIndicator":  {51, 1, Alphanumeric, Applicable},
		"FirstPayerNameLine":      {52, 40, Alphanumeric, Required},
		"SecondPayerNameLine":     {92, 40, Alphanumeric, Applicable},
		"TransferAgentIndicator":  {132, 1, Alphanumeric, Required},
		"PayerShippingAddress":    {133, 40, Alphanumeric, Required},
		"PayerCity":               {173, 40, Alphanumeric, Required},
		"PayerState":              {213, 2, Alphanumeric, Required},
		"PayerZipCode":            {215, 9, Numeric, Required},
		"PayerTelephoneNumber":    {224, 15, TelephoneNumber, Required},
		"Blank3":                  {239, 260, Alphanumeric, Nullable},
		"RecordSequenceNumber":    {499, 8, ZeroNumeric, Required},
		"Blank4":                  {507, 241, Alphanumeric, Nullable},
		"Blank5":                  {748, 2, Alphanumeric, Nullable},
	}
	// Payee “B” Record
	BRecordLayout = map[string]SpecField{
		"RecordType":               {0, 1, Alphanumeric, Required},
		"PaymentYear":              {1, 4, DateYear, Required},
		"CorrectedReturnIndicator": {5, 1, Alphanumeric, Applicable},
		"NameControl":              {6, 4, Alphanumeric, Applicable},
		"TypeOfTIN":                {10, 1, Numeric, Applicable},
		"TIN":                      {11, 9, Numeric, Required},
		"PayerAccountNumber":       {20, 20, Alphanumeric, Applicable},
		"PayerOfficeCode":          {40, 4, Alphanumeric, Applicable},
		"Blank1":                   {44, 10, Alphanumeric, Nullable},
		"PaymentAmount1":           {54, 12, ZeroNumeric, Applicable},
		"PaymentAmount2":           {66, 12, ZeroNumeric, Applicable},
		"PaymentAmount3":           {78, 12, ZeroNumeric, Applicable},
		"PaymentAmount4":           {90, 12, ZeroNumeric, Applicable},
		"PaymentAmount5":           {102, 12, ZeroNumeric, Applicable},
		"PaymentAmount6":           {114, 12, ZeroNumeric, Applicable},
		"PaymentAmount7":           {126, 12, ZeroNumeric, Applicable},
		"PaymentAmount8":           {138, 12, ZeroNumeric, Applicable},
		"PaymentAmount9":           {150, 12, ZeroNumeric, Applicable},
		"PaymentAmountA":           {162, 12, ZeroNumeric, Applicable},
		"PaymentAmountB":           {174, 12, ZeroNumeric, Applicable},
		"PaymentAmountC":           {186, 12, ZeroNumeric, Applicable},
		"PaymentAmountD":           {198, 12, ZeroNumeric, Applicable},
		"PaymentAmountE":           {210, 12, ZeroNumeric, Applicable},
		"PaymentAmountF":           {222, 12, ZeroNumeric, Applicable},
		"PaymentAmountG":           {234, 12, ZeroNumeric, Applicable},
		"ForeignCountryIndicator":  {246, 1, Alphanumeric, Applicable},
		"FirstPayeeNameLine":       {247, 40, Alphanumeric, Required},
		"SecondPayeeNameLine":      {287, 40, Alphanumeric, Applicable},
		"Blank2":                   {327, 40, Alphanumeric, Nullable},
		"PayeeMailingAddress":      {367, 40, Alphanumeric, Required},
		"Blank3":                   {407, 40, Alphanumeric, Nullable},
		"PayeeCity":                {447, 40, Alphanumeric, Required},
		"PayeeState":               {487, 2, Alphanumeric, Required},
		"PayeeZipCode":             {489, 9, Numeric, Required},
		"Blank4":                   {498, 1, Alphanumeric, Nullable},
		"RecordSequenceNumber":     {499, 8, ZeroNumeric, Required},
		"Blank5":                   {507, 36, Alphanumeric, Nullable},
		"Reserved":                 {543, 207, Alphanumeric, Expandable},
	}
	// End of Payer “C” Record
	CRecordLayout = map[string]SpecField{
		"RecordType":           {0, 1, Alphanumeric, Required},
		"NumberPayees":         {1, 8, ZeroNumeric, Required},
		"Blank1":               {9, 6, Alphanumeric, Nullable},
		"ControlTotal1":        {15, 18, ZeroNumeric, Applicable},
		"ControlTotal2":        {33, 18, ZeroNumeric, Applicable},
		"ControlTotal3":        {51, 18, ZeroNumeric, Applicable},
		"ControlTotal4":        {69, 18, ZeroNumeric, Applicable},
		"ControlTotal5":        {87, 18, ZeroNumeric, Applicable},
		"ControlTotal6":        {105, 18, ZeroNumeric, Applicable},
		"ControlTotal7":        {123, 18, ZeroNumeric, Applicable},
		"ControlTotal8":        {141, 18, ZeroNumeric, Applicable},
		"ControlTotal9":        {159, 18, ZeroNumeric, Applicable},
		"ControlTotalA":        {177, 18, ZeroNumeric, Applicable},
		"ControlTotalB":        {195, 18, ZeroNumeric, Applicable},
		"ControlTotalC":        {213, 18, ZeroNumeric, Applicable},
		"ControlTotalD":        {231, 18, ZeroNumeric, Applicable},
		"ControlTotalE":        {249, 18, ZeroNumeric, Applicable},
		"ControlTotalF":        {267, 18, ZeroNumeric, Applicable},
		"ControlTotalG":        {285, 18, ZeroNumeric, Applicable},
		"Blank2":               {303, 196, Alphanumeric, Nullable},
		"RecordSequenceNumber": {499, 8, ZeroNumeric, Required},
		"Blank3":               {507, 241, Alphanumeric, Nullable},
		"Blank4":               {748, 2, Alphanumeric, Nullable},
	}
	// State Totals “K” Record
	KRecordLayout = map[string]SpecField{
		"RecordType":                  {0, 1, Alphanumeric, Required},
		"NumberPayees":                {1, 8, ZeroNumeric, Required},
		"Blank1":                      {9, 6, Alphanumeric, Nullable},
		"ControlTotal1":               {15, 18, ZeroNumeric, Applicable},
		"ControlTotal2":               {33, 18, ZeroNumeric, Applicable},
		"ControlTotal3":               {51, 18, ZeroNumeric, Applicable},
		"ControlTotal4":               {69, 18, ZeroNumeric, Applicable},
		"ControlTotal5":               {87, 18, ZeroNumeric, Applicable},
		"ControlTotal6":               {105, 18, ZeroNumeric, Applicable},
		"ControlTotal7":               {123, 18, ZeroNumeric, Applicable},
		"ControlTotal8":               {141, 18, ZeroNumeric, Applicable},
		"ControlTotal9":               {159, 18, ZeroNumeric, Applicable},
		"ControlTotalA":               {177, 18, ZeroNumeric, Applicable},
		"ControlTotalB":               {195, 18, ZeroNumeric, Applicable},
		"ControlTotalC":               {213, 18, ZeroNumeric, Applicable},
		"ControlTotalD":               {231, 18, ZeroNumeric, Applicable},
		"ControlTotalE":               {249, 18, ZeroNumeric, Applicable},
		"ControlTotalF":               {267, 18, ZeroNumeric, Applicable},
		"ControlTotalG":               {285, 18, ZeroNumeric, Applicable},
		"Blank2":                      {303, 196, Alphanumeric, Nullable},
		"RecordSequenceNumber":        {499, 8, ZeroNumeric, Required},
		"Blank3":                      {507, 199, Alphanumeric, Nullable},
		"StateIncomeTaxWithheldTotal": {706, 18, Numeric, Applicable},
		"LocalIncomeTaxWithheldTotal": {724, 18, Numeric, Applicable},
		"Blank4":                      {742, 4, Alphanumeric, Nullable},
		"CombinedFederalStateCode":    {746, 2, Alphanumeric, Required},
		"Blank5":                      {748, 2, Alphanumeric, Nullable},
	}
	// End of Transmission “F” Record
	FRecordLayout = map[string]SpecField{
		"RecordType":           {0, 1, Alphanumeric, Required},
		"NumberPayerRecords":   {1, 8, ZeroNumeric, Required},
		"Zero":                 {9, 21, ZeroNumeric, Applicable},
		"Blank2":               {30, 19, Alphanumeric, Nullable},
		"TotalNumberPayees":    {49, 8, ZeroNumeric, Applicable},
		"Blank3":               {57, 442, Alphanumeric, Nullable},
		"RecordSequenceNumber": {499, 8, ZeroNumeric, Required},
		"Blank4":               {507, 241, Alphanumeric, Nullable},
		"Blank5":               {748, 2, Alphanumeric, Nullable},
	}
)

var (
	// Record Layout Positions 544-750 for Form 1097-BTC
	Sub1097BTCLayout = map[string]SpecField{
		"Blank1":             {0, 3, Alphanumeric, Nullable},
		"IssuerIndicator":    {3, 1, Numeric, Required},
		"Blank2":             {4, 8, Alphanumeric, Nullable},
		"Code":               {12, 1, Alphanumeric, Required},
		"Blank3":             {13, 3, Alphanumeric, Nullable},
		"UniqueIdentifier":   {16, 39, AlphanumericRightAlign, Applicable},
		"BondType":           {55, 3, Alphanumeric, Required},
		"Blank4":             {58, 61, Alphanumeric, Nullable},
		"SpecialDataEntries": {119, 60, Alphanumeric, Applicable},
		"Blank5":             {179, 26, Alphanumeric, Nullable},
		"Blank6":             {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1098
	Sub1098Layout = map[string]SpecField{
		"MortgageOriginationDate":           {0, 8, Date, Applicable},
		"PropertySecuringMortgageIndicator": {8, 1, Alphanumeric, Applicable},
		"PropertyADSecuringMortgage":        {9, 39, Alphanumeric, Applicable},
		"Other":                             {48, 39, Alphanumeric, Applicable},
		"Blank1":                            {87, 39, Alphanumeric, Nullable},
		"NumberMortgagedProperties":         {126, 4, ZeroNumeric, Applicable},
		"SpecialDataEntries":                {130, 49, Alphanumeric, Applicable},
		"MortgageAcquisitionDate":           {179, 8, Date, Applicable},
		"Blank2":                            {187, 18, Alphanumeric, Nullable},
		"Blank3":                            {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1098-C
	Sub1098CLayout = map[string]SpecField{
		"Blank1":                               {0, 2, Alphanumeric, Nullable},
		"TransactionIndicator":                 {2, 1, Alphanumeric, Applicable},
		"TransferAfterImprovementsIndicator":   {3, 1, Alphanumeric, Applicable},
		"TransferMarketValueIndicator":         {4, 1, Alphanumeric, Applicable},
		"Year":                                 {5, 4, DateYear, Applicable},
		"Make":                                 {9, 13, Alphanumeric, Applicable},
		"Model":                                {22, 22, Alphanumeric, Applicable},
		"VehicleIdentificationNumber":          {44, 25, Alphanumeric, Applicable},
		"VehicleDescription":                   {69, 39, Alphanumeric, Applicable},
		"DateContribution":                     {108, 8, Date, Applicable},
		"DoneeIndicator":                       {116, 1, Alphanumeric, Applicable},
		"IntangibleReligiousBenefitsIndicator": {117, 1, Alphanumeric, Applicable},
		"DeductionLessIndicator":               {118, 1, Alphanumeric, Applicable},
		"SpecialDataEntries":                   {119, 60, Alphanumeric, Applicable},
		"DateSale":                             {179, 8, Date, Applicable},
		"GoodsServices":                        {187, 16, Alphanumeric, Applicable},
		"Blank2":                               {203, 2, Alphanumeric, Nullable},
		"Blank3":                               {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1098-E
	Sub1098ELayout = map[string]SpecField{
		"Blank1":                       {0, 3, Alphanumeric, Nullable},
		"OriginationInterestIndicator": {3, 1, Alphanumeric, Applicable},
		"Blank2":                       {4, 115, Alphanumeric, Nullable},
		"SpecialDataEntries":           {119, 60, Alphanumeric, Applicable},
		"Blank3":                       {179, 26, Alphanumeric, Nullable},
		"Blank4":                       {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1098-F
	Sub1098FLayout = map[string]SpecField{
		"DateOrderAgreement":  {0, 8, Date, Applicable},
		"Jurisdiction":        {8, 39, Alphanumeric, Applicable},
		"CaseNumber":          {47, 39, Alphanumeric, Applicable},
		"MatterSuitAgreement": {86, 39, Alphanumeric, Applicable},
		"PaymentCode":         {125, 6, Alphanumeric, Applicable},
		"SpecialDataEntries":  {131, 60, Alphanumeric, Applicable},
		"Blank1":              {191, 16, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1098-Q
	Sub1098QLayout = map[string]SpecField{
		"Blank1":                       {0, 2, Alphanumeric, Nullable},
		"AnnuityStartDate":             {2, 8, Date, Applicable},
		"AcceleratedIndicator":         {10, 1, Alphanumeric, Applicable},
		"January":                      {11, 2, ZeroNumeric, Omitted},
		"February":                     {13, 2, ZeroNumeric, Omitted},
		"March":                        {15, 2, ZeroNumeric, Omitted},
		"April":                        {17, 2, ZeroNumeric, Omitted},
		"May":                          {19, 2, ZeroNumeric, Omitted},
		"June":                         {21, 2, ZeroNumeric, Omitted},
		"July":                         {23, 2, ZeroNumeric, Omitted},
		"August":                       {25, 2, ZeroNumeric, Omitted},
		"September":                    {27, 2, ZeroNumeric, Omitted},
		"October":                      {29, 2, ZeroNumeric, Omitted},
		"November":                     {31, 2, ZeroNumeric, Omitted},
		"December":                     {33, 2, ZeroNumeric, Omitted},
		"Blank2":                       {35, 1, Alphanumeric, Nullable},
		"NamePlan":                     {36, 39, Alphanumeric, Applicable},
		"PlanNumber":                   {75, 20, Alphanumeric, Applicable},
		"EmployerIdentificationNumber": {95, 9, Alphanumeric, Applicable},
		"Blank3":                       {104, 101, Alphanumeric, Nullable},
		"Blank4":                       {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1098-T
	Sub1098TLayout = map[string]SpecField{
		"IdentificationNumber":     {0, 1, Alphanumeric, Applicable},
		"Blank1":                   {1, 2, Alphanumeric, Nullable},
		"HalfTimeStudentIndicator": {3, 1, Alphanumeric, Applicable},
		"GraduateStudentIndicator": {4, 1, Numeric, Applicable},
		"AcademicPeriodIndicator":  {5, 1, Numeric, Applicable},
		"Blank2":                   {6, 1, Alphanumeric, Nullable},
		"Blank3":                   {7, 112, Alphanumeric, Nullable},
		"SpecialDataEntries":       {119, 60, Alphanumeric, Applicable},
		"Blank4":                   {179, 26, Alphanumeric, Nullable},
		"Blank5":                   {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-INT
	Sub1099INTLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 2, Alphanumeric, Nullable},
		"ForeignCountry":         {3, 40, Alphanumeric, Applicable},
		"CUSIP":                  {43, 13, AlphanumericRightAlign, Applicable},
		"FATCA":                  {56, 1, Alphanumeric, Applicable},
		"Blank2":                 {57, 62, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, ZeroNumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-MISC
	Sub1099MISCLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 2, Alphanumeric, Nullable},
		"DirectSalesIndicator":   {3, 1, Alphanumeric, Applicable},
		"FATCA":                  {4, 1, Alphanumeric, Applicable},
		"Blank2":                 {5, 114, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, ZeroNumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-OID
	Sub1099OIDLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 2, Alphanumeric, Nullable},
		"Description":            {3, 39, Alphanumeric, Applicable},
		"FATCA":                  {42, 1, Alphanumeric, Applicable},
		"Blank2":                 {43, 76, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, ZeroNumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
	// Record Layout Positions 544-750 for Form 1099-PATR
	Sub1099PATRLayout = map[string]SpecField{
		"SecondTinNotice":        {0, 1, Alphanumeric, Applicable},
		"Blank1":                 {1, 118, Alphanumeric, Nullable},
		"SpecialDataEntries":     {119, 60, Alphanumeric, Applicable},
		"StateIncomeTaxWithheld": {179, 12, ZeroNumeric, Applicable},
		"LocalIncomeTaxWithheld": {191, 12, ZeroNumeric, Applicable},
		"CombinedFSCode":         {203, 2, ZeroNumeric, Required},
		"Blank3":                 {205, 2, Alphanumeric, Nullable},
	}
)

func ToSpecifications(fieldsFormat map[string]SpecField) []SpecRecord {
	var records []SpecRecord
	for key, field := range fieldsFormat {
		records = append(records, SpecRecord{field.Start, key, field})
	}
	sort.Slice(records, func(i, j int) bool {
		if records[i].Key == records[j].Key {
			return records[i].Name < records[j].Name
		}
		return records[i].Key < records[j].Key
	})
	return records
}
