package main

type Orderinfo struct {
	Orderid         string          `json:"orderId"`
	Fullorderstatus Fullorderstatus `json:"fullOrderStatus"`
	Regionreference string          `json:"regionReference"`
	Group           []Group         `json:"group"`
	Shipfrom        Shipfrom        `json:"shipFrom"`
}
type Fullorderstatus struct {
	Status         string `json:"status"`
	Statuscolor    string `json:"statusColor"`
	Substatus      string `json:"subStatus"`
	Substatuscolor string `json:"subStatusColor"`
}

type Orderitems struct {
	Product Product `json:"product"`
}

type Product struct {
	Pid        string `json:"pid"`
	Title      string `json:"title"`
	Size       string `json:"size"`
	Stylecolor string `json:"styleColor"`
}
type Lineitemstatus struct {
	Status string `json:"status"`
}
type Lineitemtransaction struct {
	Quantity             int     `json:"quantity"`
	Lineitemchargedprice float64 `json:"lineItemChargedPrice"`
}

type Trackshipment struct {
	Weblink     string `json:"webLink"`
	Appcallback string `json:"appCallback"`
}

type Actions struct {
	Trackshipment Trackshipment `json:"trackShipment"`
}

type Group struct {
	Heading    string       `json:"heading"`
	Subheading string       `json:"subheading"`
	Orderitems []Orderitems `json:"orderItems"`
	Actions    Actions      `json:"actions"`
}

type Address struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Zipcode  string `json:"zipCode"`
}

type Shipfrom struct {
	Address            Address            `json:"address"`
	Recipient          Recipient          `json:"recipient"`
	Contactinformation Contactinformation `json:"contactInformation"`
}

type Contactinformation struct {
	Dayphonenumber string `json:"dayPhoneNumber"`
	Email          string `json:"email"`
}

type Recipient struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

type Config struct {
	Webhookcfg Webhookcfg `json:"webhook"`
}
type Webhookcfg struct {
	Webhookenabled bool   `json:"webhookEnabled"`
	Webhookurl     string `json:"webhookUrl"`
}
