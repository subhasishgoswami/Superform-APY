package subgraph

import (
	"fmt"
	"net/http"
)

type SuperformInterface struct {
	Client  http.Client
	URL     string
	API_KEY string
}

func superformreq(user string) string {
	return fmt.Sprintf("{\"query\":\"query DepositsWithdrawals($user: Bytes!){\\n\\tdeposits(first:1000, where:{\\n    owner: $user\\n  }){\\n    assets\\n    shares\\n    blockNumber\\n  }\\n  withdraws(first:1000, where:{\\n    owner: $user\\n  }){\\n    assets\\n    shares\\n    blockNumber\\n  }\\n}\",\"variables\":{\"user\":\"%s\"}}", user)
}

type Deposit struct {
	AssetAmount string `json:"assets"`
	ShareAmount string `json:"shares"`
	BlockNumber string `json:"blockNumber"`
}
type Withdrawal struct {
	AssetAmount string `json:"assets"`
	ShareAmount string `json:"shares"`
	BlockNumber string `json:"blockNumber"`
}

type UserDepositsWithdrawals struct {
	Data struct {
		Deposits    []Deposit    `json:"deposits"`
		Withdrawals []Withdrawal `json:"withdraws"`
	} `json:"data"`
}
