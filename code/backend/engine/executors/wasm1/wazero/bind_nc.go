package wazero

import "context"

func NcPut(ctx context.Context, kPtr, kLen, vPtr, vLen, respOffset, respLen int32, expire int64) int32 {
	e := getCtx(ctx)
	err := e.bindNcache.Put(e.getString(kPtr, kLen), e.getBytes(vPtr, vLen), expire)
	return e.writeFinal(respOffset, respLen, err)
}

func NcPutCAS(ctx context.Context, kPtr, kLen, vPtr, vLen, respOffset, respLen int32, version, expire int64) int32 {
	e := getCtx(ctx)
	err := e.bindNcache.PutCAS(e.getString(kPtr, kLen), e.getBytes(vPtr, vLen), version, expire)
	return e.writeFinal(respOffset, respLen, err)
}

func NcGet(ctx context.Context, kPtr, kLen, respOffset, respLen, rexpire, rversion int32) int32 {
	e := getCtx(ctx)
	data, ver, expire, err := e.bindNcache.Get(e.getString(kPtr, kLen))
	if err != nil {
		e.writeBytesNPtr(data, respOffset, respLen)
		return 0
	}

	e.writeBytesNPtr(data, respOffset, respLen)
	e.mem.WriteUint64Le(e.context, uint32(rexpire), uint64(expire))
	e.mem.WriteUint64Le(e.context, uint32(rversion), uint64(ver))

	return 1
}

func NcExpire(ctx context.Context, kPtr, kLen, respOffset, respLen int32) int32 {
	e := getCtx(ctx)
	err := e.bindNcache.Expire(e.getString(kPtr, kLen))
	return e.writeFinal(respOffset, respLen, err)
}
