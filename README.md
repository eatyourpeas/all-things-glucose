# An API for all things glucose

Written in Golang, this API is a service to bring all glucose calculations in to one place. There are plenty of calculators out there, but how many of them are available as an API? And how many of them are written in GO?!

This intial simple API is a Glucose Infusion Rate calculator. Helpful for calculating glucose requirement, especially where hyperinsulinism is suspected.

This includes a list of specialist and standard infant milks with glucose contents per 100g to all the user more accurately to calculate a GIR, when glucose comes through the drip, the PN and the milk all at the same time.

## Getting started

1. clone the repo
2. go run .
3. curl http://localhost:8080/milks

### Example

A 2.5kg baby on iv 10% dextrose running at 15.6ml/h (~150ml/kg/d)

***request***:

```curl
curl -H 'Content-Type: application/json' \
      -d '{ "weight":2.5,"rate":15.6, "percentage": 10}' \
      -X POST \
      http://localhost:8080/dextrose-infusion-rate
```

***response***:

```console
{
    "result": 10.400001
}
```

and

A 3kg infant on 180 ml/kg/d Cow&Gate First infant milk (7.4g/100ml)

***request***

```curl
curl -H 'Content-Type: application/json' \
      -d '{ "weight":3,"volume":180, "id": "4"}' \
      -X POST \
      http://localhost:8080/milk-glucose-infusion-rate
```

***response***

```console
{
    "result": 3.0833333
}
```

and

***request***

```curl
curl -H 'Content-Type: application/json' \
-X GET \
http://localhost:8080/milks
```

***response***

```console
[
    {
        "id": "1",
        "name": "Aptamil 1 First Milk",
        "grams": 7
    },
    {
        "id": "2",
        "name": "Aptamil Profutura 1 First Infant Milk",
        "grams": 7
    },
    {
        "id": "3",
        "name": "Breast Milk (mature)",
        "grams": 7.2
    },
    {
        "id": "4",
        "name": "Cow \u0026 Gate 1 First Infant Milk",
        "grams": 7.4
    },
    {
        "id": "5",
        "name": "Hipp Organic Combiotic First Infant Milk",
        "grams": 7.3
    },
    {
        "id": "6",
        "name": "Holle Organic Infant Formula 1",
        "grams": 7.4
    },
    {
        "id": "7",
        "name": "SMA Pro First Infant Milk",
        "grams": 7.1
    },
    {
        "id": "8",
        "name": "Holle Organic Infant Goat Milk Formula 1",
        "grams": 7.5
    },
    {
        "id": "9",
        "name": "Kabrita Gold 1",
        "grams": 7.3
    },
    {
        "id": "10",
        "name": "NANNYcare First Infant Milk",
        "grams": 7.4
    },
    {
        "id": "11",
        "name": "Aptamil Hungry Milk",
        "grams": 7.8
    },
    {
        "id": "12",
        "name": "Cow \u0026 Gate Infant Milk for Hungrier Babies",
        "grams": 7.8
    },
    {
        "id": "13",
        "name": "Hipp Organic Combiotic Hungry Infant Milk",
        "grams": 7.3
    },
    {
        "id": "14",
        "name": "SMA Extra Hungry",
        "grams": 7
    },
    {
        "id": "15",
        "name": "Aptamil Anti-reflux",
        "grams": 6.8
    },
    {
        "id": "16",
        "name": "Cow \u0026 Gate Anti-reflux",
        "grams": 6.8
    },
    {
        "id": "17",
        "name": "SMA Staydown",
        "grams": 7
    },
    {
        "id": "18",
        "name": "SMA Wysoy",
        "grams": 6.9
    },
    {
        "id": "19",
        "name": "Aptamil Lactose Free",
        "grams": 7.3
    },
    {
        "id": "20",
        "name": "SMA LF",
        "grams": 7.2
    },
    {
        "id": "21",
        "name": "Aptamil Comfort",
        "grams": 7.2
    },
    {
        "id": "22",
        "name": "Cow \u0026 Gate Comfort",
        "grams": 7.2
    },
    {
        "id": "23",
        "name": "SMA Comfort",
        "grams": 7.1
    },
    {
        "id": "24",
        "name": "SMA HA",
        "grams": 7.8
    },
    {
        "id": "25",
        "name": "Aptamil 2 Follow-on Milk",
        "grams": 8.6
    },
    {
        "id": "26",
        "name": "Aptamil Profutura 2 Follow-on Milk",
        "grams": 8.8
    },
    {
        "id": "27",
        "name": "Cow \u0026 Gate 2 Follow-on Milk",
        "grams": 8.6
    },
    {
        "id": "28",
        "name": "Hipp Organic Combiotic Follow-on Milk",
        "grams": 7.9
    },
    {
        "id": "29",
        "name": "Holle Organic Infant Follow-on Formula",
        "grams": 8.2
    },
    {
        "id": "30",
        "name": "SMA Pro Follow-on Milk",
        "grams": 7.9
    },
    {
        "id": "31",
        "name": "Holle Organic Infant Goat Milk Follow-on Formula 2",
        "grams": 8
    },
    {
        "id": "32",
        "name": "Kabrita Gold 2",
        "grams": 8
    },
    {
        "id": "33",
        "name": "NANNYcare Follow-on Milk",
        "grams": 7.4
    },
    {
        "id": "34",
        "name": "Hipp Organic Good Night Milk",
        "grams": 8
    },
    {
        "id": "35",
        "name": "Full-fat cows milk",
        "grams": 4.6
    },
    {
        "id": "36",
        "name": "Semi-skimmed cows Milk",
        "grams": 4.7
    },
    {
        "id": "37",
        "name": "Aptamil 3 Growing Up Milk 1-2 Years",
        "grams": 8.5
    },
    {
        "id": "38",
        "name": "Aptamil Profutura 3 Growing Up Milk",
        "grams": 8.4
    },
    {
        "id": "39",
        "name": "Cow \u0026 Gate 3 Growing Up Name 1-2 Years",
        "grams": 8.5
    },
    {
        "id": "40",
        "name": "Hipp Organic Combiotic Growing-up Milk",
        "grams": 8.2
    },
    {
        "id": "41",
        "name": "Holle Organic Growing-up Milk 3",
        "grams": 8.2
    },
    {
        "id": "42",
        "name": "PaediaSure Shake",
        "grams": 13.2
    },
    {
        "id": "43",
        "name": "SMA Pro Toddler Milk 3",
        "grams": 7
    },
    {
        "id": "44",
        "name": "Kabrita Gold 3",
        "grams": 7.9
    },
    {
        "id": "45",
        "name": "NANNYcare Growing Up Milk",
        "grams": 6.7
    },
    {
        "id": "46",
        "name": "Alpro Soya +1 Complete Care",
        "grams": 8.3
    },
    {
        "id": "47",
        "name": "Aptamil 4 Growing Up Milk 2-3 Years",
        "grams": 6.5
    },
    {
        "id": "48",
        "name": "Cow \u0026 Gate 4 Growing Up Milk 2-3 Years",
        "grams": 6.5
    },
    {
        "id": "49",
        "name": "Aptamil Pepti 1",
        "grams": 7.1
    },
    {
        "id": "50",
        "name": "Cow \u0026 Gate Pepti-junior",
        "grams": 6.8
    },
    {
        "id": "51",
        "name": "Mead Johnson Nutramigen LIPIL 1",
        "grams": 7.5
    },
    {
        "id": "52",
        "name": "SHS Nutricia Pepdite",
        "grams": 7.8
    },
    {
        "id": "53",
        "name": "SHS Nutricia Pepdite MCT",
        "grams": 8.8
    },
    {
        "id": "54",
        "name": "SHS Nutricia Infatrini Peptisorb",
        "grams": 10.3
    },
    {
        "id": "55",
        "name": "Aptamil Pepti 2",
        "grams": 8
    },
    {
        "id": "56",
        "name": "Mead Johnson Pregestimil LIPIL",
        "grams": 6.9
    },
    {
        "id": "57",
        "name": "Mead Johnson Nutramigen LIPIL 2",
        "grams": 8.6
    },
    {
        "id": "58",
        "name": "SHS Nutricia Pepdite 1+",
        "grams": 13
    },
    {
        "id": "59",
        "name": "SHS Nutricia Pepdite MCT 1+",
        "grams": 11.8
    },
    {
        "id": "60",
        "name": "Mead Johnson Nutramigen AA",
        "grams": 7
    },
    {
        "id": "61",
        "name": "SHS Nutricia Neocate LCP",
        "grams": 7.2
    },
    {
        "id": "62",
        "name": "SHS Nutricia GA1 Anamix Infant",
        "grams": 7.4
    },
    {
        "id": "63",
        "name": "SHS Nutricia HCU Anamix Infant",
        "grams": 7.4
    },
    {
        "id": "64",
        "name": "SHS Nutricia IVA Anamix Infant",
        "grams": 7.4
    },
    {
        "id": "65",
        "name": "SHS Nutricia MMA/PA Anamix Infant",
        "grams": 7.4
    },
    {
        "id": "66",
        "name": "SHS Nutricia MSUD Anamix Infant",
        "grams": 7.4
    },
    {
        "id": "67",
        "name": "SHS Nutricia NKH Anamix Infant",
        "grams": 7.3
    },
    {
        "id": "68",
        "name": "SHS Nutricia TYR Anamix Infant",
        "grams": 7.3
    },
    {
        "id": "69",
        "name": "Vitaflo PKU start",
        "grams": 8.3
    },
    {
        "id": "70",
        "name": "Abbott Nutrition Similac High Energy",
        "grams": 10.3
    },
    {
        "id": "71",
        "name": "SHS Nutricia Infatrini",
        "grams": 10.3
    },
    {
        "id": "72",
        "name": "SMA High Energy",
        "grams": 9.8
    },
    {
        "id": "73",
        "name": "Mead Johnson Enfamil AR",
        "grams": 7.6
    },
    {
        "id": "74",
        "name": "Cow \u0026 Gate Infasoy",
        "grams": 7
    },
    {
        "id": "75",
        "name": "SMA Wysoy",
        "grams": 6.9
    },
    {
        "id": "76",
        "name": "Mead Johnson Enfamil O-Lac",
        "grams": 7.2
    },
    {
        "id": "77",
        "name": "SMA LF",
        "grams": 7.2
    },
    {
        "id": "78",
        "name": "Aptamil Preterm",
        "grams": 8.4
    },
    {
        "id": "79",
        "name": "Cow \u0026 Gate Nutriprem 1",
        "grams": 8.4
    },
    {
        "id": "80",
        "name": "SMA Gold Prem 1",
        "grams": 8.4
    },
    {
        "id": "81",
        "name": "Cow \u0026 Gate Nutriprem 2",
        "grams": 7.5
    },
    {
        "id": "82",
        "name": "SMA Gold Prem 2",
        "grams": 7.5
    },
    {
        "id": "83",
        "name": "SHS Nutricia Caprilon",
        "grams": 7
    },
    {
        "id": "84",
        "name": "SHS Nutricia Monogen",
        "grams": 12
    },
    {
        "id": "85",
        "name": "Vitaflo Lipistart",
        "grams": 8.3
    },
    {
        "id": "86",
        "name": "SHS Nutricia Galactomin 17",
        "grams": 7.3
    },
    {
        "id": "87",
        "name": "SHS Nutricia Galactomin 19",
        "grams": 6.4
    },
    {
        "id": "88",
        "name": "SHS Nutricia Kindergen",
        "grams": 11.8
    },
    {
        "id": "89",
        "name": "Vitaflo Renastart",
        "grams": 12.5
    }
]
```
