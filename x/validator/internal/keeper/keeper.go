package keeper

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/validator/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// keeper of the validator store
type Keeper struct {
	// Unexposed key to access store from sdk.Context
	storeKey sdk.StoreKey

	// The wire codec for binary encoding/decoding
	cdc *codec.Codec
}

func NewKeeper(key sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

/*
	Validator by Validator address
*/

// Gets the entire Validator record associated with a validator address
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(types.GetValidatorKey(addr))
	if value == nil {
		panic(fmt.Sprintf("validator record not found for address: %X\n", addr))
	}

	validator = types.MustUnmarshalBinaryBareValidator(k.cdc, value)
	return validator
}

// Sets the entire Validator record for a validator address
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, validator)
	store.Set(types.GetValidatorKey(validator.OperatorAddress), bz)
}

// Check if the Validator record associated with a validator address is present in the store or not
func (k Keeper) IsValidatorPresent(ctx sdk.Context, addr sdk.ValAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorKey(addr))
}

// get the set of all validators
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator, total int) {
	k.IterateValidators(ctx, func(validator types.Validator) (stop bool) {
		validators = append(validators, validator)
		total++
		return false
	})
	return validators, total
}

// iterate over validators and apply function
func (k Keeper) IterateValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.ValidatorKey)
	defer iter.Close()

	for {
		if !iter.Valid() {
			return
		}

		validator := types.MustUnmarshalBinaryBareValidator(k.cdc, iter.Value())

		if process(validator) {
			return
		}

		iter.Next()
	}
}

/*
	Validator Address by Consensus address
*/
// Gets the entire Validator record associated with a consensus address
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	opAddr := store.Get(types.GetValidatorByConsAddrKey(consAddr))
	if opAddr == nil {
		panic(fmt.Errorf("validator with consensus-OperatorAddress %s not found", consAddr))
	}
	return k.GetValidator(ctx, opAddr)
}

// Sets a consensus address to validator address mapping record
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	consAddr := sdk.ConsAddress(validator.GetConsPubKey().Address())
	store.Set(types.GetValidatorByConsAddrKey(consAddr), validator.OperatorAddress)
}

// Check if the Validator record associated with a consensus address is present in the store or not
func (k Keeper) IsValidatorByConsAddrPresent(ctx sdk.Context, consAddr sdk.ConsAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorByConsAddrKey(consAddr))
}

/*
	Last state Validator Index
*/
// Gets validator power in the last state by the given validator address
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, address sdk.ValAddress) (power types.LastValidatorPower) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetValidatorLastPowerKey(address))
	if bz == nil {
		return power
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &power)
	return power
}

// Sets validator power
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, validator types.LastValidatorPower) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(validator)
	store.Set(types.GetValidatorLastPowerKey(validator.OperatorAddress), bz)
}

// Check if the validator power record associated with validator address is present in the store or not
func (k Keeper) IsLastValidatorPowerPresent(ctx sdk.Context, address sdk.ValAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorLastPowerKey(address))
}

// Get active validator set
func (k Keeper) GetLastValidatorPowers(ctx sdk.Context) []types.LastValidatorPower {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorLastPowerKey)
	defer iter.Close()

	var lastValidators []types.LastValidatorPower

	for {
		if !iter.Valid() {
			break
		}

		var validator types.LastValidatorPower

		k.cdc.MustUnmarshalBinaryLengthPrefixed(iter.Value(), &validator)

		lastValidators = append(lastValidators, validator)

		iter.Next()
	}

	return lastValidators
}

func (k Keeper) GetAllLastValidators(ctx sdk.Context) (validators []types.Validator) {
	k.IterateLastValidators(ctx, func(validator types.Validator) (stop bool) {
		validators = append(validators, validator)
		return false
	})
	return validators
}

// Iterate through the active validator set and perform the provided function
func (k Keeper) IterateLastValidators(ctx sdk.Context, process func(validator types.Validator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorLastPowerKey)

	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}

		addr := sdk.ValAddress(iter.Key()[1:])
		validator := k.GetValidator(ctx, addr)

		if process(validator) {
			return
		}

		iter.Next()
	}
}
