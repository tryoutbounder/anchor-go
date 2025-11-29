package anchorutils

type AnchorError struct {
	AnchorErr string `json:"error"` // Error message
}

func (ae AnchorError) Error() string {
	return ae.AnchorErr
}

type Sep9Fields struct {
	Sep9PersonFields
	Sep9FinancialAccountFields
	// Sep9OrganizationalFields
	// Sep9CardFields
}

type Sep9PersonFields struct {
	FamilyName              string `json:"family_name,omitempty" form:"family_name"`
	LastName                string `json:"last_name,omitempty" form:"last_name"`
	GivenName               string `json:"given_name,omitempty" form:"given_name"`
	FirstName               string `json:"first_name,omitempty" form:"first_name"`
	AdditionalName          string `json:"additional_name,omitempty" form:"additional_name"`
	AddressCountryCode      string `json:"address_country_code,omitempty" form:"address_country_code"`
	StateOrProvince         string `json:"state_or_province,omitempty" form:"state_or_province"`
	City                    string `json:"city,omitempty" form:"city"`
	PostalCode              string `json:"postal_code,omitempty" form:"postal_code"`
	Address                 string `json:"address,omitempty" form:"address"`
	MobileNumber            string `json:"mobile_number,omitempty" form:"mobile_number"`
	MobileNumberFormat      string `json:"mobile_number_format,omitempty" form:"mobile_number_format"`
	EmailAddress            string `json:"email_address,omitempty" form:"email_address"`
	BirthDate               string `json:"birth_date,omitempty" form:"birth_date"`
	BirthPlace              string `json:"birth_place,omitempty" form:"birth_place"`
	BirthCountryCode        string `json:"birth_country_code,omitempty" form:"birth_country_code"`
	TaxID                   string `json:"tax_id,omitempty" form:"tax_id"`
	TaxIDName               string `json:"tax_id_name,omitempty" form:"tax_id_name"`
	Occupation              string `json:"occupation,omitempty" form:"occupation"`
	EmployerName            string `json:"employer_name,omitempty" form:"employer_name"`
	EmployerAddress         string `json:"employer_address,omitempty" form:"employer_address"`
	LanguageCode            string `json:"language_code,omitempty" form:"language_code"`
	IDType                  string `json:"id_type,omitempty" form:"id_type"`
	IDCountryCode           string `json:"id_country_code,omitempty" form:"id_country_code"`
	IDIssueDate             string `json:"id_issue_date,omitempty" form:"id_issue_date"`
	IDExpirationDate        string `json:"id_expiration_date,omitempty" form:"id_expiration_date"`
	IDNumber                string `json:"id_number,omitempty" form:"id_number"`
	PhotoIDFront            []byte `json:"photo_id_front,omitempty" form:"photo_id_front"`
	PhotoIDBack             []byte `json:"photo_id_back,omitempty" form:"photo_id_back"`
	NotaryApprovalOfPhotoID []byte `json:"notary_approval_of_photo_id,omitempty" form:"notary_approval_of_photo_id"`
	IPAddress               string `json:"ip_address,omitempty" form:"ip_address"`
	PhotoProofResidence     []byte `json:"photo_proof_residence,omitempty" form:"photo_proof_residence"`
	Sex                     string `json:"sex,omitempty" form:"sex"`
	ProofOfIncome           []byte `json:"proof_of_income,omitempty" form:"proof_of_income"`
	ProofOfLiveness         []byte `json:"proof_of_liveness,omitempty" form:"proof_of_liveness"`
	ReferralID              string `json:"referral_id,omitempty" form:"referral_id"`
}

type Sep9FinancialAccountFields struct {
	BankName             string `json:"bank_name,omitempty" form:"bank_name"`
	BankAccountType      string `json:"bank_account_type,omitempty" form:"bank_account_type"`
	BankAccountNumber    string `json:"bank_account_number,omitempty" form:"bank_account_number"`
	BankNumber           string `json:"bank_number,omitempty" form:"bank_number"`
	BankPhoneNumber      string `json:"bank_phone_number,omitempty" form:"bank_phone_number"`
	BankBranchNumber     string `json:"bank_branch_number,omitempty" form:"bank_branch_number"`
	ExternalTransferMemo string `json:"external_transfer_memo,omitempty" form:"external_transfer_memo"`
	ClabeNumber          string `json:"clabe_number,omitempty" form:"clabe_number"`
	CBUNumber            string `json:"cbu_number,omitempty" form:"cbu_number"`
	CBUAlias             string `json:"cbu_alias,omitempty" form:"cbu_alias"`
	MobileMoneyNumber    string `json:"mobile_money_number,omitempty" form:"mobile_money_number"`
	MobileMoneyProvider  string `json:"mobile_money_provider,omitempty" form:"mobile_money_provider"`
	CryptoAddress        string `json:"crypto_address,omitempty" form:"crypto_address"`
	CryptoMemo           string `json:"crypto_memo,omitempty" form:"crypto_memo"`
}

type Sep9OrganizationalFields struct {
	OrganizationName                 string `json:"organization.name,omitempty" form:"organization.name"`
	OrganizationVATNumber            string `json:"organization.VAT_number,omitempty" form:"organization.VAT_number"`
	OrganizationRegistrationNumber   string `json:"organization.registration_number,omitempty" form:"organization.registration_number"`
	OrganizationRegistrationDate     string `json:"organization.registration_date,omitempty" form:"organization.registration_date"`
	OrganizationRegisteredAddress    string `json:"organization.registered_address,omitempty" form:"organization.registered_address"`
	OrganizationNumberOfShareholders int    `json:"organization.number_of_shareholders,omitempty" form:"organization.number_of_shareholders"`
	OrganizationShareholderName      string `json:"organization.shareholder_name,omitempty" form:"organization.shareholder_name"`
	// OrganizationPhotoIncorporationDoc []byte `json:"organization.photo_incorporation_doc,omitempty" form:"organization.photo_incorporation_doc"`
	// OrganizationPhotoProofAddress     []byte `json:"organization.photo_proof_address,omitempty" form:"organization.photo_proof_address"`
	OrganizationAddressCountryCode string `json:"organization.address_country_code,omitempty" form:"organization.address_country_code"`
	OrganizationStateOrProvince    string `json:"organization.state_or_province,omitempty" form:"organization.state_or_province"`
	OrganizationCity               string `json:"organization.city,omitempty" form:"organization.city"`
	OrganizationPostalCode         string `json:"organization.postal_code,omitempty" form:"organization.postal_code"`
	OrganizationDirectorName       string `json:"organization.director_name,omitempty" form:"organization.director_name"`
	OrganizationWebsite            string `json:"organization.website,omitempty" form:"organization.website"`
	OrganizationEmail              string `json:"organization.email,omitempty" form:"organization.email"`
	OrganizationPhone              string `json:"organization.phone,omitempty" form:"organization.phone"`
}

type Sep9CardFields struct {
	CardNumber          string `json:"card.number,omitempty" form:"card.number"`
	CardExpirationDate  string `json:"card.expiration_date,omitempty" form:"card.expiration_date"`
	CardCVC             string `json:"card.cvc,omitempty" form:"card.cvc"`
	CardHolderName      string `json:"card.holder_name,omitempty" form:"card.holder_name"`
	CardNetwork         string `json:"card.network,omitempty" form:"card.network"`
	CardPostalCode      string `json:"card.postal_code,omitempty" form:"card.postal_code"`
	CardCountryCode     string `json:"card.country_code,omitempty" form:"card.country_code"`
	CardStateOrProvince string `json:"card.state_or_province,omitempty" form:"card.state_or_province"`
	CardCity            string `json:"card.city,omitempty" form:"card.city"`
	CardAddress         string `json:"card.address,omitempty" form:"card.address"`
	CardToken           string `json:"card.token,omitempty" form:"card.token"`
}

type Sep9Files struct {
	PhotoIDFront                      []byte `json:"photo_id_front,omitempty"`
	PhotoIDBack                       []byte `json:"photo_id_back,omitempty"`
	NotaryApprovalOfPhotoID           []byte `json:"notary_approval_of_photo_id,omitempty"`
	PhotoProofResidence               []byte `json:"photo_proof_residence,omitempty"`
	ProofOfIncome                     []byte `json:"proof_of_income,omitempty"`
	ProofOfLiveness                   []byte `json:"proof_of_liveness,omitempty"`
	OrganizationPhotoIncorporationDoc []byte `json:"organization.photo_incorporation_doc,omitempty"`
	OrganizationPhotoProofAddress     []byte `json:"organization.photo_proof_address,omitempty"`
}
