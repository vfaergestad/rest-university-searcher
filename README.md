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

### Request

This endpoint provides an overview of the current stringency level of policies regarding Covid-19 for a given country, 
in addition to the number of currently active policies.

```
Method: GET
Path: /corona/v1/policy/{:country_code}{?scope=YYYY-MM-DD}
```

```{:country_code}``` is the ISO 3166-1 alpha-3 code of the country for which the information is requested.  
```{?scope=YYYY-MM-DD}``` is optional, and if present, it is the date for which policy stringency information is 
requested. If no scope is given, the latest policy stringency information is returned.

Example request: ```/corona/v1/policy/NOR?scope=2022-04-03```

#### Response

Content type: ```application/json```
Status codes:
- 200: Success
- 400: Bad request; something was wrong with the request, like missing the country name, wrong date format, or to many parameters.
- 404: Not found; no policy stringency information available for the given country and scope.
- 405: Method not allowed; only GET is allowed.
- 500: Internal server error; something went wrong with the server.

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

