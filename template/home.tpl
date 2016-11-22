{{define "home"}}
<html>
    <head>
        <title>Event.club</title>
        <link rel="stylesheet" href="/static/bootstrap.min.css">
        <link rel="stylesheet" type="text/css" href="/static/style.css">
        <script src="/static/jquery-3.1.1.min.js" type="text/javascript"></script>
    </head>
    <body>
        <div class="container">
            <div class="home--message {{ if .Success }}success{{ end }} {{ if .Error }}error{{ end }}">
                {{ .Message }}
            </div>
            <div class="row home--header">
                <h1>Event.club</h1>
            </div>
            <div class="row home--description">
                <p>Votre page évenement avec protection par mot de passe, réservation en ligne et paiement via Stripe.</p>
            </div>

            <div class="row home--photo">
                <div class="row">
                    <div class="col-sm-12">
                        <h2 class="home--part--title">Photos</h2>
                    </div>
                </div>
                <div class="col-sm-4">
                    <img src="/static/img.png">
                </div>
                <div class="col-sm-4">
                    <img src="/static/img.png">
                </div>
                <div class="col-sm-4">
                    <img src="/static/img.png">
                </div>
            </div>
            <div class="row home--photo">
                <div class="col-sm-4">
                    <img src="/static/img.png">
                </div>
                <div class="col-sm-4">
                    <img src="/static/img.png">
                </div>
                <div class="col-sm-4">
                    <img src="/static/img.png">
                </div>
            </div>
            <div class="row home--booking">
                <div class="row">
                    <div class="col-sm-12">
                        <h2 class="home--part--title">Réservation</h2>
                    </div>
                </div>
                <form id="home--form" method="POST" action="/booking#home--form">
                    <div class="row">
                        <div class="col-sm-6">
                            <div class="row">
                                <div class="col-sm-12">
                                    <label class="home--form--title" for="booking--name">Infos:</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--name">• Nom & Prénom:</label>
                                </div>
                                <div class="col-sm-8">
                                    <input id="booking--name" name="name" type="text" value="{{ .Booking.Name }}" required>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--phone">• Téléphone:</label>
                                </div>
                                <div class="col-sm-8">
                                    <input id="booking--phone" name="phone" type="tel" value="{{ .Booking.Phone }}" required>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--email">• Email:</label>
                                </div>
                                <div class="col-sm-8">
                                    <input id="booking--email" name="email" type="email" value="{{ .Booking.Email }}" required>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-12">
                                    <label class="home--form--title" for="booking--meal1">Repas:</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--meal1">• Midi:</label>
                                </div>
                                <div class="col-sm-8">
                                    <select id="booking--meal1" name="meal1" class="booking--meal">
                                        <option {{ if eq .Booking.Meal1 0 }}selected{{ end }}>0</option>
                                        <option {{ if eq .Booking.Meal1 1 }}selected{{ end }}>1</option>
                                        <option {{ if eq .Booking.Meal1 2 }}selected{{ end }}>2</option>
                                        <option {{ if eq .Booking.Meal1 3 }}selected{{ end }}>3</option>
                                        <option {{ if eq .Booking.Meal1 4 }}selected{{ end }}>4</option>
                                        <option {{ if eq .Booking.Meal1 5 }}selected{{ end }}>5</option>
                                        <option {{ if eq .Booking.Meal1 6 }}selected{{ end }}>6</option>
                                        <option {{ if eq .Booking.Meal1 7 }}selected{{ end }}>7</option>
                                    </select>
                                    personnes
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--meal2">• Soir:</label>
                                </div>
                                <div class="col-sm-8">
                                    <select id="booking--meal2" name="meal2" class="booking--meal">
                                        <option {{ if eq .Booking.Meal2 0 }}selected{{ end }}>0</option>
                                        <option {{ if eq .Booking.Meal2 1 }}selected{{ end }}>1</option>
                                        <option {{ if eq .Booking.Meal2 2 }}selected{{ end }}>2</option>
                                        <option {{ if eq .Booking.Meal2 3 }}selected{{ end }}>3</option>
                                        <option {{ if eq .Booking.Meal2 4 }}selected{{ end }}>4</option>
                                        <option {{ if eq .Booking.Meal2 5 }}selected{{ end }}>5</option>
                                        <option {{ if eq .Booking.Meal2 6 }}selected{{ end }}>6</option>
                                        <option {{ if eq .Booking.Meal2 7 }}selected{{ end }}>7</option>
                                    </select>
                                    personnes
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label class="home--form--title" for="booking--30midi">Total:</label>
                                </div>
                                <div id="booking--price" class="col-sm-8">
                                    <span class="home--form--title">0€</span>
                                </div>
                            </div>
                        </div>
                        <div class="col-sm-6">
                            <div class="row">
                                <div class="col-sm-12">
                                    <label class="home--form--title" for="booking--card">Paiement:</label>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--card">• Numéro de carte:</label>
                                </div>
                                <div class="col-sm-8">
                                    <input id="booking--card" type="text" size="20" data-stripe="number" required>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--exp">• Date d'expiration:</label>
                                </div>
                                <div class="col-sm-8">
                                    <select id="booking--exp" data-stripe="exp_month" required>
                                        <option>1</option>
                                        <option>2</option>
                                        <option>3</option>
                                        <option>4</option>
                                        <option>5</option>
                                        <option>6</option>
                                        <option>7</option>
                                        <option>8</option>
                                        <option>9</option>
                                        <option>10</option>
                                        <option>11</option>
                                        <option>12</option>
                                    </select>
                                    /
                                    <select data-stripe="exp_year" required>
                                        <option>16</option>
                                        <option>17</option>
                                        <option>18</option>
                                        <option>19</option>
                                        <option>20</option>
                                    </select>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col-sm-4">
                                    <label for="booking--cvc">• Cryptogramme:</label>
                                </div>
                                <div class="col-sm-8">
                                    <input id="booking--cvc" type="text" size="4" data-stripe="cvc" required>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-sm-12">
                            <div class="row">
                                <div class="col-sm-12">
                                    <input type="submit" name="Réserver">
                                </div>
                            </div>
                        </div>
                    </div>
                </form>
            </div>
        </div>
        <script type="text/javascript">
            $(".booking--meal").change(function() {
                var price = 0;
                $(".booking--meal").each(function(i, el) {
                    price += parseInt($(el).val());
                })
                $("#booking--price>span").html(({{ .MealPrice }}*price)+"€");
            });
        </script>
        <script type="text/javascript" src="https://js.stripe.com/v2/"></script>
        <script type="text/javascript">
            Stripe.setPublishableKey({{ .StripeKey }});
            $(function() {
                // Grab the form:
                $('#home--form').submit(function(event) {

                    // Disable the submit button to prevent repeated clicks:
                    $('#home--form').find('#booking--submit').prop('disabled', true);

                    // Request a token from Stripe:
                    Stripe.card.createToken($('#home--form'), function(status, response) {

                        // Grab the form:
                        var $form = $('#home--form');

                        if (response.error) { // Problem!

                            // Show the errors on the form:
                            $('.home--message').text(response.error.message);
                            $('.home--message').addClass("error")
                            $form.find('.submit').prop('disabled', false); // Re-enable submission

                        } else { // Token was created!

                            // Get the token ID:
                            var token = response.id;

                            // Insert the token ID into the form so it gets submitted to the server:
                            $form.append($('<input type="hidden" name="token">').val(token));

                            // Submit the form:
                            $form.get(0).submit();
                        }
                    });

                    // Prevent the form from being submitted:
                    return false;
                });
            });
        </script>
    </body>
</html>
{{end}}