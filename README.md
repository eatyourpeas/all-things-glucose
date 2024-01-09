# An API for all things glucose

Written in Golang, this API is a service to bring all glucose calculations in to one place. There are plenty of calculators out there, but how many of them are available as an API? And how many of them are written in GO?!

This intial simple API is a Glucose Infusion Rate calculator. Helpful for calculating glucose requirement, especially where hyperinsulinism is suspected.

This includes a list of specialist and standard infant milks with glucose contents per 100g to all the user more accurately to calculate a GIR, when glucose comes through the drip, the PN and the milk all at the same time.

## Getting started

1. clone the repo
2. go run .
3. curl http://localhost:8080/milks

### Example

request:

```curl
curl -H 'Content-Type: application/json' \
      -d '{ "weight":2.5,"rate":5, "percentage": 10}' \
      -X POST \
      http://localhost:8080/glucose-infusion-rate
```

response:

```console
{
    "data": 3.3333333
}
```
