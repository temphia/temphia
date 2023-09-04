export const load = ({ url }) =>
    ({ get: (name) => (url.searchParams.get(name)) })
