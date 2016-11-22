package controller

import (
	"crypto/sha1"
	"github.com/kapsteur/event.club/config"
	"github.com/kapsteur/event.club/helper"
	"github.com/kapsteur/event.club/model"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/currency"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var templates = template.Must(template.ParseGlob("template/*"))
var passwordHash string

const MealPrice = 10

func init() {
	conf := config.Conf()
	stripe.Key = conf.Stripe.Private_Key
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	conf := config.Conf()
	tpl := "login"
	data := model.Response{Booking: model.Booking{}, MealPrice: MealPrice, StripeKey: conf.Stripe.Public_Key}

	//Refresh templates only on dev
	if conf.Env == "dev" {
		templates = template.Must(template.ParseGlob("template/*"))
	}

	//Test if Password is needed
	if len(conf.Password) > 0 {

		userPasswordHash := ""

		//Get user password from form
		userPassword := r.PostFormValue("key")
		if len(userPassword) > 0 {
			log.Printf("Password:%s", userPassword)

			//Generate user password hash
			h := sha1.New()
			h.Write([]byte(userPassword))
			userPasswordHash = string(h.Sum(nil))

		}

		//Get user password from cookie
		if ck, ok := r.Cookie("key"); len(userPasswordHash) == 0 && ok == nil {
			userPasswordHash, _ = url.QueryUnescape(ck.Value)
		}

		//Init passHash at first occurence
		if len(passwordHash) == 0 {
			h := sha1.New()
			h.Write([]byte(conf.Password))
			passwordHash = string(h.Sum(nil))
		}

		if passwordHash == userPasswordHash {
			tpl = "home"
			http.SetCookie(w, &http.Cookie{Name: "key", Value: url.QueryEscape(userPasswordHash), Path: "/", Domain: "localhost", Expires: time.Now().AddDate(0, 1, 0)})
		}

	} else {
		tpl = "home"
	}

	templates.ExecuteTemplate(w, tpl, data)
	return

}

func BookingHandler(w http.ResponseWriter, r *http.Request) {
	conf := config.Conf()
	tpl := "home"
	data := model.Response{Booking: model.Booking{}, MealPrice: MealPrice, StripeKey: conf.Stripe.Public_Key}

	r.ParseForm()

	data.Booking.Name = r.Form.Get("name")
	data.Booking.Email = r.Form.Get("email")
	data.Booking.Phone = r.Form.Get("phone")
	data.Booking.Meal1, _ = strconv.Atoi(r.Form.Get("meal1"))
	data.Booking.Meal2, _ = strconv.Atoi(r.Form.Get("meal2"))
	data.Booking.Token = r.Form.Get("token")

	if len(data.Booking.Name) > 0 || (data.Booking.Meal1+data.Booking.Meal2) > 0 {
		if len(data.Booking.Name) == 0 {
			data.Error = true
			data.Message = "- Les nom & prénom sont incorrectes"
		}

		if len(data.Booking.Email) == 0 || !strings.Contains(data.Booking.Email, "@") || !strings.Contains(data.Booking.Email, ".") {
			data.Error = true
			data.Message = "- L'email est incorrecte"
		}

		if len(data.Booking.Phone) != 10 {
			data.Error = true
			data.Message = "- Le numéro de téléphone est incorrecte"
		}

		if (data.Booking.Meal1 + data.Booking.Meal2) == 0 {
			data.Error = true
			data.Message = "- Sélectionnez au moins 1 repas pour réserver"
		}

		if len(data.Booking.Token) == 0 {
			data.Error = true
			data.Message = "- Vérifiez vos informations de paiement"
		}

		if !data.Error {

			params := &stripe.ChargeParams{
				Amount:   uint64((data.Booking.Meal1 + data.Booking.Meal2) * MealPrice * 100),
				Currency: currency.EUR,
			}
			params.SetSource(data.Booking.Token)
			params.Desc = "Event.club"

			_, err := charge.New(params)
			if err != nil {
				data.Error = true
				data.Message = "Une erreur est survenu pendant le paiement"
				log.Printf("Charge New: %v", err)

			} else {
				data.Success = true
				data.Message = "Votre réservation a bien été enregistrée"

				if len(conf.Email.Email) > 0 && len(conf.Email.User) > 0 {
					//Send email
					helper.SendMail(data.Booking)
				}

				if len(conf.Sheet.Id) > 0 {
					//Append row
					helper.AppendRow(data.Booking)
				}

			}
		}
	}

	templates.ExecuteTemplate(w, tpl, data)
	return
}

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
