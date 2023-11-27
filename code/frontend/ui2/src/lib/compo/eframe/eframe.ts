export interface ExecData {
    tenant_id: string
    plug_id: string
    agent_id: string
    exec_token: string
    startup_payload?: any
}

export const BuildExecURL = (opts: object) => {
    
    let start = "";
    if (opts["auth_type"] === "auto_inject") {
      start = "/z/pages/agent/inject"
    }

    if (opts["start_page"]) {
      if (start === "" ) {
        start = `/${opts["start_page"]}` 
      } else {
        start = `${start}?&start_page=${opts["start_page"]}`        
      }
    }

    return `http://${opts["domain"]}:${
      location.port || 80
    }${start}`
}