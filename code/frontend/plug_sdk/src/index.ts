export const TEMPHIA_EXEC_KEY = "__temphia_exec_token__"

export const getTemphiaExecKey = (redirect: boolean) => {
    try {
        const token = localStorage.getItem(TEMPHIA_EXEC_KEY)
        if (!token) {
            if (redirect) {
                window.location.pathname = `/z/pages/agent/inject?redirect=${window.location}`
            }

            return null
        }
        
        return JSON.parse(token)
    } catch (error) {
        return null
    }
}