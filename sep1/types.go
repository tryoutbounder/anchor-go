package sep1

type Principal struct {
	Name                  string `toml:"name,omitempty"`
	Email                 string `toml:"email,omitempty"`
	Keybase               string `toml:"keybase,omitempty"`
	Telegram              string `toml:"telegram,omitempty"`
	Twitter               string `toml:"twitter,omitempty"`
	Github                string `toml:"github,omitempty"`
	IdPhotoHash           string `toml:"id_photo_hash,omitempty"`
	VerificationPhotoHash string `toml:"verification_photo_hash,omitempty"`
}

type Currency struct {
	Code                        string   `toml:"code,omitempty"`
	Issuer                      string   `toml:"issuer,omitempty"`
	Contract                    string   `toml:"contract,omitempty"`
	CodeTemplate                string   `toml:"code_template,omitempty"`
	Status                      string   `toml:"status,omitempty"`
	DisplayDecimals             int      `toml:"display_decimals,omitempty"`
	Name                        string   `toml:"name,omitempty"`
	Desc                        string   `toml:"desc,omitempty"`
	Conditions                  string   `toml:"conditions,omitempty"`
	Image                       string   `toml:"image,omitempty"`
	FixedNumber                 int      `toml:"fixed_number,omitempty"`
	MaxNumber                   int      `toml:"max_number,omitempty"`
	IsUnlimited                 bool     `toml:"is_unlimited,omitempty"`
	IsAssetAnchored             bool     `toml:"is_asset_anchored,omitempty"`
	AnchorAssetType             string   `toml:"anchor_asset_type,omitempty"`
	AnchorAsset                 string   `toml:"anchor_asset,omitempty"`
	AttestationOfReserve        string   `toml:"attestation_of_reserve,omitempty"`
	RedemptionInstructions      string   `toml:"redemption_instructions,omitempty"`
	CollateralAddresses         []string `toml:"collateral_addresses,omitempty"`
	CollateralAddressMessages   []string `toml:"collateral_address_messages,omitempty"`
	CollateralAddressSignatures []string `toml:"collateral_address_signatures,omitempty"`
	Regulated                   bool     `toml:"regulated,omitempty"`
	ApprovalServer              string   `toml:"approval_server,omitempty"`
	ApprovalCriteria            string   `toml:"approval_criteria,omitempty"`
}

type Documentation struct {
	OrgName                       string `toml:"ORG_NAME,omitempty"`
	OrgDBA                        string `toml:"ORG_DBA,omitempty"`
	OrgURL                        string `toml:"ORG_URL,omitempty"`
	OrgLogo                       string `toml:"ORG_LOGO,omitempty"`
	OrgDescription                string `toml:"ORG_DESCRIPTION,omitempty"`
	OrgPhysicalAddress            string `toml:"ORG_PHYSICAL_ADDRESS,omitempty"`
	OrgPhysicalAddressAttestation string `toml:"ORG_PHYSICAL_ADDRESS_ATTESTATION,omitempty"`
	OrgPhoneNumber                string `toml:"ORG_PHONE_NUMBER,omitempty"`
	OrgPhoneNumberAttestation     string `toml:"ORG_PHONE_NUMBER_ATTESTATION,omitempty"`
	OrgKeybase                    string `toml:"ORG_KEYBASE,omitempty"`
	OrgTwitter                    string `toml:"ORG_TWITTER,omitempty"`
	OrgGithub                     string `toml:"ORG_GITHUB,omitempty"`
	OrgOfficialEmail              string `toml:"ORG_OFFICIAL_EMAIL,omitempty"`
	OrgSupportEmail               string `toml:"ORG_SUPPORT_EMAIL,omitempty"`
	OrgLicensingAuthority         string `toml:"ORG_LICENSING_AUTHORITY,omitempty"`
	OrgLicenseType                string `toml:"ORG_LICENSE_TYPE,omitempty"`
	OrgLicenseNumber              string `toml:"ORG_LICENSE_NUMBER,omitempty"`
}

type Validator struct {
	Alias       string `toml:"ALIAS,omitempty"`
	DisplayName string `toml:"DISPLAY_NAME,omitempty"`
	Host        string `toml:"HOST,omitempty"`
	PublicKey   string `toml:"PUBLIC_KEY,omitempty"`
	History     string `toml:"HISTORY,omitempty"`
}

type StellarToml struct {
	Version                     string `toml:"VERSION,omitempty"`
	NetworkPassphrase           string `toml:"NETWORK_PASSPHRASE,omitempty"`
	FederationServer            string `toml:"FEDERATION_SERVER,omitempty"`
	AuthServer                  string `toml:"AUTH_SERVER,omitempty"`
	TransferServer              string `toml:"TRANSFER_SERVER,omitempty"`
	TransferServerSep0024       string `toml:"TRANSFER_SERVER_SEP0024,omitempty"`
	KYCServer                   string `toml:"KYC_SERVER,omitempty"`
	WebAuthEndpoint             string `toml:"WEB_AUTH_ENDPOINT,omitempty"`
	WebAuthForContractsEndpoint string `toml:"WEB_AUTH_FOR_CONTRACTS_ENDPOINT,omitempty"`
	WebAuthContractID           string `toml:"WEB_AUTH_CONTRACT_ID,omitempty"`
	SigningKey                  string `toml:"SIGNING_KEY,omitempty"`
	HorizonURL                  string `toml:"HORIZON_URL,omitempty"`
	// Accounts                    *[]string     `toml:"ACCOUNTS,omitempty"`
	URIRequestSigningKey string        `toml:"URI_REQUEST_SIGNING_KEY,omitempty"`
	DirectPaymentServer  string        `toml:"DIRECT_PAYMENT_SERVER,omitempty"`
	AnchorQuoteServer    string        `toml:"ANCHOR_QUOTE_SERVER,omitempty"`
	Documentation        Documentation `toml:"DOCUMENTATION,omitempty"`
	Principals           []Principal   `toml:"PRINCIPALS,omitempty"`
	Currencies           []Currency    `toml:"CURRENCIES,omitempty"`
	Validators           []Validator   `toml:"VALIDATORS,omitempty"`
}

func (s *StellarToml) String() string {
	return "VERSION = \"" + s.Version + "\"\n" +
		"NETWORK_PASSPHRASE = \"" + s.NetworkPassphrase + "\"\n" +
		"SIGNING_KEY = \"" + s.SigningKey + "\""
}
