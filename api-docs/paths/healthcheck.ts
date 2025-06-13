import {
  MediaType,
  Object,
  Operation,
  PathItem,
  Response,
  Responses,
  String,
} from "fluid-oas";
import { DEFAULT_ERROR_RESPONSE } from "./utils";

export const HEALTHCHECK_ROUTE = PathItem.addMethod({
  get: Operation.addResponses(
    Responses.addResponses({
      "200": Response.addDescription("Server is Healthy!").addContents({
        "application/json": MediaType.addSchema(
          Object.addProperties({
            message: String.addEnums(["OK"]),
          }),
        ),
      }),
    }).addDefault(DEFAULT_ERROR_RESPONSE),
  ),
});
