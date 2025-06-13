import { MediaType, Response, String } from "fluid-oas";
import { ERROR_SCHEMA } from "../schema/error";

export const DEFAULT_ERROR_RESPONSE = Response.addDescription(
  "Server request failed.",
).addContents({
  "application/json": MediaType.addSchema(ERROR_SCHEMA),
});

export const ID = String.addDescription("Unique identifier").addFormat(
  "550e8400-e29b-41d4-a716-446655440000",
);

export const BINARY_DATA = String.addDescription("Binary Data")
  .addFormat("binary")
  .addMaxLength(5242880);
