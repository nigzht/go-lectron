import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import { FmtClient } from "./proto/fmt/Fmt";

const FMT_PROTO_PATH =
  __dirname + "../../../../proto/fmt.proto";

export default class Fmt {
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

    await protoLoader.load(FMT_PROTO_PATH, {
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

  async Print(v: any) {
    if (!this.client) {
      await this.getClient();
    }

    this.client.Print({ type_url: "type.googleapis.com/google.protobuf.StringValue", value: v }, () => {});
  }
}
