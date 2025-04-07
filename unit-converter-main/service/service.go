package service

import (
	"log/slog"

	"github.com/fentezi/unit-converter/model"
)

type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{
		log: log,
	}
}

func (s *Service) Convert(request model.Request) model.Response {
	switch request.UnitForm {
	case "m":
		switch request.UnitTo {
		case "km":
			return model.Response{Result: request.Input / 1000, Unit: "km"}
		case "ft":
			return model.Response{Result: request.Input * 3.28084, Unit: "ft"}
		case "yd":
			return model.Response{Result: request.Input * 1.09361, Unit: "yd"}
		case "mi":
			return model.Response{Result: request.Input / 1609.34, Unit: "mi"}
		}
	case "km":
		switch request.UnitTo {
		case "m":
			return model.Response{Result: request.Input * 1000, Unit: "m"}
		case "ft":
			return model.Response{Result: request.Input * 3280.84, Unit: "ft"}
		case "yd":
			return model.Response{Result: request.Input * 1093.61, Unit: "yd"}
		case "mi":
			return model.Response{Result: request.Input / 1.60934, Unit: "mi"}
		}
	case "ft":
		switch request.UnitTo {
		case "m":
			return model.Response{Result: request.Input / 3.28084, Unit: "m"}
		case "km":
			return model.Response{Result: request.Input / 3280.84, Unit: "km"}
		case "yd":
			return model.Response{Result: request.Input / 3, Unit: "yd"}
		case "mi":
			return model.Response{Result: request.Input / 5280, Unit: "mi"}
		}
	case "yd":
		switch request.UnitTo {
		case "m":
			return model.Response{Result: request.Input / 1.09361, Unit: "m"}
		case "km":
			return model.Response{Result: request.Input / 1093.61, Unit: "km"}
		case "ft":
			return model.Response{Result: request.Input * 3, Unit: "ft"}
		case "mi":
			return model.Response{Result: request.Input / 1760, Unit: "mi"}
		}
	case "mi":
		switch request.UnitTo {
		case "m":
			return model.Response{Result: request.Input * 1609.34, Unit: "m"}
		case "km":
			return model.Response{Result: request.Input * 1.60934, Unit: "km"}
		case "ft":
			return model.Response{Result: request.Input * 5280, Unit: "ft"}
		case "yd":
			return model.Response{Result: request.Input * 1760, Unit: "yd"}
		}
	case "mm":
		switch request.UnitTo {
		case "cm":
			return model.Response{Result: request.Input / 10, Unit: "cm"}
		case "km":
			return model.Response{Result: request.Input / 10000000, Unit: "km"}
		}
	case "cm":
		switch request.UnitTo {
		case "mm":
			return model.Response{Result: request.Input * 10, Unit: "mm"}
		case "km":
			return model.Response{Result: request.Input / 100000, Unit: "km"}
		}
	case "kg":
		switch request.UnitTo {
		case "g":
			return model.Response{Result: request.Input * 1000, Unit: "g"}
		case "lb":
			return model.Response{Result: request.Input * 2.205, Unit: "lb"}
		}
	case "g":
		switch request.UnitTo {
		case "kg":
			return model.Response{Result: request.Input / 1000, Unit: "kg"}
		case "lb":
			return model.Response{Result: request.Input / 453.592, Unit: "lb"}
		}
	case "lb":
		switch request.UnitTo {
		case "kg":
			return model.Response{Result: request.Input / 2.205, Unit: "kg"}
		case "g":
			return model.Response{Result: request.Input * 453.592, Unit: "g"}
		}
	case "C":
		switch request.UnitTo {
		case "F":
			return model.Response{Result: request.Input*1.8 + 32, Unit: "F"}
		}
	case "F":
		switch request.UnitTo {
		case "C":
			return model.Response{Result: (request.Input - 32) / 1.8, Unit: "C"}
		}
	case "K":
		switch request.UnitTo {
		case "C":
			return model.Response{Result: request.Input - 273.15, Unit: "C"}
		}
	}

	return model.Response{Unit: "unknown"}
}
