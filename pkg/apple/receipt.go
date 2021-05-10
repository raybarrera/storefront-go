package apple

// AppStoreValidationRequest : https://developer.apple.com/documentation/appstorereceipts/requestbody
type AppstoreReceiptValidationRequest struct {
	ReceiptData            string `json:"receipt-data"`
	Password               string `json:"password"`
	ExcludeOldTransactions bool   `json:"exclude-old-transactions"`
}

// AppStoreValidationResponse : https://developer.apple.com/documentation/appstorereceipts/responsebody
type AppStoreReceiptValidationResponse struct {
	Environment        string                       `json:"environment"`
	IsRetryable        bool                         `json:"is-retryable"`
	LatestReceipt      string                       `json:"latest-receipt"`
	LatestReceiptInfo  []AppStoreLatestReceiptInfo  `json:"latest_receipt_info"`
	PendingRenewalInfo []AppStorePendingRenewalInfo `json:"pending_renewal_info"`
	Receipt            AppstoreReceipt              `json:"receipt"`
	Status             int                          `json:"status"`
}

type AppStoreLatestReceiptInfo struct {
	CancellationDate            string `json:"cancellation_date"`
	CancellationDateMs          string `json:"cancellation_date_ms"`
	CancellationDatePst         string `json:"cancellation_date_pst"`
	CancellationReason          string `json:"cancellation_reason"`
	ExpiresDate                 string `json:"expires_date"`
	ExpiresDateMs               string `json:"expires_date_ms"`
	ExpiresDatePst              string `json:"expires_date_pst"`
	IsInIntroOfferPeriod        string `json:"is_in_intro_offer_period"`
	IsTrialPeriod               string `json:"is_trial_period"`
	IsUpgraded                  string `json:"is_upgraded"`
	OriginalPurchaseDate        string `json:"original_purchase_date"`
	OriginalPurchaseDateMs      string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst     string `json:"original_purchase_date_pst"`
	OriginalTransactionId       string `json:"original_transaction_id"`
	ProductId                   string `json:"product_id"`
	PromotionalOfferId          string `json:"promotional_offer_id"`
	PurchaseDate                string `json:"purchase_date"`
	PurchaseDateMs              string `json:"purchase_date_ms"`
	PurchaseDatePst             string `json:"purchase_date_pst"`
	Quantity                    string `json:"quantity"`
	SubscriptionGroupIdentifier string `json:"subscription_group_identifier"`
	TransactionId               string `json:"transaction_id"`
	WebOrderLineItemId          string `json:"web_order_line_item_id"`
}

type AppStorePendingRenewalInfo struct {
	AutoRenewProductId        string `json:"auto_renew_product_id"`
	AutoRenewStatus           string `json:"auto_renew_status"`
	ExpirationIntent          string `json:"expiration_intent"`
	GracePeriodExpiresDate    string `json:"grace_period_expires_date"`
	GracePeriodExpiresDateMs  string `json:"grace_period_expires_date_ms"`
	GracePeriodExpiresDatePst string `json:"grace_period_expires_date_pst"`
	IsInBillingRetryPeriod    string `json:"is_in_billing_retry_period"`
	OriginalTransactionId     string `json:"original_transaction_id"`
	PriceConsentStatus        string `json:"price_consent_status"`
	ProductId                 string `json:"product_id"`
}

type AppstoreReceipt struct {
	AdamId                           int    `json:"adam_id"`
	AppItemId                        int    `json:"app_item_id"`
	AppVersion                       string `json:"application_version"`
	BundleId                         string `json:"bundle_id"`
	DownloadID                       int    `json:"download_id"`
	ExpirationDate                   string `json:"expiration_date"`
	ExpirationDateMilliseconds       string `json:"expiration_date_ms"`
	ExpirationDatePST                string `json:"expiration_date_pst"`
	OriginalAppVersion               string `json:"original_application_version"`
	OriginalPurchaseDate             string `json:"original_purchase_date"`
	OriginalPurchaseDateMilliseconds string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePST          string `json:"original_purchase_date_pst"`
	PreorderDate                     string `json:"preorder_date"`
	PreorderDateMilliseconds         string `json:"preorder_date_ms"`
	PreorderDatePST                  string `json:"preorder_date_pst"`
	ReceiptCreationDate              string `json:"receipt_creation_date"`
	ReceiptCreationDateMilliseconds  string `json:"receipt_creation_date_ms"`
	ReceiptCreationDatePST           string `json:"receipt_creation_date_pst"`
	ReceiptType                      string `json:"receipt_type"`
	RequestDate                      string `json:"request_date"`
	RequestDateMilliseconds          string `json:"request_date_ms"`
	RequestDatePST                   string `json:"request_date_pst"`
	VersionExternalIdentifier        string `json:"version_external_identifier"`
}
