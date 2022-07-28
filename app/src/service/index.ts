import { spawn, ChildProcessWithoutNullStreams } from "child_process";
import getPort from "get-port";
import path from "path"

export default class Service {
  private binaryFileName: string;
  private port: number;

  constructor(binaryFileName: string) {
    this.binaryFileName = binaryFileName
  }

  getPort(): number {
    return this.port
  }

  // TODO: handle port taken errors
  async spawn(extraArgs?: [string]): Promise<ChildProcessWithoutNullStreams> {
    this.port = await getPort({ port: 6969 });

    return spawn(path.join(process.resourcesPath, 'extraResources', this.binaryFileName), [`--port=${this.port}`, ...(extraArgs ? extraArgs : [])])
  }
}