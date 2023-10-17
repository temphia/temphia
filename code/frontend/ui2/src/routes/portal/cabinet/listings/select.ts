import { writable } from "svelte/store";

export const current = writable({
    folder: null,
    item: null
})


export const set = (folder, item) => {

    current.update((old) => {
        if (old.folder === folder && item === old.item) {
            return { folder, item: null }
        }

        return { folder, item }
    })

}

