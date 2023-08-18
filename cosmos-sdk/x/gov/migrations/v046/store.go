package v046

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v042 "github.com/cosmos/cosmos-sdk/x/gov/migrations/v042"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// migrateProposals migrates all legacy proposals into MsgExecLegacyContent
// proposals.
func migrateProposals(store sdk.KVStore, cdc codec.BinaryCodec) error {
	propStore := prefix.NewStore(store, v042.ProposalsKeyPrefix)

	iter := propStore.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var oldProp govv1beta1.Proposal
		err := cdc.Unmarshal(iter.Value(), &oldProp)
		if err != nil {
			return err
		}

		newProp, err := convertToNewProposal(oldProp)
		if err != nil {
			return err
		}
		bz, err := cdc.Marshal(&newProp)
		if err != nil {
			return err
		}

		// Set new value on store.
		propStore.Set(iter.Key(), bz)
	}

	return nil
}

// migrateVotes migrates all v1beta1 weighted votes (with sdk.Dec as weight)
// to v1 weighted votes (with string as weight)
func migrateVotes(store sdk.KVStore, cdc codec.BinaryCodec) error {
	votesStore := prefix.NewStore(store, v042.VotesKeyPrefix)

	iter := votesStore.Iterator(nil, nil)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		var oldVote govv1beta1.Vote
		err := cdc.Unmarshal(iter.Value(), &oldVote)
		if err != nil {
			return err
		}

		newVote := govv1.Vote{
			ProposalId: oldVote.ProposalId,
			Voter:      oldVote.Voter,
		}
		newOptions := make([]*govv1.WeightedVoteOption, len(oldVote.Options))
		for i, o := range oldVote.Options {
			newOptions[i] = &govv1.WeightedVoteOption{
				Option: govv1.VoteOption(o.Option),
				Weight: o.Weight.String(), // Convert to decimal string
			}
		}
		newVote.Options = newOptions
		bz, err := cdc.Marshal(&newVote)
		if err != nil {
			return err
		}

		// Set new value on store.
		votesStore.Set(iter.Key(), bz)
	}

	return nil
}

// MigrateStore performs in-place store migrations from v2 (v0.43) to v3 (v0.46). The
// migration includes:
//
// - Migrate proposals to be Msg-based.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	if err := migrateVotes(store, cdc); err != nil {
		return err
	}

	return migrateProposals(store, cdc)
}

// Migrate_V046_4_To_V046_5 is a helper function to migrate chains from <=v0.46.6
// to v0.46.7 ONLY.
//
// IMPORTANT: Please do not use this function if you are upgrading to v0.46
// from <=v0.45.
//
// This function migrates the store in-place by fixing the gov votes weight to
// be stored as decimals strings (instead of the sdk.Dec BigInt representation).
//
// The store is expected to be the gov store, and not any prefixed substore.
func Migrate_V046_6_To_V046_7(store sdk.KVStore, cdc codec.BinaryCodec) error {
	return migrateVotes(store, cdc)
}
