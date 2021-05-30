// Code generated by goa v3.4.2, DO NOT EDIT.
//
// goatfacts HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/martinohmann/goatops.farm/design

package server

import (
	"context"
	"net/http"

	goatfacts "github.com/martinohmann/goatops.farm/gen/goatfacts"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetFactResponse returns an encoder for responses returned by the
// goatfacts get-fact endpoint.
func EncodeGetFactResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*goatfacts.Fact)
		enc := encoder(ctx, w)
		body := NewGetFactResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetFactRequest returns a decoder for requests sent to the goatfacts
// get-fact endpoint.
func DecodeGetFactRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewGetFactPayload(id)

		return payload, nil
	}
}

// EncodeGetFactError returns an encoder for errors returned by the get-fact
// goatfacts endpoint.
func EncodeGetFactError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetFactNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetFactBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListFactsResponse returns an encoder for responses returned by the
// goatfacts list-facts endpoint.
func EncodeListFactsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]*goatfacts.Fact)
		enc := encoder(ctx, w)
		body := NewListFactsResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetRandomFactResponse returns an encoder for responses returned by the
// goatfacts get-random-fact endpoint.
func EncodeGetRandomFactResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*goatfacts.Fact)
		enc := encoder(ctx, w)
		body := NewGetRandomFactResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetRandomFactError returns an encoder for errors returned by the
// get-random-fact goatfacts endpoint.
func EncodeGetRandomFactError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetRandomFactNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalGoatfactsFactToFact builds a value of type *Fact from a value of type
// *goatfacts.Fact.
func marshalGoatfactsFactToFact(v *goatfacts.Fact) *Fact {
	res := &Fact{
		ID:   v.ID,
		Text: v.Text,
	}

	return res
}
