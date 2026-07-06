package keeper

import (
	"bytes"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bnb-chain/greenfield/x/sp/types"
)

func (k Keeper) GetStorageProvider(ctx sdk.Context, id uint32) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetStorageProviderKey(k.spSequence.EncodeSequence(id)))
	if value == nil {
		return sp, false
	}

	sp = types.MustUnmarshalStorageProvider(k.cdc, value)
	return sp, true
}

func (k Keeper) MustGetStorageProvider(ctx sdk.Context, id uint32) *types.StorageProvider {
	sp, found := k.GetStorageProvider(ctx, id)
	if !found {
		panic("must get storage provider, but it not exist")
	}
	return sp
}

func (k Keeper) GetStorageProviderByOperatorAddr(ctx sdk.Context, opAddr sdk.AccAddress) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.GetStorageProviderByOperatorAddrKey(opAddr))
	if id == nil {
		return sp, false
	}

	return k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(id))
}

func (k Keeper) GetStorageProviderByFundingAddr(ctx sdk.Context, fundAddr sdk.AccAddress) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.GetStorageProviderByFundingAddrKey(fundAddr))
	if id == nil {
		return sp, false
	}

	return k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(id))
}

func (k Keeper) GetStorageProviderBySealAddr(ctx sdk.Context, sealAddr sdk.AccAddress) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.GetStorageProviderBySealAddrKey(sealAddr))
	if id == nil {
		return sp, false
	}

	return k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(id))
}

func (k Keeper) GetStorageProviderByApprovalAddr(ctx sdk.Context, approvalAddr sdk.AccAddress) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.GetStorageProviderByApprovalAddrKey(approvalAddr))
	if id == nil {
		return sp, false
	}

	return k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(id))
}

func (k Keeper) GetStorageProviderByGcAddr(ctx sdk.Context, gcAddr sdk.AccAddress) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)

	id := store.Get(types.GetStorageProviderByGcAddrKey(gcAddr))
	if id == nil {
		return sp, false
	}

	return k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(id))
}

func (k Keeper) SetStorageProvider(ctx sdk.Context, sp *types.StorageProvider) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalStorageProvider(k.cdc, sp)

	store.Set(types.GetStorageProviderKey(k.spSequence.EncodeSequence(sp.Id)), bz)
}

func (k Keeper) SetStorageProviderByOperatorAddr(ctx sdk.Context, sp *types.StorageProvider) {
	operatorAddr := sp.GetOperatorAccAddress()
	store := ctx.KVStore(k.storeKey)

	store.Set(types.GetStorageProviderByOperatorAddrKey(operatorAddr), k.spSequence.EncodeSequence(sp.Id))
}

func (k Keeper) SetStorageProviderByFundingAddr(ctx sdk.Context, sp *types.StorageProvider) {
	fundAddr := sp.GetFundingAccAddress()
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetStorageProviderByFundingAddrKey(fundAddr), k.spSequence.EncodeSequence(sp.Id))
}

func (k Keeper) SetStorageProviderBySealAddr(ctx sdk.Context, sp *types.StorageProvider) {
	sealAddr := sp.GetSealAccAddress()
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetStorageProviderBySealAddrKey(sealAddr), k.spSequence.EncodeSequence(sp.Id))
}

func (k Keeper) SetStorageProviderByApprovalAddr(ctx sdk.Context, sp *types.StorageProvider) {
	approvalAddr := sp.GetApprovalAccAddress()
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetStorageProviderByApprovalAddrKey(approvalAddr), k.spSequence.EncodeSequence(sp.Id))
}

func (k Keeper) SetStorageProviderByGcAddr(ctx sdk.Context, sp *types.StorageProvider) {
	gcAddr := sp.GetGcAccAddress()
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetStorageProviderByGcAddrKey(gcAddr), k.spSequence.EncodeSequence(sp.Id))
}

func (k Keeper) GetAllStorageProviders(ctx sdk.Context) (sps []types.StorageProvider) {
	store := ctx.KVStore(k.storeKey)

	iter := storetypes.KVStorePrefixIterator(store, types.StorageProviderKey)

	for ; iter.Valid(); iter.Next() {
		sp := types.MustUnmarshalStorageProvider(k.cdc, iter.Value())
		sps = append(sps, *sp)
	}
	return sps
}

func (k Keeper) Exit(ctx sdk.Context, sp *types.StorageProvider) error {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetStorageProviderByOperatorAddrKey(sdk.MustAccAddressFromHex(sp.OperatorAddress)))
	store.Delete(types.GetStorageProviderByFundingAddrKey(sdk.MustAccAddressFromHex(sp.FundingAddress)))
	store.Delete(types.GetStorageProviderBySealAddrKey(sdk.MustAccAddressFromHex(sp.SealAddress)))
	store.Delete(types.GetStorageProviderByApprovalAddrKey(sdk.MustAccAddressFromHex(sp.ApprovalAddress)))
	store.Delete(types.GetStorageProviderByGcAddrKey(sdk.MustAccAddressFromHex(sp.GcAddress)))
	store.Delete(types.GetStorageProviderKey(k.spSequence.EncodeSequence(sp.Id)))
	store.Delete(types.GetStorageProviderByBlsKeyKey(types.GetStorageProviderByBlsKeyKey(sp.GetBlsKey())))
	return nil
}

func (k Keeper) SetStorageProviderByBlsKey(ctx sdk.Context, sp *types.StorageProvider) {
	blsPk := sp.GetBlsKey()
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetStorageProviderByBlsKeyKey(blsPk), k.spSequence.EncodeSequence(sp.Id))
}

func (k Keeper) GetStorageProviderByBlsKey(ctx sdk.Context, blsPk []byte) (sp *types.StorageProvider, found bool) {
	store := ctx.KVStore(k.storeKey)
	id := store.Get(types.GetStorageProviderByBlsKeyKey(blsPk))
	if id == nil {
		return sp, false
	}
	return k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(id))
}

// PruneStaleSecondaryIndexes removes seal/approval/gc/bls secondary-index entries
// whose key no longer matches the current operational identity of the SP they
// resolve to. It repairs stale bindings left by historical EditStorageProvider
// calls that wrote the new index without revoking the previous one, and drops
// orphan entries whose SP no longer exists. Intended to run once in an upgrade
// handler; safe to re-run (idempotent).
func (k Keeper) PruneStaleSecondaryIndexes(ctx sdk.Context) (removed int) {
	store := ctx.KVStore(k.storeKey)

	indexes := []struct {
		prefix  []byte
		current func(sp *types.StorageProvider) []byte
	}{
		{types.StorageProviderBySealAddrKey, func(sp *types.StorageProvider) []byte { return sp.GetSealAccAddress().Bytes() }},
		{types.StorageProviderByApprovalAddrKey, func(sp *types.StorageProvider) []byte { return sp.GetApprovalAccAddress().Bytes() }},
		{types.StorageProviderByGcAddrKey, func(sp *types.StorageProvider) []byte { return sp.GetGcAccAddress().Bytes() }},
		{types.StorageProviderByBlsPubKeyKey, func(sp *types.StorageProvider) []byte { return sp.GetBlsKey() }},
	}

	for _, idx := range indexes {
		var staleKeys [][]byte
		iter := storetypes.KVStorePrefixIterator(store, idx.prefix)
		for ; iter.Valid(); iter.Next() {
			keyID := append([]byte(nil), iter.Key()...)
			indexedValue := keyID[len(idx.prefix):]
			sp, found := k.GetStorageProvider(ctx, k.spSequence.DecodeSequence(iter.Value()))
			if !found || !bytes.Equal(indexedValue, idx.current(sp)) {
				staleKeys = append(staleKeys, keyID)
			}
		}
		iter.Close()

		for _, key := range staleKeys {
			store.Delete(key)
			removed++
		}
	}
	return removed
}
