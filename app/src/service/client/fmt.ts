import path from "path"
import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { FmtClient } from "../protos/generated/fmt/Fmt";
import { Any } from "../protos/generated/google/protobuf/Any";
const protoPath = require("../protos/fmt.proto")

export class Fmt {
  private fmt: any;
  private client: FmtClient;
  private port: number;

  constructor(port: number) {
    this.port = port;
  }

  async getClient() {
    if (this.client) {
      return this.client;
    }

    await protoLoader.load(path.resolve(__dirname, protoPath), {
      keepCase: true,
      longs: String,
      enums: String,
      defaults: true,
      oneofs: true,
    }).then(packageDefinition => {
      this.fmt = grpc.loadPackageDefinition(packageDefinition).fmt;
    });

    this.client = new this.fmt.Fmt(
      `localhost:${this.port}`,
      grpc.credentials.createInsecure()
    );
  }

  async Print(v: Any) {
    if (!this.client) {
      await this.getClient();
    }

    this.client.Print(v, (err) => err && console.log(err));
  }
}
