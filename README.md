# Fitbit Web API - Go client

This package allows you to create a "server" application that interacts with the [Fitbit API](https://dev.fitbit.com/reference/web-api/) in Go.

## Usage

The prerequisite is to follow the [Getting Started](https://dev.fitbit.com/build/reference/web-api/developer-guide/getting-started/) official guide.

Once you have created a Fitbit Developer account you have to register a new application, of **server** type. Then, there are 2 required steps:

1. Server application Authorization flow.

   `galeone/fitbit` gives you the basic functionalities for implementing the [Authorization Code Grant Flow with PKCE](https://dev.fitbit.com/build/reference/web-api/developer-guide/authorization/#Authorization-Code-Grant-Flow-with-PKCE). You need to be familiar with some Web Framework for implementing what's described in the [Server application Authorization flow](#server-application-authorization-flow).)
2. Client usage.

   Once the user granted the permissions to your application, you can use the authorized client for querying the Fitbit API.
   **NOTE**: only the GET queries are supported right now. It means that you can fetch everything (in a very convenient format, using annotated Go Structs), but you can't do POST/PUT/DELETE operations.

### Server application Authorization flow

You also need to be familiar with some Web Framework (not shown).

1. Create a type that implements the `fitbit.Storage` interface. You can see an implementation based on PostgreSQL, through the package [galeone/igor](https://github.com/galeone/igor) here: [galeone/fitbit-pgdb](https://github.com/galeone/fitbit-pgdb).
1. Create a `fitbit.Authorizer` object
1. Create the endpoint for the authorization flow. The content should look like
   ```go
    fitbitAuthorizer := fitbit.NewAuthorizer(_db, _clientID, _clientSecret, _redirectURL)

    authorizing := types.AuthorizingUser{
        CSRFToken: uuid.New().String(),
        // Code verifier for PKCE
        // https://dev.fitbit.com/build/reference/web-api/developer-guide/authorization/#Authorization-Code-Grant-Flow-with-PKCE
        Code: fmt.Sprintf("%s-%s", uuid.New().String(), uuid.New().String()),
    }

    fitbitAuthorizer.SetAuthorizing(&authorizing)

    // Potentially set cookie for identifying the authorizing user
    c.SetCookie(&http.Cookie{
        Name: "authorizing",
        Value: fitbitAuthorizer.CSRFToken().String(),
        // No Expires = Session cookie
        HttpOnly: true,
    })

    if err = _db.InsertAuhorizingUser(&authorizing); err != nil {
        return err
    }

    var auth_url *url.URL
    if auth_url, err = fitbitAuthorizer.AuthorizationURL(); err != nil {
        return err
    }

    c.Redirect(http.StatusTemporaryRedirect, auth_url.String())
   ```
1. Create the endpoint for the Redirect URI. The content should look like
   ```go
    state := c.QueryParam("state")
    if state != fitbitAuthorizer.CSRFToken().String() {
        return c.Redirect(http.StatusTemporaryRedirect, "/error?status=csrf")
    }

    code := c.QueryParam("code")
    var token *types.AuthorizedUser
    var err error
    if token, err = fitbitAuthorizer.ExchangeAuthorizationCode(code); err != nil {
        return c.Redirect(http.StatusTemporaryRedirect, "/error?status=exchange")
    }
    // Update the fitbitclient. Now it contains a valid token and HTTP can be used to query the API
    fitbitAuthorizer.SetToken(token)

    // Save token and redirect user to the application
    if err = _db.UpsertAuthorizedUser(token); err != nil {
        return err
    }
    // Cookie used to identify the user that authorized the application
    cookie := http.Cookie{
        Name:     "token",
        Value:    token.AccessToken,
        Domain:   _domain,
        Expires:  time.Now().Add(time.Second * time.Duration(token.ExpiresIn)),
        HttpOnly: true,
    }
    c.SetCookie(&cookie)
    // Redirect the user to your application endpoint
    c.Redirect(http.StatusTemporaryRedirect, "/app")
   ```

That's all.

### Client usage

After the user authorized the application, you can re-create the `fitbit.Authorizer` fetching the data from the database (using your `Storage` implementation) and create the authorized `fitbit.Client`.

```go
fitbitAuthorizer := fitbit.NewClient(_db, _clientID, _clientSecret, _redirectURL)

// Auhtorization token (after exhange)
cookie, err = c.Cookie("token")

var dbToken *types.AuthorizedUser
if dbToken, err = _db.AuthorizedUser(cookie.Value); err != nil {
    return err
}

// Set the valid token
fitbitAuthorizer.SetToken(dbToken)

// Create the client
var fb *client.Client
if fb, err = api.NewAPI(fitbitAuthorizer); err != nil {
    return err
}

// Use it!

var logs *types.ActivityLogList
if logs, err = fb.UserActivityLogList(&types.Pagination{
    Offset:     0,
    BeforeDate: types.FitbitDateTime{Time: time.Now()},
    Limit:      10,
    Sort:       "desc",
}); err != nil {
    return
}

for _, activity := range logs.Activities {
    if activity.TcxLink != "" {
        var tcxDB *tcx.TCXDB
        if tcxDB, err = fb.UserActivityTCX(activity.LogID); err != nil {
            return
        }
        // So something with the tcxDB
    }
   // Do something with the activity
}
```
