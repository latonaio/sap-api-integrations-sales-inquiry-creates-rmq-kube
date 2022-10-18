package sap_api_input_reader

import (
	"sap-api-integrations-sales-inquiry-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.SalesInquiry
	results := make([]requests.Item, 0, len(data.SalesInquiryItem))

	header := sdc.ConvertToHeader()

	for i := range data.SalesInquiry {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.SalesInquiry
	return &requests.Header{
		SalesInquiry:                   data.SalesInquiry,
		SalesInquiryType:               data.SalesInquiryType,
		SalesOrganization:              data.SalesOrganization,
		DistributionChannel:            data.DistributionChannel,
		OrganizationDivision:           data.OrganizationDivision,
		SalesGroup:                     data.SalesGroup,
		SalesOffice:                    data.SalesOffice,
		SalesDistrict:                  data.SalesDistrict,
		SoldToParty:                    data.SoldToParty,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
		CustomerPurchaseOrderType:      data.CustomerPurchaseOrderType,
		CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
		SalesInquiryDate:               data.SalesInquiryDate,
		TotalNetAmount:                 data.TotalNetAmount,
		TransactionCurrency:            data.TransactionCurrency,
		SDDocumentReason:               data.SDDocumentReason,
		PricingDate:                    data.PricingDate,
		HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
		BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
		HdrOrderProbabilityInPercent:   data.HdrOrderProbabilityInPercent,
		ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
		IncotermsClassification:        data.IncotermsClassification,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		OverallSDProcessStatus:         data.OverallSDProcessStatus,
		TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
		OverallSDDocumentRejectionSts:  data.OverallSDDocumentRejectionSts,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataSalesInquiry := sdc.SalesInquiry
	data := sdc.SalesInquiry.SalesInquiryItem[num]
	return &requests.Item{
		SalesInquiry:                  dataSalesInquiry.SalesInquiry,
		SalesInquiryItem:              data.SalesInquiryItem,
		SalesInquiryItemCategory:      data.SalesInquiryItemCategory,
		SalesInquiryItemText:          data.SalesInquiryItemText,
		PurchaseOrderByCustomer:       data.PurchaseOrderByCustomer,
		Material:                      data.Material,
		MaterialByCustomer:            data.MaterialByCustomer,
		RequestedQuantity:             data.RequestedQuantity,
		RequestedQuantityUnit:         data.RequestedQuantityUnit,
		ItemOrderProbabilityInPercent: data.ItemOrderProbabilityInPercent,
		ItemGrossWeight:               data.ItemGrossWeight,
		ItemNetWeight:                 data.ItemNetWeight,
		ItemWeightUnit:                data.ItemWeightUnit,
		ItemVolume:                    data.ItemVolume,
		ItemVolumeUnit:                data.ItemVolumeUnit,
		TransactionCurrency:           data.TransactionCurrency,
		NetAmount:                     data.NetAmount,
		MaterialGroup:                 data.MaterialGroup,
		Batch:                         data.Batch,
		IncotermsClassification:       data.IncotermsClassification,
		CustomerPaymentTerms:          data.CustomerPaymentTerms,
		SalesDocumentRjcnReason:       data.SalesDocumentRjcnReason,
		WBSElement:                    data.WBSElement,
		SDProcessStatus:               data.SDProcessStatus,
	}
}
