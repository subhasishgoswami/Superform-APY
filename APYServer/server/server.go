package server

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/subhasishgoswami/superform/smartcontract"
	"github.com/subhasishgoswami/superform/subgraph"
)

type Server struct {
	server        *http.Server
	url           string
	subgraph      subgraph.SuperformInterface
	log           *logrus.Entry
	smartcontract smartcontract.GethInterface
}

func NewServer(URL string, Subgraph subgraph.SuperformInterface, smartContract smartcontract.GethInterface, Log logrus.Entry) *Server {
	return &Server{url: URL, subgraph: Subgraph, smartcontract: smartContract, log: &Log}
}

func (reporter *Server) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", reporter.handleLanding).Methods(http.MethodGet)
	r.HandleFunc("/apy/{user:0x[a-fA-F0-9]+}", reporter.handleGetAPY).Methods(http.MethodGet)

	return loggingMiddleware(r, *reporter.log)
}

func (reporter *Server) StartServer() (err error) {
	reporter.log.Info("Superform APY Server")
	reporter.server = &http.Server{
		Addr:    reporter.url,
		Handler: reporter.Routes(),
	}

	err = reporter.server.ListenAndServe()
	return err
}

func (reporter *Server) RespondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	resp := HTTPError{Code: code, Message: message}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		reporter.log.WithField("response", resp).WithError(err).Error("Couldn't write error response")
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (reporter *Server) RespondOK(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		reporter.log.WithField("response", response).WithError(err).Error("Couldn't write OK response")
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func (server *Server) handleLanding(w http.ResponseWriter, req *http.Request) {

	server.RespondOK(w, "Superform APY Server")

}

func (server *Server) handleGetAPY(w http.ResponseWriter, req *http.Request) {

	indexTemplate, err := parseIndexTemplate()
	noDepositsTemplate, err := parseNoDeposits()
	if err != nil {
		server.log.WithError(err).Warn("Error Send Request")
		server.RespondError(w, http.StatusBadRequest, err.Error())
	}
	server.log.WithFields(logrus.Fields{
		"method": "Reporter Block Submissions",
	}).Info("Reporter API")
	vars := mux.Vars(req)
	user := vars["user"]
	userBytes := gethCommon.HexToAddress(user)

	deposits, withdrawals, err := server.subgraph.GetDepositsWithdrawals(user)
	if err != nil {
		server.log.WithError(err).Warn("Error Send Request")
		server.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(deposits) == 0 {
		server.log.Warn("No Deposits")

		if err := noDepositsTemplate.Execute(w, http.StatusInternalServerError); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			server.RespondError(w, http.StatusBadRequest, err.Error())
			return
		}
		return
	}
	balance, err := server.smartcontract.SmartContract.BalanceOf(&bind.CallOpts{}, userBytes)
	if err != nil {
		server.log.WithError(err).Warn("Error Getting Balance")
		server.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	assetsRedemption, err := server.smartcontract.SmartContract.PreviewRedeem(&bind.CallOpts{}, balance)
	if err != nil {
		server.log.WithError(err).Warn("Error Getting Balance")
		server.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}
	assetsRedemptionFloat := new(big.Float)
	assetsRedemptionFloat.SetInt(assetsRedemption)
	var totalDeposits big.Float
	var totalWithdrawals big.Float
	var depositList []string
	var withdrawList []string

	var depositBlockList []uint64
	var withdrawBlockList []uint64

	for _, deposit := range deposits {
		depositAmount := new(big.Float)
		depositAmount, _ = depositAmount.SetString(deposit.AssetAmount)
		_ = totalDeposits.Add(&totalDeposits, depositAmount)
		depositList = append(depositList, deposit.AssetAmount)
		depositBlock, _ := strconv.ParseInt(deposit.BlockNumber, 10, 32)
		depositBlockList = append(depositBlockList, uint64(depositBlock))
	}

	for _, withdrawal := range withdrawals {
		withdrawalAmount := new(big.Float)
		withdrawalAmount, _ = withdrawalAmount.SetString(withdrawal.AssetAmount)
		_ = totalWithdrawals.Add(&totalWithdrawals, withdrawalAmount)
		withdrawList = append(withdrawList, withdrawal.AssetAmount)
		withdrawalBlock, _ := strconv.ParseInt(withdrawal.BlockNumber, 10, 32)
		withdrawBlockList = append(withdrawBlockList, uint64(withdrawalBlock))
	}

	totalEarnings := assetsRedemptionFloat.Add(assetsRedemptionFloat, &totalWithdrawals)
	var totalEatningsFrontend big.Float
	totalEatningsFrontend.Copy(totalEarnings)
	totalProfit := totalEarnings.Sub(totalEarnings, &totalDeposits)
	var totalProfitFrontend big.Float
	totalProfitFrontend.Copy(totalProfit)
	depositTime, err := server.smartcontract.BlockTimeGeth(depositBlockList[0], req.Context())
	if err != nil {
		server.log.WithError(err).Warn("Error Getting Block")
		server.RespondError(w, http.StatusBadRequest, "RPC Is Down")
		return
	}
	totalTime := uint64(time.Now().Unix()) - depositTime
	depositsAPI := new(big.Int)
	depositsAPI, _ = totalDeposits.Int(depositsAPI)
	withdrawalsAPI := new(big.Int)
	withdrawalsAPI, _ = totalWithdrawals.Int(withdrawalsAPI)
	earningsPerETH, _ := totalProfit.Quo(totalProfit, &totalDeposits).Float64()
	apr := (earningsPerETH / float64(totalTime)) * 31536000 * 100
	statusData := struct {
		Balance          string
		User             string
		TotalDeposits    string
		TotalWithdrawals string
		Redemption       string
		Earnings         string
		Profit           string
		APR              float64
		Deposits         []subgraph.Deposit
		Withdrawals      []subgraph.Withdrawal
		DepositList      []string
		WithdrawList     []string
		DepositBlocks    []uint64
		WithdrawalBlocks []uint64
	}{
		balance.String(),
		user,
		depositsAPI.String(),
		withdrawalsAPI.String(),
		assetsRedemption.String(),
		totalEatningsFrontend.String(),
		totalProfitFrontend.String(),
		apr,
		deposits,
		withdrawals,
		depositList,
		withdrawList,
		depositBlockList,
		withdrawBlockList,
	}

	if err := indexTemplate.Execute(w, statusData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
