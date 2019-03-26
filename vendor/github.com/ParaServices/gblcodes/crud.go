package gblcodes

// CRUD - Create ...
const (
	CRUDCreateSuccess                  appCode = "crud.create_success"
	CRUDCreateFail                     appCode = "crud.create_fail"
	CRUDCreateInvalidEntity            appCode = "crud.create_invalid_entity"
	CRUDCreateUniqueViolation          appCode = "crud.create_unique_violation"
	CRUDCreateAssociatedRecordNotFound appCode = "crud.create_associated_record_not_found"
	CRUDCreateCheckViolation           appCode = "crud.create_check_violation"
)

// CRUD - Read ...
const (
	CRUDRecordFound         appCode = "crud.record_found"
	CRUDReadFail            appCode = "crud.read_fail"
	CRUDReadRecordNotFound  appCode = "crud.read_record_not_found"
	CRUDReadRecordsNotFound appCode = "crud.read_records_not_found"
)

// CRUD - Update ...
const (
	CRUDUpdateSuccess             appCode = "crud.update_success"
	CRUDUpdateFail                appCode = "crud.update_fail"
	CRUDUpdateFailInvalidEntity   appCode = "crud.update_fail_invalid_entity"
	CRUDUpdateFailDuplicateRecord appCode = "crud.update_fail_duplicate_record"
	CRUDUpdateFailRecordNotFound  appCode = "crud.update_fail_record_not_found"
)

// CRUD - Delete ...
const (
	CRUDDeleteSuccess appCode = "crud.delete_success"
	CRUDDeleteFail    appCode = "crud.delete_fail"
)

// CRUD Misc ...
const (
	CRUDRecordNotFound appCode = "crud.record_not_found"
	CRUDInvalidEntity  appCode = "crud.invalid_entity"
)

var (
	crudMessages = map[AppCode]string{
		CRUDCreateFail:                     "Failed to create record",
		CRUDCreateAssociatedRecordNotFound: "Failed to create record, associated record not found",
		CRUDCreateInvalidEntity:            "Failed to create record because not all required fields are met",
		CRUDCreateUniqueViolation:          "Failed to create record, unique record violation",
		CRUDCreateSuccess:                  "Successfully created record",
		CRUDDeleteFail:                     "Failed to delete record",
		CRUDDeleteSuccess:                  "Successfully deleted record",
		CRUDReadFail:                       "Failed to retrieve record",
		CRUDRecordFound:                    "Successfully retrieved record",
		CRUDReadRecordNotFound:             "Record not found",
		CRUDReadRecordsNotFound:            "No records found",
		CRUDUpdateFail:                     "Failed to update record",
		CRUDUpdateFailDuplicateRecord:      "Failed to update record, duplicate detected",
		CRUDUpdateFailInvalidEntity:        "Failed to update record because not all required fields are met",
		CRUDUpdateSuccess:                  "Successfully updated record",
		CRUDRecordNotFound:                 "Record not found",
		CRUDInvalidEntity:                  "Invalid entity. Check JSON schema validation or query parameter validation",
	}
)
