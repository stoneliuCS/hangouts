import {
  MediaType,
  Operation,
  PathItem,
  Response,
  Responses,
  String,
} from "fluid-oas";

export const API_DOCS = PathItem.addMethod({
  get: Operation.addSummary("API documentation.").addResponses(
    Responses.addResponses({
      "200": Response.addDescription("API Documentation Page.").addContents({
        "text/html": MediaType.addSchema(String),
      }),
    }),
  ),
});
