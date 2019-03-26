package gblcodes

const (
	// HTTPInternalServerError ...
	HTTPInternalServerError appCode = "http.internal_server_error"
	// HTTPBadRequest ...
	HTTPBadRequest appCode = "http.bad_request"
	// HTTPRecordCreated ...
	HTTPRecordCreated appCode = "http.record_created"
	// HTTPRecordFound ...
	HTTPRecordFound appCode = "http.record_found"
	// HTTPRecordNotFound ...
	HTTPRecordNotFound appCode = "http.record_not_found"
	// HTTPRecordsFound ...
	HTTPRecordsFound appCode = "http.records_found"
	// HTTPRecordsNotFound ...
	HTTPRecordsNotFound appCode = "http.records_not_found"
	// HTTPRecordUpdated ...
	HTTPRecordUpdated appCode = "http.record_updated"
	// HTTPJSONSchemaValidationFailed ...
	HTTPJSONSchemaValidationFailed appCode = "http.json_schema_validation_failed"
	// HTTPHeaderRequired ...
	HTTPHeaderRequired appCode = "http.header_required"
)

var httpMessages = map[AppCode]string{
	HTTPBadRequest:                 "The HTTP request could not be processed",
	HTTPInternalServerError:        "An error occured during the request, please note the error ID given",
	HTTPRecordCreated:              "Record Created",
	HTTPRecordFound:                "Record Found",
	HTTPRecordNotFound:             "Record Not Found",
	HTTPRecordsFound:               "Records Found",
	HTTPRecordsNotFound:            "Records Not Found",
	HTTPRecordUpdated:              "Records Updated",
	HTTPJSONSchemaValidationFailed: "The entity/resource that was part of the payload failed the JSON schema validation. Check response for validation errors",
	HTTPHeaderRequired:             "A header is required",
}
