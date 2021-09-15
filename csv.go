package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
)

type Task struct {
	Orderid string `csv:"Order Number"`
	Email   string `csv:"Email Address"`

	Client *http.Client
}

type Results struct {
	Itemname   string `csv:"Item Name"`
	Size       string `csv:"Size"`
	Sku        string `csv:"SKU"`
	Status     string `csv:"Status"`
	Tracking   string `csv:"Tracking"`
	Fullname   string `csv:"Name"`
	Orderid    string `csv:"Order Number"`
	Email      string `csv:"Email"`
	Addressone string `csv:"Address 1"`
	Addresstwo string `csv:"Address 2"`
	City       string `csv:"City"`
	Postalcode string `csv:"Postal Code"`
}

func ReadCsv() ([]Task, error) {
	file, err := os.OpenFile("orders.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("Fatal Error Opening CSV: %s", err)
		err := errors.New("Fatal CSV Error")
		return nil, err
	}
	defer file.Close()

	var tasks []Task

	if err := gocsv.UnmarshalFile(file, &tasks); err != nil {
		log.Fatalf("Fatal Error Unmarshalling CSV: %s", err)
		err := errors.New("Fatal Error Unmarshalling CSV")
		return nil, err
	}
	return tasks, nil
}

func writeExport(o *Orderinfo) error {
	defer handleError()
	itemName := o.Group[0].Orderitems[0].Product.Title
	itemSize := o.Group[0].Orderitems[0].Product.Size
	itemSku := o.Group[0].Orderitems[0].Product.Stylecolor
	orderStatus := o.Group[0].Heading
	trackingLink := o.Group[0].Actions.Trackshipment.Weblink
	fullName := fmt.Sprintf("%s %s", o.Shipfrom.Recipient.Firstname, o.Shipfrom.Recipient.Lastname)
	orderId := o.Orderid
	eMail := o.Shipfrom.Contactinformation.Email
	addressOne := o.Shipfrom.Address.Address1
	addressTwo := o.Shipfrom.Address.Address2
	city := o.Shipfrom.Address.City
	postalCode := o.Shipfrom.Address.Zipcode

	file, err := os.OpenFile("results.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fatal Error: %s", err)
		err := errors.New("Error Writing Results File")
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "\n"+itemName+","+itemSize+","+itemSku+","+orderStatus+","+trackingLink+","+fullName+","+orderId+","+eMail+","+addressOne+","+addressTwo+","+city+","+postalCode)
	return nil
}
