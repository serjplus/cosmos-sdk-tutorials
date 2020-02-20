package rest

import (
	"net/http"

	"github.com/serjplus/cosmos-sdk-tutorials/haytservice/x/haytservice/internal/types"
	"github.com/serjplus/cosmos-sdk/client/context"

	sdk "github.com/serjplus/cosmos-sdk/types"
	"github.com/serjplus/cosmos-sdk/types/rest"
	"github.com/serjplus/cosmos-sdk/x/auth/client/utils"
)

type buyHaytReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Hayt          string       `json:"hayt"`
	Amount        string       `json:"amount"`
	Buyer         string       `json:"buyer"`
	HaytOwnerName string       `json:"haytownername"`
}

func buyHaytHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyHaytReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Buyer)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		coins, err := sdk.ParseCoins(req.Amount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgBuyHayt(req.Hayt, coins, addr, req.HaytOwnerName)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setHaytReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Hayt          string       `json:"hayt"`
	Value         string       `json:"value"`
	Owner         string       `json:"owner"`
	HaytOwnerName string       `json:"haytownername"`
}

func setHaytHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setHaytReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgSetHayt(req.Hayt, req.Value, addr, req.HaytOwnerName)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteHaytReq struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	Hayt          string       `json:"hayt"`
	Owner         string       `json:"owner"`
	HaytOwnerName string       `json:"haytownername"`
}

func deleteHaytHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteHaytReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.Owner)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// create the message
		msg := types.NewMsgDeleteHayt(req.Hayt, addr, req.HaytOwnerName)
		err = msg.ValidateBasic()
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
