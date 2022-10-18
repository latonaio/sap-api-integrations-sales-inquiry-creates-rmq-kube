package requests

type Item struct {
	SalesInquiry                  string  `json:"SalesInquiry"`
	SalesInquiryItem              string  `json:"SalesInquiryItem"`
	SalesInquiryItemCategory      *string `json:"SalesInquiryItemCategory"`
	SalesInquiryItemText          *string `json:"SalesInquiryItemText"`
	PurchaseOrderByCustomer       *string `json:"PurchaseOrderByCustomer"`
	Material                      *string `json:"Material"`
	MaterialByCustomer            *string `json:"MaterialByCustomer"`
	RequestedQuantity             *string `json:"RequestedQuantity"`
	RequestedQuantityUnit         *string `json:"RequestedQuantityUnit"`
	ItemOrderProbabilityInPercent *string `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               *string `json:"ItemGrossWeight"`
	ItemNetWeight                 *string `json:"ItemNetWeight"`
	ItemWeightUnit                *string `json:"ItemWeightUnit"`
	ItemVolume                    *string `json:"ItemVolume"`
	ItemVolumeUnit                *string `json:"ItemVolumeUnit"`
	TransactionCurrency           *string `json:"TransactionCurrency"`
	NetAmount                     *string `json:"NetAmount"`
	MaterialGroup                 *string `json:"MaterialGroup"`
	Batch                         *string `json:"Batch"`
	IncotermsClassification       *string `json:"IncotermsClassification"`
	CustomerPaymentTerms          *string `json:"CustomerPaymentTerms"`
	SalesDocumentRjcnReason       *string `json:"SalesDocumentRjcnReason"`
	WBSElement                    *string `json:"WBSElement"`
	SDProcessStatus               *string `json:"SDProcessStatus"`
}
