/**
 *
 * Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.
 *
 */
package src

import (
	"example.com/gitbugtestgo1/lib/util/date"
)

type Supplier struct {
	AssetType string `json:"AssetType" final:"gitbugtestgo1.Supplier"`

	SupplierId           string      `json:"SupplierId" validate:"string" id:"true" mandatory:"true"`
	RawMaterialAvailable int         `json:"RawMaterialAvailable" validate:"int,min=0"`
	License              string      `json:"License" validate:"string,min=10,max=18"`
	ExpiryDate           date.Date   `json:"ExpiryDate" validate:"date,before=2020-06-26"`
	Active               bool        `json:"Active" validate:"bool" default:"true"`
	Metadata             interface{} `json:"Metadata,omitempty"`
}

type Manufacturer struct {
	AssetType string `json:"AssetType" final:"gitbugtestgo1.Manufacturer"`

	ManufacturerId       string      `json:"ManufacturerId" validate:"string" id:"true" mandatory:"true"`
	RawMaterialAvailable int         `json:"RawMaterialAvailable" validate:"int,max=8"`
	ProductsAvailable    int         `json:"ProductsAvailable" validate:"int,min=3"`
	CompletionDate       date.Date   `json:"CompletionDate" validate:"date,after=2020-06-26T02:30:55Z,before=2020-06-28T02:30:55Z"`
	Metadata             interface{} `json:"Metadata,omitempty"`
}

type Distributor struct {
	AssetType string `json:"AssetType" final:"gitbugtestgo1.Distributor"`

	DistributorId       string      `json:"DistributorId" validate:"string" id:"true" mandatory:"true"`
	ProductsToBeShipped int         `json:"ProductsToBeShipped" validate:"int"`
	ProductsShipped     int         `json:"ProductsShipped" validate:"int,min=3"`
	ProductsReceived    int         `json:"ProductsReceived" validate:"int,max=0"`
	MailId              string      `json:"MailId" validate:"string,email"`
	DistributionDate    date.Date   `json:"DistributionDate" validate:"date"`
	Metadata            interface{} `json:"Metadata,omitempty"`
}

type Retailer struct {
	AssetType string `json:"AssetType" final:"gitbugtestgo1.Retailer"`

	RetailerId        string      `json:"RetailerId" validate:"string" id:"true" mandatory:"true"`
	ProductsOrdered   int         `json:"ProductsOrdered" validate:"int" mandatory:"true"`
	ProductsAvailable int         `json:"ProductsAvailable" validate:"int" default:"1"`
	ProductsSold      int         `json:"ProductsSold" validate:"int"`
	Remarks           string      `json:"Remarks" validate:"string" default:"open for business"`
	Items             []int       `json:"Items" validate:"array=int,range=2-5"`
	Domain            string      `json:"Domain" validate:"string,url,min=30,max=50"`
	Metadata          interface{} `json:"Metadata,omitempty"`
}

type Customer struct {
	AssetType string `json:"AssetType" final:"gitbugtestgo1.Customer"`

	CustomerId     string      `json:"CustomerId" validate:"string" id:"true" mandatory:"true"`
	Name           string      `json:"Name" validate:"string" mandatory:"true"`
	ProductsBought int         `json:"ProductsBought" validate:"int"`
	OfferApplied   int         `json:"OfferApplied" validate:"int,max=0"`
	PhoneNumber    string      `json:"PhoneNumber" validate:"string,regexp=^\(?([0-9]{3})\)?[-. ]?([0-9]{3})[-. ]?([0-9]{4})$"`
	Received       bool        `json:"Received" validate:"bool"`
	Metadata       interface{} `json:"Metadata,omitempty"`
}
