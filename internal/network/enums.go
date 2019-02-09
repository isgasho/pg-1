// Copyright 2019 MQ, Inc. All rights reserved.
//
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file in the root of the source
// tree.

package network

import "log"

type TransactionStatus byte

const (
	TransactionStatusIdle                = 'I'
	TransactionStatusIdleInTransaction   = 'T'
	TransactionStatusInFailedTransaction = 'E'
)

func (t TransactionStatus) String() string {
	switch t {
	case TransactionStatusIdle:
		return "idle"
	case TransactionStatusIdleInTransaction:
		return "idle in transaction"
	case TransactionStatusInFailedTransaction:
		return "in a failed transaction"
	default:
		log.Panicf("unknown transactionStatus %d", t)
		return ""
	}
}

type Identifies byte

const (
	IdentifiesAuth                 = 'R'
	IdentifiesBackendKeyData       = 'K'
	IdentifiesBind                 = 'B'
	IdentifiesBindComplete         = '2'
	IdentifiesCancelRequest        = 0
	IdentifiesClose                = 'C'
	IdentifiesCloseComplete        = '3'
	IdentifiesCopyData             = 'd'
	IdentifiesCopyDone             = 'c'
	IdentifiesCopyFail             = 'f'
	IdentifiesCopyInResponse       = 'G'
	IdentifiesCopyOutResponse      = 'H'
	IdentifiesCopyBothResponse     = 'W'
	IdentifiesDataRow              = 'D'
	IdentifiesDescribe             = 'D'
	IdentifiesEmptyQueryResponse   = 'I'
	IdentifiesErrorResponse        = 'E'
	IdentifiesExecute              = 'E'
	IdentifiesFlush                = 'H'
	IdentifiesFunctionCall         = 'F'
	IdentifiesFunctionCallResponse = 'V'
	IdentifiesGSSResponse          = 'p'
	IdentifiesNoData               = 'n'
	IdentifiesNoticeResponse       = 'N'
	IdentifiesNotificationResponse = 'A'
	IdentifiesParameterDescription = 't'
	IdentifiesParameterStatus      = 'S'
	IdentifiesParse                = 'P'
	IdentifiesParseComplete        = '1'
	IdentifiesPasswordMessage      = 'p'
	IdentifiesPortalSuspended      = 's'
	IdentifiesQuery                = 'Q'
	IdentifiesReadyForQuery        = 'Z'
	IdentifiesRowDescription       = 'T'
	IdentifiesSASLInitialResponse  = 'p'
	IdentifiesSASLResponse         = 'p'
	IdentifiesSSLRequest           = '8'
	IdentifiesStartupMessage       = 0
	IdentifiesSync                 = 'S'
	IdentifiesTerminate            = 'X'
)
