package budget

import (
	"fmt"
	"terra/types/assets"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Params oracle parameters
type Params struct {
	ActiveThreshold sdk.Dec       `json:"active_threshold"` // threshold of vote that will transition a program open -> active budget queue
	LegacyThreshold sdk.Dec       `json:"legacy_threshold"` // threshold of vote that will transition a program active -> legacy budget queue
	VotePeriod      time.Duration `json:"vote_period"`      // vote period
	MinDeposit      sdk.Coin      `json:"min_deposit"`      // Minimum deposit in TerraSDR
}

// NewParams creates a new param instance
func NewParams(activeThreshold sdk.Dec, legacyThreshold sdk.Dec, votePeriod time.Duration, minDeposit sdk.Coin) Params {
	return Params{
		ActiveThreshold: activeThreshold,
		LegacyThreshold: legacyThreshold,
		VotePeriod:      votePeriod,
		MinDeposit:      minDeposit,
	}
}

// DefaultParams creates default oracle module parameters
func DefaultParams() Params {

	defaultVotePeriod, _ := time.ParseDuration("730h") // 1 month
	return NewParams(
		sdk.NewDecWithPrec(1, 1), // 10%
		sdk.NewDecWithPrec(0, 2), // 0%
		defaultVotePeriod,
		sdk.NewInt64Coin(assets.SDRDenom, 100),
	)
}

func validateParams(params Params) error {
	if params.ActiveThreshold.LT(sdk.ZeroDec()) {
		return fmt.Errorf("budget active threshold should be greater than 0, is %s", params.ActiveThreshold.String())
	}
	if params.LegacyThreshold.LT(sdk.ZeroDec()) {
		return fmt.Errorf("budget legacy threshold should be greater than or equal to 0, is %s", params.LegacyThreshold.String())
	}
	if params.VotePeriod < 0 {
		return fmt.Errorf("budget parameter VotePeriod must be > 0, is %s", params.VotePeriod.String())
	}
	if params.MinDeposit.Amount.LTE(sdk.ZeroInt()) {
		return fmt.Errorf("budget parameter MinDeposit must be > 0, is %v", params.MinDeposit.String())
	}
	return nil
}
