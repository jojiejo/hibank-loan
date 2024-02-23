# HiBank Loan API

[https://github.com/jojiejo/hibank-loan](https://github.com/jojiejo/hibank-loan)

This is a repository for loan module in HiBank

## Installation

Ensure you have Go installed on your machine. Clone the repository:

```bash
git clone https://github.com/yourusername/yourprojectname.git
```

Install dependencies:

```bash
go mod download
```

## Usage

Run the application:

```bash
go run main.go
```

The application will start running at `http://localhost:8080`.

## Usage

Run using docker:

```bash
docker-compose up
```

## Endpoints

### [POST] /api/hello

Create a new loan

#### Request

- Method: GET
- Path: `/api/loan`

```json
{
    "plafond":10000000,
    "duration":12,
    "interest_rate":5,
    "start_date":"2023-01-02T00:00:00.000Z"
}
```

#### Response

- Status: 200 OK
- Body:

```json
[
    {
        "installment_number": 1,
        "date": "2023-01-02",
        "annuity": 856074.82,
        "installment_principle": 814408,
        "installment_interest": 41667,
        "installment_balance": 9185592
    }
]
```