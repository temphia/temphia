export const load = ({ url }) => {
    const params = url.searchParams
    return { user: params.get("id") }
}
