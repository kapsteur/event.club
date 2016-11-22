{{define "login"}}
<html>
    <head>
        <title>Event.club</title>
        <link rel="stylesheet" href="/static/bootstrap.min.css">
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body class="login">
        <div class="container">
            <div class="row login--header">
                <h1>Event.club</h1>
                <p>Votre page évenement avec protection par mot de passe, réservation en ligne et paiement via Stripe.</p>
            </div>
            <div class="row login--form">
                <form method="POST">
                    <label for="password">Mot de passe:</label>
                    <input id="password" type="password" name="key">
                </form>
            </div>
        </div>
    </body>
</html>
{{end}}