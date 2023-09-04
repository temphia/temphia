import { derived, writable } from "svelte/store";
import { getStores, navigating, page, updated } from '$app/stores';

export const params = derived([navigating], ([ndata]) => {
    let params
    if (ndata) {
        params = ndata.to.url.searchParams
    } else {
        params = new URLSearchParams(location.search)
    }

    let o = {}
    
    params.forEach((v, k) => { o[k] = v })
    return o
})


params.subscribe((v) => {
    console.log("@params", v)
})