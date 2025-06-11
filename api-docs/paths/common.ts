import { MediaType, Response } from "fluid-oas";
import { COMPONENT_NAME_MAPPINGS } from "../schema";
import { ERROR_SCHEMA } from "../schema/error";

export const DEFAULT_ERROR_RESPONSE = Response.addDescription(
  "Server request failed.",
).addContents({
  "application/json": MediaType.addSchema(
    COMPONENT_NAME_MAPPINGS.get(ERROR_SCHEMA)!,
  ),
});
