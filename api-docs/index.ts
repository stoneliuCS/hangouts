import { Info, OpenApiV3 } from "fluid-oas";
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
  oas.writeOASSync(OAS_PATH);
}

main();
