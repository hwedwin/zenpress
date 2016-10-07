// Copyright 2015 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

const (
	// DefaultMySQLState is default state of the mySQL
	DefaultMySQLState = "HY000"
)

// MySQLState maps error code to MySQL SQLSTATE value.
// The values are taken from ANSI SQL and ODBC and are more standardized.
var MySQLState = map[uint16]string{
	ErrDupKey:                              "23000",
	ErrOutofMemory:                         "HY001",
	ErrOutOfSortMemory:                     "HY001",
	ErrConCount:                            "08004",
	ErrBadHost:                             "08S01",
	ErrHandshake:                           "08S01",
	ErrDBaccessDenied:                      "42000",
	ErrAccessDenied:                        "28000",
	ErrNoDB:                                "3D000",
	ErrUnknownCom:                          "08S01",
	ErrBadNull:                             "23000",
	ErrBadDB:                               "42000",
	ErrTableExists:                         "42S01",
	ErrBadTable:                            "42S02",
	ErrNonUniq:                             "23000",
	ErrServerShutdown:                      "08S01",
	ErrBadField:                            "42S22",
	ErrWrongFieldWithGroup:                 "42000",
	ErrWrongSumSelect:                      "42000",
	ErrWrongGroupField:                     "42000",
	ErrWrongValueCount:                     "21S01",
	ErrTooLongIdent:                        "42000",
	ErrDupFieldName:                        "42S21",
	ErrDupKeyName:                          "42000",
	ErrDupEntry:                            "23000",
	ErrWrongFieldSpec:                      "42000",
	ErrParse:                               "42000",
	ErrEmptyQuery:                          "42000",
	ErrNonuniqTable:                        "42000",
	ErrInvalidDefault:                      "42000",
	ErrMultiplePriKey:                      "42000",
	ErrTooManyKeys:                         "42000",
	ErrTooManyKeyParts:                     "42000",
	ErrTooLongKey:                          "42000",
	ErrKeyColumnDoesNotExits:               "42000",
	ErrBlobUsedAsKey:                       "42000",
	ErrTooBigFieldlength:                   "42000",
	ErrWrongAutoKey:                        "42000",
	ErrForcingClose:                        "08S01",
	ErrIpsock:                              "08S01",
	ErrNoSuchIndex:                         "42S12",
	ErrWrongFieldTerminators:               "42000",
	ErrBlobsAndNoTerminated:                "42000",
	ErrCantRemoveAllFields:                 "42000",
	ErrCantDropFieldOrKey:                  "42000",
	ErrBlobCantHaveDefault:                 "42000",
	ErrWrongDBName:                         "42000",
	ErrWrongTableName:                      "42000",
	ErrTooBigSelect:                        "42000",
	ErrUnknownProcedure:                    "42000",
	ErrWrongParamcountToProcedure:          "42000",
	ErrUnknownTable:                        "42S02",
	ErrFieldSpecifiedTwice:                 "42000",
	ErrUnsupportedExtension:                "42000",
	ErrTableMustHaveColumns:                "42000",
	ErrUnknownCharacterSet:                 "42000",
	ErrTooBigRowsize:                       "42000",
	ErrWrongOuterJoin:                      "42000",
	ErrNullColumnInIndex:                   "42000",
	ErrPasswordAnonymousUser:               "42000",
	ErrPasswordNotAllowed:                  "42000",
	ErrPasswordNoMatch:                     "42000",
	ErrWrongValueCountOnRow:                "21S01",
	ErrInvalidUseOfNull:                    "22004",
	ErrRegexp:                              "42000",
	ErrMixOfGroupFuncAndFields:             "42000",
	ErrNonexistingGrant:                    "42000",
	ErrTableaccessDenied:                   "42000",
	ErrColumnaccessDenied:                  "42000",
	ErrIllegalGrantForTable:                "42000",
	ErrGrantWrongHostOrUser:                "42000",
	ErrNoSuchTable:                         "42S02",
	ErrNonexistingTableGrant:               "42000",
	ErrNotAllowedCommand:                   "42000",
	ErrSyntax:                              "42000",
	ErrAbortingConnection:                  "08S01",
	ErrNetPacketTooLarge:                   "08S01",
	ErrNetReadErrorFromPipe:                "08S01",
	ErrNetFcntl:                            "08S01",
	ErrNetPacketsOutOfOrder:                "08S01",
	ErrNetUncompress:                       "08S01",
	ErrNetRead:                             "08S01",
	ErrNetReadInterrupted:                  "08S01",
	ErrNetErrorOnWrite:                     "08S01",
	ErrNetWriteInterrupted:                 "08S01",
	ErrTooLongString:                       "42000",
	ErrTableCantHandleBlob:                 "42000",
	ErrTableCantHandleAutoIncrement:        "42000",
	ErrWrongColumnName:                     "42000",
	ErrWrongKeyColumn:                      "42000",
	ErrDupUnique:                           "23000",
	ErrBlobKeyWithoutLength:                "42000",
	ErrPrimaryCantHaveNull:                 "42000",
	ErrTooManyRows:                         "42000",
	ErrRequiresPrimaryKey:                  "42000",
	ErrKeyDoesNotExits:                     "42000",
	ErrCheckNoSuchTable:                    "42000",
	ErrCheckNotImplemented:                 "42000",
	ErrCantDoThisDuringAnTransaction:       "25000",
	ErrNewAbortingConnection:               "08S01",
	ErrMasterNetRead:                       "08S01",
	ErrMasterNetWrite:                      "08S01",
	ErrTooManyUserConnections:              "42000",
	ErrReadOnlyTransaction:                 "25000",
	ErrNoPermissionToCreateUser:            "42000",
	ErrLockDeadlock:                        "40001",
	ErrNoReferencedRow:                     "23000",
	ErrRowIsReferenced:                     "23000",
	ErrConnectToMaster:                     "08S01",
	ErrWrongNumberOfColumnsInSelect:        "21000",
	ErrUserLimitReached:                    "42000",
	ErrSpecificAccessDenied:                "42000",
	ErrNoDefault:                           "42000",
	ErrWrongValueForVar:                    "42000",
	ErrWrongTypeForVar:                     "42000",
	ErrCantUseOptionHere:                   "42000",
	ErrNotSupportedYet:                     "42000",
	ErrWrongFkDef:                          "42000",
	ErrOperandColumns:                      "21000",
	ErrSubqueryNo1Row:                      "21000",
	ErrIllegalReference:                    "42S22",
	ErrDerivedMustHaveAlias:                "42000",
	ErrSelectReduced:                       "01000",
	ErrTablenameNotAllowedHere:             "42000",
	ErrNotSupportedAuthMode:                "08004",
	ErrSpatialCantHaveNull:                 "42000",
	ErrCollationCharsetMismatch:            "42000",
	ErrWarnTooFewRecords:                   "01000",
	ErrWarnTooManyRecords:                  "01000",
	ErrWarnNullToNotnull:                   "22004",
	ErrWarnDataOutOfRange:                  "22003",
	WarnDataTruncated:                      "01000",
	ErrWrongNameForIndex:                   "42000",
	ErrWrongNameForCatalog:                 "42000",
	ErrUnknownStorageEngine:                "42000",
	ErrTruncatedWrongValue:                 "22007",
	ErrSpNoRecursiveCreate:                 "2F003",
	ErrSpAlreadyExists:                     "42000",
	ErrSpDoesNotExist:                      "42000",
	ErrSpLilabelMismatch:                   "42000",
	ErrSpLabelRedefine:                     "42000",
	ErrSpLabelMismatch:                     "42000",
	ErrSpUninitVar:                         "01000",
	ErrSpBadselect:                         "0A000",
	ErrSpBadreturn:                         "42000",
	ErrSpBadstatement:                      "0A000",
	ErrUpdateLogDeprecatedIgnored:          "42000",
	ErrUpdateLogDeprecatedTranslated:       "42000",
	ErrQueryInterrupted:                    "70100",
	ErrSpWrongNoOfArgs:                     "42000",
	ErrSpCondMismatch:                      "42000",
	ErrSpNoreturn:                          "42000",
	ErrSpNoreturnend:                       "2F005",
	ErrSpBadCursorQuery:                    "42000",
	ErrSpBadCursorSelect:                   "42000",
	ErrSpCursorMismatch:                    "42000",
	ErrSpCursorAlreadyOpen:                 "24000",
	ErrSpCursorNotOpen:                     "24000",
	ErrSpUndeclaredVar:                     "42000",
	ErrSpFetchNoData:                       "02000",
	ErrSpDupParam:                          "42000",
	ErrSpDupVar:                            "42000",
	ErrSpDupCond:                           "42000",
	ErrSpDupCurs:                           "42000",
	ErrSpSubselectNyi:                      "0A000",
	ErrStmtNotAllowedInSfOrTrg:             "0A000",
	ErrSpVarcondAfterCurshndlr:             "42000",
	ErrSpCursorAfterHandler:                "42000",
	ErrSpCaseNotFound:                      "20000",
	ErrDivisionByZero:                      "22012",
	ErrIllegalValueForType:                 "22007",
	ErrProcaccessDenied:                    "42000",
	ErrXaerNota:                            "XAE04",
	ErrXaerInval:                           "XAE05",
	ErrXaerRmfail:                          "XAE07",
	ErrXaerOutside:                         "XAE09",
	ErrXaerRmerr:                           "XAE03",
	ErrXaRbrollback:                        "XA100",
	ErrNonexistingProcGrant:                "42000",
	ErrDataTooLong:                         "22001",
	ErrSpBadSQLstate:                       "42000",
	ErrCantCreateUserWithGrant:             "42000",
	ErrSpDupHandler:                        "42000",
	ErrSpNotVarArg:                         "42000",
	ErrSpNoRetset:                          "0A000",
	ErrCantCreateGeometryObject:            "22003",
	ErrTooBigScale:                         "42000",
	ErrTooBigPrecision:                     "42000",
	ErrMBiggerThanD:                        "42000",
	ErrTooLongBody:                         "42000",
	ErrTooBigDisplaywidth:                  "42000",
	ErrXaerDupid:                           "XAE08",
	ErrDatetimeFunctionOverflow:            "22008",
	ErrRowIsReferenced2:                    "23000",
	ErrNoReferencedRow2:                    "23000",
	ErrSpBadVarShadow:                      "42000",
	ErrSpWrongName:                         "42000",
	ErrSpNoAggregate:                       "42000",
	ErrMaxPreparedStmtCountReached:         "42000",
	ErrNonGroupingFieldUsed:                "42000",
	ErrForeignDuplicateKeyOldUnused:        "23000",
	ErrCantChangeTxCharacteristics:         "25001",
	ErrWrongParamcountToNativeFct:          "42000",
	ErrWrongParametersToNativeFct:          "42000",
	ErrWrongParametersToStoredFct:          "42000",
	ErrDupEntryWithKeyName:                 "23000",
	ErrXaRbtimeout:                         "XA106",
	ErrXaRbdeadlock:                        "XA102",
	ErrFuncInexistentNameCollision:         "42000",
	ErrDupSignalSet:                        "42000",
	ErrSignalWarn:                          "01000",
	ErrSignalNotFound:                      "02000",
	ErrSignalException:                     "HY000",
	ErrResignalWithoutActiveHandler:        "0K000",
	ErrSpatialMustHaveGeomCol:              "42000",
	ErrDataOutOfRange:                      "22003",
	ErrAccessDeniedNoPassword:              "28000",
	ErrTruncateIllegalFk:                   "42000",
	ErrDaInvalidConditionNumber:            "35000",
	ErrForeignDuplicateKeyWithChildInfo:    "23000",
	ErrForeignDuplicateKeyWithoutChildInfo: "23000",
	ErrCantExecuteInReadOnlyTransaction:    "25006",
	ErrAlterOperationNotSupported:          "0A000",
	ErrAlterOperationNotSupportedReason:    "0A000",
	ErrDupUnknownInIndex:                   "23000",
}
