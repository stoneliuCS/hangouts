import {
  MediaType,
  Operation,
  PathItem,
  RequestBody,
  Response,
  Responses,
} from "fluid-oas";
import { USER_REQUEST_SCHEMA, USER_SCHEMA } from "../schema/user";
import { ERROR_SCHEMA } from "../schema/error";
import { DEFAULT_ERROR_RESPONSE } from "./utils";

export const USER_ROUTE = PathItem.addMethod({
  post: Operation.addRequestBody(
    RequestBody.addContents({
      "application/json": MediaType.addSchema(USER_REQUEST_SCHEMA),
    }),
  ).addResponses(
    Responses.addResponses({
      201: Response.addDescription("Successfully created user!").addContents({
        "application/json": MediaType.addSchema(USER_SCHEMA),
      }),
      400: Response.addDescription("Bad request.").addContents({
        "application/json": MediaType.addSchema(ERROR_SCHEMA),
      }),
    }).addDefault(DEFAULT_ERROR_RESPONSE),
  ),
});
