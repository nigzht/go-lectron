import { spawn } from "child_process";
import getPort from "get-port";

export default class Service {
  private binaryPath: string;
  private port: number;

  constructor(binaryPath: string) {
    this.binaryPath = binaryPath
  }

  async Exec(): Promise<number> {
    this.port = await getPort({ port: 6969 });

    let ls = spawn(this.binaryPath, [`--port=${this.port}`])

    ls.stdout.on('data', function (data) {
      console.log('stdout: ' + data.toString());
    });
    
    ls.stderr.on('data', function (data) {
      console.log('stderr: ' + data.toString());
    });
    
    ls.on('exit', function (code) {
      console.log('child process exited with code ' + code);
    });

    return this.port
  }
}
