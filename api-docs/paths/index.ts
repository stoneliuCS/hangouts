import { Path, type OpenApiV3_1 } from "fluid-oas";
import { HEALTHCHECK_ROUTE } from "./healthcheck";
import { API_DOCS } from "./api-docs";
import { USER_ROUTE } from "./user";

export function addOpenApiRoutes(oas: OpenApiV3_1): OpenApiV3_1 {
  const paths = Path.addEndpoints({
    // Health check routes.
    "/healthcheck": HEALTHCHECK_ROUTE,
    // API Documentation lives in the root page for ease of use.
    "/": API_DOCS,
    // User
    "/api/v1/user": USER_ROUTE,
  });
  return oas.addPaths(paths);
}
