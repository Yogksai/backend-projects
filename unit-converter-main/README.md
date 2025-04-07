# Unit Conversion API

  

This project provides a simple API for unit conversion between various units of measurement, such as length, weight, and temperature. It is built using the Go programming language with the Gin framework.
https://roadmap.sh/projects/unit-converter


## Features

  

- Convert between units of length (`m`, `km`, `ft`, `mm`, `cm`)

- Convert between units of weight (`kg`, `g`, `lb`)

- Convert between temperature units (`C`, `F`, `K`)

- JSON-based API for unit conversion

  

## Installation

  

1. Clone the repository:
```bash
git clone https://github.com/fentezi/unit-conversion-api.git
```
2. Change into the project directory: 
```bash
cd unit-conversion-api
```
3. Install the dependencies: 
```bash
go mod download
```
4. Build the project:
```bash
go build
```
5.  Run the server:
```bash
/unit-conversion-api
```

  

## API Endpoints

**Convert Units**
 - **Endpoint**: `/convert`
 - **Method**: `POST`
 - **Description**: Converts a given input value from one unit to another.

Request Body
```json
{
  "input": 100,
  "unitForm": "m",
  "unitTo": "km"
}
```
|Field|Type|Desciption
|`input`|float64|The numeric value to convert
|`unitForm`|string|The unit of the input value (e.g., `m`, `kg`)
|`unitTo`|string|The target unit for conversion (e.g., `km`, `lb`)

Response Body

```json
{
  "result": 0.1
}
```
|Field|Type|Desciption
|`result`|float64|The converted value

## Supported Units
**Length:**
 - `m` (Meters)
 - `km` (Kilometers)
 - `ft` (Feet)
 - `mm` (Millimeters)
 - `cm` (Centimeters)
 - `yd` (Yards)
 - `mi` (Miles)
 
 **Weight**
 
 - `kg` (Kilograms)
 - `g` (Grams)
 - `lb` (Pounds)

**Temperature**

 - `C` (Celsius)
 - `F` (Fahrenheit)
 - `K` (Kelvin)

## Logging
This application uses the `slog` logging package for error handling and logging.

## Example Usage
You can test the API using tools like `curl` or Postman. Here's an example using `curl`:
```bash
curl -X POST http://localhost:8080/convert \
-H "Content-Type: application/json" \
-d '{
  "input": 1000,
  "unitForm": "m",
  "unitTo": "km"
}'
```
The response will be:
```json
{
  "result": 1
}
```

## License

This project is licensed under the MIT License.
