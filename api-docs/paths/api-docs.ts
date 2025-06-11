import {
  MediaType,
  Operation,
  PathItem,
  Response,
  Responses,
  String,
} from "fluid-oas";
import { DEFAULT_ERROR_RESPONSE } from "./common";

export const API_DOCS = PathItem.addMethod({
  get: Operation.addSummary("API documentation.").addResponses(
    Responses.addResponses({
      "200": Response.addDescription("API Documentation Page.").addContents({
        "text/html": MediaType.addSchema(String),
      }),
    }).addDefault(DEFAULT_ERROR_RESPONSE),
  ),
});
