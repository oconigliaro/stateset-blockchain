package loan

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints
const (
	QueryLoan                 = "loan"
	QueryLoans                = "loans"
	QueryLoansByIDs           = "loans_ids"
	QueryLenderLoan	          = "lender_loans"
	QueryLendersLoans         = "lenders_loans"
	QueryMarketplaceLoan 	  = "invoice_loan"
	QueryMarketplacesLoans	  = "invoices_loans"
	QueryLoansIDRange         = "loans_id_range"
	QueryLoansBeforeTime      = "loans_before_time"
	QueryLoansAfterTime       = "loans_after_time"
	QueryParams               = "params"
)

// QueryLoanParams for a single account
type QueryLoanParams struct {
	ID uint64 `json:"id"`
}

// QueryLoansParams for many account
type QueryAccountsParams struct {
	IDs []uint64 `json:"ids"`
}

// QueryAccountsIDRangeParams for accounts by an id range
type QueryLoansIDRangeParams struct {
	StartID uint64 `json:"start_id"`
	EndID   uint64 `json:"end_id"`
}

// QueryLoansTimeParams for accounts by time
type QueryLoansTimeParams struct {
	CreatedTime time.Time `json:"created_time"`
}

// NewQuerier returns a function that handles queries on the KVStore
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case QueryLoan:
			return queryLoan(ctx, req, keeper)
		case QueryLoans:
			return queryLoans(ctx, req, keeper)
		case QueryLoansByIDs:
			return queryLoansByIDs(ctx, req, keeper)
		case QueryLoansIDRange:
			return queryLoansByIDRange(ctx, req, keeper)
		case QueryLoansBeforeTime:
			return queryLoansBeforeTime(ctx, req, keeper)
		case QueryLoansAfterTime:
			return queryLoansAfterTime(ctx, req, keeper)
		case QueryParams:
			return queryParams(ctx, keeper)
		}

		return nil, sdk.ErrUnknownRequest("Unknown account query endpoint")
	}
}

func queryLoan(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params QueryLoanParams
	codecErr := ModuleCodec.UnmarshalJSON(req.Data, &params)
	if codecErr != nil {
		return nil, ErrJSONParse(codecErr)
	}

	loan, ok := keeper.Loan(ctx, params.ID)
	if !ok {
		return nil, ErrUnknownLoan(params.ID)
	}

	return mustMarshal(loan)
}

func queryLoans(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	loans := keeper.Loans(ctx)

	return mustMarshal(loans)
}

func queryLoansByIDs(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params QueryLoansParams
	codecErr := ModuleCodec.UnmarshalJSON(req.Data, &params)
	if codecErr != nil {
		return nil, ErrJSONParse(codecErr)
	}

	var loans Loans
	for _, id := range params.IDs {
		loan, ok := keeper.Loan(ctx, id)
		if !ok {
			return nil, ErrUnknownLoan(id)
		}
		loans = append(loans, loan)
	}

	return mustMarshal(loans)
}


func queryLoansByIDRange(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params QueryLoansIDRangeParams
	codecErr := ModuleCodec.UnmarshalJSON(req.Data, &params)
	if codecErr != nil {
		return nil, ErrJSONParse(codecErr)
	}
	loans := keeper.LoansBetweenIDs(ctx, params.StartID, params.EndID)

	return mustMarshal(loans)
}

func queryLoansBeforeTime(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params QueryLoansTimeParams
	codecErr := ModuleCodec.UnmarshalJSON(req.Data, &params)
	if codecErr != nil {
		return nil, ErrJSONParse(codecErr)
	}
	loans := keeper.LoansBeforeTime(ctx, params.CreatedTime)

	return mustMarshal(loans)
}

func queryLoansAfterTime(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params QueryLoansTimeParams
	codecErr := ModuleCodec.UnmarshalJSON(req.Data, &params)
	if codecErr != nil {
		return nil, ErrJSONParse(codecErr)
	}
	loans := keeper.LoansAfterTime(ctx, params.CreatedTime)

	return mustMarshal(loans)
}

func queryParams(ctx sdk.Context, keeper Keeper) (result []byte, err sdk.Error) {
	params := keeper.GetParams(ctx)

	result, jsonErr := ModuleCodec.MarshalJSON(params)
	if jsonErr != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marsal result to JSON", jsonErr.Error()))
	}

	return result, nil
}

func mustMarshal(v interface{}) (result []byte, err sdk.Error) {
	result, jsonErr := codec.MarshalJSONIndent(ModuleCodec, v)
	if jsonErr != nil {
		return nil, ErrJSONParse(jsonErr)
	}

	return
}