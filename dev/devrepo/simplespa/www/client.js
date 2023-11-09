

function RegisterXYZ(params) {
    
}

const getTemphiaExecKey = (redirect) => {
    try {
        const token = localStorage.getItem("__temphia_exec_token__");
        if (!token) {
            if (redirect) {
                window.location.pathname = `/z/pages/agent/inject?redirect=${window.location}`;
            }
            return null;
        }
        return JSON.parse(token);
    }
    catch (error) {
        return null;
    }
};