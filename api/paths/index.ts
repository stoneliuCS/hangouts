import { Path, type OpenApiV3_1 } from "fluid-oas";
import { HEALTHCHECK_ROUTE } from "./healthcheck";

export function addOpenApiRoutes(oas: OpenApiV3_1): OpenApiV3_1 {
  const paths = Path.addEndpoints({
    "/api/v1/healthcheck": HEALTHCHECK_ROUTE,
  });
  return oas.addPaths(paths);
}
