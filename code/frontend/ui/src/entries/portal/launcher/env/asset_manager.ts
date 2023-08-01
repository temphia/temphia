export class EnvAssetManager {
  baseURL: string;
  plugId: string;
  agentId: string;
  executor: string;

  constructor(baseURL: string, plugId: string, agentId: string, exec: string) {
    this.baseURL = baseURL;
    this.plugId = plugId;
    this.agentId = agentId;
    this.executor = exec;
  }

  GetAgentAssetURL(name: string): string {
    return this.agent_url(name);
  }

  GetExecutorAssetURL(name: string): string {
    return `${this.baseURL}/engine/plug/${this.plugId}/agent/${this.agentId}/executor/${this.executor}/${name}`;
  }

  ImportDyanmic(name: string): Promise<any> {
    // fixme => impl
    return Promise.resolve();
  }

  SheduleWorker(name: string): Worker {
    return new Worker(this.agent_url(name));
  }

  private agent_url = (name: string) =>
    `${this.baseURL}/engine/plug/${this.plugId}/agent/${this.agentId}/serve/${name}`;
}
