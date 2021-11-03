/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
	"time"
)

// InvoiceReadResponse The invoice data stored in HubSpot
type InvoiceReadResponse struct {
	// The invoice number. Note that this is _not_ the ID of the invoice, but the number that the billed customer will see.
	ExternalInvoiceNumber *string `json:"externalInvoiceNumber,omitempty"`
	// The total amount that this invoice is for.
	TotalAmountBilled float32 `json:"totalAmountBilled"`
	// The remaining balance due for the invoice.
	BalanceDue float32 `json:"balanceDue"`
	// The ISO 4217 currency code that represents the currency of the invoice.
	CurrencyCode string `json:"currencyCode"`
	// The due date of the invoice
	DueDate string `json:"dueDate"`
	// The id of the customer in the external accounting system that the invoice was sent to.
	ExternalRecipientId string `json:"externalRecipientId"`
	// The datetime that that invoice was sent to the customer.
	ReceivedByRecipientDate *int64 `json:"receivedByRecipientDate,omitempty"`
	// The datetime that the invoice was created in the external accounting system.
	ExternalCreateDateTime *int64 `json:"externalCreateDateTime,omitempty"`
	// Indicated is the invoice has been voided or not.
	IsVoided bool `json:"isVoided"`
	// The datetime that the invoice was created in HubSpot.
	CreatedAt time.Time `json:"createdAt"`
	// The datetime that the invoice was last updated in HubSpot.
	UpdatedAt  time.Time  `json:"updatedAt"`
	ArchivedAt *time.Time `json:"archivedAt,omitempty"`
	Archived   bool       `json:"archived"`
	// The id of the account in the external accounting system that this invoice is related to.
	ExternalAccountId string `json:"externalAccountId"`
	// The status of the invoice
	InvoiceStatus string `json:"invoiceStatus"`
	// The id of this invoice in the external accounting system.
	Id string `json:"id"`
}

// NewInvoiceReadResponse instantiates a new InvoiceReadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceReadResponse(totalAmountBilled float32, balanceDue float32, currencyCode string, dueDate string, externalRecipientId string, isVoided bool, createdAt time.Time, updatedAt time.Time, archived bool, externalAccountId string, invoiceStatus string, id string) *InvoiceReadResponse {
	this := InvoiceReadResponse{}
	this.TotalAmountBilled = totalAmountBilled
	this.BalanceDue = balanceDue
	this.CurrencyCode = currencyCode
	this.DueDate = dueDate
	this.ExternalRecipientId = externalRecipientId
	this.IsVoided = isVoided
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Archived = archived
	this.ExternalAccountId = externalAccountId
	this.InvoiceStatus = invoiceStatus
	this.Id = id
	return &this
}

// NewInvoiceReadResponseWithDefaults instantiates a new InvoiceReadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceReadResponseWithDefaults() *InvoiceReadResponse {
	this := InvoiceReadResponse{}
	return &this
}

// GetExternalInvoiceNumber returns the ExternalInvoiceNumber field value if set, zero value otherwise.
func (o *InvoiceReadResponse) GetExternalInvoiceNumber() string {
	if o == nil || o.ExternalInvoiceNumber == nil {
		var ret string
		return ret
	}
	return *o.ExternalInvoiceNumber
}

// GetExternalInvoiceNumberOk returns a tuple with the ExternalInvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetExternalInvoiceNumberOk() (*string, bool) {
	if o == nil || o.ExternalInvoiceNumber == nil {
		return nil, false
	}
	return o.ExternalInvoiceNumber, true
}

// HasExternalInvoiceNumber returns a boolean if a field has been set.
func (o *InvoiceReadResponse) HasExternalInvoiceNumber() bool {
	if o != nil && o.ExternalInvoiceNumber != nil {
		return true
	}

	return false
}

// SetExternalInvoiceNumber gets a reference to the given string and assigns it to the ExternalInvoiceNumber field.
func (o *InvoiceReadResponse) SetExternalInvoiceNumber(v string) {
	o.ExternalInvoiceNumber = &v
}

// GetTotalAmountBilled returns the TotalAmountBilled field value
func (o *InvoiceReadResponse) GetTotalAmountBilled() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalAmountBilled
}

// GetTotalAmountBilledOk returns a tuple with the TotalAmountBilled field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetTotalAmountBilledOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalAmountBilled, true
}

// SetTotalAmountBilled sets field value
func (o *InvoiceReadResponse) SetTotalAmountBilled(v float32) {
	o.TotalAmountBilled = v
}

// GetBalanceDue returns the BalanceDue field value
func (o *InvoiceReadResponse) GetBalanceDue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BalanceDue
}

// GetBalanceDueOk returns a tuple with the BalanceDue field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetBalanceDueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceDue, true
}

// SetBalanceDue sets field value
func (o *InvoiceReadResponse) SetBalanceDue(v float32) {
	o.BalanceDue = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *InvoiceReadResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *InvoiceReadResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetDueDate returns the DueDate field value
func (o *InvoiceReadResponse) GetDueDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetDueDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DueDate, true
}

// SetDueDate sets field value
func (o *InvoiceReadResponse) SetDueDate(v string) {
	o.DueDate = v
}

// GetExternalRecipientId returns the ExternalRecipientId field value
func (o *InvoiceReadResponse) GetExternalRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalRecipientId
}

// GetExternalRecipientIdOk returns a tuple with the ExternalRecipientId field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetExternalRecipientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalRecipientId, true
}

// SetExternalRecipientId sets field value
func (o *InvoiceReadResponse) SetExternalRecipientId(v string) {
	o.ExternalRecipientId = v
}

// GetReceivedByRecipientDate returns the ReceivedByRecipientDate field value if set, zero value otherwise.
func (o *InvoiceReadResponse) GetReceivedByRecipientDate() int64 {
	if o == nil || o.ReceivedByRecipientDate == nil {
		var ret int64
		return ret
	}
	return *o.ReceivedByRecipientDate
}

// GetReceivedByRecipientDateOk returns a tuple with the ReceivedByRecipientDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetReceivedByRecipientDateOk() (*int64, bool) {
	if o == nil || o.ReceivedByRecipientDate == nil {
		return nil, false
	}
	return o.ReceivedByRecipientDate, true
}

// HasReceivedByRecipientDate returns a boolean if a field has been set.
func (o *InvoiceReadResponse) HasReceivedByRecipientDate() bool {
	if o != nil && o.ReceivedByRecipientDate != nil {
		return true
	}

	return false
}

// SetReceivedByRecipientDate gets a reference to the given int64 and assigns it to the ReceivedByRecipientDate field.
func (o *InvoiceReadResponse) SetReceivedByRecipientDate(v int64) {
	o.ReceivedByRecipientDate = &v
}

// GetExternalCreateDateTime returns the ExternalCreateDateTime field value if set, zero value otherwise.
func (o *InvoiceReadResponse) GetExternalCreateDateTime() int64 {
	if o == nil || o.ExternalCreateDateTime == nil {
		var ret int64
		return ret
	}
	return *o.ExternalCreateDateTime
}

// GetExternalCreateDateTimeOk returns a tuple with the ExternalCreateDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetExternalCreateDateTimeOk() (*int64, bool) {
	if o == nil || o.ExternalCreateDateTime == nil {
		return nil, false
	}
	return o.ExternalCreateDateTime, true
}

// HasExternalCreateDateTime returns a boolean if a field has been set.
func (o *InvoiceReadResponse) HasExternalCreateDateTime() bool {
	if o != nil && o.ExternalCreateDateTime != nil {
		return true
	}

	return false
}

// SetExternalCreateDateTime gets a reference to the given int64 and assigns it to the ExternalCreateDateTime field.
func (o *InvoiceReadResponse) SetExternalCreateDateTime(v int64) {
	o.ExternalCreateDateTime = &v
}

// GetIsVoided returns the IsVoided field value
func (o *InvoiceReadResponse) GetIsVoided() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsVoided
}

// GetIsVoidedOk returns a tuple with the IsVoided field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetIsVoidedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsVoided, true
}

// SetIsVoided sets field value
func (o *InvoiceReadResponse) SetIsVoided(v bool) {
	o.IsVoided = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InvoiceReadResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InvoiceReadResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *InvoiceReadResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *InvoiceReadResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *InvoiceReadResponse) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *InvoiceReadResponse) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *InvoiceReadResponse) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetArchived returns the Archived field value
func (o *InvoiceReadResponse) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *InvoiceReadResponse) SetArchived(v bool) {
	o.Archived = v
}

// GetExternalAccountId returns the ExternalAccountId field value
func (o *InvoiceReadResponse) GetExternalAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalAccountId
}

// GetExternalAccountIdOk returns a tuple with the ExternalAccountId field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetExternalAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalAccountId, true
}

// SetExternalAccountId sets field value
func (o *InvoiceReadResponse) SetExternalAccountId(v string) {
	o.ExternalAccountId = v
}

// GetInvoiceStatus returns the InvoiceStatus field value
func (o *InvoiceReadResponse) GetInvoiceStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InvoiceStatus
}

// GetInvoiceStatusOk returns a tuple with the InvoiceStatus field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetInvoiceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvoiceStatus, true
}

// SetInvoiceStatus sets field value
func (o *InvoiceReadResponse) SetInvoiceStatus(v string) {
	o.InvoiceStatus = v
}

// GetId returns the Id field value
func (o *InvoiceReadResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InvoiceReadResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InvoiceReadResponse) SetId(v string) {
	o.Id = v
}

func (o InvoiceReadResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalInvoiceNumber != nil {
		toSerialize["externalInvoiceNumber"] = o.ExternalInvoiceNumber
	}
	if true {
		toSerialize["totalAmountBilled"] = o.TotalAmountBilled
	}
	if true {
		toSerialize["balanceDue"] = o.BalanceDue
	}
	if true {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if true {
		toSerialize["dueDate"] = o.DueDate
	}
	if true {
		toSerialize["externalRecipientId"] = o.ExternalRecipientId
	}
	if o.ReceivedByRecipientDate != nil {
		toSerialize["receivedByRecipientDate"] = o.ReceivedByRecipientDate
	}
	if o.ExternalCreateDateTime != nil {
		toSerialize["externalCreateDateTime"] = o.ExternalCreateDateTime
	}
	if true {
		toSerialize["isVoided"] = o.IsVoided
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.ArchivedAt != nil {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["externalAccountId"] = o.ExternalAccountId
	}
	if true {
		toSerialize["invoiceStatus"] = o.InvoiceStatus
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableInvoiceReadResponse struct {
	value *InvoiceReadResponse
	isSet bool
}

func (v NullableInvoiceReadResponse) Get() *InvoiceReadResponse {
	return v.value
}

func (v *NullableInvoiceReadResponse) Set(val *InvoiceReadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceReadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceReadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceReadResponse(val *InvoiceReadResponse) *NullableInvoiceReadResponse {
	return &NullableInvoiceReadResponse{value: val, isSet: true}
}

func (v NullableInvoiceReadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceReadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
