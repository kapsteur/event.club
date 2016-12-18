#Event Club
Your event page with some features:
- Limited access with password
- Google Sheets edition (append row at each booking)
- Payment with stripe

You only need to add some style.

------- FRENCH -------

Votre page évenement avec
- Protection par mot de passe
- Édition d'un Google Sheets (ajout d'une ligne à chaque réservation)
- Paiement via Stripe

Vous n'avez plus qu'à ajouter du style.

## Installation

```
export EVENT_ENV=prod
export EVENT_PORT=8082
export EVENT_STRIPE_PUBLIC=
export EVENT_STRIPE_PRIVATE=

#Only if you need password on front page
export EVENT_PASSWORD=

#Only if you need email at each booking
export EVENT_EMAIL=
export EVENT_EMAIL_USER=
export EVENT_EMAIL_PASSWORD=

#Only if you need sheet update
export EVENT_SHEET_ID=
```

Generating a service account credential : `client_secret.json` (like `client_secret.sample.json`) more infos : https://cloud.google.com/storage/docs/authentication#service_accounts.

