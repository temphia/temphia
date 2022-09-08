export const actionFetch = (actionUrl: string, token: string) => async (name: string, data: string): Promise<Response> => {
    const response = await fetch(`${actionUrl}/${name}`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: token,
        },
        redirect: "follow",
        referrerPolicy: "strict-origin-when-cross-origin",
        body: data,
    });
    return response;
}
