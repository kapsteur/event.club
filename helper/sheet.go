package helper

import (
	"github.com/kapsteur/event.club/config"
	"github.com/kapsteur/event.club/model"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
)

//That's function name
func AppendRow(booking model.Booking) {
	conf := config.Conf()

	//Append sheet
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Printf("Unable to read client secret file: %v", err)
	}

	jwtconfig, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Printf("Unable to parse client secret file to config: %v", err)
	}

	client := jwtconfig.Client(ctx)

	srv, err := sheets.New(client)
	if err != nil {
		log.Printf("Unable to retrieve Sheets Client %v", err)
	}

	values := new(sheets.ValueRange)

	values.Values = make([][]interface{}, 1)
	values.Values[0] = make([]interface{}, 0)
	values.Values[0] = append(values.Values[0], booking.Name, booking.Email, booking.Phone, booking.Meal1, booking.Meal2)

	_, err = srv.Spreadsheets.Values.Append(conf.Sheet.Id, "A1", values).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Printf("Unable to append document:", err)
	}

}
