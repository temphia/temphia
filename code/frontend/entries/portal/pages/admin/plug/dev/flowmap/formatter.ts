export interface FormatedPlug {
  plug: any;
  agents: any[];
  resources: { [_: string]: any[] };
  exts: { [_: string]: any[] };
}


export const formatFlowData = (rawdata: {
  [_: string]: any;
}): FormatedPlug[] => {
  const { plugs, agents, agent_resources, exts } = rawdata;

  const formattedPlugs: FormatedPlug[] = plugs.map(
    (plug: { [_: string]: any }) => {
      const { id, name, bprint_id } = plug;

      const plugAgents = agents.filter(
        (agent: { [_: string]: any }) => agent.plug_id === id
      );
      const plugResources = (agent_resources || [])
        .filter((resource: { [_: string]: any }) => resource.plug_id === id)
        .reduce((obj: { [_: string]: any }, resource: { [_: string]: any }) => {
          const resarray = obj[resource.agent_id] || [];
          obj[resource.agent_id] = resarray;
          return obj;
        }, {});

      const plugExts = (exts || [])
        .filter((ext: { [_: string]: any }) => ext.plug_id === id)
        .reduce((obj: { [_: string]: any }, ext: { [_: string]: any }) => {
          const extarray = obj[ext.agent_id] || [];
          obj[ext.agent_id] = extarray;

          return obj;
        }, {});

      return {
        plug: { id, name, bprint_id },
        agents: plugAgents,
        resources: plugResources,
        exts: plugExts,
      };
    }
  );

  return formattedPlugs;
};
