import {
  Info,
  OpenApiV3,
  SecurityRequirement,
  Server,
  type OpenApiV3_1,
} from "fluid-oas";
import { addOpenApiRoutes } from "./paths";
import { addOpenApiComponents } from "./schema";
const CWD = import.meta.dir;

const OAS_PATH = CWD + "/../openapi.json";

async function main() {
  const info = Info.addTitle("Hangouts.ai")
    .addVersion("0.0.1")
    .addDescription(
      "Plan hangouts with your friends quickly and easily with AI!",
    );
  let oas = OpenApiV3.addOpenApiVersion("3.1.1").addInfo(info);
  oas = addOpenApiRoutes(oas);
  oas = addOpenApiComponents(oas);
  oas = addOpenApiServers(oas);
  oas = oas.addSecurity([
    SecurityRequirement.addSecurityRequirement({ bearerAuth: [] }),
  ]);
  oas.writeOASSync(OAS_PATH);
}

function addOpenApiServers(oas: OpenApiV3_1) {
  return oas.addServers([Server.addUrl("http://localhost:8081")]);
}

main();
