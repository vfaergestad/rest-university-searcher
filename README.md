# Assignment 2: REST-Covid-Info

This project is a submission to the second assignment in PROG-2005: Cloud Technologies. 
It is a REST web application in Golang that provides the client with the ability to retrieve information 
about Corona cases occurring in different countries, as well as the number and stringency of current policies in place.

[TOC]

## Endpoints

The service is provided by the following endpoints:
```
/corona/v1/cases/
/corona/v1/policy/
/corona/v1/status/
/corona/v1/notifications/
```

### Covid-19 Cases per Country

This endpoint returns the latest number of cases and deaths for a given country, alongside growth rate of cases.

#### Request

```
Method: GET
Path: /corona/v1/cases/{:country}
```
```{:country}``` is the country for which the information is requested. 
This can be the name of a country, or a country code.

Example request: ```/corona/v1/cases/Norway```

#### Response

Content type: ```application/json```  
Status codes:
- 200: Success
- 400: Bad request; something was wrong with the request, like missing the country name, or to many parameters.
- 404: Not found; no data was found for the requested country.
- 405: Method not allowed; only GET is allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```
{
    "country": <contry_name>,        (string)
    "date": <scope>,                 (string)
    "confirmed": <confirmed_cases>,  (int)
    "recovered": <recovered_cases>,  (int)
    "deaths": <deaths>,              (int)
    "growth_rate": <growth_rate>     (float)
}
```

#### Example requests and responses

Request: ```/corona/v1/cases/Norway```

Response:
```json
{
    "country": "Norway",
    "date": "2022-04-03",
    "confirmed": 1408708,
    "recovered": 0,
    "deaths": 2518,
    "growth_rate": 0.0003202540445387454
}
```

Request: ```/corona/v1/cases/SWE```
```json
{
    "country": "Sweden",
    "date": "2022-04-03",
    "confirmed": 2487852,
    "recovered": 0,
    "deaths": 18365,
    "growth_rate": 0
}
```

### Covid Policy Stringency per Country

#### Request

This endpoint provides an overview of the current stringency level of policies regarding Covid-19 for a given country, 
in addition to the number of currently active policies.

```
Method: GET
Path: /corona/v1/policy/{:country_code}{?scope=YYYY-MM-DD}
```

```{:country_code}``` is the ISO 3166-1 alpha-3 code of the country for which the information is requested.  
```{?scope=YYYY-MM-DD}``` is optional, and if present, it is the date for which policy stringency information is 
requested. If no scope is given, the latest policy stringency information is returned. If there is no avalable data
for the last 7 days, an error is returned.

Example request: ```/corona/v1/policy/NOR?scope=2022-04-03```

#### Response

Content type: ```application/json```  
Status codes:
- 200: Success
- 400: Bad request; something was wrong with the request, like missing the country name, wrong date format, or to many parameters.
- 404: Not found; no policy stringency information available for the given country and scope.
- 405: Method not allowed; only GET is allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```
{
    "country": <country_name>,        (string)
    "date": <scope>,                  (string)
    "stringency": <stringency_level>, (float)
    "policies": <policies>            (int)
}
```

#### Example requests and responses

Request: ```/corona/v1/policy/NOR?scope=2021-04-03```
```json
{
    "country_code": "NOR",
    "scope": "2021-04-03",
    "stringency": 71.3,
    "policies": 23
}
```

Request: ```/corona/v1/policy/SWE```
```json
{
    "country_code": "SWE",
    "scope": "2022-04-02",
    "stringency": 19.44,
    "policies": 0
}
```

### Status Interface

The status interface indicates the availability of all individual services this service depends on.

#### Request

```
Method: GET
Path: /corona/v1/status/
```

#### Response

Content type: ```application/json```
Status codes:
- 200: Success
- 405: Method not allowed; only GET is allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```
{
   "cases_api": "<http status code for *Covid 19 Cases API*>",              (int)
   "policy_api": "<http status code for *Corona Policy Stringency API*>",   (int)
   "country_api": "<http status code for *Country API*>"                    (int)
   "webhooks": <number of registered webhooks>,                             (int)
   "version": <version>,                                                    (string)
   "uptime": <time from the last service restart>                           (string)
}
```

#### Example request and response

Request: ```/corona/v1/status/```

Response:
```json
{
    "cases_api": 200,
    "policy_api": 200,
    "country_api": 200,
    "webhooks": 1,
    "version": "v1",
    "uptime": "24s"
}
```

### Notifications endpoint

Users can register webhooks that are triggered by the service based on how frequent a given country is being invoked, 
where the minimum frequency is specified by the user.

#### Registration of Webhook

##### Request

```
Method: POST
Path: /corona/v1/notifications/
```

Content type: ```application/json```

Body:
```
{
    "url": {url},                   (string)
    "country": {country},           (string)
    "calls": {minimum frequency}    (int)
}
```

```{url}``` is the URL to be triggered upon event (the service that should be invoked).  
```{country}``` is the country for which the trigger applies. Can be the country code (ISO 3166-1 alpha-3) or the country name.  
```{minimum frequency}``` is the minimum number of repeated invocations before notification should occur (i.e., "greater equals").


##### Response

Content type: ```application/json```
Status codes:
- 201: Success; Webhook created.
- 400: Bad request; something was wrong with the request, like missing the country name, wrong date format, or to many parameters.
- 405: Method not allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```
{
    "webhook_id": <webhook_id>, (string)
}
```

##### Example request and response

Request: ```/corona/v1/notifications/```  
Body: 
```
{
    "url": "https://example.com/notification", 
    "country": "NOR", 
    "calls": 3
}
```

Response:
```json
{
    "webhook_id": "d99f52fa1eac2495f8b980e33a848e04823af7926b9363c036ae096be3dc37a5"
}
```

#### Deletion of Webhook

##### Request

```
Method: DELETE
Path: /corona/v1/notifications/{webhook_id}
```

```{webhook_id}``` is the ID of the webhook to be deleted. This is returned during the webhook registration.

##### Response

Content type: ```text/plain```  
Status codes:
- 200: Success; Webhook deleted.
- 400: Bad request; to long path.
- 404: Not found; no webhook with the given ID was found.
- 405: Method not allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```text/plain
confirm/error message
```

##### Example request and response

Request: ```/corona/v1/notifications/d99f52fa1eac2495f8b980e33a848e04823af7926b9363c036ae096be3dc37a5```

Response:
```text/plain
webhook deleted
```

#### View registered webhook

##### Request

```
Method: GET
Path: /corona/v1/notifications/{webhook_id}
```

```{webhook_id}``` is the ID of the webhook to view. This is returned during the webhook registration.

##### Response

Content type: ```application/json```
Status codes:
- 200: Success; Webhook found.
- 400: Bad request; to long path.
- 404: Not found; no webhook with the given ID was found.
- 405: Method not allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```
{
    "webhook_id": <webhook_id>,     (string)
    "url": <url>,                   (string)
    "country": <country>,           (string)
    "calls": <minimum frequency>    (int)
}
```

##### Example request and response

Request: ```/corona/v1/notifications/d99f52fa1eac2495f8b980e33a848e04823af7926b9363c036ae096be3dc37a5```

Response:
```json
{
    "webhook_id": "d99f52fa1eac2495f8b980e33a848e04823af7926b9363c036ae096be3dc37a5",
    "url": "https://example.com/notification",
    "country": "Norway",
    "calls": 3
}
```

##### View all registered webhooks

##### Request

```
Method: GET
Path: /corona/v1/notifications/
```

##### Response

Content type: ```application/json```
Status codes:
- 200: Success; Webhooks retrieved.
- 400: Bad request; to long path.
- 405: Method not allowed.
- 500: Internal server error; something went wrong with the server.

Body:
```
[
    {
        "webhook_id": <webhook_id>,     (string)
        "url": <url>,                   (string)
        "country": <country>,           (string)
        "calls": <minimum frequency>    (int)
    },
    ...
]
```

##### Example request and response

Request: ```/corona/v1/notifications/```

Response:
```json
[
    {
        "webhook_id": "2c5c00f9a7ec9f32b06a3d28b1866120d4af9fa04011206df1e0f1089f40d414",
        "url": "https://example2.com/notification",
        "country": "Sweden",
        "calls": 5
    },
    {
        "webhook_id": "d99f52fa1eac2495f8b980e33a848e04823af7926b9363c036ae096be3dc37a5",
        "url": "https://example.com/notification",
        "country": "Norway",
        "calls": 3
    }
]
```

## Deployment


## Design choices


## Extra features

- Caching of country names and country codes.
- Caching of policies
- Cache entries has expire date
- Logging


## Edge cases

### Resolved

- Policy API data unavailable for today's date.
- Policy API responding with an empty policy when no policy is active.

### Not resolved

- Cases API using other country names than the country API.

## Further improvements

