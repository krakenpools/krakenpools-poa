package types

import "fmt"

// NewGenesisState creates a new GenesisState object
func NewGenesisState(validators []Validator) GenesisState {
	return GenesisState{
		Validators: validators,
	}
}

// DefaultGenesisState - default GenesisState used by Cosmos Hub
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Validators: []Validator{},
	}
}

// ValidateGenesis validates the poa genesis parameters
func ValidateGenesis(data GenesisState) error {
	if err := validateGenesisStateValidators(data.Validators); err != nil {
		return err
	}

	return nil
}

// Validate the validator set in genesis
func validateGenesisStateValidators(validators []Validator) (err error) {
	addrMap := make(map[string]bool, len(validators))

	for i := 0; i < len(validators); i++ {
		val := validators[i]
		key, err := val.ConsPubKey()
		if err != nil {
			panic(err)
		}
		strKey := string(key.Bytes())

		if _, ok := addrMap[strKey]; ok {
			value, err := val.GetConsAddr()
			if err != nil {

				panic(err)
			}
			return fmt.Errorf("duplicate validator in genesis state: moniker %v, address %v", val.Description.Moniker, value)
		}

		addrMap[strKey] = true
	}
	return
}
