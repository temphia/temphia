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

export const hashedPosCalc = (
  hash: string,
  maxWidth: number,
  maxHeight: number,
  middlePoint: { left: number; top: number },
  radius: number
) => {
  const hashValue = hash.split("").reduce((acc, char) => {
    acc = (acc << 5) - acc + char.charCodeAt(0);
    return acc & acc;
  }, 0);

  const angle = (Math.abs(hashValue) % 360) * (Math.PI / 180); // Convert angle to radians
  const left = middlePoint.left + radius * Math.cos(angle);
  const top = middlePoint.top + radius * Math.sin(angle);

  // Limit the position within the max width and height
  const limitedLeft = Math.min(Math.max(0, left), maxWidth);
  const limitedTop = Math.min(Math.max(0, top), maxHeight);

  return { top: limitedTop, left: limitedLeft };
};

/*
id="tail-{plug.id}-{agent.id}"
id="agent-out-port-{plug.id}-{agent.id}"
id="agent-in-port-{plug.id}-{agent.id}"

  {
    "agent_links": [
      {
        "id": 1,
        "name": "link2adaper",
        "from_plug_id": "chc50lom4q7efu3enuq0",
        "from_agent_id": "default",
        "to_plug_id": "adapter-1",
        "to_agent_id": "default"
      }
    ]
  }
*/
export const generateAgentLinkIds = (rawdata) => {
  return (rawdata["agent_links"] || []).map((link) => [
    `agent-out-port-${link.from_plug_id}-${link.from_agent_id}`,
    `agent-in-port-${link.to_plug_id}-${link.to_agent_id}`,
  ]);
};
