# googleapiauth

Package `googleapiauth` contains helper functions for easy authentication with Google APIs using the Google Go Client Library.

## How to use

1. Login to your Google Cloud Platform, and go to the [API Console](https://console.developers.google.com/?pli=1).
2. Click on `Credentials`, `Create Credentials` and choose `OAuth Client ID`.
3. Choose `TVs and Limited Input devices` and give an suitable name for the OAuth Client instead of the default name.
4. This is will create a new Client ID for an installed application in the following format.

```json
{
  "installed": {
    "client_id": "****.apps.googleusercontent.com",
    "project_id": "random-project",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_secret": "********",
    "redirect_uris": ["urn:ietf:wg:oauth:2.0:oob", "http://localhost"]
  }
}
```

5. Download the json credentials file and name it appropriately (such as `credentials.json`)
6. The credentials JSON file can be used with this package as follows.
7. Upon successful authentication, a `token.json` file is created in the same folder as the credentials JSON file.

```go
// The following code snippet creates a token.json which is refreshed automatically.
// The scope in this snippet is CalendarReadonlyScope ("https://www.googleapis.com/auth/calendar.readonly)
// The scope enables viewing your calendars.

// 1. Create a new context
ctx := context.Background()
// 2. Get token
ts, err := googleapiauth.GetTokenSource(ctx, calendar.CalendarReadonlyScope)
if err != nil {
  // handle error
}
// 3. Create Google Calendar Service
calService, err := calendar.NewService(ctx, option.WithTokenSource(ts))
if err != nil {
  // handle error
}
// 4. User calService to make API calls
```
